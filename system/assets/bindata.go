// Code generated by go-bindata.
// sources:
// views/auswahlliste.html
// views/base.html
// views/index.html
// views/index_ig.html
// views/base/footer.html
// views/base/header.html
// DO NOT EDIT!

package assets

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

var _viewsAuswahllisteHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x56\xdb\x6e\x1b\x37\x13\xbe\xf6\x0f\xfc\xef\x30\xe5\x45\x73\x51\x30\x1b\x1f\x0a\x18\xe9\x7a\x01\x39\x50\x1a\x01\x8e\x0d\xd8\xae\x5b\xf4\xc6\xd8\x03\x77\x97\x35\x45\x2e\x48\xae\x23\x57\xf0\xdb\xf4\x19\xfa\x02\x79\xb1\x0e\xb9\x67\x59\x8a\xdc\x3a\x81\xc4\xa1\x66\x3e\xce\x37\x27\x72\xbd\xce\x58\xce\x25\x03\x92\xa8\xec\x89\x3c\x3f\xff\xff\x7f\xee\xdf\x41\xf8\x1d\xa5\x30\xfb\xe5\xe6\xd7\xd9\xa7\x8b\x8b\xc5\xcd\xed\x1c\xce\xe7\xb3\xeb\xf3\xf9\xe2\x76\x7e\x09\x94\x46\xe0\x74\x32\xfe\x08\xa9\x88\x8d\x39\x23\xa9\x92\x36\x46\x1c\x4d\x22\x6f\x0f\x00\xe3\x9f\xab\x58\x32\x01\xfe\x93\xe2\x81\x71\x2d\xac\x53\x3c\x38\x68\x40\x26\x6a\xb4\x64\x71\xc6\x65\x41\x22\x04\xf1\x3a\x07\x61\x79\x38\x55\xb1\xdc\x0a\x46\x20\x16\xbc\x90\x78\x38\x93\xd6\x9d\x1c\x1a\xfc\x35\x0a\x93\xe8\x7a\x46\x3f\x2f\x3e\x5c\x5f\x81\xdf\x01\x63\x9f\x04\x3b\x7b\x93\x2a\xa1\xf4\x7b\xd0\x2c\x7b\x13\xfd\xac\xc2\xc0\x6b\xc3\xac\xaa\x60\x4e\x2f\xe3\xb4\xd4\x3c\x2d\x2d\x84\x41\x12\x8d\xcd\x48\x8e\xd4\xa8\xe1\x7f\xb2\xf7\x70\x78\x5a\xad\x48\xb4\x5e\xbf\xbd\x63\xda\x70\x25\x9f\x9f\x5b\x94\xfe\xab\x3c\xf4\xb4\xc0\xf1\x0f\x90\x5a\x04\xce\xff\x26\x20\xfe\xef\x45\x54\xa8\x0f\x7c\xd4\x12\x7d\x55\xcc\x1a\x78\x97\x21\xaf\xcf\xb3\x33\x12\xd7\xe6\x4b\x5c\x8a\x7b\x97\x07\xad\x84\x21\xdb\xce\x70\x89\x6b\x00\xbc\x4b\xfe\x2f\xb4\x71\x22\x58\xa7\xdd\x08\x89\xd2\x19\xd3\x82\x19\x03\x36\x11\xef\x0e\x49\xd4\xe8\x42\x9f\x10\x6f\xe8\x40\xb7\xfd\xd2\x2e\xdc\x1e\x6a\xe9\x0e\x7b\x40\x25\xbd\x17\x0e\x26\x1b\x49\xd3\x9a\x2a\x59\xfa\x90\xa8\x15\x74\x0b\x6a\xea\x34\x75\x5e\xf9\x0d\xe7\xd7\xd8\x92\xcb\xaa\xb6\x60\x9f\x2a\x36\x98\x12\x1f\x9c\x34\xb9\xff\x9d\xe7\x39\xd3\x92\x40\x30\xb1\x11\x71\x82\x31\xce\x95\xf6\x4a\x46\xe1\xca\xca\xae\x14\x88\x2b\xa5\xef\x65\x62\xaa\x9f\x5a\x73\x17\x62\x5f\x1f\x81\x37\x9c\x40\xf9\x6c\xf7\xe2\x88\x60\x80\x0c\xb7\xed\xff\x77\xe2\x47\x13\xe2\x7b\x98\xcf\x9a\xca\xd8\x60\xfe\x7a\xea\xad\xbd\xe0\xc6\x32\xc0\x32\x7b\x10\x71\x55\x31\xf9\xed\x30\x6c\xd0\x9f\x94\xc7\x0b\xe6\x55\x5f\xab\x77\x0a\xbb\x18\x3f\x4c\x5a\x8a\xb8\x08\x83\x2a\xda\x00\xda\x0d\x62\x98\x60\xa9\xf5\xa4\x7f\x6b\xd6\x7d\x0b\x20\xb3\x25\x6d\x1b\x03\x54\x65\x37\xea\x06\x77\xb0\x95\xbd\xe5\x3b\x12\x2d\xd2\x12\x12\x2e\x21\x61\xe8\xc4\xd7\xbf\x72\xcb\x0b\x0b\x14\xce\xb9\x45\xfa\xd2\x4f\x08\x63\xbf\xfe\xad\x5d\x04\x1a\xcb\x5d\x60\x87\x03\x18\xfe\x67\x6e\x38\xc2\x39\x33\x95\x66\x69\x59\xcb\x62\x9f\x39\x26\xb9\x39\x34\xd3\x38\x0c\x99\xcc\xc0\x35\x4f\xc6\x8b\xfd\x27\x1f\x77\xa6\xb7\x4c\x2f\xf1\xec\x3b\x67\x89\x76\x81\x64\x35\x93\xdd\xee\x23\xd3\xe8\x54\x12\xbf\x82\xca\x49\x07\xa8\xeb\x1c\x01\x6e\x38\x83\x25\x86\x02\xbe\x70\x86\x2d\x0d\xf1\x5e\x80\x1f\x49\x44\xa9\x60\x4c\x53\xba\x55\x35\x68\x52\xf6\xfa\x74\x27\xb5\xb5\x88\xdd\x16\x7d\x23\xf5\x09\x4f\x2c\xe6\xcf\xca\x6e\x6a\xe2\x1a\x67\x05\x78\x47\x70\xfb\x3e\xce\x32\x12\xfd\x80\x05\xec\xad\x76\x9e\x89\x92\x9e\x48\x7e\xde\xf5\x73\x33\xf0\xb3\xd2\x77\x76\x33\xf1\xfd\x2e\x0e\xe5\x5e\x63\xda\x1c\x21\x5f\x16\xde\x05\xfc\xbe\xff\xc4\x44\xd5\xbb\x8b\x1b\xa5\x97\x8d\x4e\xcf\x48\x80\x62\xe0\x64\x7a\x7c\xb4\x3a\x3e\x7a\x5b\xe1\x4d\x88\x17\x9d\x3d\x23\xde\x68\xd4\xc5\x13\xc4\x85\xcc\xd5\x18\x91\x7b\x79\x40\x74\xf2\xfd\xc9\xe9\xea\xe4\x74\x84\xe8\x8d\x82\x9e\x43\xc3\x6b\xea\xf6\x70\x65\xf8\x0b\xa7\xbb\xd5\xfa\xed\x46\x6e\x55\x46\x97\x98\xf3\xcb\xd4\xc9\x1f\x98\x55\xe3\x66\x47\xef\x5b\xa9\xb0\xf0\x6e\x5d\xec\xfa\x36\x84\xb0\x16\xde\xa0\x16\x5b\x4d\x9c\x40\x0b\xad\xea\x0a\x2c\x33\x76\xb0\x6b\x2e\xc1\xcf\xf8\xf0\x80\x0b\xd4\x81\x12\x4b\x7a\x74\xc9\x85\x41\x2d\x5e\xba\x39\x5a\x0e\x2f\x90\xc9\x6d\x99\x2b\x85\xef\x09\xdc\x63\xb1\xce\xf9\xca\x9d\xd7\x5d\xdf\x1b\xea\xb5\x10\x54\xf3\xa2\xb4\x63\x9d\x56\xaf\x2b\xd1\x1d\x45\xe9\x05\x51\x90\xa1\x2c\xdb\x2b\xdc\xc4\x8f\x43\x68\x0e\x9a\xa7\x48\x8b\x51\x88\xa7\xaa\xe4\x38\xca\xa0\x5f\x51\xf5\x40\xba\xe7\x47\x4f\xbb\xaf\xec\xed\x2e\xed\x6c\x93\xed\x0e\x65\xd8\x9b\xf6\xdf\xb8\xa4\xd9\x52\x39\x0e\x7b\xdd\x1a\x92\x31\x95\xda\xd4\xf4\x62\xb7\xe8\x1e\xa6\xf3\xab\x8f\xf0\xe1\xea\xf2\x76\xb6\xb8\x9c\x5f\xc3\xc7\xab\xeb\x6f\x3d\x55\x9d\xd9\x7a\x8d\x03\xd4\x3d\x6e\xff\x09\x00\x00\xff\xff\xab\x77\xc9\x3d\xf0\x0a\x00\x00")

func viewsAuswahllisteHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewsAuswahllisteHtml,
		"views/auswahlliste.html",
	)
}

func viewsAuswahllisteHtml() (*asset, error) {
	bytes, err := viewsAuswahllisteHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/auswahlliste.html", size: 2800, mode: os.FileMode(438), modTime: time.Unix(1452602757, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsBaseHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\x4a\x2c\x4e\x55\xaa\xad\xe5\xe5\xb2\x51\x74\xf1\x77\x0e\x89\x0c\x70\x55\xc8\x28\xc9\xcd\xb1\x03\x0a\x80\x68\x85\x9c\xc4\xbc\x74\x5b\xa5\xd4\x3c\x25\xa0\x88\x02\x10\xd8\x64\xa4\x26\xa6\x40\xd9\x20\x50\x5d\x5d\x92\x9a\x5b\x90\x93\x58\x02\x34\x0b\x24\xa5\xa4\xa0\x57\x5b\xcb\xc9\x09\x55\xac\x8f\xa4\xda\x26\x29\x3f\xa5\xd2\x0e\xaa\x0d\x22\x84\xac\x19\x24\x0b\xd6\x0c\x94\xe3\x44\x96\x48\xcb\xcf\x2f\x81\x49\x40\x0c\x05\x1b\x04\x74\xa0\x3e\xd4\xa5\xd5\xd5\xa9\x79\x29\x20\x3f\x00\x02\x00\x00\xff\xff\xd0\x62\xc2\x46\xd7\x00\x00\x00")

func viewsBaseHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewsBaseHtml,
		"views/base.html",
	)
}

func viewsBaseHtml() (*asset, error) {
	bytes, err := viewsBaseHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/base.html", size: 215, mode: os.FileMode(438), modTime: time.Unix(1452676680, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x57\x4b\x6f\xe3\x36\x10\x3e\x2b\x40\xfe\xc3\x54\x87\xee\xa5\x5c\xd5\x6d\x0f\xdb\x54\x11\xe0\xa6\xde\x6e\x80\x36\x59\x24\xe9\x02\x3d\x19\x94\x44\x5b\x44\x68\x51\x20\x69\xe7\x61\xe4\xbf\x77\x48\xea\x6d\x6b\x37\xdd\x1e\x6a\x20\xb1\x86\x1c\xce\xe3\x9b\x6f\x46\xf4\x7e\x9f\xb3\x15\x2f\x19\x84\xa9\xcc\x9f\xc2\x97\x97\xd3\x93\xd3\x13\x80\x78\xc3\x0c\x85\xac\xa0\x4a\x33\x73\x1e\x6e\xcd\x8a\xbc\x0b\x13\xc0\x8f\xdb\x35\xdc\x08\x96\xec\xf7\x6f\xef\xec\xc3\xcb\x4b\x1c\xf9\x95\x7a\xdb\x1d\x2e\xe9\x86\x9d\x87\x3b\xce\x1e\x2a\xa9\x4c\x08\x99\x2c\x0d\x2b\xd1\xd8\x03\xcf\x4d\x71\x9e\xb3\x1d\xcf\x18\x71\xc2\x77\xc0\x4b\x6e\x38\x15\x44\x67\x54\xb0\xf3\xd9\xdb\xef\xc3\xe4\xf4\x24\x8e\x0a\x46\x73\xfb\x60\x63\xc3\xef\x00\x1f\xbf\x21\x04\x2e\xae\xaf\xee\xe6\x97\x57\x8b\x1b\x78\x7f\x7d\x03\x0b\x72\x35\xbf\xf8\x70\x73\x79\xf1\xe1\x0e\x80\x90\x04\x63\x88\x73\xbe\x83\x4c\x50\xad\xcf\x43\xeb\x97\x62\x86\x2a\xf4\x16\x30\x87\xfe\x76\x45\x4b\x26\xc0\xfd\x27\x08\x05\xdd\x0a\xe3\x14\x83\xc0\x69\x0d\xd4\x88\x8d\x87\x97\x6b\x87\x84\xd3\x09\xe2\x62\x36\x54\x71\x40\x84\x40\x05\x5f\x97\xe8\x1c\x33\xb6\x9e\x63\x8d\xbb\x49\x9c\x26\x37\x73\xf2\xe7\xe5\xc5\xcd\x35\xb8\x15\xd0\xe6\x09\xf3\x7d\x93\x49\x21\xd5\x19\x28\x96\xbf\x49\x7e\x97\x71\xe4\xb4\x61\x5e\x55\x36\x39\x9a\x15\x8a\x67\x85\x81\x38\x4a\x93\xfe\xb1\x70\x85\xa9\x11\xcd\x9f\xd9\x19\xcc\xde\x55\x8f\xa1\xad\xc8\x27\xa6\x34\x97\xa5\xad\x89\xf7\xd9\x7c\x15\x33\x97\x96\xad\x61\x1c\x61\x6a\x09\xd8\xf8\x71\x09\xea\xcf\x01\x2a\xc4\x51\xc2\x83\x01\x41\xbc\x92\x6a\x03\x4a\x7a\xc7\x6a\x13\x36\xba\x56\x20\x85\x54\xfc\xd9\x22\x2d\x42\xe0\x39\x2e\x6e\x96\x68\xa2\x92\xda\xdc\x4b\x95\xd6\x46\x82\x1a\xb3\x15\x67\x22\x47\x62\xd5\xab\x41\x2c\xd8\x9a\x95\x79\x32\x2f\x31\xdf\x8f\xf5\x99\x38\xaa\x57\x1b\x25\x43\x53\xc1\x1a\xa7\x5e\x48\xa5\xca\x99\x12\x4c\xeb\xc6\x83\xd5\x73\xac\x41\xac\x1d\x5e\x91\x17\xfd\x1e\xb4\x75\xf3\x38\x18\x47\xac\x76\xaf\xb7\x85\x66\x54\xe3\xeb\x98\x17\xeb\x27\x4f\x3c\x70\xdd\x5a\xd0\xc7\xd0\xe1\xb2\x56\x72\x5b\xf5\x8f\xd5\x3a\x16\xa3\xac\x60\xd9\x7d\x2a\x1f\x6f\x99\x60\x99\xf9\x4b\x63\xe5\xc2\x8e\xb5\xd8\x0c\x1b\xf2\x63\x98\x74\x07\x83\xbe\x95\x20\xd6\xee\xd8\xc0\x9b\xe5\x3a\x56\x08\x9a\xd3\xd0\xb8\xf0\x45\xd9\x5a\x17\x82\x6b\xec\xc6\x41\x44\x07\xa6\x23\x6f\x7b\xca\xb7\xe7\xcf\x41\x4e\xc3\xd0\x67\xc8\xfa\xd7\xe8\xf5\x82\x1c\x06\x15\xf3\xb2\xda\x1a\x30\x4f\x15\xeb\xb0\xf2\x89\x64\xe9\x52\x4b\xcc\xd8\x94\x4d\x73\x84\x10\x8d\x4e\x0b\x9a\x62\x63\xa3\xd2\x51\x75\xcb\x8f\x6f\xcb\x54\x57\xbf\xcc\x85\x86\xdb\xe1\x2e\xd0\xf2\x99\x71\x24\x9f\x67\x90\xb3\x94\x1c\x02\x70\x14\x9a\xe3\x40\xf4\x82\xe1\xe5\x92\xd1\x7b\xc3\xc2\x64\x41\xe6\xf8\x0d\x67\x8d\x87\xbe\x55\x98\x82\xfe\xd0\xfe\x4f\x68\xbf\x0f\x95\x61\x8f\x26\x3c\x46\x0b\x8f\x5d\xeb\x3f\x1a\x38\xac\xc7\x40\xdf\xd9\xb8\x7c\xd8\x4b\x79\x27\xa2\xa4\x92\x51\x84\xb8\xd6\x8c\xea\x5a\xb4\x6d\x5a\x8b\x71\x34\xea\x7b\x3b\xcb\x17\xd7\xef\xe1\x16\xa9\x86\x23\xcb\xc2\x03\xf3\x92\x2c\x9a\x09\x60\x87\xf9\x60\x6e\xd8\xd9\xdf\x57\xee\x46\x63\xab\x39\x39\x5b\x5a\x5d\x3b\x57\xca\xaf\x1a\x2c\xe8\xbe\x1b\x2e\x3d\x7b\xdd\x94\xc1\x30\x1a\x28\xff\x97\x31\xf3\x6a\xce\xfc\x00\xe3\x66\x1b\x32\x74\x27\x95\xa0\x08\x52\xf3\xe6\xfa\x54\xcb\x67\xed\xeb\x64\xaa\x29\xfa\x2b\x5f\x08\xe2\xe7\xc1\x70\x0b\x8e\x35\xff\x71\x1a\xf7\x79\x5e\x53\xba\x0d\xb8\xbe\x77\x0c\x73\x88\x9c\xb5\xff\x16\xed\x6c\x88\x18\x96\x2e\xdd\x1a\x83\x4c\xf4\xd1\x78\xa1\xed\xbb\xd4\x94\x80\x7f\xed\x95\xc2\xc5\x89\x0b\x4b\xba\xd5\x0f\xb4\x10\x76\x0a\xb3\x51\x0d\x02\xff\x72\xef\x2c\x8c\x2e\x24\x7d\x74\x5a\x60\x28\xac\x28\x61\x68\xaf\xd2\x5c\x93\xc2\x25\x3b\x36\xeb\x6b\x36\x1a\xf1\x3e\xe0\x29\xca\x7c\x0e\x9d\x57\xcc\x85\xaf\x1a\x08\x28\x63\x99\x93\x60\x7a\x3a\x1c\x69\xf8\x89\xd1\xd0\xdc\x9b\x50\x95\xe5\xdc\xe0\x4a\x37\x22\x9a\xf7\x2f\xcb\xcd\x72\x71\xd5\x7b\x25\xf8\x71\xd8\xb3\x38\xf6\xff\x19\xab\x75\x06\xd6\x42\x7b\xf4\xe3\xfc\x6a\xf1\x07\xfc\x7a\xfd\xdb\xdf\xad\xde\xe1\x25\x6b\x25\x25\x5e\x11\x71\x8d\x51\xb5\xe2\xee\x25\xd8\xdc\xc8\x46\xea\x5b\x21\x88\xe2\xeb\xc2\xf4\x75\x6a\xbd\x9a\x8b\x13\xf4\x03\x27\x88\x75\x47\xc3\x0d\x5e\x89\x97\x9a\xee\x7a\x1c\x1c\xb0\x6f\x2d\x9e\xaa\x82\x63\xc7\x41\xfb\x44\xe4\x7d\x98\x0c\xb9\xd4\xb2\x68\x2a\x9e\x89\x70\x8e\x46\x93\xd1\x32\x63\xe2\x5f\xc4\xa3\xd8\x46\xda\x04\xbe\x18\x53\x47\xd9\xa1\x54\x3f\xc0\xb0\x60\xae\x54\xd0\xe8\xb4\x7b\xd3\x3f\x3a\xdc\x81\xd3\x93\xfd\x1e\x5f\x2a\xfe\x07\xd4\x3f\x01\x00\x00\xff\xff\xd8\x67\x9f\x6e\x56\x0d\x00\x00")

func viewsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewsIndexHtml,
		"views/index.html",
	)
}

func viewsIndexHtml() (*asset, error) {
	bytes, err := viewsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/index.html", size: 3414, mode: os.FileMode(438), modTime: time.Unix(1452676680, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsIndex_igHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x56\xcd\x6e\xdb\x38\x10\x3e\x3b\x40\xde\x81\xcb\xc3\xf6\xb2\xaa\x92\xd3\x16\xa9\x4c\x40\x9b\x2a\xdd\x00\x89\xb5\x70\xdc\x00\x3d\x05\x94\x34\xaa\x88\xb0\xa4\x20\x32\x6e\x52\x23\xef\xbe\x24\x25\x4a\x94\xe3\x64\x7f\xe4\x83\x34\xc3\x99\x6f\x66\x3e\xce\x90\xde\xed\x2a\xa8\x99\x00\x84\x0b\x59\x3d\xe1\xe7\xe7\xe3\x23\x84\x92\xef\xa0\x29\x2a\x1b\xda\x29\xd0\x4b\xfc\xa0\xeb\xe8\x03\x26\xc8\x3c\x6e\x55\x33\xcd\x81\xec\x76\xef\x37\xf6\xe3\xf9\x39\x89\x7b\xcd\xb0\xec\x9c\x05\xfd\x0e\x4b\xbc\x65\xf0\xa3\x95\x9d\xc6\xa8\x94\x42\x83\x30\x60\x3f\x58\xa5\x9b\x65\x05\x5b\x56\x42\xe4\x84\xdf\x10\x13\x4c\x33\xca\x23\x55\x52\x0e\xcb\xd3\xf7\x27\x98\x1c\x1f\x25\x71\x03\xb4\xb2\x1f\x36\x33\xf3\x5e\x24\xbf\x44\x11\xba\x4a\xbf\xe6\x5f\x36\x28\x8a\x9c\xa6\x62\x5b\xc4\xaa\x25\xe6\xf4\x49\x3e\x98\x30\x4a\x3f\x19\x04\xdc\x00\xfb\xd6\xe8\xb3\xdf\x4f\x4e\xda\xc7\x8f\x16\x6c\xb1\x18\xbc\xb3\x8b\x0d\xba\x4e\xd7\x9f\x2f\x57\x3d\x04\x0a\x1e\x87\x56\x72\xaa\x94\x01\x84\x5a\x9b\x92\xe7\x06\xfd\xb3\xe7\x14\x1b\xaf\x20\x42\x96\x5f\xfc\xfb\x28\x9d\x4d\x13\x93\xff\x14\xe5\x75\x34\xcb\x17\x74\x43\xb9\x8b\x44\xb5\x54\x90\xa4\x20\xeb\x34\xba\xbe\x3c\x5f\xe7\xc8\x69\x06\x8a\xde\x95\x92\xcb\xee\x0c\x75\x50\xbd\x23\x9f\x65\x12\x3b\x6b\x94\xb6\x2d\xca\xa2\x15\x2d\x9b\x8e\x95\x8d\x36\x71\x8b\x10\xce\xf3\x5b\x9b\xdd\x8c\x14\xfb\x09\x67\xe8\xf4\x43\xfb\x88\x6d\x33\xdc\x42\xa7\x98\x14\xb6\x1d\xfa\xc8\xfe\xd5\x9c\x0e\xf4\x84\x35\x78\xba\xfe\xc8\x3f\x7d\x45\x37\xd9\xf9\xe6\x32\x5f\xbd\x49\x55\x69\xba\xc7\x16\x77\x80\x28\x97\x5e\xcd\x80\x57\xa6\x5d\x87\x74\x17\x09\x87\x6f\x20\x2a\x92\x0a\x53\xd0\x5f\x52\xe9\x7b\xd9\x15\x49\x3c\x68\x17\xde\x4a\xd3\x82\x83\x8f\xe1\x04\x4c\xdc\xca\x62\xc2\xee\xbf\x13\xed\xfa\x70\x54\x84\x6b\x06\xa7\xf3\x20\x85\xec\xcc\x2e\x70\x50\xca\xef\xc4\x10\xa9\x22\x28\x90\xa7\xde\x2d\x1b\x28\xef\x0b\xf9\x78\x03\x1c\x4a\xfd\x45\x19\x1e\x31\x09\xc8\xf2\xf6\x4c\xb4\x0f\x1a\xe9\xa7\x16\x26\x1f\xdc\x23\x14\x77\x4a\xd6\x66\xce\x84\xdf\x38\x8c\x62\x92\x70\x5a\x00\x47\x46\x7f\xd0\x82\xa4\x5c\xa1\x9b\xb9\x12\x51\xf1\xd3\x8c\x0e\x08\x43\x94\x75\x9e\x27\xe0\xf6\xf3\x57\x51\xa8\xf6\xa3\xef\x17\x14\x96\x64\xce\x81\x89\xd9\xb1\xe6\xb9\xcd\x41\x94\x7d\x8c\x40\xfe\x1f\x68\x6f\xc1\x25\x73\x71\x11\x72\xc4\xc4\x1d\xd0\x7b\x6d\x1a\x20\x8b\x52\xf3\x46\x67\x07\x59\x08\xb7\x41\xc3\xa3\xee\xb7\x60\x74\x36\xc4\xcf\xb3\x8e\x5f\xa6\x3d\x7d\x9a\xd5\x6e\x82\x37\x92\x3f\xeb\x06\xd1\x36\x24\xf1\x2e\xa6\xcf\x46\xdf\xa1\xeb\xe3\x59\xdb\xff\xc3\x28\x8c\x63\x3d\x8e\x81\x37\xf0\xad\xb8\x95\x1d\xa7\x66\x05\xfb\xf3\xdb\xcb\x63\x8e\x7b\x84\x4d\x06\xb7\xc3\xd7\x4b\xd6\x5e\x67\x6c\x3f\xdc\x0c\x30\x89\x9d\xdf\x04\x53\x3c\x68\x2d\x85\xf3\x2d\xf4\x4b\xe7\x99\x2e\xe0\x94\xf9\xc1\xac\x29\xaa\x69\x04\x9c\xb3\x56\x31\x15\x35\x2e\xc4\x04\x1f\xf7\xf8\x13\xf9\x76\x00\x5f\x63\x7a\x54\xdb\x33\xec\xcf\xcd\xf5\x15\xca\x3e\x5d\x6e\xf2\xf5\x70\x2f\x85\x9c\x42\xc5\xb4\xec\xee\xb2\xd5\x34\x79\x03\x76\x00\x61\x6f\x8d\x43\x30\xa1\xe5\x68\x18\x9e\x99\xd6\x72\x6a\x88\xc9\xee\x22\xcf\x37\xd9\x3a\xb4\x7a\xfd\x68\xad\xa5\xd4\xc1\xbd\xb1\xc7\xb3\xa2\x5b\x08\x39\x76\x32\xb9\x49\x6f\xb3\x3d\xca\xf6\x1d\x4b\x2a\x4a\xe0\xa1\xeb\xa0\x21\xe7\xe9\xea\x3c\xbb\x9a\xbb\x1f\xba\x47\x87\x22\x6c\x89\x53\xf6\xa3\xe1\x74\xdd\x06\x7f\x09\xfa\xdf\x6e\x67\xba\xdb\xfe\x97\x39\x3e\xfa\x3b\x00\x00\xff\xff\xf1\xe2\xe4\xf0\xe1\x08\x00\x00")

