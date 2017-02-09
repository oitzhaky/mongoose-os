// Code generated by go-bindata.
// sources:
// data/stub_flasher_esp32.json
// data/stub_flasher_esp8266.json
// DO NOT EDIT!

package esp

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _dataStub_flasher_esp32Json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x57\xbd\xb2\xac\x27\xae\xcd\xef\x53\xec\xda\xf1\x09\x96\x04\x12\xc2\x99\x10\x1f\xaf\xe1\x3a\x75\xed\xd0\x9e\x29\xfb\x4c\x30\x35\x35\xef\x3e\x25\xf8\xfa\x67\xdb\x1d\x34\xdd\x48\x80\x90\x96\x96\xc4\x7f\x3e\xff\xff\x1f\xbf\xfc\xfa\xf3\x9f\x3f\xbe\xff\xf1\xe3\xf3\xa7\x0f\x42\xab\xa5\x90\x0a\x7f\xfb\xd8\x92\xcf\x9f\x3e\x3e\x61\xc0\xaa\x65\x51\x9c\x11\x78\x1b\x67\x59\x30\xc3\xf9\x9c\x91\x6e\x79\xd1\x4a\x00\xd3\xb2\xb5\x9c\x96\xaf\x65\xfb\x7b\xd0\x1a\x6b\x05\x18\xc0\x34\xf0\x45\x4b\xd6\x5a\xfb\xdb\xdc\x7c\xf8\xf0\x23\xed\x86\xab\xa3\xa3\xd5\x19\x57\x87\xa2\x9f\x73\xc2\x60\x47\x83\x0d\x8b\x46\x89\xd2\x61\xcc\x35\xb6\x6c\xd0\x88\x88\x8e\x91\x7a\xa3\x95\x00\x79\x87\x61\xf4\xae\x8a\x3e\x57\xd5\x55\xd2\x9a\xeb\x65\x47\x87\xd1\x5c\x78\x7e\x08\x88\xe0\x80\x56\xd4\x7b\xbc\xef\x03\xc1\xa8\x75\x82\x4b\x30\x8b\x9a\xa0\x95\x3d\x57\xa2\xb0\x28\x04\xed\x12\x15\xa0\x44\x4d\xb9\x43\x8b\xd7\x0a\x45\x3d\x5b\x3b\x3c\x2a\x19\xad\xba\xd6\xb5\x9d\x26\x3a\x1c\xb5\x70\x3a\xb3\xc0\x6b\x20\x42\x0b\x4f\x94\xe2\xe2\x91\x7b\xee\x75\x8c\x55\x45\x27\xaf\x69\x9a\x3e\x74\x38\x96\x18\x5d\xf1\xd8\xe9\x0a\xf7\x02\xe6\x40\xa1\xab\xad\xc5\x70\x09\x70\x2b\xe5\x32\xf7\xca\x55\x51\x5e\x16\x5c\x4f\x0b\x66\xb0\x33\x4f\x5e\xe0\x7d\xbe\xe8\xe0\x75\x1d\x8f\x94\x10\xce\x5f\x25\xf4\x1e\xdb\x1e\x51\x4f\xbc\xd3\x7f\xe8\x15\xb8\xf1\xd1\x62\x3a\xa4\xe2\x0a\x6d\xc8\x4b\x9b\x5a\x8e\xfd\x9e\x17\xf0\x4c\x7f\x0e\x1c\x79\xd4\x33\xae\x5b\x9f\x62\x8e\xd4\xdb\xfe\x26\x04\x44\x58\x89\xdc\x27\xe8\x78\xdf\x28\xe3\x77\x2c\x2f\xc1\x2a\xda\x85\xe4\x44\x41\xd9\x51\x4a\xc6\x81\xc4\x45\x3a\x20\x13\x5e\xa2\xaa\xa8\x3a\x55\x25\xef\x19\xfb\xfc\x6e\xe7\xb7\x9b\x6a\x87\x19\x5d\xba\x56\xf3\xa6\x39\x33\xe8\x92\xb5\x72\x3e\x26\xd4\x3b\x9a\xb3\x03\x62\x69\xcb\xf5\xee\xbb\x4e\x17\xad\xe5\xec\x04\x3a\x68\x1a\x86\xee\xf0\xc1\x38\xff\xbd\xa3\xe7\x6a\x4e\x90\x5d\x8f\x75\xce\x41\x6a\x88\x3a\x2e\xa9\x0d\x09\x60\xa2\x4b\x27\xcc\x94\xaf\xed\x54\xf2\x6b\x2d\x81\x95\x80\xb5\xd1\xa9\x33\x0c\xab\xce\x28\xe6\xa9\x1d\x46\x1e\x2f\x14\x75\x47\xce\x58\xde\x49\x04\x7a\x12\x72\x6b\x4e\x23\xd7\x2f\x9a\x43\x78\x0a\xe1\x20\xa2\x47\xa9\x8c\x8e\x25\x0c\x8b\x4a\x3d\x6f\x19\xfb\xae\x73\xdb\xe5\xc6\x99\xab\xd6\x8a\x8c\xd2\xd9\xe9\x91\x39\x66\x68\x69\xdb\x6a\xc3\xa0\x04\x27\x46\x82\x09\x30\x0a\x7e\x9c\x98\x32\x4e\x4b\x9e\xb2\xb1\x5e\xd6\x98\xa3\x19\x8d\xa7\x67\x5e\x3b\x27\x26\xc1\x79\xf6\xba\xee\x59\x4e\x94\x18\x8d\xfe\xd0\xfe\x62\xdf\x46\x8a\x19\xf4\xfc\x6f\xdb\x2a\xeb\xd0\xd0\x4d\x4c\xdb\xab\xda\xd7\xea\xac\xfa\x88\xc1\xc8\xc8\x56\x97\x32\x1e\xbb\x95\x5c\xc7\x55\x67\xac\xc5\x01\x7f\xf7\x73\xae\x11\x23\x0c\xca\x78\xba\xff\x65\x3e\x30\x54\x4a\x32\x55\x40\xb1\xbd\x5b\x42\xee\xf3\x4b\xa8\x9a\x1e\x3b\x4a\x34\xcd\x4c\x78\xf1\xce\x31\x50\xed\x8c\x72\x8f\xed\x9e\x9f\xc6\xd7\xe6\x23\xf4\x92\x79\x92\xf9\xc1\xa4\xb4\x62\x2d\x04\x1a\xb9\xaa\x74\x68\xa7\xd5\xd6\x52\x9e\x44\x68\xac\x4a\x2e\xc1\x56\xdb\xe8\x60\x28\xe7\xad\x67\x1e\x62\xe3\x61\x39\x94\xda\x39\x29\x91\x56\x21\xa6\xc5\x27\x78\x4c\x90\x20\x12\x13\xb4\xf0\x8a\x97\x3b\x34\xb3\x71\xcc\x34\xf2\x7a\x21\x7a\x42\x73\x4d\xec\x79\x7b\x46\x5f\x9c\x59\xbc\xb5\xd2\x4a\xc3\x92\x9d\x3b\x89\x7b\x71\x27\x95\x92\xac\xaa\xd5\x49\x4c\xf3\xa6\x42\x17\xaf\x65\xcc\x4a\x4e\x2e\x22\x06\x31\x58\x0b\x48\x2b\xe6\x01\xcd\xfb\x1d\xb9\xb5\xd2\xf8\x3a\xbf\x4b\x2b\x16\xd3\x69\xf6\xb5\x46\xfa\xd5\xdd\x8d\xda\x13\xf5\x4e\x53\x57\x56\xb1\x71\x24\xb3\x3d\xef\x1f\xb0\xd0\x84\x97\x44\x31\xdb\x91\x92\xa8\x56\x77\xa4\x24\xc4\x78\xbe\xea\xc3\xb8\xe3\x51\x34\x88\x49\x28\x92\x43\xc2\x37\xf3\x28\xd7\x36\x04\x6c\xca\x35\x63\x24\x95\xa1\x14\x19\x0f\x57\x3a\x7e\x97\x27\x92\xea\xf6\xbb\xef\x48\x19\xc9\x13\xd1\x59\x52\x5e\x11\x88\x99\xd9\x11\x4f\xff\x5b\xa8\xef\xb3\x4a\xdc\x71\xfb\xe2\xf9\x66\x24\x4f\x76\xf4\x68\xf5\xec\xff\x92\xd6\xf5\xf2\x48\x32\xde\xd8\x72\x82\xbb\xa5\x15\x78\x93\x96\x77\x7f\x0d\x7a\x48\x9a\x33\x37\x94\x92\x75\xe2\x2a\xab\x3b\xcd\x99\x75\x3d\x73\x21\x75\x87\x0f\x32\xaa\xed\xcd\xf3\xfd\x7d\xa7\x64\xa8\x23\xd1\x89\xf2\xca\x81\xf2\xec\x2b\x32\x17\xec\xfe\x7f\x3e\xe9\xef\x0d\x8f\x00\x73\x07\x9d\x5e\x82\xf7\xf7\xdd\x0b\x74\x58\x9f\xe0\x67\x67\x60\x1d\xcc\x04\xd1\xca\xbb\xf7\xd0\xb5\x12\xb5\xcf\xde\xc2\x4e\x0d\x4e\x24\xa6\x32\xe5\x02\x32\xea\x4f\x66\xbf\x23\x1e\xb6\xfb\x1c\xaa\x8f\x3e\x20\x13\x33\xb3\x9d\x68\x33\x07\xb1\xd0\x48\x0c\x93\xef\xc8\x2f\x80\xd9\xd7\x5a\xcc\x95\x08\xcc\x48\x4d\x35\x06\x87\x72\x4f\xf6\xda\xbd\x4e\x6a\x81\x99\x39\x2b\xf2\x29\xa3\x79\xef\x64\xd0\xdc\x97\xa0\x42\xbb\xc2\x11\xab\x9a\x4b\x06\x88\x86\x11\xb9\x8c\xb1\x56\x55\xea\x3b\xf2\x94\x9e\x65\x49\x2c\xd4\xad\xeb\x5e\xeb\x43\x37\x8c\x93\x00\x64\xeb\x97\xbf\xed\x2f\x89\xef\x5c\x33\x9c\xfd\x7d\xcd\xb4\x42\x97\x24\xff\x55\xcd\xca\xfe\xbe\xce\x25\x77\x4a\xeb\x95\xf2\x34\xcc\xe7\xca\x74\x38\xb3\x64\xef\x51\xef\xbe\xe5\x78\xc1\x34\x2f\x79\xc9\xa2\xb5\xa0\x8f\xdb\xa6\x2f\xa6\xbc\x18\x34\x35\x5f\xeb\x1c\xb7\x9f\x71\xfb\x79\xcb\x55\x70\xe1\xf6\xea\x69\x08\x8f\x0e\x82\x86\x0b\xe6\x43\x66\x1c\xbc\x12\xd3\x68\x75\x28\x67\x75\xd0\xb0\xb5\xbe\xf4\x79\xbd\xa2\xae\x6a\x95\x6a\x7f\xf6\xb3\x77\xbc\x1f\xa4\x7c\xd0\x97\x1e\xe0\x9d\x6d\x82\x28\x9b\x2d\xe3\xf4\x35\x35\x62\xe7\xba\x46\x3a\xa4\xa2\x58\xe1\x3d\x56\x93\x23\x37\x3e\x7c\x9d\xf8\x35\x5a\xbc\x71\x8b\xcd\xda\x9e\x98\xdb\xb6\x6d\x7c\x3e\x33\xef\x5a\x2b\xe8\x5a\x5b\x67\x38\x22\x02\x81\xf4\xee\x59\xff\xe4\x2d\x77\x0c\xff\xca\xcf\x83\xae\x9e\xab\xbd\xad\x6b\x52\xd9\xbd\x75\xc6\x63\xf7\x2f\xb8\xb2\x4f\xc1\xe9\x7d\x6f\x94\x67\x34\xc8\x37\x2b\xdb\x5e\xfb\x88\x46\xea\x8f\x99\xf1\xa4\x63\x15\x6d\x8d\x67\xbc\x5e\x98\x40\xbd\x24\x69\xf1\x44\xde\x91\x45\xde\x69\xee\x9e\x7f\x3e\xf9\x86\x55\x39\x2b\xc3\xab\x07\xda\xfd\xe4\xee\xab\x6f\xb6\x2b\x46\xfc\x97\xfc\xf3\xda\xf9\x51\xef\x36\x63\x70\x22\x23\xf2\xce\xf3\xc9\x98\x13\xce\xad\xfb\xaa\xc9\x0f\x6d\xcb\x0f\x6e\x0f\xc6\xd6\x93\x8b\xea\x84\xb7\x46\x5e\xcc\xd8\x31\x87\xb5\xee\xd4\xde\x3b\xca\x3d\x1f\xd6\xc8\xd1\xf7\xef\xb9\x7f\x4b\xec\x5a\xbe\x77\x10\x76\xd3\xca\xd9\x00\xd3\x60\x2e\xad\x73\xc4\xa9\x16\x79\x7e\xb9\xed\x7b\xf9\xb2\xa4\x6d\x37\xcb\x2d\xe3\xc3\x6b\x56\xde\xde\x49\x7d\xdf\xe6\xd9\xf9\x24\x3d\x81\xab\x68\xb5\xb5\xfa\x1b\xcb\xfd\x55\xf6\xde\x3d\xf0\xcd\x9f\xfc\xc6\xa3\xf7\xa3\x6c\xf7\xe1\xaf\xf9\xba\xab\xfd\xc5\xcb\xcb\xaa\xc6\x35\xf1\xba\x32\x6f\xf4\x81\x53\x95\xfd\xc2\x98\x9b\x69\xe3\x85\x38\x62\xd1\xc8\xe2\x63\x99\x8f\xd9\xcb\x31\xb8\x10\xca\x46\x60\x32\xb3\xca\xea\xc9\x7d\x8e\x47\x27\xc7\xac\x65\x63\xee\x4a\xe4\xe7\x6b\xe5\xd9\xcb\xf1\x5b\xb7\xc7\xa2\xcd\xde\xee\x59\xd0\x2e\x9b\xad\xd0\xf5\x9c\x2d\x7b\x05\xa3\xf0\xa3\xb7\x66\x43\x11\x6d\xbc\x56\x79\x20\xb8\x3e\xb4\xbb\xa1\x9c\x2a\x96\xfc\xfe\x58\x61\x1d\x65\x18\xa8\x6c\x4e\xb1\x8d\xea\x1b\xf1\xe0\xd2\xcb\xc3\xa3\x66\x27\x7f\xf3\x29\xf9\xc2\xdf\x3b\xea\x32\xfd\x88\x0a\x8c\x2a\xe5\xbb\xb5\x60\x8c\x60\x86\xa3\xe5\x1b\xa5\x7f\xc5\xf1\xe7\xb7\x8f\xcf\x5f\x7f\xff\xf1\xc7\xbf\x9f\x0f\xed\x82\xae\xdf\x3e\x3e\x7f\xff\xd7\x6f\x3f\xff\xf3\xfb\x1f\xdf\x7f\xfb\x33\x25\xdf\x3e\x3e\xcf\x9f\xbf\xbd\xca\xab\x7d\xfb\xf8\xfc\xe5\xfb\x8f\xef\xf9\x2a\x1f\x23\x1b\xc1\x8a\x59\xce\x78\xdd\xff\xd1\x20\x39\xb2\x9f\xb1\x94\x33\x56\x7c\x1d\x3f\xef\xbd\xde\x4f\x29\xd2\x5a\xef\xfa\xdf\xff\xfb\x5f\x00\x00\x00\xff\xff\x18\xd8\x85\x75\x17\x10\x00\x00")

