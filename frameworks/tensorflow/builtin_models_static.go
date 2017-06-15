// Code generated by go-bindata.
// sources:
// frameworks/tensorflow/builtin_models/caffenet.yml
// DO NOT EDIT!

package tensorflow

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _caffenetYml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\x4d\x4f\xdb\x40\x10\x86\xef\xfb\x2b\xe6\xc2\x0d\x7f\x24\x04\x0a\x3e\x54\xaa\xe8\xa5\x6a\x0b\x52\x8b\x7a\xa9\xaa\x68\xb2\x1e\xc7\x4b\xd7\xb3\xab\x9d\x71\x42\x2e\xfd\xed\xd5\xda\x81\x02\xbd\x70\x49\x9c\xcc\x33\x5f\xef\x3b\x66\x1c\xa8\x81\x6b\xec\x3a\xba\x21\x35\x5d\xc2\x81\xf6\x21\xfd\x6e\x0c\xc0\x1c\xbb\x23\x96\x90\x3a\x1f\xf6\x06\x60\x47\x49\x5c\xe0\x06\xfe\x2c\xca\x45\xf9\x60\x6c\x60\x45\xc7\x94\x32\x8f\x43\x7b\xb1\xca\x0f\x00\xdb\x38\x36\x90\xd0\xc5\x14\xee\xc9\x6a\x65\x31\x0d\xbe\xd0\xa7\x52\xcd\xc4\x16\x36\x8e\x13\x6e\xdf\x86\x6f\x27\x3c\x46\x7b\xb1\xf2\xd4\xbc\x2d\xf3\x48\x1f\x73\xdf\x30\xd9\xf3\x84\x96\xc4\x26\x17\x75\x5a\xf9\xbd\x01\xf8\xc0\xe0\x06\xdc\x52\x61\x3d\x8a\xb8\xce\x59\xcc\x41\x68\x89\x22\xd8\xc0\xbb\xe0\xc7\xfc\x07\x7a\x60\xd2\x2c\x24\x68\xca\x02\xb5\x10\x18\xb4\x27\xf8\xf2\xfd\xc7\xb7\xeb\x62\x59\x2f\x6a\xf8\x94\x2b\xdd\x90\xce\x88\xe3\x2d\x08\xe9\xa9\x01\xd8\xf7\xce\xf6\x80\xb6\x77\xb4\xa3\x16\x34\xc4\x62\x01\xc8\xf3\xd3\x39\x50\x4a\x21\x41\x42\x25\x81\xd0\xc1\xd9\x55\xf9\xee\x64\x0a\x2f\x2e\xcb\xab\x93\xd2\x00\xdc\xf5\xf4\xd4\xdf\x06\x16\x27\x3a\xa1\x17\xf5\x57\x88\x98\x3d\x56\x4a\x32\xe5\x9c\xd7\xf5\x67\x60\x1a\x53\x60\x01\xb4\x29\x88\xc0\xf9\xab\x55\x3c\x1e\x32\xae\xfb\x00\xdd\xe8\xfd\xa1\xb0\x81\x99\xac\x52\x7b\x0c\x95\x26\x51\x47\x89\xd8\x92\x64\x5f\x0a\xe8\x55\x63\x53\x55\x11\x63\x0e\xb3\x8b\x52\x5a\x3b\xff\xac\x56\x97\xcb\x55\x31\xc9\xc8\xa4\xaf\x94\x2c\xf6\x4e\xfb\x22\xcb\x59\xbc\x98\xa1\xc8\x23\x4e\x5f\xd3\x56\x62\xbc\xb3\xc4\x92\xef\xf3\xf6\xe3\xad\x71\x1c\x47\x3d\xb6\xd6\x43\xa4\x66\xb6\x69\x72\xfc\x85\x89\xd9\x82\x09\x7e\x06\xfc\x53\x64\xbe\x29\x80\xd6\x0d\xc4\xf9\xd0\xa5\x81\x9f\x8b\x53\x38\x3b\x85\xe5\x72\x35\x7d\xfc\x32\x61\xd4\x38\x6a\x46\xe7\x4e\x1d\xa1\x8e\x29\x97\x7a\xd1\x09\x19\x66\x72\xee\x94\xfd\x00\x8f\x1b\xf2\xe6\xff\x8e\xc7\x12\xb2\x1e\x93\x6f\x1e\xb5\x6b\x51\xb1\x6c\x07\x6f\xcb\xc1\x57\xc3\x03\x93\x56\x43\x68\xc9\x4b\xf5\xa8\x5d\x25\x07\x16\xd2\x52\x1f\xd4\x4c\xa1\x5c\x6d\x83\x42\x73\x1d\x0d\x6d\x30\x00\xdb\x84\xb1\x5f\x47\xd4\xbe\x01\x9b\x5f\xf4\xac\xba\x1c\x86\x4d\xf0\xe5\xbd\x04\xce\x07\x47\x6e\xdb\xab\xbc\x86\xea\xba\xae\xcb\x69\x56\x31\x00\x4e\xd6\x98\x6c\xef\x76\x79\x67\xf4\x42\x06\x55\x93\xdb\x8c\x3a\x9b\x3e\x20\xbb\x8e\x44\xd7\x38\x6a\x1f\x52\x03\xb8\x69\x47\xdf\x66\x9d\x12\x3a\xef\x78\xbb\xce\x2b\x09\x69\xf3\x74\xfa\xe6\x6f\x00\x00\x00\xff\xff\x3b\x75\xd7\x54\x80\x04\x00\x00"

func caffenetYmlBytes() ([]byte, error) {
	return bindataRead(
		_caffenetYml,
		"caffenet.yml",
	)
}

func caffenetYml() (*asset, error) {
	bytes, err := caffenetYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "caffenet.yml", size: 1152, mode: os.FileMode(420), modTime: time.Unix(1497478652, 0)}
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
	"caffenet.yml": caffenetYml,
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
	"caffenet.yml": &bintree{caffenetYml, map[string]*bintree{}},
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

