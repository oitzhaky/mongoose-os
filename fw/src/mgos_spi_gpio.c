/*
 * Copyright (c) 2014-2017 Cesanta Software Limited
 * All rights reserved
 *
 * SPI implementation using GPIO (bit-banging).
 */

#if MGOS_ENABLE_SPI && MGOS_ENABLE_SPI_GPIO

#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#include "fw/src/mgos_gpio.h"
#include "fw/src/mgos_hal.h"
#include "fw/src/mgos_spi.h"
#include "fw/src/mgos_sys_config.h"

#include "common/cs_dbg.h"

struct mgos_spi {
  int miso_gpio;
  int mosi_gpio;
  int sclk_gpio;
  /* Clock polarity, CPOL = 0 -> active high */
  unsigned int cpol : 1;
  /* Clock phase: CPHA = 0 -> sample on inactive to active transition */
  unsigned int cpha : 1;
  unsigned int msb_first : 1;
  unsigned int debug : 1;
};

inline void mgos_spi_clk_idle(struct mgos_spi *c) {
  mgos_gpio_write(c->sclk_gpio, c->cpol); /* CPOL = 0 -> idle at 0 */
}

inline void mgos_spi_clk_active(struct mgos_spi *c) {
  mgos_gpio_write(c->sclk_gpio, !c->cpol);
}

struct mgos_spi *mgos_spi_create(const struct sys_config_spi *cfg) {
  struct mgos_spi *c = (struct mgos_spi *) calloc(1, sizeof(*c));
  if (c == NULL) goto out_err;

  if (cfg->sclk_gpio < 0 || (cfg->miso_gpio < 0 && cfg->mosi_gpio < 0)) {
    goto out_err;
  }

  if (!mgos_spi_configure(c, cfg)) {
    goto out_err;
  }

  LOG(LL_INFO, ("SPI GPIO init ok (MISO: %d, MOSI: %d, SCLK: %d; mode %d)",
                c->miso_gpio, c->mosi_gpio, c->sclk_gpio, cfg->mode));
  return c;

out_err:
  free(c);
  LOG(LL_ERROR, ("Invalid SPI GPIO settings"));
  return NULL;
}

bool mgos_spi_configure(struct mgos_spi *c, const struct sys_config_spi *cfg) {
  if (!mgos_spi_set_mode(c, cfg->mode)) {
    return false;
  }

  if (cfg->miso_gpio >= 0 &&
      mgos_gpio_set_mode(cfg->miso_gpio, MGOS_GPIO_MODE_INPUT) &&
      mgos_gpio_set_pull(cfg->miso_gpio, MGOS_GPIO_PULL_UP)) {
    c->miso_gpio = cfg->miso_gpio;
  } else {
    return false;
  }
  if (cfg->mosi_gpio >= 0 &&
      mgos_gpio_set_mode(cfg->mosi_gpio, MGOS_GPIO_MODE_OUTPUT)) {
    c->mosi_gpio = cfg->mosi_gpio;
  } else {
    return false;
  }
  if (cfg->sclk_gpio >= 0 &&
      mgos_gpio_set_mode(cfg->sclk_gpio, MGOS_GPIO_MODE_OUTPUT)) {
    c->sclk_gpio = cfg->sclk_gpio;
  } else {
    return false;
  }

  if (!mgos_spi_set_msb_first(c, cfg->msb_first)) {
    return false;
  }

  c->debug = cfg->debug;

  mgos_spi_clk_idle(c);

  return true;
}

bool mgos_spi_set_freq(struct mgos_spi *c, int freq) {
  /* ignored */
  (void) c;
  (void) freq;
  return true;
}

bool mgos_spi_set_mode(struct mgos_spi *c, int mode) {
  if (mode < 0 || mode > 3) {
    return false;
  }
  c->cpol = ((mode == 0 || mode == 1) ? 0 : 1);
  c->cpha = ((mode == 0 || mode == 2) ? 0 : 1);
  return true;
}

bool mgos_spi_set_msb_first(struct mgos_spi *c, bool msb_first) {
  c->msb_first = !!msb_first;
  return true;
}

/* This function delays for half of a clock pulse, i.e. quarter of a period. */
inline void mgos_spi_half_delay(struct mgos_spi *c) {
  (void) c;
  /* This is ~70 KHz. TODO(rojer): Make speed configurable. */
  mgos_usleep(1);
}