func viewsIndex_igHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewsIndex_igHtml,
		"views/index_ig.html",
	)
}

func viewsIndex_igHtml() (*asset, error) {
	bytes, err := viewsIndex_igHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/index_ig.html", size: 2273, mode: os.FileMode(438), modTime: time.Unix(1452676680, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsBaseFooterHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x94\xc1\x6e\xe2\x30\x10\x86\xcf\xac\xb4\xef\xe0\xcd\xdd\xf8\x05\x80\xc3\xee\x69\xab\xde\x2a\xf5\x5a\x39\xf6\x00\x13\x25\x76\x3a\x9e\x44\x8d\x22\xde\xbd\x36\xb8\x28\x22\x3d\x40\xe1\x66\x86\x99\xef\xff\x67\x32\x9a\x71\xb4\xb0\x45\x07\xa2\xd8\x7a\xcf\xc5\xe1\xf0\xfb\xd7\x2a\xbd\x80\x36\x2b\x95\x1f\x31\xa4\x2c\xf6\x1b\xb1\xfa\x23\xa5\x50\xc6\x3b\xd6\xb1\x84\x84\x94\xe9\xbf\x14\xfc\xdf\xb4\x9e\x18\xdd\x4e\x3c\x63\x49\x9a\x10\x82\x50\xc2\x42\x0b\xce\x82\x33\xe9\xe7\x31\x39\xa7\x3f\xbd\x77\x40\x43\xce\x1d\x32\x27\x18\xc2\x96\x05\x0f\x2d\xac\x0b\x86\x0f\x56\x95\xee\xf5\x29\x5a\x88\x40\x66\x5d\xa8\x2a\xa8\xea\x58\xbb\x6c\xd0\x2d\xab\x50\x44\x97\xa7\x8c\x8d\xb8\x15\x21\x3b\xbc\x20\x88\xb3\xc1\xbf\xb1\xf3\xc0\xa4\xdb\x49\x3f\xd1\xe5\x0d\x1a\xe5\x17\xe1\x2e\xa7\x67\x8a\x6c\xba\x9a\x31\x40\x0d\x86\x67\xb4\x6c\xfa\xa5\x45\x37\x9d\xe9\x62\x71\xbd\x50\x88\xb5\x73\xa7\x19\xdc\xeb\x1a\xad\x66\xf4\xee\x8c\x8d\xe1\x9b\x7b\x78\x3d\x61\x3c\xcd\x85\xa6\xcb\x81\x6e\x4b\x7a\x87\x81\xd1\x84\x6f\x56\xe8\x5a\xd9\x29\x66\x19\x3f\xf5\x9e\x9b\x1a\x2c\x46\x79\x69\x61\xd6\xe6\x4f\xa0\xc6\xd3\x63\x40\xb6\x7f\x08\xa6\xf6\xe5\x25\x27\x8f\xf4\x5f\x17\xd8\x37\x77\x0c\x13\x9c\x36\x7b\x42\xb3\xe7\xb7\x46\xcf\xb7\x24\x5d\x88\xd2\xdb\xe1\x78\x2a\xd2\xa0\xe3\x63\x1c\xa3\x56\xba\x27\x9f\x01\x00\x00\xff\xff\x32\xf7\x0d\xee\x63\x04\x00\x00")

func viewsBaseFooterHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewsBaseFooterHtml,
		"views/base/footer.html",
	)
}

func viewsBaseFooterHtml() (*asset, error) {
	bytes, err := viewsBaseFooterHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/base/footer.html", size: 1123, mode: os.FileMode(438), modTime: time.Unix(1452676680, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsBaseHeaderHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x93\x4f\x6b\x1b\x31\x10\xc5\xef\x81\x7c\x07\x55\xe7\x8e\x4d\x6f\x3d\x64\x03\x26\x98\x92\x43\x28\x04\x5a\x7a\x1d\x4b\x6f\xb3\xd3\xea\xcf\x46\x9a\xb5\x1b\x96\x7c\xf7\x2a\x0d\x21\xae\x4f\x5d\xa3\x93\x86\xf7\x7b\xf3\x47\x9a\x79\xf6\xe8\x25\xc1\xd8\x01\xec\xed\xf3\xf3\xe5\xc5\xe5\xc5\x55\x84\xb2\x71\x03\x97\x0a\xed\xec\xa4\x3d\x7d\xb6\xd7\x6f\xf1\x41\x75\x24\x3c\x4e\xb2\xef\xec\x0f\xfa\xb6\xa1\x9b\x1c\x47\x56\xd9\x05\x58\xe3\x72\x52\xa4\x06\xdd\x6e\x3b\xf8\x07\xbc\x63\x89\x23\x3a\xbb\x17\x1c\xc6\x5c\xf4\x48\x79\x10\xaf\x43\xe7\xb1\x17\x07\xfa\x7b\xf9\x68\x24\x89\x0a\x07\xaa\x8e\x03\xba\x4f\xa7\x2e\x1e\xd5\x15\x19\x55\x72\x3a\x32\x3a\x55\xf1\xa4\x43\x2e\xa7\x02\x15\x0d\xb8\xbe\xdf\xd0\xdd\xed\xcd\xfd\x57\xf3\x25\x9b\xcd\x38\x9a\x2d\xdd\xb1\x04\xb3\xfd\xfd\x52\x1c\xca\xd5\xfa\x55\xd6\xf4\x41\xd2\x2f\x53\x10\x3a\x5b\xf5\x29\xa0\x0e\x40\xab\x7e\x28\xe8\x3b\xbb\x76\xb5\xae\x77\x39\x6b\xd5\xc2\xe3\x2a\x4a\x5a\xb5\x88\x5d\x2f\xe3\x28\x4e\x41\xa5\x22\xc0\xe9\x02\xbe\x6f\x6d\x11\x1f\x50\x73\xc4\xc2\xd4\x3f\x1f\x27\x94\x27\x9a\x64\x01\xc3\x49\x22\x2b\xde\x89\x0f\x44\x8b\xda\xfc\xce\x41\x3c\x6b\x2e\x67\xcf\xc9\xe5\xb8\x6b\x7f\xd5\xaf\x52\x26\x69\xaf\x5a\xcf\x76\x2a\xa8\x63\xe3\x65\xff\xef\xe4\x88\xfe\xcb\x27\xf2\xa2\xb4\x92\xfa\xc2\x0f\x52\x55\x5c\x5d\xe9\x80\x88\x73\xe1\x37\xec\xf5\xcc\x33\x92\x7f\xd9\xd8\x3f\x01\x00\x00\xff\xff\xa1\x5e\x0e\x46\xc5\x03\x00\x00")

func viewsBaseHeaderHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewsBaseHeaderHtml,
		"views/base/header.html",
	)
}

func viewsBaseHeaderHtml() (*asset, error) {
	bytes, err := viewsBaseHeaderHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/base/header.html", size: 965, mode: os.FileMode(438), modTime: time.Unix(1452676680, 0)}
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
	"views/auswahlliste.html": viewsAuswahllisteHtml,
	"views/base.html": viewsBaseHtml,
	"views/index.html": viewsIndexHtml,
	"views/index_ig.html": viewsIndex_igHtml,
	"views/base/footer.html": viewsBaseFooterHtml,
	"views/base/header.html": viewsBaseHeaderHtml,
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
	"views": &bintree{nil, map[string]*bintree{
		"auswahlliste.html": &bintree{viewsAuswahllisteHtml, map[string]*bintree{}},
		"base": &bintree{nil, map[string]*bintree{
			"footer.html": &bintree{viewsBaseFooterHtml, map[string]*bintree{}},
			"header.html": &bintree{viewsBaseHeaderHtml, map[string]*bintree{}},
		}},
		"base.html": &bintree{viewsBaseHtml, map[string]*bintree{}},
		"index.html": &bintree{viewsIndexHtml, map[string]*bintree{}},
		"index_ig.html": &bintree{viewsIndex_igHtml, map[string]*bintree{}},
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

