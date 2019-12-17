// Package crds Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// installer.stash.appscode.com_stashoperators.yaml
package crds

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _installerStashAppscodeCom_stashoperatorsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\xff\x8f\xdc\xb6\x72\xff\xfd\xfe\x8a\xc1\xb5\x80\xef\x8a\xdd\xb5\xdd\x14\x0f\xed\x25\x48\x71\xb5\x9d\x07\x23\x8e\x63\xe4\xdc\xa4\x68\xf0\x0a\x70\xa5\xd9\x5d\xbe\x93\x48\x99\xa4\xf6\xbc\x29\xfa\xbf\x17\xfc\xa2\xef\x22\x45\xad\xcf\x71\xd3\x27\xfe\x62\x9f\xbe\x8c\x86\xc3\xe1\x70\xf8\x99\x8f\xb4\xa4\xa0\x3f\xa3\x90\x94\xb3\x1b\x20\x05\xc5\x8f\x0a\x99\xfe\x4b\x6e\xee\xff\x59\x6e\x28\x7f\x7a\x7c\xbe\x45\x45\x9e\x5f\xdc\x53\x96\xde\xc0\x8b\x52\x2a\x9e\xff\x84\x92\x97\x22\xc1\x97\xb8\xa3\x8c\x2a\xca\xd9\x45\x8e\x8a\xa4\x44\x91\x9b\x0b\x80\x44\x20\xd1\x07\xdf\xd3\x1c\xa5\x22\x79\x71\x03\xac\xcc\xb2\x0b\x80\x8c\x6c\x31\x93\xfa\x1a\x00\x52\x14\x37\x20\x15\x91\x87\x0b\x00\x46\x72\x74\x7f\xf1\x02\x05\x51\x5c\xc8\x0d\x65\x52\x91\x2c\x43\xb1\x31\x27\x36\xa4\x28\x64\xc2\x53\xdc\x24\x3c\xbf\x90\x05\x26\x5a\xd0\x5e\xf0\xb2\xb8\x81\xe0\xb5\x56\xbe\x7b\x6e\x42\x14\xee\xb9\xa0\xd5\xdf\xeb\x5a\x09\xfd\xff\xea\x3e\xf3\xa7\xed\xf3\x9d\x3e\xfd\xa3\xd3\xca\x1c\xcf\xa8\x54\xdf\x0f\xcf\xbd\xa1\x52\x99\xf3\x45\x56\x0a\x92\xf5\xfb\x63\x4e\x49\xca\xf6\x65\x46\x44\xef\xe4\x05\x40\x21\x50\xa2\x38\xe2\xbf\xb3\x7b\xc6\x1f\xd8\x77\x14\xb3\x54\xde\xc0\x8e\x64\x52\x6b\x23\x13\x5e\xe0\x0d\xbc\xd5\x1d\x29\x48\x82\xe9\x05\xc0\x91\x64\x34\x35\xa6\xb6\x5d\xe1\x05\xb2\xdb\x77\xaf\x7f\xfe\xea\x2e\x39\x60\x4e\xec\x41\x2d\x59\x3f\x46\xd5\x3d\xb6\xd6\xaf\xc7\xbd\x3e\x06\x90\xa2\x4c\x04\x2d\x8c\x44\x78\xa2\x45\xd9\x6b\x20\xd5\x23\x8d\x12\xd4\x01\xe1\x68\x8f\x61\x0a\xd2\x3c\x06\xf8\x0e\xd4\x81\x4a\x10\x68\xfa\xc0\x94\x51\xa9\x25\x16\xf4\x25\x84\x01\xdf\xfe\x15\x13\xb5\x81\x3b\xdd\x4f\x21\x41\x1e\x78\x99\xa5\x90\x70\x76\x44\xa1\x40\x60\xc2\xf7\x8c\xfe\x56\x4b\x96\xa0\xb8\x79\x64\x46\x14\x3a\xdb\x56\x8d\x32\x85\x82\x91\x4c\x1b\xa1\xc4\x15\x10\x96\x42\x4e\x4e\x20\x50\x3f\x03\x4a\xd6\x92\x66\x2e\x91\x1b\xf8\x81\x0b\x04\xca\x76\xfc\x06\x0e\x4a\x15\xf2\xe6\xe9\xd3\x3d\x55\x95\xa7\x27\x3c\xcf\x4b\x46\xd5\xe9\x69\xc2\x99\x12\x74\x5b\xea\x51\x7b\x9a\xe2\x11\xb3\xa7\x92\xee\xd7\x44\x24\x07\xaa\x30\x51\xa5\xc0\xa7\xa4\xa0\x6b\xa3\x38\x53\x66\xba\xe4\xe9\xdf\x09\x37\x2d\xe4\x93\x96\xa6\xea\x54\x18\xcf\x16\x94\xed\xeb\xc3\xc6\xb1\xbc\x76\xd7\xae\x05\x54\x02\x71\xb7\x59\xfd\x1b\xf3\xea\x43\xda\x2a\x3f\xbd\xba\x7b\x0f\xd5\x43\xcd\x10\x74\x6d\x6e\xac\xdd\xdc\x26\x1b\xc3\x6b\x43\x51\xb6\x43\x61\x07\x6e\x27\x78\x6e\x24\x22\x4b\x0b\x4e\x99\x32\x7f\x24\x19\x45\xd6\x35\xba\x2c\xb7\x39\x55\x7a\xa4\x3f\x94\x28\x95\x1e\x9f\x0d\xbc\x20\x8c\x71\x05\x5b\x84\xb2\x48\x89\xc2\x74\x03\xaf\x19\xbc\x20\x39\x66\x2f\x88\xc4\xcf\x6e\x76\x6d\x61\xb9\xd6\x26\x9d\x36\x7c\x3b\x4c\x55\x6d\x6c\x7a\xe8\x66\x62\x52\xe7\x08\x40\x4e\x3e\xbe\x41\xb6\x57\x87\x1b\xf8\xd3\x57\xbd\x73\x05\x51\xda\x25\x6f\xe0\xbf\x7e\x25\xeb\xdf\xfe\x72\xf5\xeb\x9a\xac\x7f\x7b\xb6\xfe\x97\xbf\xfc\xc3\xaf\xee\x3f\xd7\xff\xfa\xf7\xbd\x7b\x46\x95\xac\x0e\xdb\x01\xac\x0f\x57\x11\x6f\xd4\x69\x3a\x81\xe8\xae\xc0\x44\x3b\x90\x1e\x45\x7d\x17\xec\xb8\x00\x81\x29\x95\xd5\xe4\x8d\xe8\x3e\x49\x53\x13\xda\x49\xf6\x8e\xa7\x77\x98\x94\x82\xaa\xd3\x3b\x9e\xd1\x64\x70\x29\x00\x55\x98\x0f\x0e\x7a\xbb\xd7\x9c\x22\x42\x90\x53\xf7\xb1\x3b\xb3\xa6\x9c\xfa\xc2\x3a\xbd\x7d\xbd\x33\xfd\xa2\x3b\x8a\xe9\xca\x74\xb3\xe0\xe9\x13\x69\xc2\x46\x5a\x66\x7a\x82\x24\x9c\x49\x25\x08\x65\x4a\xf6\xc7\xc9\xd3\x61\xdd\x18\x4f\xf1\xd6\xa3\xc1\x40\x8b\x97\xe6\x8f\x2d\x4a\x73\x5b\xad\x79\x5b\x0b\x51\x66\x28\x8d\xf9\x9d\x92\x9b\x11\xa1\x21\x85\xec\x79\xdc\xa1\x10\x98\xbe\x2c\xb5\x21\xef\x6a\xf1\xaf\xf7\x8c\xd7\x87\x5f\x7d\xc4\xa4\x54\xbd\x80\xee\xd5\xfd\xbd\x76\x0d\x2b\x08\x05\x3c\xd0\x2c\x73\x8f\xd1\x21\xb7\x3a\xa1\x15\x36\x31\x58\xf7\xaf\x6f\xc6\xa6\xa9\x03\x51\x20\x89\xa2\x72\x77\x32\xfd\xac\x2d\x81\x1f\x75\xec\x31\xf9\x44\x33\x60\xb0\x3d\xb9\xb0\xa3\x97\xb8\x95\x57\xec\xb6\x54\x40\x95\x89\x55\xc9\x81\x73\x89\x40\xac\xa1\xcd\xf3\x8e\x94\x9b\x55\x01\x38\x43\xe0\x02\x72\x1d\x64\xcc\x4a\x84\x5e\x89\x2d\x75\x36\xc6\x02\x8d\x38\x2a\x21\xe7\x52\x35\xb6\xae\xe6\x8f\x16\xff\x40\xd5\x21\xd0\x7b\x84\xbd\xce\x78\x50\x2a\x90\x65\xae\x95\x78\x40\xba\x3f\x28\xb9\x02\xba\xc1\x8d\x19\x7e\x24\xc9\xa1\xf5\xb8\x1c\x71\xe0\x97\x4d\x23\x59\xe6\xba\xd2\xf1\x25\xfc\x50\x52\x81\xb9\x0e\xe5\x70\x55\xc7\x7d\x17\x8b\x57\xd5\xf9\x81\x97\xf8\x1f\x33\x32\x4c\x2b\x40\x95\x6c\xae\x57\x90\xf0\xbc\x28\x95\xb6\xb9\xee\xd3\xf6\xa4\xa7\xb8\x20\x6e\xed\x11\xbc\xdc\x87\x2d\x82\x99\x53\xb4\x4a\x0e\xcc\x60\x9b\x55\x5a\x07\x16\xb6\x87\x4b\x6b\xa4\xcb\x6a\x8d\x97\x65\xee\x95\x48\xad\x31\x8c\xfd\x72\xa2\x92\x83\x4b\x45\x12\x2e\x04\xca\x82\x33\x23\xd1\x9c\x79\xd5\xf4\xe5\xeb\xa0\x33\x68\x61\x57\xf2\xda\x0c\xae\x11\x76\xa0\xfb\x43\x35\x86\x44\xa0\x39\xd6\xf5\x89\xb1\xc9\x0b\xfe\xe8\x57\xb5\xce\xc4\xbb\x65\x80\x79\xa1\x4e\x2d\x4f\x6b\x8d\xb1\x42\x91\xd7\x3d\x24\x26\x67\xf6\x35\xbb\x3a\x48\xab\x3f\xcd\x0b\x1d\x98\x95\xf3\x3c\x78\x06\x57\xc6\xf5\xa8\x7a\x22\xcd\xb4\x59\xf3\xe2\x7a\x03\xb7\x55\x22\xee\x6b\xd3\x4a\x31\x5e\x3f\xd9\x3d\x42\x2b\x2a\x79\x40\x68\xfd\x7c\xef\x35\x53\x11\xb0\xad\x1c\xb2\x64\xb0\x2c\x77\x5b\xd7\xde\xd6\x6b\x24\x66\x98\x28\x1d\x87\x51\xe4\x2b\x20\x52\xf2\x84\xea\x64\xa5\x1e\xff\xa0\x48\xe8\xb9\x9a\x35\xb3\xbf\x43\xf1\x9d\x02\x93\x55\x74\x1d\x77\xea\xfa\x41\x17\xf5\x86\x44\xcf\xb4\x6e\x57\xdb\x01\x63\x52\x22\xe8\x39\xae\xef\x7f\x22\xdd\x36\x2d\xdc\x3b\x98\xf6\x7b\xaf\xba\x5e\x35\x5d\xd6\xeb\xce\x44\x08\x76\x8b\x8f\xce\x1c\x09\x65\xd2\x65\xfa\x2b\x20\x70\x8f\x27\xbb\x29\xd0\xfb\x0e\x97\x17\x99\x8b\xa3\xa4\x0a\xb4\x8b\x8b\x8e\x01\xf7\x78\x32\x82\xdc\x2e\x22\xe2\xfe\xf8\x91\xb7\xed\x1e\x47\x93\x8d\xb1\x36\x58\xc4\xcd\x58\x19\x1d\x8d\x25\x4c\x24\x9d\x63\x3f\xb0\x5b\xf1\x8c\xa2\xc9\xe6\x23\xef\x09\x24\x76\xe3\xad\x1a\x82\xb3\xfa\xf9\x53\xbd\x85\xb1\x03\xfb\x44\xda\x01\xd2\x73\xe5\x40\x8b\xe8\x7e\x2a\x6e\xbc\xcb\x4c\x95\x6a\x4f\xf8\xb3\xde\x43\xd7\xea\x49\x13\xf9\x5f\x33\x7f\x56\xd2\x6f\x6f\xb9\x7a\xcd\x56\xf0\xea\x23\x95\x7a\xc1\x7f\xc9\x51\xbe\xe5\xca\xfc\xb9\x81\x3f\x2b\xeb\x83\x6f\x26\x42\x45\x4b\xc5\xb9\x86\xb5\xfd\x38\xcb\xac\xb7\xcc\xe6\xdf\xda\x1c\xed\x9d\xa6\xdc\xe8\x04\x7b\x3a\x24\x36\xad\x9e\x60\x54\xea\xbd\x1f\x17\x95\x59\x0c\x5e\x60\x64\x8e\xa4\xfa\xa1\x96\x97\xd2\x6c\x29\x19\x67\x6b\xb3\x5e\x56\x3a\x75\x9e\x65\xad\x1e\xaf\xa6\xe8\x8c\xcf\x50\xbd\xea\xb1\xd1\x12\xfd\xaa\xfd\x59\xe9\xc7\xbd\xe9\x3c\x24\x7e\x42\x36\xca\x1c\xc8\xd1\x24\x61\x94\xed\xb3\x3a\xad\x5a\xc1\xc3\x81\x26\x07\x93\xb7\x47\x0b\xdd\xa2\x05\x4d\x0a\x81\x7a\xdd\x23\x52\x87\x46\x7d\x64\x8f\x42\xa7\xc3\xb4\x32\x02\x8d\x57\x54\x60\x91\x91\x04\x53\x48\x4d\xd2\x69\x21\x0b\xa2\x70\x4f\x13\xc8\x51\xec\x51\xef\x8a\x93\x43\xac\xf7\x47\x2f\x28\xb6\xcd\x9e\x2c\xfe\x6d\xe7\x78\xab\x52\xea\x18\x95\xd6\x3a\x32\x45\x5d\xc7\xdb\x58\x62\x8c\xba\x3d\x20\x20\x7c\x71\x4c\xdf\x4c\xc2\xe1\x10\xc6\x2f\x9c\x6b\x98\x7d\xc1\x92\x6b\x2c\xb9\x86\xb7\x2d\xb9\x46\xd5\x96\x5c\x63\xc9\x35\x96\x5c\x63\xc9\x35\xfe\x40\xb9\x46\xa4\x50\x8b\xa7\xcc\x80\x75\x7e\xb1\x38\x57\x1f\xc7\x31\x89\x4d\x55\x1f\xeb\x40\x36\x13\x3d\xd2\x69\xc2\x9d\x5b\xcb\xde\x1b\x88\x88\x32\x23\x44\x10\xb6\x47\x78\xbe\x7e\xfe\xec\x59\xd8\xb3\x76\x5c\xe4\x44\xdd\x68\x2f\xff\xea\x1f\x23\x6c\xe2\x66\x83\xf7\xca\x69\x7f\x58\xb7\x10\xb1\xc0\x45\xd6\xb6\x7e\xb4\x76\x7a\x84\xa6\x06\xdb\x87\x3c\x7f\x42\x7d\xc2\x45\xb9\x1a\xa2\xee\x80\xdf\x83\x52\x82\xb7\x73\x0e\x75\x16\x3a\xb8\x2b\xc8\x51\x01\x51\x1d\x68\x93\xe6\x58\x17\x90\x6c\x19\xc4\xd6\x32\xbd\x12\xab\xda\x48\x0a\x9c\x39\xe4\x5a\xfb\xce\x26\x52\x63\x7f\xb5\xa3\x5d\x14\x81\x04\x89\x44\x9d\x43\x6c\xb1\xd6\x9a\xe7\x5a\x4b\xca\x54\x15\x00\xb5\xca\x58\x59\xd5\x2b\xf8\x0a\x37\xfb\x0d\xa4\xa5\x11\x47\x98\x2b\xd2\x5e\xdb\x5e\xcb\x93\x54\x98\x9b\x1a\x0b\x17\xe6\x1f\xdd\x7d\x25\x4e\xa0\xfc\x88\x2e\x1e\x91\xa9\x92\x64\xd9\x09\xf0\x48\x13\x55\xdb\xcf\xd4\x91\xa9\xb2\xf5\x30\xdf\x6c\x89\x49\x58\xfb\xb3\x31\x18\xa7\x7b\xe9\x9b\x75\xc5\x8d\x77\xa7\xa2\xb4\x3c\x53\xfe\x09\x4f\x52\x7d\x99\xf1\x9c\x1f\x7f\xf2\x23\xff\x10\xb7\x90\xf4\xf7\x24\x65\x96\x69\x7b\xdb\x42\xc0\x50\xbd\x0a\x6c\x9f\x8c\x59\x15\x14\x6f\xab\x59\x1d\x8f\xb3\xf5\x23\x5b\xc9\xb8\x7d\xfb\x52\x5b\x64\xaa\xcb\x00\xef\x79\xc1\x33\xbe\x3f\xb5\x6d\x6f\x66\xbf\x29\x30\x38\xc9\x04\x64\xb9\x75\x99\xed\x74\xe2\xf6\xb6\x37\x94\x0b\x66\xbe\xec\x63\xc7\xda\xb2\x8f\x1d\xb4\x65\x1f\x1b\xa9\xe2\xb2\x8f\x35\x6d\xd9\xc7\x2e\xfb\xd8\xc9\xb6\xec\x63\x47\x2e\x5e\x30\xf3\x25\xd7\x08\xb4\x25\xd7\x18\xb4\x25\xd7\x58\x72\x8d\x25\xd7\x58\x72\x8d\x60\x5b\x72\x8d\x91\x8b\x1f\x0d\x33\x9f\x16\x37\x65\x9e\xf5\x10\x68\x0b\x22\xc0\x5e\x95\x82\xa7\x0b\x9e\x9e\x41\xa9\x2f\x78\x1a\x60\xd4\x5b\x50\x33\xe1\xeb\x8c\x27\x44\x8d\xc7\x03\x03\xa7\x6a\x31\x0e\xc9\x97\x24\xb7\x58\xed\x0a\x7e\xe3\x0c\x2d\xd3\x59\x4f\x33\x83\xac\x72\x75\x40\xa1\x2f\xbf\x92\xd7\xa3\x4c\xd5\x85\xa5\x3f\xda\x16\x96\xfe\xc2\xd2\x77\xad\xcd\xd2\x3f\x10\x69\xfd\xd2\x2e\x84\x7e\xd2\x7e\x2b\x3a\xe8\x00\xf4\x75\x50\xdf\x2f\xc4\xd9\xd7\x4e\xe8\x9c\xc5\xbc\xc9\xd8\x0c\xbc\xed\x57\xea\xca\x91\x98\xbe\xeb\xf6\x26\x10\xbd\xed\x1e\xce\x28\x4d\xd2\x14\x53\x28\x50\xac\xad\xeb\x71\xd8\x51\x96\x8e\xf4\xa5\xea\xbf\x57\x6c\x24\x8f\xbe\xab\xe4\x8c\xd2\x45\xbb\xba\xd2\x09\xd0\x7d\x56\xfd\xc4\x5a\x58\x8f\xdf\xe7\x64\xd5\x9b\x9d\x57\xb5\xb8\xcd\xdf\xb2\x9b\x7d\xdb\x87\x12\xc5\x09\xf8\x11\x45\xb3\x33\xa9\x5f\xf3\x8c\xd9\x84\x98\xb5\x87\x4a\x48\x88\xb4\x81\x7a\x3a\xd5\x9a\xb7\x3b\x9d\x5f\x07\x19\x74\xb6\x2f\xc2\xee\xf2\x2b\xcc\xc2\x18\x22\x32\x7b\x1b\x85\x36\x46\x8a\x53\x44\xc4\xa6\xf0\xb6\x74\x15\x75\xf1\xac\xe4\x74\x74\xb4\x3d\x90\x47\xfc\xb6\xa0\x55\xc6\x9b\x80\x3d\xe2\x65\xf6\xe0\x91\x4f\x84\x3e\xe0\x0c\xf8\x03\xe6\x41\x20\xd0\x37\xaf\xd6\xd2\xad\xd3\x43\x34\x64\x86\xd0\x96\x7f\xcd\x47\x44\xe0\xbc\xfd\xc8\x7c\x64\x04\xfa\xdd\xaf\x87\x4f\x0c\x60\x92\x59\x9d\x6f\x43\x2a\x7e\xa8\x64\x96\xc8\x01\xac\xd2\x85\x4b\x8c\x6f\x75\x10\x93\xcf\x6d\xec\x79\x68\x09\xf4\x4d\xed\xb0\x02\x6a\xb6\xce\x3d\xec\x64\x96\x61\xba\x38\x8b\x17\x3f\x99\x25\xd3\x07\x66\x74\x31\x94\xd9\x22\x87\x78\xcb\x00\x47\x79\x1c\x35\x9d\x8a\x0d\x10\x31\x4b\xac\xfd\x3e\xc4\x63\x82\x11\x30\x1f\x90\x80\x73\xfd\x72\x2e\x30\x01\x33\xc1\x09\x98\x01\x50\xc0\x5c\x90\x02\xe6\x02\x15\x30\xbb\xbf\x26\x85\x78\xd3\xfa\xba\xcb\x54\x6b\x7d\x5d\x60\xf6\x6a\x34\x7b\x04\x87\xd9\x8e\x55\xd5\x26\x3a\x39\x29\x74\x94\xf8\x6f\xbd\x34\x1b\xc7\xff\x9f\xd8\x75\x94\x50\x21\x75\x2a\xec\xc0\xbf\x96\x84\x0a\x73\x68\x3d\x2c\x52\xa8\xd6\x86\x4a\xd0\xbe\x73\x24\x99\x4e\x40\x2c\x6d\xcb\x6d\xd5\xb4\xa6\xfd\x7c\x2d\x76\x7e\x3f\x1c\xf4\xf6\x5c\x2f\xbe\x76\x9b\x47\x25\x5c\xde\xe3\xe9\x72\x35\x88\x23\x97\xaf\xd9\x65\xac\x54\xe2\xb6\x2a\x9d\x98\x51\x67\x3e\x9c\x65\x27\xb8\x34\xe7\x2e\x63\x27\xf6\x58\xba\x38\x27\x11\x3c\x03\x94\x8b\xba\x98\x55\xdf\xde\x99\x5b\x00\x6c\x6e\xac\xf1\x95\x6a\x63\xdc\x9c\x8a\x41\x1b\xab\x0c\xea\x6e\x98\x07\xc1\x55\xfd\xda\xf8\x5e\x5b\x5e\x5d\xfb\xb7\xd2\xad\x2e\x75\x98\x68\x26\xe5\xcf\x91\x30\x09\x97\x15\x7a\xf6\x44\x36\x3a\x5e\x3e\x5e\xc5\x71\xd6\x1c\x8e\x8f\x45\xca\x11\xd8\xbe\x8f\x49\x57\x7b\x7b\x7c\x87\x16\xba\x8f\x12\x6d\xb1\x81\x17\x53\xb8\xaa\x76\xba\xfe\xbd\x77\xd3\xb8\x30\x2c\xca\xce\xed\x4c\xd1\x75\x2d\xa3\xd9\xff\xea\x1d\x61\x6c\x78\xad\x68\xcd\x5d\x0f\xa8\xc0\xcd\x1a\xb7\x6b\x3c\x2a\x66\x06\x3f\x1c\x50\x74\x7a\x4a\xa5\xfb\xd8\x93\xa9\x40\x88\x92\x31\xfd\x5c\xce\x1c\xac\x17\x25\x52\x87\x19\xfb\xcd\x22\x07\x93\xd8\xb4\xdf\xf4\xda\xe4\xfe\xcd\x28\x45\x52\x1d\xa1\x02\x30\xcd\x87\xa4\x1c\x67\x92\x33\x37\x89\xf4\x91\x0a\x89\x33\x76\xc1\x34\xd6\xb2\xb4\xee\xe3\x06\x5e\x99\x49\xd0\x56\x8e\x4a\x33\x92\x24\xcb\xf8\x43\x4c\xf4\x89\xf6\xea\xb8\xdc\x60\xdd\x56\xe6\x31\x4a\x06\xb3\x69\xf6\x0f\x8f\x4c\xb3\xef\x41\x4f\x7f\x10\x96\x7d\x24\xa8\xb7\x50\xed\x17\xaa\x7d\x8b\x6a\x6f\x6e\xb2\x91\x6f\x9a\x73\xef\xf7\x19\xc3\xc5\x8f\xe5\xdc\xc3\x2f\x07\x34\x33\x2a\x00\xb0\xe9\x21\xca\xcb\x4c\xd1\xa2\x29\x58\x4b\xab\x5a\x66\xb7\x8f\x96\xa8\x24\x7b\xe8\x6c\xe8\x8d\x00\x92\x1c\xfa\xd3\xc4\x3c\xc7\x14\xb4\xa5\x89\xc8\xae\xcc\x42\xb2\xcc\x71\xeb\xf5\xbe\xd2\x3f\x46\xe8\x6a\x55\xf4\x71\x20\xfc\x97\xee\x03\x86\x35\x68\x62\x8a\x13\x57\x7a\xb1\xcc\xb4\x3b\xe8\x25\xab\x8a\x6a\xa1\x9a\xeb\x60\xfd\xb5\xa8\xcc\x11\xab\x02\xc9\x9e\x1e\x91\x35\x8b\xf0\x95\xbc\xbe\x9e\xa2\x35\xa9\xc8\xd4\x63\x98\x58\x04\x84\x8e\xa5\x1c\xab\xc8\xe5\x3e\x20\xb6\x4e\x04\x22\x96\xf9\x6f\x5a\xab\xd7\xb7\x01\x99\x4d\x71\xc8\xbb\xc0\x1b\xf3\xd4\x4b\x7c\x3d\x80\x01\xa1\x74\xba\x37\x71\x38\xe8\x8c\x32\xc2\x19\x25\x04\xa0\xfe\x70\x62\xdb\x9c\xf2\xc1\xef\xf6\xfa\x44\x44\xc9\x60\x0e\xcd\x6d\xba\x5c\x10\xbb\xff\x3b\x97\xf2\x18\x2c\x00\x2c\x9c\xc7\x60\x8b\x07\xfb\xff\xff\x51\x1f\x03\xe0\xfe\xff\x51\x0e\xe4\xd9\xa0\xfe\xef\x49\x7d\x0c\x01\xf9\x33\xab\x5d\x30\x05\xe2\x7f\x22\x01\x70\x8a\x04\x19\x2d\xd3\x03\xde\x8f\x03\xf2\xd1\x52\xc7\x80\xfb\x51\x30\x3e\x5a\xe2\xc2\x20\x9c\xbc\xee\x4b\x33\x08\x67\x02\xf2\xe7\x82\xf1\xb3\x46\x67\x2e\x08\xef\xe0\xf5\x08\x35\x22\x01\xf8\x21\xb4\x1e\xd3\xc5\x49\xf0\xbd\x0f\xab\xc7\x81\x4e\x21\xe0\x7d\x14\x52\x8f\x10\x3b\x0e\xba\x7f\x52\x3a\x15\xed\x9d\x91\x17\xc6\x42\xe8\xf3\xe1\xf3\x08\x2e\xc1\x0c\xe8\xbc\x02\xc6\x27\x24\x3e\x06\x6c\x1e\x15\x11\xa3\x67\x5a\x5c\x84\x88\x86\xc9\x3f\x07\x44\x3e\x13\x1e\x8f\xd9\x96\xc3\xe8\xd6\x3c\x04\x8d\xdb\x9d\xf0\x84\xc8\x78\x58\xbc\xbd\x1b\x9e\xea\x7e\x2c\x24\xde\xde\x0f\x4f\x55\xa6\xa2\xe0\xf0\x21\xd8\x1d\x5f\x4d\x99\x05\x85\x47\x79\x6b\x0c\xf2\x1a\x03\x7f\x7f\x32\xa8\x3a\x49\x5e\x67\x8a\x9e\x4b\x60\x6f\xfb\xb5\x87\xc5\x3e\xaa\x33\x39\x72\x9a\x42\x51\x2a\x47\xe5\x9d\xcd\x64\x1f\x95\xfa\x37\xc5\x6e\xef\x98\x3e\x48\x71\x0f\x43\xda\xab\x33\x28\xee\x5e\x89\x6e\x5a\x9e\x41\x71\xf7\x8b\x74\xd4\xf7\xb3\x28\xee\x5e\xa9\x86\xfa\x7e\x1e\xc5\x7d\x72\xc6\xf7\x5d\xc8\x3f\x56\x15\xcf\xdd\x2b\x72\x9a\xff\x1e\xe0\xb9\xfb\x11\xf2\x20\xff\x3d\xc0\x73\xf7\x9b\x33\x9a\xff\x3e\xe0\xb9\x07\x5c\x7e\xe1\xbf\xf7\xda\xc2\x7f\x6f\xb5\x85\xff\x1e\xd9\xd9\x85\xff\xbe\xf0\xdf\xa7\xda\xc2\x7f\x5f\xf8\xef\x0b\xff\x7d\xe1\xbf\x2f\xfc\xf7\x85\xff\x3e\xd2\x16\xfe\xfb\xc2\x7f\x5f\xf8\xef\xad\xb6\xf0\xdf\x27\xba\xb2\xf0\xdf\x17\xfe\xfb\xc2\x7f\x5f\xf8\xef\x0b\xff\x7d\xe4\x92\x2f\xc2\x7f\xef\x80\xd0\x5e\x12\x7c\x00\x8e\x6d\xbe\x9f\x32\x93\x04\xef\x95\xb9\xc5\x69\x12\xbc\x57\x6d\xaf\x54\xcf\x37\x7e\xa2\x98\xf0\x7e\xe8\xb5\xcd\x90\x9f\xc5\x84\x0f\x80\xe6\x23\x5f\xa5\xff\xc4\xaf\xcf\x43\x8b\x21\x7f\x2e\x13\xde\xef\x02\x7c\x61\xc2\x2f\x4c\xf8\x85\x09\xbf\x30\xe1\x17\x26\xbc\x6d\x0b\x13\x7e\x61\xc2\x2f\x4c\xf8\x85\x09\xbf\x30\xe1\x07\x6d\x61\xc2\x8f\xaa\xbb\x30\xe1\x17\x26\xfc\xc2\x84\x6f\xda\xc2\x84\xef\xb6\x85\x09\xbf\x30\xe1\x03\x6d\x61\xc2\x7f\x1e\x26\xbc\xf7\x14\x61\x8c\x2b\x9b\xdc\xf7\xf5\x8f\x5b\x4c\x03\x26\xf2\x3f\xb4\xa0\x12\xc5\x11\x07\x3b\x95\xd0\xde\x6e\x7b\x2a\x88\x94\x66\x07\x61\x18\xc2\xbf\xe0\xf6\xc0\xf9\xfd\x7f\x08\x32\x3a\xf5\xed\xc3\xb7\x9c\x67\x48\x86\xc8\x44\x42\xfc\xf7\x78\x86\x1b\x19\xd9\x66\xf8\x43\xa9\xda\x4f\x9f\xff\x64\x2b\x66\xd0\x8d\xf9\x82\xf6\x82\x97\xc5\x3b\x41\xb9\xa0\xea\xf4\x03\x65\x34\x2f\x47\xb9\xb0\x53\x25\x87\x70\xa1\xe1\x80\x24\x53\x87\xe4\x80\xc9\xa8\x8a\x53\x9b\x71\xdb\x5b\xef\xd4\x08\xf7\x70\x62\x4a\x94\x12\xbf\x2f\xb7\x58\x3b\xd3\x77\x1f\x52\xf6\x1d\x17\xb7\xf7\xa3\xba\x84\x1f\x75\x44\xa1\x93\xa0\xca\x9c\x8f\x6d\x47\x5f\x88\x58\x43\xd2\xdf\xae\xac\xc7\x1d\xcd\x73\xd5\xc0\x8f\x06\xd7\x8d\xb9\xc9\xe0\xa2\xd6\x28\x0f\xce\xf9\xcd\x3c\xb8\xb4\x67\xc5\xd8\x70\x90\xe8\x21\x99\x17\x0c\x04\xee\xa9\x54\x22\x30\xf3\x3d\xb3\x58\x60\xc1\x25\x55\xfc\x8c\x5b\x15\xd9\xcf\xbc\xc7\x3f\xec\x95\xfe\x23\x27\x2a\xfd\x06\xa7\x14\x89\x0e\xb0\x89\xa0\x8a\x26\x24\xbb\x4d\xd3\x61\x15\xcd\x3f\x13\xac\x4f\xdd\x32\x92\x9d\x14\x4d\x06\x66\xf7\xdf\x48\x73\xb2\xc7\x77\x65\x96\xbd\xe3\x19\x4d\x06\xa6\xf5\x9a\xa8\xbe\xef\x0e\x13\x81\x6a\xf0\x44\x4f\x4a\x39\xb9\xe6\x0c\x57\xc8\x8c\xef\xdf\xe0\x11\xb3\xbe\xb0\xd0\xac\xf6\xcf\xe8\x9c\x33\x3d\x46\x94\x0d\x3c\x22\xe4\xb4\x64\x8f\x6c\xb4\x7a\x1d\x74\xbb\x2d\x49\xee\xcb\x62\x7e\x48\x0b\xa1\x81\xe1\x3b\x0b\xc1\x73\x54\x07\x2c\x47\x03\xe9\x54\xc8\xaf\x73\xd3\x70\xd0\xf7\x26\x55\xa1\x74\x6a\xdd\x48\x9f\xbb\x5a\xe8\xe0\x45\x13\xfc\xc1\x8e\xdc\x39\x1d\xcb\x82\x20\xc4\x3c\xe8\x21\x22\xb1\x9c\x48\x07\xc3\x66\xca\x7c\x9c\xbd\x80\x54\x7f\xb4\x32\x8e\x3b\x38\x6a\x1d\x73\x70\xd8\x03\x19\xad\x5b\x6e\x35\x38\xd5\x1d\x9c\xd8\x28\xd7\xfe\xf5\xa3\xdf\x2f\x79\xf5\xcd\xab\x65\xb9\xea\x9e\x9a\xb1\x5c\x15\x19\x51\x3a\x10\xcf\x31\x29\x2f\x90\xc9\x03\xdd\x05\x02\xea\x78\x7c\xf3\xf7\xb2\x16\x19\xad\x77\x29\x0f\x7b\xa2\xf0\x61\xb8\x05\x59\xbc\xa1\x7b\x6a\x86\x37\x08\x2c\x32\x9a\x90\x17\xbc\x1c\xae\x96\xe7\x2d\xd7\x2e\xbc\xdc\x26\xc9\x98\xcc\xd0\x50\x25\xe6\x85\xdc\xf9\x4b\xa8\x5e\xa6\x1e\xcd\xde\x56\x89\xc1\x61\xfd\x8c\x58\x9b\x2a\x9e\x99\x17\x6b\x47\xb6\xf9\x7d\x32\x55\x8d\x32\xd5\x44\xa7\x27\xb2\x7d\x7f\x1f\x31\xf1\xe4\x68\x83\x37\x50\x0b\x9e\x5a\x0a\xc0\xfb\x5a\x96\x41\xcc\x95\x22\xe6\x1d\x54\xc5\xab\xa7\xa0\xf4\x20\x45\x8a\x50\xa6\xdc\x6f\x7b\xb5\xde\xc8\x55\xc2\x30\x80\xbe\xa9\x41\xf3\x15\xee\x76\x98\xa8\x6f\xa1\x94\x15\xcb\x2f\x80\xbd\xd5\xb0\xf5\x37\xd5\xff\xbe\x1d\xa2\x42\xe1\xe4\xc0\x3e\x6f\x7c\xb9\xef\x18\xe2\x95\xb9\x10\x28\x4b\x69\x52\xd7\xc0\x6d\xb7\xac\x0c\x6d\x06\xa3\x6b\x18\xd8\xb2\x20\xad\xb9\xd0\x12\x91\x5a\x22\xa4\xfb\x8c\x68\x6b\x20\x1d\xd8\x15\x2e\x0e\x12\x81\xf0\x96\x3b\x4e\x1e\xae\xe0\x9d\x79\x6f\xb7\x39\x62\xd0\xfc\xb7\xdc\xb2\xf3\x3c\x0c\xad\x89\xbc\xc6\x5b\x8b\xef\x18\xe9\xfb\xa6\xf2\x6e\xfb\xd5\xa9\xbc\x37\xae\x58\x21\xdf\xbe\x04\x8a\x57\x9c\xd8\x71\x6b\xdd\xe3\xa9\x79\x45\xca\x55\xfb\x0d\x0e\xbe\x9a\x2a\x7b\x55\xe5\x52\x5b\xe8\xfc\xda\x31\x5b\x78\xbe\xa5\xcc\x2a\x66\x1f\x58\x0d\xa5\x79\x66\x55\x02\xf1\xbe\x3b\xa9\x2f\x32\x2a\x9d\x63\xd8\x70\xe9\xbf\x63\xdd\x1f\x23\x0b\xfd\x15\x0f\xcc\x28\xee\x83\x49\xc7\x0a\xfa\xad\xea\xfc\xab\x0f\x25\xc9\x36\xf0\x12\x77\xa4\xcc\x94\xb1\x88\x3b\x14\xfc\x69\xd9\xc1\x0b\x24\x0f\x34\x4b\x13\x22\x52\xf3\x09\x04\x3b\xc5\x41\x72\xeb\x13\x96\x52\x95\x10\x56\x87\x8f\x80\x81\xcd\xc8\xdb\xb7\xe0\xa1\x20\x42\xd1\xa4\xcc\x88\x00\x3d\x17\xf7\x5c\x9c\xce\xb2\x7d\xe3\x90\x77\x98\x70\xe6\xfb\x7d\xf1\x6e\x40\xec\xdf\xd3\x1e\x0d\x13\x76\x51\x50\xc7\x14\xa3\x39\x06\x6a\x42\xad\xe9\x70\xe5\x7e\xac\xcf\x79\x27\xdf\x55\x31\xa5\x9e\xb4\x2b\xfb\x75\x91\x07\x2a\x31\xf4\xe9\x8c\xfa\xfd\x1a\x6a\x69\xb9\xd7\xad\xc8\x5c\xcf\xca\x0d\xfc\xdb\x09\x52\x3b\xb2\x2b\xa0\xca\x01\xeb\x1e\x91\x12\xeb\x5f\xd5\xad\x26\x87\x95\xd8\x9a\xe6\x3b\x2e\xf0\x88\x02\xae\x52\x6e\x20\x7a\xc3\x0f\x1d\xfd\xdc\x89\x6e\xff\x89\x82\x1b\x27\x63\xb8\xb7\xc4\x45\x37\xc5\x0c\x2d\x77\xab\x97\x05\x24\xee\x97\x78\x9f\xc1\x95\x25\x9b\xd2\x3c\xc7\x94\x12\x85\x99\xb7\xba\xb3\x3d\xb5\xc8\xad\xe3\x0f\x6f\x65\x21\x7f\xfa\xa7\x80\xcb\xf8\xd9\xdb\x46\xd5\x08\x3f\xf9\xd9\xb0\x12\x3b\xc1\xd0\x12\x15\x7b\x91\xb0\x5e\x0e\xbd\xec\x22\xef\xcf\x3b\xb7\x78\x19\xad\x7a\x58\x15\x08\x2b\x77\xf1\x48\xfd\xab\xf6\x35\xa2\x33\x42\x33\x8f\xec\x1c\x39\x63\x16\x4d\x96\x2e\xfa\x88\xce\x58\xb6\xb4\x6e\xef\xfa\x78\x7a\x87\x49\x29\xa8\x3a\x19\x4c\xaa\xbb\x4c\xac\x9b\x12\x44\xe7\xa8\x43\x22\x3b\xc7\x7a\xa0\x58\xe7\x5c\x0f\xf7\x1a\x3f\xe7\xb0\xad\xce\xc9\x0a\x8e\xea\x1c\x6c\x50\xa5\xce\xe1\x91\xc5\x68\x5d\x6f\x99\xba\x07\x9b\xfd\x48\xe7\x78\x3b\xa3\xee\x9c\xe8\x26\xc6\x17\xc1\xf1\xe8\x1d\x72\x58\xef\x0d\x1c\x9f\x93\xac\x38\x90\xe7\xcd\x31\x13\x00\x6d\x7e\xda\x39\x6d\x13\x71\x4c\x6f\x40\x09\x57\x5d\x94\x8a\x0b\xb2\x47\x77\xe4\x7f\x03\x00\x00\xff\xff\x32\xa7\xa1\x56\xce\xae\x00\x00")

func installerStashAppscodeCom_stashoperatorsYamlBytes() ([]byte, error) {
	return bindataRead(
		_installerStashAppscodeCom_stashoperatorsYaml,
		"installer.stash.appscode.com_stashoperators.yaml",
	)
}

func installerStashAppscodeCom_stashoperatorsYaml() (*asset, error) {
	bytes, err := installerStashAppscodeCom_stashoperatorsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "installer.stash.appscode.com_stashoperators.yaml", size: 44750, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
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
	"installer.stash.appscode.com_stashoperators.yaml": installerStashAppscodeCom_stashoperatorsYaml,
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
	"installer.stash.appscode.com_stashoperators.yaml": {installerStashAppscodeCom_stashoperatorsYaml, map[string]*bintree{}},
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