bool mgos_spi_txn(struct mgos_spi *c, const void *tx_data, void *rx_data,
                  size_t len) {
  size_t i, j;
  if (c->debug) {
    LOG(LL_DEBUG, ("len %d", (int) len));
  }
  if (len > 0) {
    if (c->mosi_gpio < 0 && (rx_data != NULL && c->miso_gpio < 0)) return false;
    for (i = 0; i < len; i++) {
      uint8_t tx_byte = ((uint8_t *) tx_data)[i], rx_byte = 0;
      if (c->debug) LOG(LL_DEBUG, ("write 0x%02x", tx_byte));
      if (c->cpha == 0) {
        for (j = 0; j < 8; j++, tx_byte <<= 1) {
          uint8_t tx_bit = (tx_byte & 0x80 ? 1 : 0);
          mgos_gpio_write(c->mosi_gpio, tx_bit);
          mgos_spi_half_delay(c);
          mgos_spi_clk_active(c);
          rx_byte <<= 1;
          rx_byte |= mgos_gpio_read(c->miso_gpio);
          mgos_spi_half_delay(c);
          mgos_spi_clk_idle(c);
        }
      } else {
        for (j = 0; j < 8; j++, tx_byte <<= 1) {
          uint8_t tx_bit = (tx_byte & 0x80 ? 1 : 0);
          mgos_spi_half_delay(c);
          mgos_spi_clk_active(c);
          mgos_gpio_write(c->mosi_gpio, tx_bit);
          mgos_spi_half_delay(c);
          mgos_spi_clk_idle(c);
          rx_byte <<= 1;
          rx_byte |= mgos_gpio_read(c->miso_gpio);
        }
        mgos_spi_half_delay(c);
      }
      if (c->debug) LOG(LL_DEBUG, ("read 0x%02x", rx_byte));
      if (rx_data != NULL) ((uint8_t *) rx_data)[i] = rx_byte;
    }
  }
  return true;
}

bool mgos_spi_txn_hd(struct mgos_spi *c, const void *tx_data, size_t tx_len,
                     void *rx_data, size_t rx_len) {
  size_t i, j;
  if (c->debug) {
    LOG(LL_DEBUG, ("tx_len %d rx_len %d", (int) tx_len, (int) rx_len));
  }
  if (tx_len > 0) {
    if (c->mosi_gpio < 0) return false;
    for (i = 0; i < tx_len; i++) {
      uint8_t byte = ((uint8_t *) tx_data)[i];
      if (c->debug) LOG(LL_DEBUG, ("write 0x%02x", byte));
      if (c->cpha == 0) {
        for (j = 0; j < 8; j++, byte <<= 1) {
          uint8_t bit = (byte & 0x80 ? 1 : 0);
          mgos_gpio_write(c->mosi_gpio, bit);
          mgos_spi_half_delay(c);
          mgos_spi_clk_active(c);
          /* Sampling takes place here */
          mgos_spi_half_delay(c);
          mgos_spi_clk_idle(c);
        }
      } else {
        for (j = 0; j < 8; j++, byte <<= 1) {
          uint8_t bit = (byte & 0x80 ? 1 : 0);
          mgos_spi_half_delay(c);
          mgos_spi_clk_active(c);
          mgos_gpio_write(c->mosi_gpio, bit);
          mgos_spi_half_delay(c);
          mgos_spi_clk_idle(c);
          /* Sampling takes place here */
        }
        mgos_spi_half_delay(c);
      }
    }
  }
  if (rx_len > 0) {
    if (c->miso_gpio < 0) return false;
    for (i = 0; i < rx_len; i++) {
      uint8_t byte = 0;
      if (c->cpha == 0) {
        for (j = 0; j < 8; j++) {
          byte <<= 1;
          mgos_spi_clk_active(c);
          byte |= mgos_gpio_read(c->miso_gpio);
          mgos_spi_half_delay(c);
          mgos_spi_clk_idle(c);
          mgos_spi_half_delay(c);
        }
      } else {
        for (j = 0; j < 8; j++) {
          byte <<= 1;
          mgos_spi_clk_active(c);
          mgos_spi_half_delay(c);
          mgos_spi_clk_idle(c);
          byte |= mgos_gpio_read(c->miso_gpio);
          mgos_spi_half_delay(c);
        }
        mgos_spi_half_delay(c);
      }
      if (c->debug) LOG(LL_DEBUG, ("read 0x%02x", byte));
      ((uint8_t *) rx_data)[i] = byte;
    }
  }
  return true;
}

void mgos_spi_close(struct mgos_spi *spi) {
  free(spi);
}

#endif /* MGOS_ENABLE_SPI && MGOS_ENABLE_SPI_GPIO */
