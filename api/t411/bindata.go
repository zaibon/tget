// Code generated by go-bindata.
// sources:
// ../../scripts/mapping.json
// DO NOT EDIT!

package t411

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

var _mappingJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x98\xcd\x6a\xdb\x4c\x18\x85\xf7\xb9\x0a\xa1\x55\x3e\x08\x64\x46\x9a\x1f\xe9\xdb\x75\xa3\x5d\x29\x69\x43\x36\xa5\x0b\xd5\x55\x5d\x53\x23\x17\x2b\x2a\x94\x90\xeb\x29\xb9\x0e\xdf\x58\xc7\xc5\x6d\x52\x38\x63\x3f\xd8\x1b\x13\x39\x67\xe6\xfd\x39\xef\x39\x33\x7a\xb8\x28\x8a\xa2\x9c\x86\x7e\xda\x8c\x53\xf9\x7f\xf1\x7e\xff\x77\xfa\x3c\x1c\xbe\xf7\xbf\x7e\x1d\x7e\xa4\x5f\xaa\xf6\xea\xc5\xb3\xef\xfd\x7a\x1e\xd2\xd3\xb6\x0d\x87\xa7\x8f\x57\x79\xa8\x44\x86\x16\x20\x35\xb4\x01\x50\xab\x37\x6d\x4e\x23\x9d\x44\x46\x7b\x1a\x19\x34\xb2\x06\x89\x36\xba\xba\xfe\x34\x54\xf7\x25\x92\xbe\x18\x5d\xdd\x08\xaa\xeb\x35\xb4\x02\x50\x9d\x6a\x03\x52\xad\x75\xc0\x2d\x09\x58\x33\x29\x12\x12\xd6\x7a\x57\x03\x76\xd5\x01\x47\x10\x70\xa5\x09\xdc\x00\x02\x5b\x1d\x70\x03\x02\xae\xa2\xce\x15\x50\xd8\x6a\xf6\x37\x00\xaa\x19\x11\x01\x23\xac\xae\x52\x04\x55\xd2\x0c\x8e\x84\xc1\x7a\xe4\x1a\x30\x72\x56\x8b\x4b\x03\xc4\xa5\xd2\x01\xb7\x0e\xec\xaa\xbb\xda\x00\xa8\xa6\x52\x24\x54\xd2\x7c\x68\x41\x81\x75\xb8\x11\x84\x5b\xe9\xfa\xb6\x7f\xeb\xbb\xff\xfe\xf0\xfb\x7f\xca\xe1\xdb\x6a\xda\x7c\x1a\x8e\xdb\x9e\x97\x0b\xda\xf4\x39\xd7\xf7\x6a\xc0\xcd\x5a\x4e\x84\x35\xc4\x0f\x6a\xc9\xce\x84\x05\xf4\x74\x52\xaf\x12\x16\x08\x56\x2d\x47\xd1\x1a\xa4\x58\xba\x52\x60\x57\x97\xe9\x8f\x01\xfd\xa9\x25\xcf\x52\xb6\x80\x68\x5e\x4e\x63\xa2\x05\xd1\x0f\xed\x0a\x0e\xa4\xeb\x75\x73\xad\x25\x76\xaf\x65\x36\x10\x99\xd5\x8a\xe7\xc1\xae\x19\xcf\x0e\xc4\x02\x35\xd4\x13\x5a\xc8\xd9\x4b\xad\x25\x27\x47\xad\x21\x1e\x30\x4a\x0b\x97\x03\x16\xe8\xf5\xf4\xa4\xf1\x39\x57\xa3\x6b\x90\xaa\xd7\x13\x60\x2d\x91\x5a\xed\x47\x01\x0c\x80\xd3\xf2\x66\x0d\xa0\x62\xe6\x34\x16\x80\x23\xe9\xbe\xba\xf3\xbd\x2c\x90\xc6\xea\x4d\x89\x7d\x68\x36\x05\xa2\x4e\x5a\xc7\xad\x21\xa3\x93\x51\x36\x03\x1a\xeb\x65\x9d\x12\x9f\x40\xa1\xb4\xc4\x38\x22\x31\xda\x03\x4c\x4b\x3c\x40\x47\x6c\xc8\x89\x37\x73\x25\x0e\x20\x64\x2d\xc5\x0e\xb4\xd6\x65\x6c\xb6\x25\x42\x21\x67\x27\x61\xc9\x55\x46\x9f\x66\x1c\x11\x54\x3d\xb2\x9e\x8c\x6c\x86\x50\x06\xb4\xa7\xce\x78\x00\x79\x7b\xe0\x32\x3e\x6b\xc8\x19\x5f\x8b\x85\x27\x73\xab\x1b\x94\xf6\x05\xdb\x6a\xb9\xf0\xe4\xea\xa6\xcf\x5f\x8e\x9c\xdc\xb4\xbe\xf9\xf3\x5f\x04\x78\x60\x02\x5e\xb7\x36\x31\x19\x4c\xad\x4e\xd6\x83\x64\x33\xaf\xa3\x3c\xd9\x55\x9b\x4f\x00\x12\xe5\x32\x26\x6d\x00\xa1\x34\x15\x1d\xf1\xe8\x0c\x15\x0d\xa0\xa2\xcf\xf8\xbb\x05\xad\x0d\x19\xeb\xb2\xe4\x0a\xa2\xad\xcb\x3c\xdf\x39\xf7\xdf\x87\xeb\xdf\xba\x1f\x97\x73\xbf\x3c\x71\xff\x2b\xef\xde\xbc\xbb\xed\xde\x96\x6a\xe1\x58\x81\xee\x95\xaf\xe7\xe1\x3e\x03\x07\x6d\x48\xf0\xf5\xfd\xaa\xb8\xbc\x99\x77\x4f\x1f\x77\x4f\x8b\xcd\x6a\x2a\x56\xe3\x62\x3d\x4f\xff\xc9\x45\xd3\x4d\x00\x34\xa8\x7c\x35\x2e\xd7\xfd\x6a\x92\x4b\x78\x72\x22\x2a\xbb\x6d\x3f\xee\x7e\xa6\x35\x8a\xcb\xbb\xae\xbb\xbe\xdd\xce\xc3\xe7\xed\x30\x2e\xbe\xe8\xb8\x3c\x39\xf1\xfc\x49\xf6\x79\xed\x63\xa9\x7a\x42\xe3\xf2\x65\xe1\x52\xa4\x37\xd7\xdd\x91\x28\xa3\x05\xa3\x5c\xde\x75\x47\x28\x61\xfe\xa1\xda\xc5\xe3\xaf\x00\x00\x00\xff\xff\x67\x38\x1f\xcd\x62\x17\x00\x00")

func mappingJsonBytes() ([]byte, error) {
	return bindataRead(
		_mappingJson,
		"mapping.json",
	)
}

func mappingJson() (*asset, error) {
	bytes, err := mappingJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mapping.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"mapping.json": mappingJson,
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
	"mapping.json": &bintree{mappingJson, map[string]*bintree{}},
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
