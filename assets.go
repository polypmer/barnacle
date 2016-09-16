// Code generated by go-bindata.
// sources:
// templates/index.html
// templates/playlist.html
// DO NOT EDIT!

package main

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

var _templatesIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x54\x41\x6f\xe3\x36\x13\x3d\x4b\xbf\x82\x1f\x81\x0f\x3d\x64\x29\xda\xde\x64\xeb\x2e\x24\x03\x4d\x82\x6d\x9b\x2d\xd0\x6c\x12\x74\x37\xb9\x8d\xc9\x91\x44\x87\x22\x15\x72\xe4\xd8\x35\xfc\xdf\x0b\x49\x71\xd2\x6c\xd1\x16\xeb\xd3\x0c\x3d\xf3\xe6\xbd\x37\x22\xf3\x9a\x1a\xbb\x48\x19\x63\x2c\xaf\x11\xf4\x22\x4d\xf2\xff\x09\xc1\xae\xf0\xa1\x33\x01\x35\x6b\x90\x80\x11\x54\x91\x81\x7d\x84\x6d\x64\xca\x37\xc8\x4a\x13\x22\x31\x21\xfa\xf2\xa1\x42\xd5\x10\x22\x52\xc1\x3b\x2a\xc5\x9c\x3f\x9f\x3b\x68\xb0\xe0\x6b\x83\x8f\xad\x0f\xc4\x99\xf2\x8e\xd0\x51\xc1\x1f\x8d\xa6\xba\xd0\xb8\x36\x0a\xc5\x90\xbc\x61\xc6\x19\x32\x60\x45\x54\x60\xb1\x98\xbe\x61\xb1\x0e\xc6\xdd\x0b\xf2\xa2\x34\x54\x38\xff\x82\x5b\x13\xb5\xa2\xe7\xb8\x2e\xf8\x46\x74\x20\x94\x6f\x5a\x20\xb3\xb4\xf8\x97\x21\x06\x0b\xd4\x15\xf6\x6d\x69\x92\x93\x21\x8b\x8b\xdd\x8e\x65\xbf\x3d\x3a\x0c\x6c\xbf\xff\x2e\xb2\x33\x6f\x2d\x2a\x32\xde\xe5\x72\x2c\x48\xd3\x24\xb7\xc6\xdd\xb3\x3a\x60\x59\xf0\x7e\x52\x7c\x2f\x65\x03\x1b\xa5\x5d\xb6\xf4\x9e\x22\x05\x68\xfb\x44\xf9\x46\x3e\x1f\xc8\xe3\x6c\x92\x4d\x04\xd8\xb6\x86\x6c\x26\x55\x8c\x2f\xff\x65\x8d\x71\x99\x8a\x91\xb3\x80\xb6\xe0\x91\xb6\x16\x63\x8d\x48\x7c\x91\x8c\xe6\xcb\xd1\xfd\x21\x5e\x7a\xbd\xed\x95\x6a\xb3\x66\xca\x42\x8c\x05\xef\x25\x81\x71\x18\x38\x5b\xa4\xc9\xb8\xae\x29\x1b\x70\x0a\xde\x40\xa8\x8c\x13\xe4\xdb\xf7\x27\xff\xe7\xff\xa6\xb0\x9e\x1e\xba\x3b\x7b\x80\xb6\x26\x92\xa8\x82\xef\xda\xde\xa7\x64\xb7\x63\x01\x5c\x85\x2c\xfb\xd1\x2e\xbb\x26\xb2\xfd\x3e\x4d\x92\x1c\xfe\x5e\x2e\x0c\x61\xc3\x47\x9b\x64\x7f\x8c\x4e\xee\x76\xd9\x4d\xef\xe2\x7e\xdf\x63\x0d\x93\x62\x0b\xee\xd0\x4c\x50\xf5\x1f\x93\xd0\x58\x42\x67\x69\x88\x5b\x63\x2d\x6b\x3b\x6b\xc5\x26\x8a\x60\xaa\x9a\x06\x1e\x3d\x91\xec\xcc\x77\x8e\x46\x06\xa3\x4b\x3d\xd8\x01\xb9\x2f\x18\x86\xb1\xfd\x3e\x97\xf0\x44\x1e\x9d\x1e\x1a\xc6\xfa\xce\x1e\x04\xc3\x57\xfb\xac\x0c\xd5\xdd\x72\x58\x61\xeb\xed\xb6\x6d\x30\xc8\x25\x04\x07\xca\x22\x5f\x5c\xfb\x2e\x28\x64\x67\x5e\xe3\x08\x9d\x4b\x6d\xd6\x4f\xfb\x91\xe3\x82\x86\xb8\xbf\x2c\xa7\x87\x35\x33\xe5\x03\xb2\x0b\x58\xc3\xb5\x0a\xa6\xa5\x34\x61\xc5\x37\xff\x86\x6b\xf5\x8c\x7d\x69\x41\xa1\x66\x40\x8c\x6a\x1c\xc4\xf9\x72\x08\xb5\x57\x5d\x83\x8e\x58\xf4\x43\xde\x42\x85\x91\x59\x0f\x9a\x95\x10\x09\xc3\x0b\x4e\x7c\x22\x13\x83\x7a\xd1\xaf\xbc\xc6\x6c\xf5\xd0\x61\xd8\x0e\x26\x8c\xa1\x98\x65\xb3\xec\x78\xf8\x5c\x57\x91\xa7\x89\x71\x84\x55\x30\xb4\x2d\x78\xac\x61\x76\xf2\x4e\x9c\x2e\x6b\x6d\xd7\x9f\x4a\xb9\xb9\xb9\xfd\xa1\x5a\xc1\xe4\xfc\xe1\xed\xcf\xe6\xf1\xd3\x87\xf9\xaf\x70\x76\x75\xf3\x65\x73\xf7\xf1\xaa\x23\xb4\x37\xc7\xc7\x05\x4f\x13\x15\x7c\x8c\x3e\x98\xca\xb8\x82\x83\xf3\x6e\xdb\xf8\x2e\xf2\x45\x2e\x47\x56\xaf\x28\xb2\xd7\x0c\xb5\x5b\xc5\x4c\x59\xdf\xe9\xd2\x42\xc0\x81\x26\xac\x60\x23\xad\x59\x46\x49\x48\x35\x06\x39\xcd\x66\xd9\x44\xae\x0e\xf9\x81\x3a\x7b\xcd\xfc\xed\xfc\x58\x5c\xda\x65\x53\xcd\x2f\x6e\x67\xf3\x8f\x1f\xd0\xae\x2f\x7e\x07\x33\x99\xda\xf9\xe7\xed\xdd\x1f\xe1\xf6\xf3\x4f\xf3\xd9\x49\x73\xa4\xee\x26\x78\x7e\x7e\x3d\x2d\xbf\xd7\x72\x15\xdf\x99\xfb\xf5\x76\x7a\xf4\xe5\xa8\xea\x2e\x7f\x39\xe5\xec\x3f\xc5\x24\xff\xac\xe6\x5b\xdf\x8f\xd5\xd7\xcf\xc7\xea\x95\x6b\x69\x2e\xc7\xa7\xfb\xcf\x00\x00\x00\xff\xff\x58\x70\xa7\x91\xc2\x05\x00\x00")

func templatesIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesIndexHtml,
		"templates/index.html",
	)
}

func templatesIndexHtml() (*asset, error) {
	bytes, err := templatesIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/index.html", size: 1474, mode: os.FileMode(420), modTime: time.Unix(1474040479, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPlaylistHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x56\x4d\x73\xdb\x38\x12\x3d\x4b\xbf\xa2\x17\xeb\x54\xe4\xb2\x49\x4a\x8a\x9d\xf5\xda\xa2\xb6\x62\x7b\x93\xc9\xc7\x4c\x1c\xc7\x35\x89\x93\xf2\x01\x22\x5a\x24\x24\x10\x60\x00\x50\x36\x47\xa5\xff\x3e\x05\x50\xa4\x14\x25\xa9\xa9\xcc\x61\x74\x11\xd8\xec\x8f\xf7\xba\x1f\x40\x8c\x32\x9b\x8b\x71\x17\x00\x60\x94\x21\x65\xe3\x6e\x67\xf4\xaf\x20\x80\x6b\xfc\x52\x72\x8d\x0c\x72\xb4\x14\x2c\x4d\x0d\x50\x71\x4f\x2b\x03\x89\xca\x11\xa6\x5c\x1b\x0b\x41\xe0\xdc\xbd\x47\x92\x51\x6d\xd0\xc6\xa4\xb4\xd3\xe0\x84\xb4\x76\x49\x73\x8c\xc9\x82\xe3\x7d\xa1\xb4\x25\x90\x28\x69\x51\xda\x98\xdc\x73\x66\xb3\x98\xe1\x82\x27\x18\xf8\x87\x43\xe0\x92\x5b\x4e\x45\x60\x12\x2a\x30\x1e\x1c\x82\xc9\x34\x97\xf3\xc0\xaa\x60\xca\x6d\x2c\xd5\x26\x6f\x66\x6d\x11\x38\x8c\x8b\x98\x3c\x04\x25\x0d\x12\x95\x17\xd4\xf2\x89\xc0\xad\x22\x1c\x63\x64\x29\xba\xb0\x6e\x67\x64\xb9\x15\x38\x5e\x2e\x21\xbc\x71\x2b\x58\xad\x46\x51\x6d\xeb\x76\x3b\x23\xc1\xe5\x1c\x32\x8d\xd3\x98\xb8\xe4\xe6\x34\x8a\x72\xfa\x90\x30\x19\x4e\x94\xb2\xc6\x6a\x5a\xb8\x87\x44\xe5\x51\x6b\x88\x8e\xc2\x7e\xd8\x0f\xa8\x28\x32\x1a\x0e\xa3\xc4\x98\xcd\xbb\x30\xe7\x32\x4c\x8c\x21\xa0\x51\xc4\xc4\xd8\x4a\xa0\xc9\x10\x2d\x19\x77\xea\x7e\x47\x75\xc3\xfd\x7a\xa2\x58\xe5\xc8\x31\xbe\x80\x44\x50\x63\x62\xe2\x58\x50\x2e\x51\x3b\xf8\xf5\x80\x06\xe0\xd3\xc4\x24\xa7\x3a\xe5\x32\xb0\xaa\x38\x3d\x7e\x44\x76\x38\x65\x83\x26\x80\x36\xb9\x26\x56\xc2\xc4\xca\xa0\xd0\x3c\xa7\xba\x22\x6b\xa2\x11\x19\x5f\xa3\x2d\xb5\x04\xab\xe0\x42\x09\x81\x89\xe5\x4a\x8e\x22\xda\x96\xd4\xcd\x6a\x0b\x99\x56\xf7\x04\xc6\xdd\xce\x0e\x5c\x11\xe4\x2c\x38\x76\x68\xeb\x88\x49\x69\xad\x92\xbb\x10\x18\x4e\x69\x29\x2c\x01\x5b\x15\x18\x93\xda\x89\x80\x92\x89\xe0\xc9\x3c\x26\x1a\x85\xa2\xac\xb7\xef\xa0\xb9\xd5\x28\xaa\x5d\xda\xb4\xb4\x64\x5c\x01\x67\x75\x87\xb4\x12\xf5\xc0\xb5\x12\xc6\xf9\x74\x46\x46\x95\x3a\x41\x30\x3a\x89\x49\x53\xc6\x07\x45\x79\x81\xa9\xc7\xd7\xb9\x55\xa5\x86\x89\x56\xf7\x06\x35\x30\x85\x06\xa4\xb2\x60\xca\xc2\xc9\x14\x6c\x86\x50\x97\x41\x81\x39\x4a\x1b\x36\xc5\x23\x6f\x6e\xb1\xf8\xf6\xd4\xcb\x52\x34\x4c\x05\x37\x36\x48\xb5\x2a\x0b\xe2\x61\x1a\x25\x53\x67\xab\x2b\x2f\x97\xa0\xa9\x4c\x11\xf6\xf8\x21\xec\x19\x38\x8d\x21\x7c\xaf\x64\x6a\x60\xb5\xf2\xf0\xe9\x7a\x38\xff\xf6\xd1\xcb\xe5\x1e\x5f\xad\xbe\x4d\x1d\x70\x8b\x39\x71\x01\x00\x9b\xe6\x19\xc4\x79\x6f\xb9\x84\x3d\x0e\xab\xd5\xbe\x17\xc6\x9e\xf1\xa2\xa0\x4d\x71\x94\xac\xae\x54\xf3\x29\xc5\xb8\x5e\xba\x71\x46\x8c\x2f\xfe\x7a\xae\x3c\x4f\x1b\x1d\xfa\x8d\x7b\x3a\xe8\xf7\x1f\x11\xdf\x6f\x27\xc4\x0b\xb5\x40\x0d\x1b\xcc\x3c\x4f\x03\xad\x4a\xc9\xdc\x69\x12\x3c\x04\xb4\xb4\x0a\x58\x30\x11\x2a\x99\x13\xa0\xc2\xba\x1a\x2e\x84\x6a\xb7\x37\x76\xc1\xd4\x30\xeb\xf5\x28\x6a\x37\x85\xdf\xb3\x26\xd1\xbc\xb0\xce\x29\x8a\xe0\x8a\x1a\x03\x54\x6b\x5a\x81\x9a\x82\xf1\x2d\x75\x02\x42\x06\x93\x0a\x5e\x28\xef\xb5\xa0\x1a\x0a\x41\x2b\xd7\x48\x88\xe1\x33\xc9\x91\x71\x1a\x19\x9a\x17\x02\x07\x61\x5e\x3c\x21\x87\xf0\x95\x71\xe8\x8d\x77\xdd\x0e\xec\xc6\xde\x9d\x35\x46\x65\xb8\xdb\x37\x10\x43\xbf\xb1\x35\x43\x87\x18\x98\x4a\x4a\x2f\xa2\x14\xed\xff\x6b\x3d\x99\xf3\xea\xc2\x75\xe7\x37\x9a\x63\xef\x9b\xa9\xee\x37\x49\x68\xc9\xbe\x1f\x7f\x5e\xbd\x64\xbd\x56\xff\xde\xbf\xd5\x55\x78\x45\x6d\x56\xab\xa9\x85\x1b\x16\xa5\xc9\x9c\x30\x42\xa7\x8b\xb5\x7b\xa3\x04\xe8\x76\x60\x5a\x4a\xbf\xf5\xc1\xa0\x7d\x96\x58\xbe\xc0\x1e\x67\x0f\xfb\xb0\xac\x07\x00\x53\xa5\xa1\xe7\x20\x71\x4f\x12\x38\x8c\x5a\x8a\xa1\x40\x99\xda\xec\x0c\xf8\xc1\x81\x8f\xe8\xb4\xaf\x3e\xf3\xbb\xd0\xcb\xe0\x8d\xf3\xd3\x98\xab\x05\xf6\x1e\x53\x5f\xe0\xb1\xc7\xe1\x7e\xab\xf5\xff\x8f\x88\x3a\x24\x5b\x69\x28\x63\x5f\xe7\x58\xd5\x02\xb8\x79\x7b\xf9\xf6\x14\x9e\x31\x06\x66\xce\x8b\xf6\x50\xe8\x76\x5c\x1f\x43\xa3\x13\x88\xdb\x86\x7c\xee\xfb\xe9\xb9\x17\xce\xd4\xf3\x79\x36\xdc\x9b\x89\x7a\x73\x14\xc1\xaf\x74\x8e\xe0\xc7\xc9\xb8\x71\x01\x90\x94\x5a\xa3\xb4\x9e\xe9\x3a\x91\x92\xe8\x45\x1e\xb7\xdd\xec\x6d\x1a\xd8\x64\x3c\x38\x68\x58\x47\x51\x3b\x1d\x93\xf1\xa9\xed\xed\x6f\xde\xac\xb9\xbc\x46\x74\x3c\x8a\xca\xf7\x5f\xa3\x0b\xe0\x32\xfd\xdf\xda\xef\x3b\xb4\x9a\x32\x77\x67\x5b\x3e\xf5\x81\xba\x6d\xd9\x70\x76\xbf\x1f\xf0\x5e\x9d\x6d\x0b\xa3\x39\x97\xd7\x8c\x84\x4a\xa8\xb3\x87\x8d\xfd\xcc\xed\x5a\x3f\x8a\x2d\x2d\xe1\xbc\xc7\xd9\xb7\x4d\x80\x18\x38\x3b\xfb\x87\x58\xb8\x83\x22\x6a\x4f\x0a\x7f\xbb\x39\x6f\x3e\xd2\x90\x28\x8d\xf0\x8a\x2e\xe8\x7b\xef\xb0\x4e\x15\xff\xf4\x6f\x7d\x17\x72\xc9\xaf\x04\x4d\x90\x01\xad\x3f\x23\x6e\x93\xa9\xa9\x5f\x36\xf2\x06\xa3\xfc\x73\x41\x53\xac\x8f\x28\x98\x52\x63\x51\xaf\x93\x98\x2d\x28\xfe\x23\xd6\xdc\x47\x12\xc5\x30\x9c\x7d\x29\x51\x57\xfe\x1e\x52\x2f\x83\x61\x38\x0c\x8f\xfc\x75\x63\x66\x48\x1d\xc6\xa5\xc5\x54\x73\x5b\xc5\xc4\x64\x74\x78\xfc\x34\x38\x9f\x64\x4c\x2c\xde\x4d\xa3\x87\x9b\xdb\xff\xa6\x33\xda\xbf\xfc\xf2\xe4\x17\x7e\xff\xee\xf9\xc9\x1b\x7a\x71\x7d\xf3\xf1\xe1\xd3\xeb\xeb\xd2\xa2\xb8\x39\x3a\x8a\xd7\x49\x12\xad\x8c\x51\x9a\xa7\x5c\xc6\x84\x4a\x25\xab\x5c\x95\x86\x8c\xdb\x6e\x76\x5a\xb0\x3b\x38\x99\x9c\x99\x30\x11\xaa\x64\x53\x41\x35\x7a\xb0\x74\x46\x1f\x22\xc1\x27\x26\xb2\x68\x33\xd4\xd1\x20\x1c\x86\xfd\x68\xd6\x3c\x37\x04\x76\xb0\x3f\x39\x39\x0a\xae\xc4\x24\x4f\x4f\x5e\xdd\x0e\x4f\x5e\x3f\x47\xb1\x78\xf5\x3b\xe5\xfd\x81\x38\xf9\x50\x7d\xfa\x43\xdf\x7e\x78\x71\x32\x3c\xce\x0f\x92\x4f\x7d\xbc\xbc\x7c\x3f\x98\xfe\x87\x45\x33\xf3\x94\xcf\x17\xd5\xe0\xe0\xe3\x41\x5a\x5e\xbd\x3c\x27\x7f\x9f\xca\xcf\x5e\x01\x67\xbb\x37\xc0\xd9\x76\x95\xf5\xbd\xaf\xbe\xec\x8d\xa2\xfa\xee\xfd\x67\x00\x00\x00\xff\xff\x2c\x47\xf9\xc4\x83\x0b\x00\x00")

func templatesPlaylistHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesPlaylistHtml,
		"templates/playlist.html",
	)
}

func templatesPlaylistHtml() (*asset, error) {
	bytes, err := templatesPlaylistHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/playlist.html", size: 2947, mode: os.FileMode(420), modTime: time.Unix(1474040453, 0)}
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
	"templates/index.html": templatesIndexHtml,
	"templates/playlist.html": templatesPlaylistHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{templatesIndexHtml, map[string]*bintree{}},
		"playlist.html": &bintree{templatesPlaylistHtml, map[string]*bintree{}},
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

