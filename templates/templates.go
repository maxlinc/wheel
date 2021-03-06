// Code generated by go-bindata.
// sources:
// templates/single-master.cloudformation.json
// DO NOT EDIT!

package templates

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

var _templatesSingleMasterCloudformationJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x3d\x6b\x73\xe2\x3a\xb2\x7f\xc5\xeb\xbd\x1f\xf6\x11\x03\x36\x90\x04\xaa\xf8\xc0\x10\x66\x86\x9d\x3c\xa8\x40\x32\xbb\x77\xc8\xa5\x64\x5b\x80\x36\xb6\xe5\x95\x64\x92\x9c\x39\xf9\xef\xb7\x24\xd9\xe0\x27\x8f\x84\x4c\x72\xb6\x3c\xe7\xc3\x49\xe4\x96\xd4\x6a\xf5\x4b\x52\x77\xe7\xa7\x3a\x04\x04\xb8\x90\x41\x42\xd5\xb6\xf2\x53\xbd\xea\x06\x6c\xd1\xf7\x80\xe9\x40\x5b\x34\x74\x1d\x07\x3f\x40\xfb\x16\x38\x01\xe4\x20\x3f\x54\x46\x02\xa8\x1e\x29\xea\x0c\x38\x14\xaa\x77\x47\x8a\x7a\x06\xa9\x45\x90\xcf\x10\xf6\xd4\xb6\xa2\x4e\x3c\x39\x80\x22\x06\x53\x40\xc0\x16\xd0\x63\xc8\x02\x02\xe0\x48\x51\xc7\x4f\x3e\xe4\x80\x23\x46\x90\x37\x57\xc5\x08\x33\x10\x38\x8c\x37\x8a\xe1\x9f\x8f\x14\xf5\x1b\x7c\xba\x04\x2e\x14\x58\xa4\x66\xb8\x86\xff\x09\x10\x81\x76\x5b\x19\xf9\xd0\x42\xb3\x27\xe5\x09\x07\x44\xe9\x7e\x1f\x29\xfd\x9e\xa1\x7c\x83\x4f\xca\x10\x20\x52\x89\x4f\xd6\xfd\x3e\x6a\xb7\xfb\x3d\xa3\xdd\xfe\x06\x9f\xf8\x57\xf1\x83\x98\x81\xcf\x36\x72\xc0\x12\x0e\x3c\xca\x80\x67\xc1\x1e\x0e\x3c\xb6\xdb\xc4\x6c\x01\x15\x2f\x70\x4d\x48\x14\x3c\x53\x7c\x82\x96\x80\x41\x05\xcc\xa1\xc7\x14\x0f\xdb\x90\x2a\x98\x28\xc0\xb2\xa0\xcf\x04\xac\x2d\x57\x9a\x40\xed\x52\xf4\x4f\xd1\xa1\x29\xd0\x1a\x06\xa6\x83\xac\x03\x21\x27\xc6\x3a\x04\x6e\xba\xc0\xad\x6b\xbb\xc8\x3b\xc7\xe1\xbe\x4a\x84\x56\x20\xb5\x8a\xf8\xaf\x5a\xcb\xdf\xf0\x90\xad\x86\x80\x31\x48\x04\xf6\xff\xf7\x97\x1f\x35\xad\x75\xf7\xf7\xc9\xa4\xf2\xd7\x9f\xf5\xe7\xe8\x97\xaa\xfc\xe1\x7f\x78\xa7\x0b\xe4\x9d\x43\x6f\xce\x16\x1c\xbe\x25\x5a\xc0\xe3\xba\x45\x3f\x55\xb3\xcc\x78\x25\x7e\x02\x4e\x92\x28\x83\xa1\x42\x80\x37\x87\x0a\xc3\xca\xc3\x02\x31\xe8\x20\xca\x94\x59\x48\x0e\x4a\x79\x3b\x87\x03\x7c\x89\xca\x6f\xd8\x83\x15\xe5\x22\xa0\x4c\x31\xa1\x02\x94\x25\x70\x90\xad\xf4\x06\x67\xd7\x82\x54\x3d\xec\x51\x46\x00\xf2\x58\x6a\x6e\x37\xb7\xc7\xf3\xb3\x40\xdc\xf7\x91\x37\x97\x42\x97\x92\xc1\xa1\x64\xa2\x51\x60\x7a\x90\x5d\x73\x34\x45\xb3\x1d\xa3\x7f\x44\x5d\xc3\x90\xcc\xcb\x80\x75\xdf\x23\x50\xec\xc4\x18\xb9\x10\x07\x2c\xdd\x67\x38\x6e\x34\x2f\xe2\x3c\xb5\x65\xf8\x46\x6c\xf8\x38\xfb\x85\x7b\x99\xe8\xe0\xd6\x2b\x8f\x0e\x20\x73\x29\x4a\xb7\xc3\xde\x2e\xb8\xeb\xc7\x45\x1c\xbe\x7d\x8a\x0b\x40\x19\x24\x3b\xc2\xf3\x0e\x97\xdd\x71\xd7\x45\x02\x08\x06\x9a\x05\x3d\x46\x80\xa3\xe9\xe9\x5e\xc0\x45\x9a\x51\x6b\x58\x27\xa0\x6e\x8b\x89\x60\xa0\x3d\x40\xca\xf2\x21\xeb\x27\xc7\x35\xb3\xd6\xa8\x09\xc8\x80\x4a\x48\x23\x0f\xd2\x34\x8f\x5b\xba\x71\x6a\x26\x20\xf3\x67\x37\x0d\xd3\x68\x1d\xcb\x65\x52\xa0\x41\x50\x04\x69\xb6\x4e\x0c\xdb\x04\x0d\x01\x09\x7c\x8d\x62\xae\x68\x8b\xc1\x6b\xa7\x86\x0d\xa0\x11\xa1\x30\xc7\xcb\x0d\x68\xc0\x53\x60\xea\x8d\xd3\x56\x04\x5d\x3c\x6e\xc3\x6a\xc1\x86\x69\xe4\xa0\x91\x4b\x89\x56\xeb\xb8\x51\x33\x40\x3d\x02\xf7\x30\xd9\x84\x75\xb3\x69\x19\x2d\xd8\x6c\xc8\x5d\xbc\x86\x73\xce\xe0\xb8\x70\x2b\x29\xe3\x86\x67\xb5\x3f\x00\xd6\xf5\x66\xb3\x99\xb3\x93\x49\x40\xf3\xc4\x32\x41\xdd\x6a\x14\xd0\x26\x0d\xac\x1b\xc0\xb2\x8f\x73\xf6\x32\x09\x08\x61\xf3\x44\x6f\x9c\xc2\x1c\xf6\x48\x02\xda\xd6\xb1\x09\xea\xa6\x95\xb3\xe7\x49\xc0\x63\x1d\xd6\x4f\x9a\x35\xbb\x68\xcb\x53\xcb\xd7\x8d\xda\x0c\x36\x8d\x9c\x3d\x4c\x0d\x6b\xeb\xf5\xd3\xd9\x09\x28\xda\xc2\xd4\xfa\x75\xa3\xa5\xdb\xb6\x51\xb4\x83\x49\xe8\xd6\x71\xf3\xb4\xd5\x9a\x9d\xa8\xcf\x62\x07\x7b\xd8\xb3\x11\xd7\x52\x52\xd5\xc9\x0d\x1d\xd0\x1b\xda\x07\x94\xc9\xfe\x9f\xbd\x76\xbb\xff\x9f\x00\x38\xc2\xdd\xe0\x30\xb3\x95\xfd\x96\xf0\xa9\x05\xdd\xad\x59\x83\x8f\xf4\x05\x2f\xbf\xc3\xfd\x07\x8b\xed\xf9\x9d\x40\xf5\x2a\x60\x7e\xc0\x24\x9e\xfd\xc7\x05\x32\x11\xc3\x64\x54\xff\x14\x58\xf7\x50\x6a\x57\xe1\x14\x85\xeb\x10\xc3\x66\xc1\x9e\xb3\x16\x69\x05\xa4\x8c\xea\x8a\x29\xc0\x14\x2f\x72\x43\x62\xda\xf0\xcc\xa3\x5d\xdb\x26\x90\xd2\xd4\x64\x7c\x4d\x5f\x20\xeb\x32\x26\xfc\xb1\x58\x97\x73\x0c\xec\x4f\xc0\xe1\x3a\x51\x9a\xec\xcb\x91\x70\x70\xee\x72\xd0\x90\xdd\x14\xca\xfb\x51\x89\xe7\x8e\x13\xf6\x1d\x40\x19\xb2\xf6\x99\xec\x02\x52\x4c\x15\xa9\xb5\x43\x03\x08\x19\xb0\x01\x03\xd2\x6f\xb0\x30\x1d\xb8\x60\x0e\x7b\xd8\x75\x91\x34\x13\x66\xa3\x3e\x9b\x9d\x80\x9a\xd9\xd2\x8d\x86\x6d\xb6\x1a\xf5\x96\xd1\x6a\x99\x27\xa7\xad\x99\x01\x0d\xdb\xaa\x5b\x56\xed\xf4\xd8\x12\xee\x05\x74\x7d\x07\x30\xf8\x05\x7a\x90\x08\x1b\x78\x06\x98\xe0\x40\xa3\xa6\x1f\x6b\xba\xae\xe9\x4d\x45\x3f\x6d\xd7\x4e\xda\x8d\x93\x8a\x7e\x52\xd7\x6b\x92\x7b\xbb\xdf\x47\x51\xdf\xcf\x98\xb8\x80\xdd\x42\x42\x43\x94\x8d\x9a\x5e\xd3\x6a\x2d\xad\xd6\xca\xf1\x2a\xce\x7a\xd5\xab\x91\xf0\x37\x7b\x0e\x0e\x6c\xd9\x19\x61\x4f\x89\x86\x53\x05\x4b\x52\x1c\x10\x0b\xe6\x59\x75\x1c\x30\x38\xe6\x72\xd2\xa5\x14\x5b\x68\xed\x42\x0d\x09\xf6\x21\x61\x28\xec\xb5\x06\x1c\xd8\x71\x4e\x0b\x47\x5b\x7f\x96\x86\x5a\x0c\x9e\x0b\x29\x3f\x49\xda\x67\x5d\xe2\x4d\x58\x65\x8c\xf4\x9c\xb3\xc8\x78\x41\x20\xcc\xc3\x78\xe0\x0f\x09\x66\xd8\xc2\x8e\xf0\xe6\x2d\x5f\xb8\x48\xc8\x26\x03\x3f\xe3\x17\x7e\x21\x38\xf0\x53\xe8\xae\x67\x1a\x41\x2b\x20\x88\x3d\x09\x28\x81\xc6\x67\x82\xdd\x21\x26\xd2\x3f\xae\x35\x0d\xb1\xf9\x38\x6a\xa9\x1b\xb5\x9a\xb4\xc5\x39\x0b\x8c\x8f\x15\x2e\x41\x0d\xa5\xdc\xc4\x81\x67\x5f\x42\xf6\x80\xc9\x7d\xd7\x72\xfa\x1e\x23\x4f\x79\x2b\xe3\x8b\xf8\xe4\x60\xeb\x3e\xb3\x8e\xeb\xc0\x81\xa1\xa3\x2c\x3c\x1c\xd1\xb8\x1e\x31\x6f\x85\xeb\xaf\x02\x8d\xfe\x3c\x94\xbb\xd5\xf1\x2a\x4e\x45\x4d\x8f\x66\xe9\x5a\x11\x0b\x02\xee\x41\x0b\x40\x4c\x62\x9e\x16\x27\x91\xc0\x4f\xd2\x86\xff\x78\xdc\x6c\xd6\x9b\x45\x3b\x9f\x5e\xb7\xd8\xed\x04\xcb\xe4\x50\xe2\xd6\xb7\x92\x6b\xba\xf5\xa5\xf1\x8a\x93\x48\xea\x8d\xcf\xc8\xb3\x07\xde\x05\xf0\xa5\xae\x5a\x7b\xba\xe9\x79\xe4\x12\x8e\x94\x95\x07\x20\xf4\xc8\x18\xcc\x43\xcd\x9d\x51\xb6\x62\x0d\xc2\xf3\x5d\x1d\xe2\xbe\xc1\x27\xf1\xc5\xf7\x9d\xe8\xb8\xf9\x7c\x14\x53\x65\xd1\x8c\xea\x1a\x36\x5c\xbf\xfa\x1c\xaa\x2d\x1f\x7a\x36\xbd\xf2\xa2\x45\x6d\x90\x96\x38\xa9\x84\xe4\xe4\x51\xea\x0c\x52\x86\x3c\x81\xca\x06\xf6\xd9\x4b\xc8\x23\xaf\x37\x09\x7b\xd9\x1d\x47\x1f\x8a\xb6\x5a\x22\x29\xb8\x2d\x47\x83\xe7\x20\xff\x15\x02\x87\x2d\x7a\x0b\x18\xee\xe7\x98\xfb\xd3\x42\xd4\xc6\xbd\x61\xbb\x59\x6b\x0a\xec\x6f\xbc\x85\x80\x7b\xe2\x2a\x81\x2e\xb0\x63\x0b\xf5\x29\x48\xb7\x3a\x8a\xa8\x4d\x55\xa0\xce\x20\x59\x02\xc1\xd3\x75\xd1\xf9\x6b\x6e\x57\xa1\xcb\xe2\x12\x9b\xb4\xde\xe7\x66\x46\x35\xac\xb9\x82\x9f\xd8\x52\x9f\xef\x8e\x14\xf5\x1c\x51\xc6\xcd\x43\x38\x52\x7c\xe9\x91\xfe\x38\xad\xa5\xe5\x6e\xdc\x1b\xaa\x31\x8a\x6f\xfa\xb4\x1e\x43\x60\x93\x37\x7e\xa3\x51\x7f\xfd\x04\x7c\x10\xb1\x20\xc9\x84\x49\xc2\xc4\x4f\x76\x79\x1c\xfd\x05\x30\xf8\x00\x9e\xc6\x58\x6c\x04\x07\xca\x30\x4a\x9a\x33\x90\x37\x6f\xb7\x13\x8c\x92\xb2\x07\x23\x48\x96\x90\x48\x4a\x0b\x2d\x19\x1e\x43\x87\xd8\x41\xd6\x53\xc8\xa2\xd2\x10\x8e\xd0\xdc\x13\x9b\xff\x53\x8d\x1f\x52\xb7\xa9\x89\xdc\xd3\x6d\x5a\x51\xac\x6f\x45\xb2\x86\x24\x79\x73\x22\x1d\xd1\xac\x9c\x22\x02\xed\x1e\xf0\x81\x85\xd8\xd3\x6e\x03\x71\xb6\x8a\x51\x86\xab\xa1\xdc\xfd\xc8\xb8\x66\x62\x03\x2f\x90\x37\x42\xbf\xc1\x9d\xa7\xba\x00\x8f\x7b\xc1\x77\x97\x00\x39\xc0\x44\x0e\x62\x4f\xff\x8b\xbd\x08\xb5\x7c\xef\x51\xb2\x4c\x4e\x2f\xf5\x4e\x20\x9b\xd5\xc3\xaa\xcb\x7d\x3a\x4d\xf8\x8f\x31\x7d\x4a\xb0\x13\x99\x2f\x1f\xcc\x01\x83\x5d\x76\x0e\x02\xcf\x5a\xac\xef\xf6\x84\x38\x8a\xb6\x1e\xf6\x66\x68\x1e\x48\x9f\x6d\x75\xd9\x97\x43\xbc\x18\x74\x74\xc5\xc0\x91\x1b\xd8\xd0\x63\x68\x86\x84\xfa\x3a\xa4\x18\x74\x03\x86\x47\x16\x70\x04\xf7\xc7\x7e\x59\xbb\x22\x59\x8f\xe8\x33\x0e\x72\xb5\x68\xd2\x21\x0a\xec\xb7\x72\x88\x6a\x49\x6f\xc8\xd0\xf7\x75\x85\xb2\x4b\xba\xf2\xde\xd5\xc5\x7b\xf5\x8a\xb2\x1b\xbd\x8f\x33\x13\x75\x0a\x47\x49\x02\xa5\x3e\x4a\x9b\x9b\xe0\xb2\x34\x44\x2e\xea\xb7\xc3\x5e\xf8\xbd\xcb\x18\xb0\x16\x2e\x0c\xa5\x37\x6e\xcf\x73\x90\x0e\xdd\x16\x31\xc7\x0c\x58\x91\x70\x9f\x41\x07\x32\x78\xe5\x8d\x21\x71\x43\xb7\x23\xee\x53\x9e\xc1\x25\xb2\xe0\xc0\xb3\xe1\xe3\x8a\xbe\xb9\x67\x86\x84\x00\x71\xad\x10\x9e\x07\xa0\xfc\x32\xf0\xd7\x07\xc5\xd5\xe0\x82\xfa\x23\x41\xe3\xb5\x28\xe6\xef\xf2\xea\xb3\x3c\x11\xee\x69\xcd\x85\x02\x89\xbf\x0e\x84\xc0\xf1\xeb\xfc\xd4\xd5\xa0\xea\xd6\x2b\x2e\xb4\x51\xe0\x0a\x9b\xca\x8f\x9a\xe1\x8a\xb3\xc6\x27\xbc\x2a\x3c\x52\x0a\x6f\x0b\x12\x86\x67\x24\x4c\x1b\xf7\xf3\x22\x47\x29\x7c\x19\x79\x89\xf5\xe5\x2c\xb1\x76\xe3\x8e\xd6\x04\xfa\x50\x36\x56\xdf\xcd\x82\xea\x5b\xed\x63\x9e\x1f\x1a\xdf\x7e\x49\x28\xe0\x48\x2a\x6c\x32\xa1\x7a\xc2\x40\xea\xbf\xcc\xfc\xb9\xf2\x4a\xe3\x2d\xec\x5f\xb8\xe8\x8f\x67\xfa\x84\x4c\x8f\x71\x74\x07\x3f\x5f\xdd\x18\xa5\x38\x42\x4a\x46\x52\x39\x27\xd4\x4c\x81\x09\x48\x9a\x17\x79\xf6\xcd\xe9\x5d\xa0\x3b\x36\x59\x90\xf0\x24\xfc\x5a\xb3\xf8\x19\x2d\x77\xb0\x8b\x6f\x67\xe9\x8d\x7a\x72\x59\xe2\x24\xf6\xf2\x55\x25\x58\x2c\x67\x59\x7d\x93\x5e\xf9\x0c\xb9\xe8\x37\xf1\x26\xbc\x52\xf8\x5b\x54\xf0\x86\x03\x5c\xf1\x4a\xef\xb6\xaa\xe7\xf8\x1b\xc0\x06\x1d\x1d\xde\x42\xdf\xed\x6a\xbd\x6e\x28\x24\x67\xd1\x7d\x24\x9f\xf5\x13\xa0\xf0\xb8\xb1\xfa\xf5\x1f\x18\x79\x62\x7e\xf5\x48\xf9\xa1\xfe\xd9\x72\x70\x60\x6b\x96\xa0\xd9\x44\x3c\x6a\x4f\x54\x0b\x13\x88\xe9\x44\x6d\xcb\x06\x45\x99\xa8\x81\x87\x58\xbc\x45\x53\x38\x98\xeb\x02\xcf\x9e\xa8\x6d\xe5\x77\x2d\xfa\xc0\xff\x51\x06\x08\x5b\x37\x70\x48\x8f\x41\x8f\x09\xc8\x38\xe0\x8f\x1b\x0f\xb1\xbb\x78\x4b\xec\x7a\xb2\xd3\xfd\x3e\x52\x46\x90\x05\x7e\x5b\x91\xf7\x92\x54\x3c\x65\x56\x97\x80\x54\x1d\x64\x2a\xd0\x5f\x40\x17\x12\xe0\x28\x36\x41\x4b\x18\x1f\xe6\x13\x9c\x61\x02\x3b\x4b\x40\x34\x07\x99\x15\x97\xeb\x7b\xc5\x36\x03\x5a\xa1\x90\x70\xd7\x21\x81\xc5\x48\xb6\x25\x10\xe1\x1c\xd8\xe1\xda\x76\x81\x59\xbc\xfd\x1a\xba\x00\x79\xdd\x19\x83\xa4\xff\x88\x58\xe7\x09\xd2\xf8\xe7\xfe\x23\xb4\x46\x7c\xf9\x9d\xaa\x89\xbc\xaa\x09\xe8\x42\xd1\x2c\x65\xa2\xfe\xc5\x74\xee\x91\xad\x68\x4c\x19\xff\x6b\xd8\xef\xc0\x47\xd6\x50\x7e\x57\xe6\x04\xfa\xca\xe3\xd2\x36\xff\xaa\xfc\xfe\xbb\xf2\x97\x6a\x40\x49\x95\xf2\x8e\xee\xfd\x8c\x56\x04\x90\xf6\x59\xa9\xda\x70\x59\x15\x50\x13\x35\x4e\x56\x0f\xb8\x30\x4b\xfd\x99\xa0\x95\x16\xae\x5d\x5b\x51\x29\xbd\xf4\x5f\xbc\x85\x17\x62\x0f\xa2\xcd\xcb\xd9\xac\xc2\xed\x11\x3d\x13\x53\x7c\x5f\x00\xd6\x59\x51\x25\xf9\x05\x12\xd8\xc9\x9b\x45\x6c\x28\xa7\xe8\x76\x0a\x26\xd8\x66\x47\x72\x61\x3f\x3e\xae\x0b\xe8\x3d\x87\xfa\xd3\x9f\x4c\x8c\x9d\x14\x30\x97\xd4\xed\x48\x40\x66\xd9\x7b\x6e\xd9\xc1\x71\x08\x7c\x1b\x30\xa8\x41\x6f\x8e\x3c\xf8\xde\xc8\x38\xd8\xba\xa7\x2e\x62\x8b\xd7\x91\x25\x7f\x70\xfa\x44\x19\x74\x6d\x8d\x40\x8a\x9d\x25\xdc\x6f\x0a\x02\x33\xe2\xb2\x79\x96\x7f\xe3\x80\x7b\x83\x6f\x34\x8b\x8d\xad\x7b\x48\x7e\x95\xb8\x47\xd2\x6b\x61\x5a\x61\xe2\x5e\xf5\xc5\xba\x75\xc4\x80\x67\x03\x62\xcb\x07\xcb\x4e\x48\xa5\xbf\x5b\xd8\xa3\xd8\x81\x79\x90\x7d\x42\x30\xd9\x04\xb8\xd2\xc7\x43\xae\x16\xb8\x7a\x95\xda\xd5\x46\x44\xd1\x7c\xa5\x0a\x99\x55\xf5\x09\x9e\x21\x07\x56\xec\x02\x45\x1e\xf5\x72\x3c\x45\xa3\x33\xa5\x8a\x7d\x56\x15\x4e\xb3\xcf\x95\x4d\x15\x7a\x4b\x44\xb0\xc7\x0f\xbc\x15\xf8\xe8\x63\xc2\x52\xc3\x56\x05\x6d\xe8\x62\x87\x9d\xb3\x30\xd5\x1c\xe4\xdd\x6b\xd0\x5b\xe6\x6f\xe0\x4b\xb4\xf0\xf0\x7e\xee\x73\x7a\xb5\x95\x33\xfc\xe0\x39\x18\xd8\x8a\x7c\xf7\x13\xc1\x41\x88\x2a\x0b\x4c\x59\x25\xde\x5b\xd8\xb7\x8e\x27\xcf\xe6\x1a\xf6\x1c\x2e\xff\xd9\xdd\xfd\x0e\x3c\x46\xb7\x83\xad\xde\xca\x87\x80\x2d\xfa\x8f\x88\x32\xda\xf9\x53\x9a\x8c\x1f\x9f\x6b\xac\x80\x38\x8a\xa6\xdd\x43\xe8\x03\x07\x2d\xa1\xc6\x90\x0b\x15\x43\xd1\x66\xe7\x74\xb4\x54\x34\x8d\x40\x46\x9e\x14\xa3\xa6\x68\xff\x52\xf4\x1a\xff\xa7\x68\x4f\xca\x71\x4d\xd1\xb0\x52\x65\xae\x5f\x35\x31\x66\x94\x11\xe0\x73\x22\x55\x1e\x7f\x53\x16\x8c\xf9\xb4\x5d\xad\xda\xe1\xbe\xd0\x8a\xe0\x15\x84\x05\xcf\x54\xa5\xef\xb7\xee\x56\x85\x27\x75\x13\x18\xa6\x6e\xd9\xfa\xc9\x49\xab\x09\x1b\xb6\x65\xd6\xed\xe3\xe3\xc6\x89\xad\xeb\xc0\x68\x99\x2d\xab\xde\xac\x9d\x36\x2a\xe9\x99\xf6\x91\x89\xe4\xc6\x6c\x11\x0a\x06\x88\xa2\x81\xc7\x59\xc1\x02\xb5\xde\x4e\xe3\x0d\x31\x65\x1d\x6d\x35\x28\x71\x15\xad\x60\xc4\x1d\x85\x28\x22\xe8\x7b\x38\x3d\x6b\x71\x13\x91\x7a\xc0\x41\xbf\xc1\x50\xe0\x66\x98\x14\x48\x5c\x18\xf3\x48\x3b\x1b\xd1\x5f\xcb\xe6\x56\xb0\x77\x97\xa1\xb5\x5a\xfc\x8c\x1c\xd8\xd9\xa0\x36\x0b\x58\x2c\xd5\x81\x33\x86\x1f\x92\x56\xa1\xdc\xa9\x54\x34\xcd\xc3\x9a\xc9\x1d\x03\x2d\x34\xad\x09\x0a\x88\x0b\x28\xc7\xb9\x4b\xab\x2c\x68\x7f\x7a\xea\xb8\x81\xc3\x90\x16\x50\x48\x32\x1a\x6b\xa2\x42\x11\xf7\xfb\x4a\x87\x45\x6c\x91\x40\x34\xbb\x3f\x13\xd5\xc3\x53\x81\xf9\x4e\x93\xbc\x9c\x6d\xb5\x17\x39\xeb\xf2\xf2\x2d\x1d\x13\x32\x0a\x44\x4c\xe9\x76\xd5\x1e\xba\xdf\x42\x89\x69\xd6\xcc\xd3\xa8\x18\x6f\x37\xf6\xa4\xc8\xf5\x93\xac\x74\x2d\x3d\x9f\x0e\xf6\xb4\x19\x40\x4e\x40\xd2\x1c\x49\xd8\x39\x72\x11\x8b\x5e\x65\x3b\xb5\x9c\xde\x23\x68\x75\xf4\x26\x3d\x04\x8b\x6e\xeb\xc6\xac\xaa\x35\xf3\xa6\x72\xd1\x53\x37\x0c\x10\x2a\x18\xa2\x33\xe1\x47\xfe\x69\xef\xf3\xe5\x74\x34\xf8\x72\xd9\x3d\x9f\x8e\xbf\x0e\x46\xd3\xeb\xfe\xe8\xea\xe6\xba\xd7\xef\xe4\xbf\x50\xc6\x0f\x82\x69\x8d\x2e\x04\x05\x79\x73\x45\xb3\x74\xc5\x81\xc0\x86\xa4\x22\xd0\x2b\xee\x92\x23\x6a\xf9\xdb\x96\x67\x01\x70\x60\x2d\x94\xed\x7b\xbe\x51\x4c\xd6\x1d\x5e\x2d\x2b\x13\x55\x1e\x5b\x62\xb7\x15\xbc\x91\x40\x6e\x4a\x34\x6e\x4b\x18\x9c\x3f\x65\xf1\xc0\xb3\x99\xfc\x75\xa2\x3e\x10\xc4\xe0\x94\x7b\x6f\xeb\x3b\x8f\x62\xf7\x6b\x27\x23\xbe\xc6\x0e\x3f\x78\x90\xa4\xa7\x27\x78\xad\x90\x27\xaa\x0f\xd8\x22\x0d\x21\xd8\x2a\xb6\x47\x42\xb1\x68\x33\x07\xcc\x69\x95\x40\x1f\x53\xc4\x30\x79\xd2\x02\xe2\xc4\xc6\x81\xc4\x45\x94\x22\xec\xd1\xf4\x70\xb5\xe3\x46\x63\xdb\xc2\x7e\x4c\x54\xb9\x39\xe2\x76\x48\x93\xba\x6c\x5a\x6f\xd6\x8c\x63\xfd\x44\xaf\xc1\x46\xd3\x00\xb6\x79\x5a\xab\x37\x5b\x16\x6c\x9e\x5a\x46\xbd\x71\x7c\x0c\x8d\x93\x56\x0d\xe8\x7a\x6b\xa2\x1e\x29\xe1\x00\x91\x08\xec\x3f\xc4\xdd\x9b\xd2\xcd\x72\x02\xca\x20\xd1\x7c\x60\xdd\x83\x39\xa4\x95\x7f\x53\xec\x1d\x8e\x7c\xff\x90\xd6\x31\xa6\xe2\x44\x98\xff\x12\x3a\x3d\x69\x2f\x3b\x0f\x80\x78\xc8\x9b\xaf\x01\xae\x01\x83\x49\x5d\xa6\xd3\x9c\xaf\x9f\x02\x42\x59\xc7\xe0\x2e\xe6\x61\x08\x14\x9a\xcf\xea\xea\x64\xca\xf7\x3c\x3a\xba\xf0\x9f\x0f\x46\x15\x02\x1f\x09\x78\x6a\x27\x4f\xf6\x73\x87\x53\xa5\xad\x20\x6f\x86\xe3\x5f\x5c\x6c\x07\x0e\xa4\x09\x68\x25\x4a\xef\xd0\x44\x52\x43\xea\x9b\x22\x5c\xab\xb6\xc2\x2c\xbf\x5d\xad\xea\xc6\x89\xb8\xab\xd6\xdb\xc7\x7a\xad\x56\x4f\xdd\x0b\x10\x30\x87\x67\x04\x2d\x21\x49\xce\xa0\x29\xd0\x32\xe2\x0d\x4b\xec\x04\x2e\x4c\xcd\x14\x78\xe2\x7e\x28\x33\x3f\x9a\x7b\x98\xc0\x80\x42\xdb\x12\xdf\x53\xca\xa9\x70\x0b\x24\x61\xaa\x52\xd8\x2a\x4f\xee\xe1\xe4\xf8\xa2\x3f\xba\x1a\x4d\x7b\xe7\x37\xa3\x71\xff\xba\x93\xb9\x6c\x4e\x86\x9e\x09\x0d\xf8\x02\x89\x8a\x64\x48\x1a\x00\x9f\xe0\x25\xb2\x21\xd1\xc0\x03\x0d\xa5\x7e\xdd\x2d\x7c\x6c\x5a\x01\x6d\xc3\xbf\x7b\x76\x31\xb8\xbc\xbe\xba\x19\xf7\xaf\xa7\xdd\xde\x78\x70\xdb\x1d\xf7\xa7\xdd\x9b\xf1\xd7\xe9\xc5\xd5\xd9\xcd\x79\x3f\xb1\xa4\x44\xe2\xd7\xdb\x2f\x48\xf0\x20\xc1\x01\x83\xa4\x02\xbd\xe5\xd6\xad\xe8\xf2\x3d\x98\x86\xb6\x1d\x46\x11\xd4\xd3\x80\xa0\x35\x50\xff\x9f\x5f\x07\x9f\x06\xe3\xab\xeb\xe9\xcd\xf5\xa0\xc3\xcd\x4b\xbb\x5a\x15\x6b\x4c\x3d\xfa\x6d\x78\x56\x4c\xc7\x31\xb7\x4f\xf5\x53\xbd\xba\x9a\xb0\xba\xd4\x23\xed\xc7\x0d\x14\x0b\x68\x9c\x4c\x71\x14\xba\x67\x67\xd7\xfd\xd1\xa8\xf3\xca\xf9\x93\xc3\x73\xef\xe6\xfc\xb6\x7f\x3d\xea\xe8\xc7\xad\x8a\xd1\x6c\x54\xe4\xff\xeb\x6f\xb7\x51\xb6\x47\xa7\xf1\x67\x8e\xe2\x3d\x5a\xaf\xfd\x53\xb7\xf7\xad\x7f\x79\xc6\x3d\xe3\xe9\x28\xa6\x3d\xf8\xef\xd7\xfd\x2f\x83\xab\xcb\xac\x34\xc5\x9e\x6e\x92\x6b\x96\x83\x4c\x3f\xdd\xf4\xbe\xf5\xc7\x89\x7e\xf9\xd1\xf6\xb9\xbd\x87\xd7\xfd\xcf\x83\x7f\xbe\xaf\x0c\xaf\x98\x68\x1b\x21\x7f\x4e\xd4\x00\x25\x9e\x86\x27\x6a\xfb\xe7\x44\xf5\x9d\x60\x8e\x84\x2a\xfb\x39\x51\x4d\xe0\x49\xb3\xf5\x33\x3a\x76\xf1\x23\x8e\x88\x3c\x38\xe2\x5f\xad\xfb\x39\xc1\x81\x67\xf7\xb0\x83\x39\xd8\x44\xfd\xb3\xde\x37\xea\xc6\xe7\x89\x7a\x34\x51\x67\x98\xc0\xcc\xf7\xcf\xe2\x9f\xf8\xbe\x10\xee\xef\x18\x31\x71\x9a\xf3\x02\xc7\x59\x35\xf6\x56\x38\x87\xcd\x33\x8c\x59\x4e\x33\x72\xc1\x1c\x0e\x25\x0d\xc3\x26\x1b\x51\xa1\x8f\xcd\xd5\xa8\xcf\x1c\x57\x02\x3c\x9b\x5b\xf3\xbc\xb5\x70\x00\xf8\x28\xc5\x45\x5c\xf1\x09\x4d\x9e\x80\x53\x22\x40\x2f\xb9\xf5\x13\x35\x99\x79\x5a\x34\x7e\xa6\x1b\xe6\xfd\xd2\xd0\x1b\x75\xe5\x91\x9c\xeb\x2b\xa6\x4c\x90\x72\xe5\xda\x72\x37\x80\x7f\xa9\x55\x2c\xec\x4e\xd4\xe7\xa3\xf5\x5c\xd9\x79\x19\x01\xd6\x7d\x0e\x21\xb8\x21\x7c\x7e\x7e\x7e\x7e\x3b\xd6\x0c\x50\xe8\xac\xc6\x3c\xb9\x0d\x76\xe5\xc5\x62\x3c\xee\xf6\xbe\x4d\x07\x67\x05\x72\x38\xc8\x1a\x9e\x78\xbf\xcb\xee\x45\x7f\x1f\x09\x96\x7d\x07\xdd\x8b\x69\x68\x3d\xae\xaf\xce\xfb\xd9\x51\xa4\x12\xbe\xc6\xce\xa6\x11\x46\xe7\xdd\xdb\x7e\xc1\x00\xe2\x8c\x99\xdb\xff\xf0\x1b\x55\x78\x3e\x2e\xde\xad\xc1\xe5\xb8\x7f\xcd\xcf\xc7\x21\x15\xce\x3f\x4d\xb9\x85\x59\xad\xe2\x50\xb6\x69\xb7\xe1\x77\xc9\x1d\x7a\x63\xcf\xe3\x81\x4e\x6d\x8f\xf2\x43\x35\xcd\xa7\x5d\x6c\xa2\x9f\x6f\x24\x74\xfe\xfd\x9c\xfb\xec\xc5\xe2\x36\x51\x27\x7b\xd2\x80\x60\x07\xd2\xaa\x08\xc7\x9d\xca\x4c\xef\x83\x0e\x0c\x1e\x38\xb5\xee\x64\x6a\x9c\x48\x63\x90\xc1\x84\xb1\x44\xe6\x1f\x3f\xd5\x5b\x44\x58\x00\x9c\x30\x2c\x44\x5d\x3d\xa2\x8b\x90\x17\xd9\x21\xfa\x26\xde\xa1\xa9\x6d\x86\x31\x1f\xe9\x54\xde\xad\xb9\x23\x05\x69\xc3\xe9\x98\xb5\x01\x70\x63\xd1\xf5\x33\xe4\xc0\x4c\x04\x52\xfa\x7b\x36\x7d\x22\x11\x13\x95\x13\xba\x95\x08\xd5\x8b\xc7\xb3\xec\x13\x6e\x2a\x7a\x6c\xc8\x9b\x93\x71\x7a\x79\xf1\x3c\x82\xf4\xb9\x41\xb9\x5b\x43\xb0\x72\x52\x29\xd2\x59\x56\xa9\x18\xaa\x28\xd8\xe8\xad\x66\x3c\x4d\xcd\x77\xfa\xa6\xb3\x85\x39\x19\xeb\xe9\x44\x7e\xc5\x5b\xae\x2e\xbb\xbe\xb7\x5d\x21\x3f\xc0\xa4\x66\xe4\x2d\x6f\x38\xa3\x91\x99\x51\xb4\xc8\x68\xc4\x2d\x21\x6a\x99\x02\x05\xeb\xfc\xb4\x2d\x99\x92\x3b\xc5\x31\xef\x9e\x93\xb7\x29\x4f\x32\x1f\xa7\x75\xc8\x78\xae\xf5\xfc\x70\xf9\x55\x16\xd7\xcd\xfc\x77\x14\x22\x9d\xd1\x2f\xaf\xcb\xba\x4a\xbb\x48\xc5\x9f\x37\x05\x3e\x6e\x0e\xf7\xdc\x31\xb3\x2b\x22\x65\x9c\xdb\xbf\x8e\xc7\x85\xb9\x57\x99\x6f\x19\xdd\x97\x37\x4b\xc4\xf7\xaf\x4b\xf0\x32\x56\xd2\x99\x9b\xa3\x96\x33\xc7\x4b\x56\x72\xba\x79\x96\x3f\x4c\x26\x5c\xf1\x0a\xb2\x6b\x78\x11\x9d\x84\x76\xde\x23\xdf\xee\x95\xd9\x74\x87\x08\xab\xde\x18\x47\xfc\xc1\x82\xab\x3f\x74\x0a\xe1\x01\x92\x07\x8b\x72\xf9\xb2\xb9\x7f\xbb\x67\xfd\xbd\x2e\xdf\x2f\x91\xdb\xfd\xa1\x12\xfe\x5e\x94\xea\x97\xac\x6e\xf0\x96\xb9\x7e\x31\xef\x60\x1f\xb7\xfe\xd0\x19\xec\x02\x17\xb5\x28\x81\x7d\x53\x82\x7f\x3c\x61\xfd\xbf\x6a\x31\x87\xf0\x14\x93\x7c\xb4\xc9\x55\xcc\x10\xf0\xc5\xbe\xe2\xc7\x2a\x3c\x21\xb3\xd6\xde\xb9\xf2\x44\xdc\xa0\xbe\x61\xe1\x89\x4c\x05\xb4\xb7\xae\x3b\xb1\x99\xd1\xf7\x2e\x3b\x91\xe3\x6f\xbf\xf6\xaa\x21\xac\xd9\x28\x2b\xde\xad\xab\xe0\x51\x61\x9b\x5f\x72\xf7\xb0\x4a\x7c\x4a\x9e\x12\x56\x05\x03\x33\xc7\xc6\x54\x39\x17\x59\x7c\xee\x40\x83\xef\x75\xab\xb0\xef\xe0\xf9\x97\x08\x7b\x1e\x77\xdf\xa8\x84\xc8\x7a\xec\x75\x05\x91\x17\x26\x34\xef\x98\xbf\xba\x2e\x33\x92\xf0\x68\xcf\x3f\xbd\x79\xaa\x60\xb8\x7d\x07\xbc\xb2\x78\x81\x33\x3b\xc6\xf1\x54\xc1\x77\xc8\x8b\xdc\x3f\xb1\xfe\x15\x69\x91\x31\xce\xfa\xb8\x7e\xc4\x6e\xea\x35\x5d\x4b\x2b\xef\x6a\x38\x67\x8d\xd7\xd8\x49\x65\x35\xc7\x1e\x84\xb8\xe7\x3b\x04\xb2\x3c\x69\x35\x4b\xe1\x41\xf7\x62\x9d\xef\xbd\xba\x7e\x4e\x9f\x08\xb3\x09\xa8\xef\x76\x2c\x7c\x69\xe9\x86\x03\x25\xde\x8e\xd0\xe3\x7b\xe6\xdd\x1e\xa8\xe4\x58\xce\x4c\x87\x79\x27\x18\xad\x6b\xf9\xed\x62\x77\xf2\x8b\x1a\x8a\x12\x12\xf1\xc3\xb7\x7a\x0d\x19\x40\x5e\x46\x68\x46\xf5\x76\x3b\x16\x77\x71\x3b\xec\x9d\x7d\xed\x0d\x65\xbd\x5d\xba\xc5\x05\x2f\x5c\xdc\xd9\xc2\xf2\xc3\x21\x92\x00\xb1\xb1\x73\x4c\x53\xa1\x50\x17\x63\xb5\xda\x09\x21\xa9\x39\x38\x0a\x02\xa0\x48\xb6\x25\x39\xa2\xc7\xab\x90\xd2\xc2\x07\xe6\xed\x67\xd8\x0a\x44\xe9\x0e\x21\x89\x0c\x30\x18\xfe\xf6\xe3\xa7\xda\x9f\xcd\xa0\x25\xb8\xa5\x1b\x39\xce\x2b\x5f\xfa\x87\x2a\x12\x89\x67\x51\xac\x78\xfb\x6f\xa2\xa0\x76\x74\xdb\x91\xad\x59\x19\x7b\x28\xcf\xc9\x4b\x2e\x7e\x54\xaf\xfe\x4d\xbd\xbb\x13\xda\x70\x1b\x42\xd0\x32\xda\xe2\xa6\x04\x0a\x25\x7d\xa4\x88\x16\xc9\x6f\xe6\x4a\x21\xae\x3e\x48\xd0\x5b\x11\xe9\xb7\x06\x76\x60\xba\x4d\x96\x37\x49\xc3\xe5\xb5\xc9\x89\x64\x6b\x66\x7e\xd9\x3c\x5a\x85\x5e\x65\xbf\x75\x19\x23\xc8\x0c\x18\x4c\x62\x38\xf2\x80\x4f\x17\x98\xad\x5a\xb1\xff\x94\x6e\x93\x78\x67\x5b\xe5\xf0\x51\x7b\x66\xde\xe8\x43\x62\x66\x10\x30\x4c\xc3\x5b\x85\x08\x30\x7d\xb9\x20\x46\x12\x0c\xf0\x00\x98\xb5\x68\x0f\x03\x76\x01\x19\x41\x96\x48\x44\x4f\x31\x82\xfa\x37\x69\x4e\x92\x25\x2a\x0d\x4d\xaf\x69\xfa\x89\xfa\x2c\xbe\x75\x29\x0d\x5c\xc1\xd0\x2f\x62\xcb\x21\x41\x9e\x85\xfc\xf0\x7e\x2d\xcc\x34\x88\x98\xa2\x02\x5c\xf0\x1b\xf6\xc0\x03\xad\x58\xd8\x0d\x13\xeb\xd7\x6c\x43\x19\x6d\xaf\xa7\x0f\x2f\x91\x8a\x70\xcd\xb5\x83\xab\x10\x8a\x58\x44\xc6\x9e\x62\x49\xb0\xdc\xb6\xd7\x0a\x25\xad\xb7\xbb\x26\x26\xec\x22\x70\x18\xf2\x01\x61\x37\xbe\x83\x81\xcd\xa1\x68\x3d\x64\x93\x2b\xf3\xdf\x7c\x04\xd9\xf4\x05\x32\xa9\x0a\xf9\xb1\x3a\xd5\x14\x2a\xd2\x55\x6b\xba\xa3\xfc\x7d\xdd\xf1\x1c\xd1\xb0\x67\xa6\x21\x85\x0f\x8d\x01\xa4\x3e\x0d\x01\x61\xd1\xe7\x61\x90\x9c\x73\xf5\x3b\x9f\x33\xab\x6e\x72\xb4\x0a\x6f\x1a\xcc\x62\xe5\x17\x12\x75\x76\x39\xb7\x13\xaf\x0d\x1e\xa8\x26\xeb\xe8\xb6\x69\xbd\xdd\x6e\xc7\xda\xc3\x86\xbb\xe7\xad\xd1\x7b\x52\x4f\x15\x28\xb7\x5f\x84\xc6\x8e\x8a\xf2\x43\x6a\xee\x48\xd5\x84\x7f\xe3\x20\xab\xab\xc2\x57\x8c\x02\x0d\x95\x73\x4b\x9b\x81\xbd\x11\x69\x22\x99\x8b\xd2\x7d\x74\x5e\x1e\x60\x08\xc4\xd7\xb3\x44\x42\xda\x39\xea\xf2\x01\x85\x73\xb4\xb9\x7a\x40\x59\x21\x1b\x7b\x47\xa1\xff\xc5\xea\x72\x87\x33\x44\x4a\x77\xbe\xe8\xd8\x12\x0f\x84\x7b\xe9\xb9\x45\x8e\x71\xb0\x23\x4b\xd1\x2b\xd4\x87\x3c\xad\x64\x8a\xb5\x7e\xdc\x03\xf1\xd6\xc3\x6f\x51\xe1\xca\xbd\x82\x18\xbe\x8e\xc7\xc3\x76\xab\xd6\xaa\x55\xa7\x0b\xe0\x13\xfc\xf8\x34\x95\x91\x0c\x53\x4b\x40\xef\x1c\xdb\x60\xa4\x63\x1b\x9a\x5b\x42\x1b\x5e\x5a\x79\xa8\x2c\x12\x7b\xc8\x22\xb1\xe9\x1b\xc5\x1c\xe6\xf9\x88\xf7\x40\x99\x8b\xd0\xf4\x83\xc1\xab\x6b\xb7\xe7\xdd\xc9\x6e\x0f\x5f\xda\xf4\xca\xb4\xa1\x72\xbb\x54\x4b\xef\x50\xe6\x3c\xfb\x4e\xf6\xd1\xea\x9c\x47\xd6\xea\xf5\x21\x17\xaf\x34\x54\xbf\x2a\xda\x22\x7d\x0d\x7e\xc8\xf0\x54\xe5\xfc\xd3\xce\x37\x4f\x39\x7f\x4e\xe0\x01\x6f\xbf\xd9\x7b\xbb\x4a\xb3\x07\xa9\xa8\x27\x19\xfe\xd7\x3f\xeb\x6e\x11\xb4\x8f\xf0\xae\xfb\x4e\x75\x06\x3f\x44\x85\xc1\x15\xfd\xcb\x12\x83\x65\x89\xc1\xb2\xc4\x60\x59\x62\xb0\x2c\x31\x98\x8f\x49\x59\x62\xb0\x2c\x31\x58\x96\x18\x2c\x4b\x0c\x96\x25\x06\xcb\x12\x83\x65\x89\xc1\x37\x94\xa1\xb2\xc4\x60\x59\x62\xb0\x2c\x31\xf8\xea\x12\x83\x65\x71\xc1\xb2\xb8\x60\x59\x5c\xb0\x2c\x2e\x58\x16\x17\x2c\x8b\x0b\x96\xc5\x05\xcb\xe2\x82\x65\x71\xc1\xb2\xb8\x60\x59\x5c\xb0\x2c\x2e\x58\x16\x17\x2c\x8b\x0b\x96\xc5\x05\xcb\xe2\x82\xb9\xb4\xfb\xc3\x17\x17\x3c\xe8\x88\x1f\xad\xaa\xe0\x47\xaa\x27\x58\x18\xbc\x74\xc8\x64\x41\x25\x0c\x99\xdc\x35\x72\xeb\xb0\x29\xa9\x7f\xcc\x64\xd4\x03\x64\x78\x1f\x22\x17\xf7\x17\x2d\x3c\x9e\x78\x99\x57\x9d\x00\xbb\x00\xad\x4b\xf9\xe4\xa6\xe8\xf4\x41\x98\x9f\x03\x2d\xa3\x12\xab\x36\xb7\x35\x0f\x26\x66\xf0\xb9\x8f\xe3\x07\x0c\xae\xfb\x87\x59\x31\x31\x0c\x46\x61\xad\x0a\x3e\x58\x57\xa4\x5f\x0c\xa5\x7a\xb4\xcf\x2e\x47\x6a\x41\xe4\x7b\x22\xb1\xf4\x48\x4a\xce\xbb\x46\x2a\xef\x5d\xcd\xe4\x76\xd8\xdb\x54\xca\x44\xfa\x92\x67\x1e\xe5\x1e\xa4\x17\xfe\x0d\xee\x55\x70\xdd\xea\xeb\x28\xf0\xfd\x90\x4f\x64\xd5\xa6\xa2\x7c\xda\x44\xe8\xec\x7b\xa7\x78\xfc\x22\x21\xc8\xf9\x5b\xdc\xbf\x24\x80\xf1\x45\x7f\x1f\xbf\xfc\x13\xca\xea\x2b\xdf\x89\xcb\xf8\xc6\x32\xbe\xb1\x8c\x6f\x2c\xe3\x1b\x57\x6f\x21\x65\x7c\x63\x19\xdf\x98\x84\x2c\xe3\x1b\x77\xd1\xc2\x65\x7c\xe3\x41\xb8\xa6\x8c\x6f\x2c\xe3\x1b\xf7\x13\xb7\x32\xbe\xb1\x8c\x6f\x2c\xe3\x1b\xcb\xf8\xc6\xd7\xc4\x37\x46\x97\x0f\x65\x80\x63\x19\xe0\x58\x06\x38\x96\x01\x8e\x65\x80\x63\x19\xe0\xb8\x3b\xc9\xcb\x00\xc7\xb7\x5b\x50\x19\xe0\x58\x06\x38\x96\x01\x8e\x65\x80\x63\x19\xe0\xb8\x9e\xb7\x0c\x70\x2c\x03\x1c\xcb\x00\xc7\x32\xc0\xf1\x35\xe1\x88\xd2\x45\x3c\xe8\x90\x9c\x60\x6f\x33\xec\x87\x0a\x9c\x4c\xd6\x9d\x7d\x41\xe4\x64\x7e\xe1\xda\x97\x84\x4e\x3e\x3f\xff\x7f\x00\x00\x00\xff\xff\x9b\xf2\x91\x59\xd6\xba\x00\x00")

func templatesSingleMasterCloudformationJsonBytes() ([]byte, error) {
	return bindataRead(
		_templatesSingleMasterCloudformationJson,
		"templates/single-master.cloudformation.json",
	)
}

func templatesSingleMasterCloudformationJson() (*asset, error) {
	bytes, err := templatesSingleMasterCloudformationJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/single-master.cloudformation.json", size: 47830, mode: os.FileMode(420), modTime: time.Unix(1479233377, 0)}
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
	"templates/single-master.cloudformation.json": templatesSingleMasterCloudformationJson,
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
		"single-master.cloudformation.json": &bintree{templatesSingleMasterCloudformationJson, map[string]*bintree{}},
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
