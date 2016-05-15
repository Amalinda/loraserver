// Code generated by go-bindata.
// sources:
// ../../migrations/0001_initial.sql
// ../../migrations/0002_join_accept_params.sql
// DO NOT EDIT!

package migrations

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

var __0001_initialSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x90\xc1\x4e\xc3\x30\x10\x44\xcf\xdd\xaf\xd8\x63\x2b\xa8\x54\xce\xbd\xf2\x0b\x9c\xa3\xad\x3d\x14\x0b\x67\x6d\x6d\x9c\x42\xfe\x1e\x27\x04\x88\x25\x2a\xe5\x10\xef\x8c\xdf\xcc\xfa\x78\xe4\x87\x3e\x5c\x4d\x0a\xf8\x25\x93\x33\xcc\x7f\x45\x2e\x11\x2c\x39\xc7\xe0\xa4\x84\xa4\xbc\xa7\x5d\x3d\x76\x18\x03\x5f\xa6\x02\xe1\x6c\xa1\x17\x9b\xf8\x1d\xd3\x23\xed\x54\x7a\xb0\x7b\x13\x13\x57\x60\x7c\xab\x4a\xd0\x2b\xef\x9f\x4e\xa7\x03\x6b\x2a\xac\x63\x8c\x74\x38\x53\x9b\xa0\xc9\x63\x46\x7b\xdc\xee\xa2\xdb\x58\xc3\x2b\x0c\xea\x30\x34\xf5\xea\xe7\x11\x51\xc1\x4e\x06\x27\x1e\xbf\xa1\x2b\xa1\xb2\x56\xc2\x46\x18\x07\xf8\x6e\xce\xd6\xb4\x10\x17\xc3\xb6\x65\x50\x8f\xcf\xa5\x65\xf7\x53\xa3\x26\x7d\xb7\x5e\x07\xb3\x9b\xb6\xaf\xf8\x9c\x3e\x94\xbc\xa5\xfc\xcf\xed\xea\x5d\x94\xbf\xed\xdb\xc9\x66\xa5\x33\x7d\x05\x00\x00\xff\xff\xbe\x13\xa2\xf9\x9b\x01\x00\x00")

func _0001_initialSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_initialSql,
		"0001_initial.sql",
	)
}

func _0001_initialSql() (*asset, error) {
	bytes, err := _0001_initialSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_initial.sql", size: 411, mode: os.FileMode(420), modTime: time.Unix(1457868859, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __0002_join_accept_paramsSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x52\xcd\x6e\xe2\x30\x10\x3e\xe3\xa7\x98\x23\xd1\x06\x29\xec\x1e\xf6\x40\xe1\xd4\x57\xe8\x39\x1a\xec\x49\xb0\x70\xc6\x74\xe2\xb4\xe5\xed\xeb\x18\x25\x84\x34\x52\xab\xde\xac\x99\xf9\x7e\xe5\xcd\x06\xfe\x34\xb6\x16\x0c\x04\x2f\x17\xa5\x85\xfa\x57\xc0\xa3\x23\xd0\x27\x64\x26\x57\x3a\xdb\x06\x58\xab\x95\x35\x70\xb4\x75\x4b\x62\xd1\xc1\x45\x6c\x83\x72\x85\x33\x5d\x73\xb5\x62\x6c\xd2\xbd\xa0\x0e\x24\xf0\x16\x37\x96\x6b\x58\x6f\x8b\x22\x03\xf6\x01\xb8\x73\x4e\x65\x3b\xb5\xa8\xf0\x0d\xf9\xd4\x47\x79\xbb\xb3\x1c\x40\xa8\x22\x21\xd6\xd4\x3e\x3a\xf5\x0c\x86\x1c\x45\x11\x8d\xad\x46\x43\xa3\x81\x3b\x17\x44\x02\xaa\xa3\xd3\xc9\xaa\x12\x7a\xed\x22\xdf\x75\x69\xa9\x4f\xa4\xcf\xb0\x1e\xe0\x87\x3d\xfc\x03\x64\x33\x26\x78\xda\xc3\xff\x34\xb8\xb3\x1c\xa0\xc8\x22\xb2\x63\x1b\x07\x23\x74\x48\x91\x0f\xd0\x2c\xd5\x82\xae\xef\xed\xd6\x0a\x7b\x43\x6a\x85\x26\xb2\x7b\xd7\x35\x0c\xf2\x51\xc6\x44\x98\x8c\xfd\x1d\x5d\xc5\x94\x15\x76\x2e\x40\x91\xcf\xae\xb7\xa5\x91\xd2\x57\x55\x4b\xe1\x67\x90\x5f\x37\xdc\x2b\xf4\xc4\xcb\x09\xd2\x64\x9e\xc1\x88\xbf\x0c\x3e\xf2\x2f\x47\x53\xeb\xd3\xcb\x28\x30\xfd\xaa\xcf\xfe\x9d\x17\x14\x13\x62\x26\x98\xcf\xc7\x13\x89\xd9\x6e\x56\x43\xd4\x4c\xdb\x87\xbf\xba\x38\x4c\x88\x9d\xfa\x0c\x00\x00\xff\xff\x38\x36\x6f\x1a\x4d\x03\x00\x00")

func _0002_join_accept_paramsSqlBytes() ([]byte, error) {
	return bindataRead(
		__0002_join_accept_paramsSql,
		"0002_join_accept_params.sql",
	)
}

func _0002_join_accept_paramsSql() (*asset, error) {
	bytes, err := _0002_join_accept_paramsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0002_join_accept_params.sql", size: 845, mode: os.FileMode(420), modTime: time.Unix(1463330460, 0)}
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
	"0001_initial.sql": _0001_initialSql,
	"0002_join_accept_params.sql": _0002_join_accept_paramsSql,
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
	"0001_initial.sql": &bintree{_0001_initialSql, map[string]*bintree{}},
	"0002_join_accept_params.sql": &bintree{_0002_join_accept_paramsSql, map[string]*bintree{}},
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