func dataStub_flasher_esp32JsonBytes() ([]byte, error) {
	return bindataRead(
		_dataStub_flasher_esp32Json,
		"data/stub_flasher_esp32.json",
	)
}

func dataStub_flasher_esp32Json() (*asset, error) {
	bytes, err := dataStub_flasher_esp32JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/stub_flasher_esp32.json", size: 4119, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataStub_flasher_esp8266Json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x58\xbb\xd2\x25\xb7\xcd\xcc\xff\xa7\xf8\xea\x8b\x37\x68\x00\x24\x86\x54\x36\x03\x12\xaf\xa1\xda\xfa\xa5\x50\xb2\x4b\x5a\x07\x2e\x97\xdf\xdd\xd5\xe4\x9c\xcb\xae\x1c\xf8\x04\xc3\x19\xde\x00\x02\x8d\x06\x78\xfe\xf5\xf9\xff\x7f\xfb\xe5\xd7\x9f\xff\xfc\xf6\xf5\x8f\x6f\x9f\x3f\x7d\x08\x8e\x72\x74\x14\x94\x2f\x1f\x6b\xe4\xf3\xa7\x8f\x4f\x34\x00\x70\x48\xec\x76\xff\xd8\x36\xe4\xb4\x44\x6b\x77\xdf\x6e\xe5\x1e\x57\xc9\x9e\x59\x24\xaf\xcc\x80\x02\x38\x1a\xb4\xae\xef\x3d\xd6\x24\x6b\x66\x97\xf4\x4c\x3d\xb5\xd4\xb3\x96\x3d\xd3\x1a\x9a\xc1\x70\x94\x11\xc3\xc2\x6f\x1d\x1a\x74\x8f\x9f\x0d\x5d\xae\xcb\xaf\x0e\x2e\xf4\x35\x56\xe5\x72\xf7\x8e\xca\x79\xf5\x30\x87\x94\x0e\xc5\x65\xe6\x0e\x1b\xd9\x3c\x8d\xb2\x67\x3c\x35\xe2\xf8\x48\x3c\x7f\x42\x2d\x4b\x03\x0a\x80\x72\xb2\xbd\x4a\xe9\x6c\x45\x43\x12\xd1\x55\x46\x17\x99\x1d\x32\x06\x14\xdd\x44\x29\xa4\xcc\x01\xb3\xd0\xa8\x1e\x8a\xc3\x56\x9f\x85\x45\xf5\xa2\x38\xb0\xce\xba\x4c\x08\x8b\x12\xd5\xbb\xc2\x63\x60\xce\x73\x8c\xe6\x30\x40\xa1\x51\x04\x92\xf2\x98\x59\xfd\x54\x94\xd0\x11\x89\x80\x8e\x40\x84\x47\x9c\x53\x87\xcc\x33\x13\x4e\x47\x28\x0c\x59\xaa\x0f\xcb\xd1\x9c\xf6\xe5\x5a\x6a\x95\x15\x32\xdb\x63\xb7\x19\x4d\x07\x22\x82\x7b\xcd\x40\x1c\x36\xe6\xa5\xe7\x88\xf2\x9d\xfc\xf9\x94\x3f\xa2\xe8\xbb\xf4\xea\x57\xe4\x0c\xa7\xf1\x2c\x6a\x34\xa7\xc5\x2c\x9c\x3b\xac\x93\x1d\x81\x66\x42\x31\xa3\x89\x44\x53\x99\x0d\x42\xdb\x09\x6e\x4b\x97\x8d\x13\x5d\xe6\x2e\xc0\x8d\x2b\x69\x7d\xd9\xbd\x35\x24\xdb\xf3\x6e\xad\xec\xfe\x12\xdb\x1f\x47\xd9\xfd\x0d\xbb\xed\xf7\x3c\xa0\x2f\x3f\x75\x3d\x71\xa1\x43\x24\x10\xea\xa2\x67\xd0\x4f\x01\x55\x75\x91\x53\xa0\xa2\x30\xbe\x07\x38\x7e\x0d\x3e\xfb\xe4\xb3\x25\x9f\x47\xe1\x68\x42\xe6\x78\x58\x22\x60\xcb\x9f\xaa\xd5\x9b\x89\xab\xaa\x48\xe8\xdd\x6b\x4a\xeb\x4b\x85\x5c\x4f\x6b\x13\x1d\x16\x45\xab\x63\x48\x9d\xd2\xcf\xcc\x94\xbe\xd0\xdf\x7b\x66\x95\x91\x0b\xf7\xb3\x9e\x75\xe6\x99\xd3\x64\x8c\xcc\xd9\x51\x67\x87\x96\x81\x39\x3a\x52\xf5\x04\x6a\xa5\x5e\x02\x19\xf1\xd8\x9f\x38\xcc\x34\x3d\x85\xa1\xa9\xdb\xb0\x6a\xb0\xb2\x0e\x4b\x6f\x74\x9e\xf9\xe4\xbb\x0c\x7f\xae\xd3\x10\x6f\x88\x62\xa5\xd6\x0b\xa8\xaa\xa2\x32\x07\x6a\xb8\x26\x60\xeb\x64\x4b\x79\x98\xf4\x91\xa9\xc7\x65\xd2\xd8\xaf\x18\xa1\xb6\xce\x3e\xb0\xc6\x1e\x38\x4d\x45\xa8\xf4\x83\x27\x1a\xa3\x38\x4e\x4a\x52\x91\x89\xd5\xfb\x98\x47\x4f\x0c\x1d\x43\x10\x8e\xba\x30\xc3\x59\xd9\xa3\x53\x1e\xb2\x2e\xcf\x8c\x60\x94\xc8\x7a\x9f\xab\xbf\xf0\x5d\xf7\x9e\xa7\xea\xe0\xeb\x61\xe3\x3c\xf6\x59\xb5\x61\xd9\x62\x6b\x9b\xc5\x1a\xa6\x0e\x48\xd1\x73\x01\xed\x7a\xea\x70\x8f\xc4\x73\xa4\xbe\xb4\xeb\x2b\x52\xaf\xf2\xf2\xfa\x63\x6f\x46\x01\x85\x0f\xcd\x79\xf7\x76\x24\x84\x4e\x7d\xd8\xf5\x4d\xb3\x85\x34\xa5\xa4\xf5\x9d\xc7\xa5\x0c\x13\xdc\x9e\xab\x2d\x53\xd5\x27\x4f\x4f\x50\x6f\x9f\x44\x2d\xf3\xb1\xcf\xf6\x82\x71\x9d\x86\x8f\x33\x73\x51\x9f\x9c\xfa\xee\xc7\xda\x04\xe4\x1f\xc8\x89\x1f\xfa\x03\x56\x48\xa7\xeb\x10\x9b\x5f\x2c\x2a\xb9\x72\xc5\xae\x2b\xfc\xa6\xea\x1d\xb7\xd4\x5a\xf4\xda\xd1\x13\x7c\x3f\x07\x9f\x7d\xf2\xd9\x92\xcf\xa3\x9f\x22\x2f\xbe\xbc\x99\xbe\xed\xb6\xde\xed\x71\xf7\x1f\xe5\x6e\x1f\xfd\x21\x77\x96\xd8\xad\x44\xb9\x76\x1c\x5b\xac\xb8\x95\x1c\x99\x64\x8a\x09\x97\xa4\xbf\xba\xcb\xec\x45\xb2\x9b\xa0\x1f\x32\x7a\x95\x1d\xd7\x72\xba\x6b\x87\xab\x90\xd1\x43\x87\x08\x4d\xa3\x44\xbc\x49\x6a\x92\xb0\x52\x4e\xb3\xda\x61\xb6\xb8\x32\xd5\x43\xce\x39\x50\x6a\xa8\x96\xe3\x32\xda\xc5\xc8\x9f\x03\x01\x39\x5e\x16\x94\x59\x99\xb9\xb8\x4a\x4f\x7a\xb2\x14\x67\x5a\x74\xd5\x7b\xbc\x64\x52\x9e\x36\x68\x62\x68\x60\x1c\xd7\x5c\x19\x61\x9a\x4c\x64\x96\x81\x41\xf9\xda\x60\x36\x40\x56\xd5\x17\xce\x44\xd1\x07\x6c\x28\x36\x1b\x61\x71\xc1\x1e\xe5\x4a\x6a\xb4\x57\xf9\x53\xab\x22\xa3\x66\x12\x09\xd4\xa7\x34\x4e\x74\x36\x32\x74\xf7\xb9\x7a\x01\x5c\x86\x64\x0a\xdc\x1b\x15\x66\x6e\x3c\x2c\xc5\x4c\x06\xd8\x6f\xd6\x88\x68\x23\x03\x44\xa5\xed\xd5\x8b\x8a\x17\x17\xdb\x72\x82\xe7\x26\x8c\xce\x87\x24\x42\x38\x50\x0e\x1e\x07\xbe\xf2\xf5\x1a\xd7\xc3\xd2\x9a\x4b\xe8\x6b\x85\x2f\xc1\x7e\x18\xc1\x24\xd7\xc8\xb4\x81\x10\xd0\x58\x90\x72\xbd\xac\x7c\xf1\x2c\x24\x99\x53\x15\x72\x3d\x2d\x40\xc4\xee\x8c\x54\xc3\x88\x53\xdd\xdd\x35\x8a\xee\x3c\x53\xa3\x6a\x5f\x31\xbd\x10\xd9\x0e\x89\xe6\x32\x5a\x95\xd9\x8a\x64\xb3\x95\x71\x74\x65\x9c\xeb\xc6\x5b\xdc\xed\x78\xe0\x6f\x21\xad\x4b\xce\x8d\x34\xa2\x8b\x68\x23\xd2\x88\xb8\x18\x24\x50\xa2\x6d\x0e\xe8\x18\x28\x26\xe4\x97\x1e\x52\xc8\xaa\x7e\x63\x08\xae\x83\xa7\xc9\x1b\x13\xac\x39\x76\x5e\xd4\x95\xc9\x1f\xb8\x54\xe9\x41\x3c\x25\xb3\x63\x42\xf4\x89\x87\xe2\xac\xad\x74\xa1\x18\x22\xf9\x86\xc3\x9e\x99\xb4\xde\x8d\xb5\x20\x1f\xb8\xca\xf4\xbd\x6f\x36\x2e\xc3\x24\x56\xca\xda\xb5\x3f\xd9\xac\x45\x5d\xec\xef\x16\x2e\x6d\xef\xd7\xae\xe5\x0d\xd9\xde\x48\x94\x24\xf6\xe4\xc9\x77\x67\x1c\x63\xeb\xf1\xda\x51\xde\xe2\x82\x59\xc4\x76\xb4\xdd\xfe\x14\x79\x1b\xb5\x77\x7f\xb6\xa7\x26\x79\xce\x99\xab\xee\xa8\x7e\x46\x36\x95\x20\x2a\xe4\x38\x6f\xeb\x9c\xc6\x0d\xdb\x6b\xa7\xe8\xef\x3b\x1d\xaf\x9c\x37\x10\x5d\xa2\xfd\x2f\x7e\xdf\x25\x86\x3f\xb9\xaa\xdc\xdf\xb9\x7e\x0f\xff\xaf\xba\xa0\x43\x4c\xf2\x5c\x75\x29\xab\x2a\x75\xc1\x41\x4c\xdc\x4c\xaf\x6e\x8f\x5c\x10\x0d\xa6\x88\x10\x54\x2f\xb1\x56\xd4\x47\x35\xc9\xc0\xda\x38\x50\xd2\x10\x33\x84\x30\xca\x74\x47\xb2\x3d\xf1\xdd\x8e\x55\x21\x51\xf7\x37\x7d\x03\xc8\xb4\x6c\x6f\xf5\x26\x8f\x7b\x0a\xfa\x45\x1b\x04\x03\x4c\x04\x21\x86\x2a\x09\xec\x84\x58\xa8\x1c\x68\x27\x16\x24\x45\x04\x97\xb7\xd5\xaf\x17\x60\x4b\x3f\xc3\xaa\x03\x1a\x98\x6d\xe8\xf2\xbb\x06\x3a\x69\xf6\x4a\xf4\x39\xd7\x74\x6d\x10\x6b\x22\xad\x9e\xf4\xc6\xaa\x2e\x2d\x84\x78\xae\x0f\x3b\xbe\x56\x46\x95\x93\x2b\x4f\xc5\x7c\xac\x2c\xf4\x6e\xbd\x58\x67\x2f\x5d\x9a\x73\x4e\xe1\x21\x9d\xfc\x4e\x5b\xd0\xb3\x20\x95\xc3\xfd\xfa\x61\x6d\x6d\x26\xad\x06\x3d\xe4\xcc\x01\x8c\x75\xca\xe7\x29\xfe\x8b\x7c\xbf\xe5\xc7\xfb\x1e\xa8\xac\x60\xe1\x73\xf9\x79\xc5\x80\x14\x67\x2d\x57\x6a\x12\xcd\x3e\xef\x1a\x99\x1b\x8e\x67\x1c\x86\xb3\x56\xde\x95\x22\x75\xa7\xa5\x75\x57\xae\x95\x4e\xda\x36\xde\x35\xf7\x23\xd3\xee\x59\xa8\x7c\xd9\xe3\xa6\xa1\x79\xee\xdb\xcb\xe5\xba\x18\x82\x88\x45\xbb\x24\xda\xb9\x3c\x6e\xf8\xee\xd6\xb1\xd0\x59\xb2\xb4\xc2\x1b\xcb\xe3\x5e\x75\xe3\xe1\x31\x65\xa3\x76\x59\xe0\x58\x1c\x88\xb0\x7b\x48\x76\xbd\x4c\x5e\x2b\xb8\xe2\xae\x8b\x31\x57\x7f\x6d\xb6\xbe\x35\xca\xca\xb3\x11\x3a\x77\x3d\xcd\xdc\x50\xc8\x52\x78\xe4\xdd\x85\xf9\x85\xe2\x39\x33\xc9\x84\x1b\xc3\x50\x56\xd2\x44\x71\x96\x37\x76\xba\x56\x96\xbc\x96\xcd\x4c\x51\x4a\x70\x5f\x65\x9d\x9a\x4f\xd6\x78\x7f\xb7\xc5\x5c\x45\xaa\xe6\xac\x92\x2d\xe7\xb1\x72\x2d\x25\x38\x2b\x59\x60\xed\xb7\x67\x8b\x93\xb4\x75\x65\x4e\x2c\x2e\x7c\xe5\x07\xd6\x87\x8a\x52\x65\x33\xcf\x5c\x33\x9e\x8c\xf3\x42\x09\x0d\xc0\xd0\x6f\x75\x8d\x0e\x9e\x6d\xd0\x1b\x6f\x19\xd7\xdd\x43\xb2\x5c\x2a\xb6\x18\x6e\xe7\x98\x72\xa9\xbe\xd5\xd5\x47\x2f\x59\x6d\x71\xdc\xab\xba\xb6\x91\x68\x8f\x3a\xea\xad\xc6\x7e\x30\xd3\x3b\x1b\xed\xfb\xe3\xe2\x16\x5b\x77\xac\xf9\x98\x6b\xdf\xdd\x8f\xbc\xd8\xf5\x8a\xff\xd9\xdb\x8a\x7f\xf2\xc0\xe8\x5d\xf2\xf6\x89\xf1\x0e\xaa\x27\x41\xb2\x38\xeb\x46\xf0\x80\xce\xa3\x6b\x16\x04\xc6\x1e\x4f\x3d\x31\xae\xe6\x18\x4f\x24\xbd\x7c\x48\x6f\x61\x1e\xa2\x05\x79\x74\x95\x60\xe5\xca\x3a\x73\x3e\xad\x53\xb9\x3e\x56\x6c\x1e\xac\x5b\xd6\xf7\x58\xef\xb5\xdc\xbc\x50\x58\xaf\x34\xc8\x38\xd5\x8a\x16\x32\xcc\x35\x06\xe3\xd1\x8e\x3e\xa2\x3e\x2a\xdf\xe5\xbb\x11\xb7\xde\xf3\xc9\x88\x41\x9d\x9d\x11\xa2\x03\x83\x51\x32\x5a\x97\xd9\x9a\x64\x3b\xfe\x12\x31\x07\xb9\xc3\x24\xa6\xd2\x4e\xa3\xd9\x0f\xf7\xf3\xed\x65\xde\x83\x0d\x45\xca\x22\xc9\x06\xde\xd4\xb1\xfe\x77\xb8\xb1\x38\x56\xee\x0c\x10\x87\xf9\xf2\x05\xef\xac\xef\xfe\xf8\xfc\xf2\xf1\xf9\xeb\xef\xdf\xfe\xf8\xe7\xf3\x0f\x12\x95\x86\x2f\x1f\x9f\xbf\xff\xe3\xb7\x9f\xff\xfe\xf5\x8f\xaf\xbf\xfd\xc9\x91\x2f\x1f\x9f\xfb\xe3\xaf\xff\xa6\x70\xf2\x2f\x5f\xbf\x7d\xfd\xfc\xe9\xe3\x73\x4e\xd4\x95\x8b\x2e\x38\x5b\xbd\xdb\x72\xb7\x7e\xee\xf6\xb0\xbb\x8d\xef\xdb\xcf\x7b\xaf\x77\x29\xe6\x07\x73\xc0\xbf\xff\xef\x3f\x01\x00\x00\xff\xff\xf2\xf3\xb9\xb9\xcf\x11\x00\x00")

func dataStub_flasher_esp8266JsonBytes() ([]byte, error) {
	return bindataRead(
		_dataStub_flasher_esp8266Json,
		"data/stub_flasher_esp8266.json",
	)
}

func dataStub_flasher_esp8266Json() (*asset, error) {
	bytes, err := dataStub_flasher_esp8266JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/stub_flasher_esp8266.json", size: 4559, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"data/stub_flasher_esp32.json":   dataStub_flasher_esp32Json,
	"data/stub_flasher_esp8266.json": dataStub_flasher_esp8266Json,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"stub_flasher_esp32.json":   &bintree{dataStub_flasher_esp32Json, map[string]*bintree{}},
		"stub_flasher_esp8266.json": &bintree{dataStub_flasher_esp8266Json, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}