// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// gui.glade
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

// Mode return file modify time
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

var _guiGlade = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x56\x4f\x73\xda\x3e\x10\xbd\xf3\x29\xf4\xdb\xeb\x6f\xc0\x04\x3a\x9d\x1e\x64\x67\xa6\x33\x4d\x2e\xbd\x35\x6d\x8e\x1e\x21\x2d\x78\x8b\x90\x5c\x69\xf9\xf7\xed\x3b\xc4\x6d\x31\x41\x14\xc7\xc9\xa9\x37\x46\x7a\x4f\xfb\x9e\xde\x6a\xb1\xbc\xdd\xad\xac\xd8\x60\x88\xe4\x5d\x0e\x37\xa3\x31\x08\x74\xda\x1b\x72\x8b\x1c\xbe\x3e\xdc\x0d\x3f\xc0\x6d\x31\x90\xff\x0d\x87\xe2\x1e\x1d\x06\xc5\x68\xc4\x96\xb8\x12\x0b\xab\x0c\x8a\xe9\x68\xfa\x7e\x34\x16\xc3\x61\x31\x90\xe4\x18\xc3\x5c\x69\x2c\x06\x42\xc8\x80\x3f\xd6\x14\x30\x0a\x4b\xb3\x1c\x16\xbc\xfc\x1f\x8e\x85\xa6\xa3\xc9\x04\xb2\x27\x9c\x9f\x7d\x47\xcd\x42\x5b\x15\x63\x0e\xf7\xbc\x7c\x24\x67\xfc\x16\x04\x99\x1c\x56\x8a\x5c\xb9\x6d\x16\x0e\x68\x21\x64\x1d\x7c\x8d\x81\xf7\xc2\xa9\x15\xe6\xa0\x95\x2b\xe7\x5e\xaf\x23\x14\x77\xca\x46\x94\xd9\x6f\x40\x1a\xdf\x1c\x56\xd6\x3e\x12\x93\x77\x50\x68\x3c\xc8\x3e\xa3\xe9\x8a\xac\x69\x7e\xa7\x44\x7e\xd1\xc1\x5b\x8b\xe6\xb1\xad\x2d\x55\x6f\x43\x91\x66\x16\xa1\x78\x08\xeb\x33\x71\x57\x0c\x75\xa5\xc4\x4a\x1d\x3c\xf1\xbe\x46\x28\xc8\x75\xa1\xac\xc8\x95\xda\x3b\x46\xc7\xe5\x96\x0c\x57\x50\xbc\x1b\x8f\x5f\xca\xac\x90\x16\x15\x43\x31\xe9\x48\x55\xbb\xbe\x45\x5b\xcc\x6b\x45\x4f\x92\x4b\xa7\xf7\x8d\x70\x5b\xfb\xc0\xd0\x86\xf5\xc8\xae\x4f\x43\x5e\xd6\x99\xd6\xfa\xd1\xef\xe0\x39\xaa\xa7\xd4\xd7\xc8\x4d\x71\x7d\x20\x74\xac\x9a\x77\xb4\xc1\xc0\xa4\x95\xfd\xeb\x01\x49\xcf\x69\xdf\x9f\x1c\x87\x7d\x33\x05\x6a\xc5\x8c\xc1\x95\xf8\xb4\x94\xa2\xbf\xe2\x42\xae\x5c\xca\x75\xba\xcc\x1a\xed\xc9\xbd\x5a\xe9\x25\xb9\x45\xb7\xb2\xb8\xab\x95\x33\x1d\x82\x48\x91\xe7\x64\x6d\x3f\xb7\xc7\x59\x98\x7c\x50\x6d\xa7\x17\xed\xc8\xec\x42\xb4\x2f\x89\xfc\x33\x45\x3e\xb4\xfb\xaf\xd1\xcf\xba\xc2\x58\x5a\x8a\x5c\xce\x92\x8f\x20\x65\xe6\xad\x72\xef\x90\xc1\x3f\x14\xfc\xcd\x5b\x07\x9f\xbe\x9b\x04\xf8\x1c\xf8\x0c\x74\x0a\x38\xd9\x6c\x7a\x4b\x1c\xfe\xf7\x72\x60\x62\x8b\x33\x15\xfe\xf4\x89\xac\xad\xd2\x58\x79\x6b\x30\x64\x67\xec\xe3\xb1\x32\x6b\x7d\xb8\xfc\x0c\x00\x00\xff\xff\x29\xaf\xf5\xab\x11\x09\x00\x00")

func guiGladeBytes() ([]byte, error) {
	return bindataRead(
		_guiGlade,
		"gui.glade",
	)
}

func guiGlade() (*asset, error) {
	bytes, err := guiGladeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gui.glade", size: 2321, mode: os.FileMode(420), modTime: time.Unix(1602961737, 0)}
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
	"gui.glade": guiGlade,
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
	"gui.glade": &bintree{guiGlade, map[string]*bintree{}},
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
