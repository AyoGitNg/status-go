// Code generated by go-bindata. DO NOT EDIT.
// sources:
// config/public-chain-accounts.json (165B)
// config/status-chain-accounts.json (330B)
// config/status-chain-genesis.json (612B)
// config/test-data.json (84B)
// keys/bootnode.key (65B)
// keys/firebaseauthkey (153B)
// keys/test-account1-status-chain.pk (489B)
// keys/test-account1.pk (491B)
// keys/test-account2-status-chain.pk (489B)
// keys/test-account2.pk (491B)
// keys/test-account3-before-eip55.pk (489B)
// keys/wnodepassword (21B)

package static

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _configPublicChainAccountsJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xcc\x31\x0e\xc2\x30\x0c\x40\xd1\x3d\xa7\xb0\x32\x33\xc4\x4e\x9d\xd8\xdd\x2a\xd4\xde\x23\xd8\xe9\x08\x12\x05\x09\x09\xf5\xee\xa8\x4c\x2c\x8c\xff\x0f\xef\x1d\x00\xe2\x64\x76\x7b\x5e\x1f\x18\x47\x38\xfa\x38\xee\xf7\xbe\x6d\x71\x84\x98\x5e\x4b\xe6\x39\x65\x62\x6f\x2e\xb5\x53\x29\x68\x43\x5f\x94\xd1\x59\x2a\xd5\x5e\x9c\x25\x3b\x5b\x0c\x00\xfb\xe9\x07\xa4\x3f\xe0\x94\x1a\x2a\x11\x52\x11\xd7\xac\x96\x9b\x56\xba\xac\x3c\x14\xf4\x33\xa6\xb5\xaa\xa4\x59\x70\xf8\x82\x61\x0f\x9f\x00\x00\x00\xff\xff\x45\x4e\x11\xf0\xa5\x00\x00\x00")

func configPublicChainAccountsJsonBytes() ([]byte, error) {
	return bindataRead(
		_configPublicChainAccountsJson,
		"config/public-chain-accounts.json",
	)
}

func configPublicChainAccountsJson() (*asset, error) {
	bytes, err := configPublicChainAccountsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/public-chain-accounts.json", size: 165, mode: os.FileMode(420), modTime: time.Unix(1531118119, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf6, 0xa7, 0xf7, 0xe4, 0xdd, 0x94, 0xb2, 0x97, 0x99, 0x13, 0xd1, 0xc0, 0x5f, 0xf9, 0x63, 0x13, 0x14, 0xc0, 0xe1, 0xde, 0x56, 0x3f, 0x14, 0x1b, 0x37, 0xf1, 0x87, 0x5d, 0xda, 0xc8, 0xd4, 0x1c}}
	return a, nil
}

var _configStatusChainAccountsJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xcd\xb1\x6a\xc4\x30\x0c\xc6\xf1\x3d\x4f\x61\x3c\xdf\xa0\x48\xb2\x62\xdf\x96\x5c\x93\xb9\xaf\x60\x59\xf6\xd8\x2b\xe7\x96\x16\xca\xbd\x7b\x09\xa5\xd0\x25\x94\xdb\xa4\x0f\xfe\xfc\xbe\x06\xe7\xfc\x5c\xca\xf5\xfd\xe5\x6d\xf4\x67\xb7\xff\xfb\x62\x76\xab\xbd\xfb\xb3\xf3\xf0\xa9\xdb\x28\x5c\x32\xf1\x48\x28\x19\x48\x03\x4f\x05\xc2\x42\x4c\x8a\x2b\x8e\x75\x9b\x2b\xb2\x26\x7f\xfa\x89\x9f\x73\xef\x1f\xd7\x9b\xed\xf5\xeb\xef\x3d\x38\x77\x3f\xfd\xc1\xf0\x00\x43\x08\x39\x18\xac\x13\xb6\x26\x01\xa6\xb4\x5d\xf2\xb2\xf0\x2c\xb4\x12\x19\xc6\x3c\xcb\x53\x6d\x78\x79\x08\xa3\x03\x8c\xb2\x11\x57\x49\xd1\x38\x82\xe4\x66\x10\x95\x42\xd2\x84\xd0\x42\x11\x15\x54\x89\xb5\xf2\xbf\xd8\x70\x1f\xbe\x03\x00\x00\xff\xff\x35\xef\xf0\x36\x4a\x01\x00\x00")

func configStatusChainAccountsJsonBytes() ([]byte, error) {
	return bindataRead(
		_configStatusChainAccountsJson,
		"config/status-chain-accounts.json",
	)
}

func configStatusChainAccountsJson() (*asset, error) {
	bytes, err := configStatusChainAccountsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/status-chain-accounts.json", size: 330, mode: os.FileMode(420), modTime: time.Unix(1531118119, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3d, 0x4a, 0x37, 0x3b, 0x7f, 0xee, 0x4b, 0x1b, 0x5, 0x9e, 0xc5, 0x5d, 0xbf, 0x31, 0x12, 0xf5, 0xa, 0xd2, 0xcf, 0xec, 0xee, 0xe4, 0x90, 0x42, 0x5, 0xba, 0x71, 0xce, 0x84, 0x98, 0xd1, 0xd6}}
	return a, nil
}

var _configStatusChainGenesisJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\x41\x6f\x13\x31\x10\x85\xef\xf9\x15\x96\xcf\x1c\xc6\xf6\x78\x66\xbd\x37\x4a\x59\x81\xc4\x8d\x5f\x30\xb6\xc7\xcd\xaa\x9b\x6c\x94\x6c\xa4\x16\xd4\xff\x8e\x92\x2c\xa1\x42\x9c\xf0\xcd\x7e\x33\xef\x7b\x7e\x3f\x37\xc6\x58\x99\xa6\xb9\xd8\xde\x5c\x2e\xc6\x18\x9b\x07\x47\x58\x24\xa0\x0b\x9e\x04\x42\x8e\xc8\x05\xe2\x43\xc0\x90\xfd\x67\xef\x74\xf8\xa8\x1e\x73\xfa\xb3\x62\x6c\x96\x49\xf6\x45\x6d\x6f\x2c\xbc\x38\x18\xe0\xd3\x00\x84\x8f\x8f\x31\x79\xb8\x1c\x7b\x1d\x7c\xfb\xb0\x22\x3c\x44\x89\x15\x94\x7d\x6b\x14\x81\x53\x2b\x92\x33\x0a\x05\x0d\xa1\xfa\x4e\x84\xaa\x36\x5f\xfe\x03\xb1\x59\x31\xb6\xcc\xfb\x36\x3e\xdd\x1d\x6c\xd9\xca\xb8\xff\x5a\x6d\x6f\x98\xf9\x16\xc4\x6e\xe7\x9d\x9e\x16\x95\xfa\x30\xcd\xe5\xd9\xf6\x06\x56\xa1\xca\x3c\xcc\xc7\xe7\xef\xe7\xc3\x61\x3e\x2e\xb6\x37\xcb\xf1\xac\xab\x96\x5f\x7f\xc8\x7e\x19\xcf\xbb\xbf\x97\x74\x3c\xb8\x08\x5f\xe4\xb4\xbd\x85\x8c\xa0\xa4\x35\xa3\x26\xa8\x89\x1c\x49\xe9\x72\x63\xe7\x92\x6a\xe0\x82\x80\x5d\x41\x27\x15\x52\xf0\x1d\x31\xa9\xcb\x21\xd5\x42\x9d\x80\x96\x96\x43\xb5\xef\x7d\xe3\x3f\x69\xdd\xfd\xf5\xfe\xef\x3a\xb6\x36\x96\xf3\xb4\xbc\xfe\xae\xea\xda\xce\x55\xd3\x97\xe5\x28\x8f\xb2\xc8\x1a\x30\x30\x92\x63\xe4\xc8\x01\x03\x75\xe4\x28\x91\x7a\x88\xc0\x9e\x12\xd3\x45\xa3\xe8\x21\x22\x45\xbe\xcc\x26\x52\x62\x0f\xa8\x14\x19\x99\xa9\xb1\xa7\x7c\xb3\x7e\x92\xd3\xb7\x71\x37\x2e\xef\xa1\x60\x37\x6f\x9b\x5f\x01\x00\x00\xff\xff\x91\xc6\xb3\x58\x64\x02\x00\x00")

func configStatusChainGenesisJsonBytes() ([]byte, error) {
	return bindataRead(
		_configStatusChainGenesisJson,
		"config/status-chain-genesis.json",
	)
}

func configStatusChainGenesisJson() (*asset, error) {
	bytes, err := configStatusChainGenesisJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/status-chain-genesis.json", size: 612, mode: os.FileMode(420), modTime: time.Unix(1531118119, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb, 0xf0, 0xc, 0x1, 0x95, 0x65, 0x6, 0x55, 0x48, 0x8f, 0x83, 0xa0, 0xb4, 0x81, 0xda, 0xad, 0x30, 0x6d, 0xb2, 0x78, 0x1b, 0x26, 0x4, 0x13, 0x12, 0x9, 0x6, 0xae, 0x3a, 0x2c, 0x1, 0x71}}
	return a, nil
}

var _configTestDataJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\x52\x50\x50\xf2\xcb\x4f\x49\x55\xb2\x52\x00\xb1\x15\x14\x94\x82\x2b\xf3\x92\x83\x53\x93\xf3\xf3\x52\x8a\x95\xac\x14\x8c\x0d\x74\x20\xc2\x1e\x21\x21\x01\x01\xf9\x45\x25\x4a\x56\x0a\x16\x66\x26\xa6\x50\xd1\xf0\x60\x84\x98\x19\x97\x82\x42\x2d\x57\x2d\x17\x20\x00\x00\xff\xff\x51\xca\x96\xb1\x54\x00\x00\x00")

func configTestDataJsonBytes() ([]byte, error) {
	return bindataRead(
		_configTestDataJson,
		"config/test-data.json",
	)
}

func configTestDataJson() (*asset, error) {
	bytes, err := configTestDataJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/test-data.json", size: 84, mode: os.FileMode(420), modTime: time.Unix(1531118119, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xce, 0x9d, 0x80, 0xf5, 0x87, 0xfa, 0x57, 0x1d, 0xa1, 0xd5, 0x7a, 0x10, 0x3, 0xac, 0xd7, 0xf4, 0x64, 0x32, 0x96, 0x2b, 0xb7, 0x21, 0xb7, 0xa6, 0x80, 0x40, 0xe9, 0x65, 0xe3, 0xd6, 0xbd, 0x40}}
	return a, nil
}

var _keysBootnodeKey = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x4b\x01\x80\x40\x08\x04\xd0\xbb\x69\x58\x96\x6f\x1c\x66\x94\xfe\x11\x7c\xbe\x65\xab\x03\x07\x43\x1b\x87\x5b\x17\x8e\x9e\x23\x02\x11\x1a\x57\x51\x9d\x32\xf7\x24\x99\x1b\x08\xd8\xcb\xcf\x2d\xf4\xba\x3e\x7f\x00\x00\x00\xff\xff\x4a\x3d\x56\xc6\x41\x00\x00\x00")

func keysBootnodeKeyBytes() ([]byte, error) {
	return bindataRead(
		_keysBootnodeKey,
		"keys/bootnode.key",
	)
}

func keysBootnodeKey() (*asset, error) {
	bytes, err := keysBootnodeKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "keys/bootnode.key", size: 65, mode: os.FileMode(420), modTime: time.Unix(1531118119, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x31, 0xcf, 0x27, 0xd4, 0x96, 0x2e, 0x32, 0xcd, 0x58, 0x96, 0x2a, 0xe5, 0x8c, 0xa0, 0xf1, 0x73, 0x1f, 0xd6, 0xd6, 0x8b, 0xb, 0x73, 0xd3, 0x2c, 0x84, 0x1a, 0x56, 0xa4, 0x74, 0xb6, 0x95, 0x20}}
	return a, nil
}

var _keysFirebaseauthkey = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xcb\x52\x85\x20\x18\x00\xe0\x7d\xef\xc2\xcc\x71\x4a\xc4\x76\x3f\x29\x43\x5c\x34\x2d\x4d\x59\x6a\xa5\x8d\x24\xde\x42\x7b\xfb\x3e\x00\x00\x9e\x87\xfa\xfe\xf3\xf9\x11\x5e\x20\x0e\x3a\xf6\x7b\xb8\xf2\x03\xc7\xd9\x4a\x10\xec\x46\x74\xa7\x26\x9a\xfd\x71\xec\xeb\xea\x98\x84\xb2\x9b\x6c\x92\x5b\xaa\x0d\x0e\x7c\x71\x71\xb8\xb1\xb1\xf2\x62\xcc\x61\xfb\x59\x91\xa2\x6f\x69\x6f\x29\x09\xdf\x07\x99\xb5\x09\x2a\x5f\x5b\x69\xa2\xa5\x91\x33\xa9\xa5\x67\x2d\x5d\xdd\xe1\xf0\xac\x82\xb0\xcf\xec\xd7\x72\x3d\x9c\x4f\x54\x4c\x3c\x2d\xad\x8b\xac\x32\x57\x8c\x06\xfc\x5d\x24\xd3\x7e\xf7\x1f\x00\x00\xff\xff\xd6\xa2\x00\x4a\x99\x00\x00\x00")

func keysFirebaseauthkeyBytes() ([]byte, error) {
	return bindataRead(
		_keysFirebaseauthkey,
		"keys/firebaseauthkey",
	)
}

func keysFirebaseauthkey() (*asset, error) {
	bytes, err := keysFirebaseauthkeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "keys/firebaseauthkey", size: 153, mode: os.FileMode(420), modTime: time.Unix(1531118119, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe, 0x69, 0x23, 0x64, 0x7d, 0xf9, 0x14, 0x37, 0x6f, 0x2b, 0x1, 0xf0, 0xb0, 0xa4, 0xb2, 0xd0, 0x18, 0xcd, 0xf9, 0xeb, 0x57, 0xa3, 0xfd, 0x79, 0x25, 0xa7, 0x9c, 0x3, 0xce, 0x26, 0xec, 0xe1}}
	return a, nil
}

var _keysTestAccount1StatusChainPk = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xd0\xdd\x8a\xe4\x40\x08\x05\xe0\x77\xf1\x3a\x01\xb5\x4c\x55\x99\xb7\x51\xcb\x62\x9b\xf9\x6b\x92\x66\xd8\x65\xe8\x77\x5f\x32\xcb\x32\x97\x82\x7e\x1c\xcf\x17\x7c\xe6\x71\xde\x3e\xde\x61\x2f\x0b\xdc\x06\xec\xd0\x29\xb2\x25\xe9\xca\xd4\x71\x15\xe4\xbe\xba\x2a\xae\x68\x52\x94\x7c\x56\x33\x84\x05\x6c\x8c\x23\xcf\x13\x76\xf0\x49\x55\xc2\x8a\x50\xe1\x6a\x58\x7c\x93\x16\xb8\x79\x91\xe2\x9c\x4c\x39\x2d\x59\x5c\x61\x81\x38\xfe\xdc\x1f\x1f\xb0\x7f\x41\xdc\xee\xbf\xf2\x78\xe4\xef\x07\xec\xb0\x35\x17\xb1\x6e\xd4\x67\x6f\x8a\xea\xce\xd1\x4c\x95\x98\x5c\x54\x8b\xa9\xb7\x31\x2c\x89\x46\xe9\x53\x23\x31\x91\x42\x9a\x8c\x51\x6d\x9b\xf5\x92\xbf\xbd\xbb\x1d\xf6\x76\x5e\xfe\xed\x13\x76\x30\xc3\xad\x17\xec\x93\x38\x2a\x73\xca\xa8\x38\x34\xb0\x59\x6a\x2f\x2d\xe1\xf9\xff\xf0\x5a\xce\x73\x25\xee\x6b\x3c\x0e\x58\xe0\x65\x4c\xd8\xe1\xfc\x0e\xfc\x6f\xfc\xb1\xc7\xcb\x6b\x5e\x85\xf1\x02\xa7\xbd\x5e\x0f\xb8\x57\xc1\x62\x9e\x8d\xb3\xb6\x9a\xcc\x64\x82\xc6\x81\x6d\x93\xd9\xe8\x4a\xae\x54\x44\xab\x78\x20\x99\x85\x52\x99\x21\x65\x23\x1a\x19\x05\x16\x78\x87\xbd\x93\xf2\x02\x07\xec\x7d\x81\x3b\xec\xf4\x5c\xe0\xcd\x02\x76\x68\x9d\xb1\x0d\xa4\x40\x2f\x51\x49\xad\x4f\x51\x47\xd6\xda\xbd\x56\xaf\xc5\xc2\x37\x66\xec\x41\x45\xc8\xb4\x8c\xd0\x4a\x6d\xe0\x4c\x95\x9e\x09\xcf\xe7\xdf\x00\x00\x00\xff\xff\xd2\xdb\x1b\x65\xe9\x01\x00\x00")

func keysTestAccount1StatusChainPkBytes() ([]byte, error) {
	return bindataRead(
		_keysTestAccount1StatusChainPk,
		"keys/test-account1-status-chain.pk",
	)
}

func keysTestAccount1StatusChainPk() (*asset, error) {
	bytes, err := keysTestAccount1StatusChainPkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "keys/test-account1-status-chain.pk", size: 489, mode: os.FileMode(420), modTime: time.Unix(1531118119, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8f, 0xba, 0x35, 0x1, 0x2b, 0x9d, 0xad, 0xf0, 0x2d, 0x3c, 0x4d, 0x6, 0xb5, 0x22, 0x2, 0x47, 0xd4, 0x1c, 0xf4, 0x31, 0x2f, 0xb, 0x5b, 0x27, 0x5d, 0x43, 0x97, 0x58, 0x2d, 0xf0, 0xe1, 0xbe}}
	return a, nil
}

var _keysTestAccount1Pk = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8f\x4b\x8e\x24\x3b\x08\x45\xf7\xc2\x38\x42\xb2\x31\xd8\xd8\xf3\x57\xfb\xc0\x80\xf5\x52\xf5\x4b\x45\xa4\x4a\xdd\x2a\xd5\xde\x5b\x91\x83\xee\x19\x20\x74\xee\x3d\xdf\xa0\xee\x47\x9c\x27\x0c\x78\x29\xfc\x5f\x2a\xc8\xae\x2e\x2d\xb0\xd6\x6c\x14\x2f\x9d\xb3\xb3\x34\x6c\x51\x9d\xa5\x38\x1b\x6c\x60\xc7\xef\xfb\xe3\x13\xc6\x37\xd8\xed\xfe\x7f\x1c\x30\x40\xe3\xdc\x33\xca\x6e\x8f\xe3\x7a\x78\x9e\x1f\xf1\xeb\x01\x03\x48\x02\x7d\x89\xb7\xb4\x64\xae\x42\xea\x1a\x9c\x66\x9a\x29\x4a\x96\xa4\xa5\x63\xd0\x9c\xb9\x89\x66\xa5\x82\xe4\x66\x92\x08\x5b\x4c\x6a\x95\xd6\x5f\xde\x5d\x0f\x7d\x3f\xaf\xd8\xdb\xd7\xc5\x2d\xc4\xa6\xb4\xb0\xd8\xcc\x19\xbb\xf6\x9e\x5b\x0e\xe9\x5a\x3b\xb2\x24\xf8\xd9\xe0\xd5\x17\x0c\x38\x9f\x85\xe1\xb9\xfe\x83\xf8\xeb\x5b\x7c\xc0\x28\xb8\xc1\x07\x0c\xac\x98\x89\x36\xb8\xc3\xc8\x1b\x1c\x30\x64\x83\x53\xdf\x2e\x81\xdc\x99\xb1\xc9\x9c\x46\xab\x68\x22\xb1\xa4\xe2\xad\x79\x70\xca\x54\x8b\xa1\x08\xf6\x59\x9e\xb3\x06\xad\x1e\xc8\x68\x58\xda\xc2\xa9\xd6\xae\x22\xef\x6a\x30\xa0\xa3\x97\x60\xd2\xde\x22\xbc\x74\xcf\x8b\x6a\xc3\xb6\x12\x13\x55\x26\x4c\xcb\x78\x95\x6c\x6d\x51\xb4\x14\x42\x55\xa7\x75\x76\x5b\x2e\x53\xf0\x49\xba\x39\x0c\x98\x35\x11\x59\x95\xbd\x74\xa3\x9d\xaa\xfb\x2e\x39\xe3\x8e\x16\x8c\xd9\x19\x6d\x31\x6c\xf0\x15\xc7\x79\xfb\xbc\x24\x7f\xfe\x04\x00\x00\xff\xff\x0d\x85\xaf\xaf\xeb\x01\x00\x00")

func keysTestAccount1PkBytes() ([]byte, error) {
	return bindataRead(
		_keysTestAccount1Pk,
		"keys/test-account1.pk",
	)
}

func keysTestAccount1Pk() (*asset, error) {
	bytes, err := keysTestAccount1PkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "keys/test-account1.pk", size: 491, mode: os.FileMode(420), modTime: time.Unix(1531118119, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9, 0x43, 0xc2, 0xf4, 0x8c, 0xc6, 0x64, 0x25, 0x8c, 0x7, 0x8c, 0xa8, 0x89, 0x2b, 0x7b, 0x9b, 0x4f, 0x81, 0xcb, 0xce, 0x3d, 0xef, 0x82, 0x9c, 0x27, 0x27, 0xa9, 0xc5, 0x46, 0x70, 0x30, 0x38}}
	return a, nil
}

var _keysTestAccount2StatusChainPk = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xd0\xdb\x8a\x1c\x41\x08\x80\xe1\x77\xf1\xba\x1b\x2c\xb5\x8e\x6f\x63\xa9\x45\x86\x3d\x0d\xdd\xc3\x92\xb0\xcc\xbb\x87\xde\x10\xf6\xd2\x8b\xff\x43\xfd\x82\xcf\x38\xce\xdb\xc7\x3b\x0c\xde\xe0\xe6\x30\x60\x9a\x69\xa3\x4c\xfb\x4a\x9a\x76\x59\xa1\xbb\xd6\x25\x3b\x17\x12\x49\x94\x8c\xd0\x60\x03\x75\x3f\xe2\x3c\x61\x00\x61\xd6\xec\x18\x95\xd6\x2a\x19\x6b\x5f\xa6\x73\x8a\x16\x0e\x66\xa7\xa6\x5a\x3c\x16\x5d\x91\x1d\x7f\xee\x8f\x0f\x18\x5f\x60\xb7\xfb\xaf\x38\x1e\xf1\xfb\x01\x03\xa4\xf6\x5c\x3d\x2c\xa8\xba\x4a\x11\x66\x77\xf4\x9c\x27\x77\x5f\x69\x61\xd7\xe0\x1e\x5c\x98\xe6\x42\xed\xc6\xce\x4b\x71\x52\x42\xd1\xa6\x34\x2f\xf9\xdb\xbb\xeb\xa1\x6f\xe7\xe5\xdf\x3e\x61\x80\x75\xae\x91\x7b\x23\x69\xa4\x8e\x3d\x73\x2f\x09\xb5\xf6\x98\xdd\x4a\x15\x78\xfe\x0f\x61\x80\xc6\xb9\x27\x6a\xbb\x3d\x0e\xd8\xe0\xc5\x17\x0c\x38\xbf\x17\xfe\x37\xfe\xd8\xfe\xf2\x1a\xd7\xc3\x68\x83\x53\x5f\xaf\x03\xd8\xb4\x18\xb7\x94\x5b\xc9\x92\x75\x22\x1b\x49\x14\x94\x42\x86\xc9\xc9\x5b\x53\xa9\xc2\x5c\xb4\x7a\x12\xd4\xa6\x4b\x94\x32\x16\x66\xce\x15\x61\x83\x77\x18\x2d\x75\xda\xe0\x80\xd1\x36\xb8\xc3\x48\xcf\x0d\xde\xd4\x60\x00\x56\x47\x67\x6a\xd6\xb1\x36\x2b\xd1\x66\x89\x55\x72\xe5\x54\xa8\xab\x72\x4c\x6a\xd1\x91\x49\x1a\x2f\xa1\x99\x45\xcd\x94\x85\x5c\x22\xb1\x34\x78\x3e\xff\x06\x00\x00\xff\xff\x4e\x39\x80\x7d\xe9\x01\x00\x00")

func keysTestAccount2StatusChainPkBytes() ([]byte, error) {
	return bindataRead(
		_keysTestAccount2StatusChainPk,
		"keys/test-account2-status-chain.pk",
	)
}

func keysTestAccount2StatusChainPk() (*asset, error) {
	bytes, err := keysTestAccount2StatusChainPkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "keys/test-account2-status-chain.pk", size: 489, mode: os.FileMode(420), modTime: time.Unix(1531118119, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9, 0xf8, 0x5c, 0xe9, 0x92, 0x96, 0x2d, 0x88, 0x2b, 0x8e, 0x42, 0x3f, 0xa4, 0x93, 0x6c, 0xad, 0xe9, 0xc0, 0x1b, 0x8a, 0x8, 0x8c, 0x5e, 0x7a, 0x84, 0xa2, 0xf, 0x9f, 0x77, 0x58, 0x2c, 0x2c}}
	return a, nil
}

var _keysTestAccount2Pk = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8f\x4d\x8a\x1c\x3d\x0c\x40\xef\xa2\x75\x15\xd8\xb2\x6c\x4b\xde\x7d\x7c\xe4\x20\xb2\x24\x93\x66\xfe\x9a\xaa\x66\x48\x18\xe6\xee\xa1\x7a\x91\x2c\x25\xc4\xd3\x7b\x5f\xa0\xee\x47\x9c\x27\x0c\xf8\x2f\x69\x16\xc4\x8c\x8d\x5d\x8a\x58\x51\xe9\x38\x57\xa5\x96\xfd\xff\x9c\x56\x17\x4e\x3f\x38\x13\x6c\x60\xc7\xef\xfb\xe3\x03\xc6\x17\xd8\xed\xfe\x33\x0e\x18\xa0\x71\xee\x19\x79\xb7\xc7\x71\x1d\x3c\xd7\x8f\xf8\xf5\x80\x01\xac\x2b\xc9\x5c\x9a\x50\x69\xf5\xa8\x2b\x17\x2d\xcc\x33\x8a\x64\xb7\x95\xc4\x6a\x36\xa6\x65\x2c\xc8\x84\x19\x31\x45\xea\xd6\x72\x4b\x84\x75\xf1\x5f\xde\x5d\x0f\x7d\x3b\xaf\xb7\xb7\x4f\x18\xd0\x75\x56\xd1\xa0\x2e\xb9\xf6\x1e\x1d\x3d\x9a\xd7\xe9\xac\x14\x48\xa1\xf0\xbd\xc1\x8b\x2f\x18\x70\x3e\x85\xe1\x39\xfe\x83\xf8\xcb\x6b\xbc\xc3\x28\xb8\xc1\x3b\x0c\x6c\x98\x89\x36\xb8\xc3\xc8\x1b\x1c\x30\x78\x83\x53\x5f\xaf\x00\xec\xd1\x6c\x72\xe4\x10\x4f\xd6\xdb\x9a\x73\x39\x91\x99\x22\xd5\xd2\x92\x2f\xa1\x24\xd3\x4c\xd4\x68\x51\x10\xdb\xa5\x35\xeb\x72\x9e\x49\xfa\x25\xf2\xa6\x76\x29\x2f\xf7\x59\x44\x3a\x7a\xcf\xb8\x26\x8a\x52\x59\x39\x72\xd4\x92\x4a\xe4\xc2\x1e\x26\x66\xa8\x2b\x66\xc2\xda\x82\xb5\xa7\x8a\xae\x54\xa4\x5b\xbb\x48\x37\x87\x01\x93\x51\xba\x4a\xde\x25\x05\xed\x24\x5e\x77\x6d\x4b\x76\xb6\x92\x2a\xf5\xc9\xd1\x14\x36\xf8\x8c\xe3\xbc\x7d\x5c\x91\xdf\x7f\x02\x00\x00\xff\xff\x80\x42\xa2\x5c\xeb\x01\x00\x00")

func keysTestAccount2PkBytes() ([]byte, error) {
	return bindataRead(
		_keysTestAccount2Pk,
		"keys/test-account2.pk",
	)
}

func keysTestAccount2Pk() (*asset, error) {
	bytes, err := keysTestAccount2PkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "keys/test-account2.pk", size: 491, mode: os.FileMode(420), modTime: time.Unix(1531118119, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9f, 0x72, 0xd5, 0x95, 0x5c, 0x5a, 0x99, 0x9d, 0x2f, 0x21, 0x83, 0xd7, 0x10, 0x17, 0x4a, 0x3d, 0x65, 0xc9, 0x26, 0x1a, 0x2c, 0x9d, 0x65, 0x63, 0xd2, 0xa0, 0xfc, 0x7c, 0x0, 0x87, 0x38, 0x9f}}
	return a, nil
}

var _keysTestAccount3BeforeEip55Pk = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x90\xdb\x8a\x1c\x31\x0c\x44\xff\x45\xcf\xdd\x20\x4b\xbe\xc8\xfe\x1b\x59\x96\xc8\xb0\xb7\xa1\x7b\x58\x12\x96\xf9\xf7\xd0\x59\xc2\x3e\xd6\x43\x1d\x4e\xd5\x17\x7c\xfa\x71\xde\x3e\xde\x61\xf0\x06\xb7\x05\x03\xd0\x96\xb9\x85\xed\x46\x61\x7b\x26\x4d\xbb\xf8\x2a\x7b\x26\x33\xec\xd4\x16\x2f\x81\x0d\x74\xad\xc3\xcf\x13\x06\xb0\x2e\xce\x5e\xbb\xac\x2c\x58\x35\x16\xca\xe4\xd2\x67\x27\x8c\x62\x75\x56\x9a\x55\xdc\x33\x6c\x60\xc7\x9f\xfb\xe3\x03\xc6\x17\xd8\xed\xfe\xcb\x8f\x87\xff\x7e\xc0\x00\xa2\xd2\xbb\x98\x77\x56\x2d\x93\x5a\x6a\xa5\x60\xc9\x9c\xd2\x52\x31\x6b\x53\xa8\x24\x6a\x2b\x23\x5a\xa7\x16\xca\xc2\x0d\x3d\x27\xa6\x4e\x66\x97\xce\x37\xef\xae\x87\xbe\x9d\x17\xff\xf6\x09\x03\xa4\x27\x6a\xd9\x49\x5b\xa3\xb6\x34\xd9\xd2\xb4\xc8\x6c\x12\xfb\x94\x89\xf0\xfc\x5f\x84\x01\xea\xe7\x9e\x48\x76\x7b\x1c\xb0\xc1\xcb\x0a\x18\x70\xfe\x13\xfe\x8e\x3f\xec\xf5\xf2\xea\xd7\x61\xb4\xc1\xa9\xaf\xd7\x80\x45\x93\x50\xb5\xd5\x2c\x11\xd8\xc9\x98\x51\x1a\x51\xc6\xf0\x5c\x83\xc2\xd8\x53\x66\xa6\xc9\x42\x9a\x05\x25\x6a\xc1\xc9\x56\xa5\x4b\x2d\x13\x36\x78\x87\x21\xa9\xd3\x06\x07\x0c\xd9\xe0\x0e\x23\x3d\x37\x78\x53\xbb\xd4\x10\x33\x63\x13\x89\xd5\x66\xa0\x93\xbb\xd7\x48\x5c\x4a\xcb\xa5\x27\x8e\xe2\x14\x25\x65\xaa\x1a\x9d\x58\x8a\x22\xba\x9b\xa6\x28\xd6\x98\x12\x3c\x9f\x7f\x03\x00\x00\xff\xff\x4c\x6d\xd5\xbe\xe9\x01\x00\x00")

func keysTestAccount3BeforeEip55PkBytes() ([]byte, error) {
	return bindataRead(
		_keysTestAccount3BeforeEip55Pk,
		"keys/test-account3-before-eip55.pk",
	)
}

func keysTestAccount3BeforeEip55Pk() (*asset, error) {
	bytes, err := keysTestAccount3BeforeEip55PkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "keys/test-account3-before-eip55.pk", size: 489, mode: os.FileMode(420), modTime: time.Unix(1531118119, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x81, 0x40, 0x56, 0xc1, 0x5e, 0x10, 0x6e, 0x28, 0x15, 0x3, 0x4e, 0xc4, 0xc4, 0x71, 0x4d, 0x16, 0x99, 0xcc, 0x1b, 0x63, 0xee, 0x10, 0x20, 0xe4, 0x59, 0x52, 0x3f, 0xc0, 0xad, 0x15, 0x13, 0x72}}
	return a, nil
}

var _keysWnodepassword = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x2e\x49\x2c\x29\x2d\xd6\xcd\x4f\x4b\xcb\xc9\xcc\x4b\xd5\xcd\xcc\x4b\xca\xaf\xe0\x02\x04\x00\x00\xff\xff\xef\xf3\x8b\x45\x15\x00\x00\x00")

func keysWnodepasswordBytes() ([]byte, error) {
	return bindataRead(
		_keysWnodepassword,
		"keys/wnodepassword",
	)
}

func keysWnodepassword() (*asset, error) {
	bytes, err := keysWnodepasswordBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "keys/wnodepassword", size: 21, mode: os.FileMode(420), modTime: time.Unix(1531118119, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x94, 0x4b, 0xc5, 0xc1, 0xf, 0x43, 0x6b, 0xab, 0xa2, 0x7d, 0x66, 0xae, 0xf8, 0x31, 0x63, 0xd0, 0x7b, 0xec, 0x4d, 0xd6, 0x91, 0x2, 0x29, 0x12, 0x9b, 0x43, 0x43, 0xba, 0x6a, 0x84, 0x4f, 0x25}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"config/public-chain-accounts.json": configPublicChainAccountsJson,

	"config/status-chain-accounts.json": configStatusChainAccountsJson,

	"config/status-chain-genesis.json": configStatusChainGenesisJson,

	"config/test-data.json": configTestDataJson,

	"keys/bootnode.key": keysBootnodeKey,

	"keys/firebaseauthkey": keysFirebaseauthkey,

	"keys/test-account1-status-chain.pk": keysTestAccount1StatusChainPk,

	"keys/test-account1.pk": keysTestAccount1Pk,

	"keys/test-account2-status-chain.pk": keysTestAccount2StatusChainPk,

	"keys/test-account2.pk": keysTestAccount2Pk,

	"keys/test-account3-before-eip55.pk": keysTestAccount3BeforeEip55Pk,

	"keys/wnodepassword": keysWnodepassword,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"config": &bintree{nil, map[string]*bintree{
		"public-chain-accounts.json": &bintree{configPublicChainAccountsJson, map[string]*bintree{}},
		"status-chain-accounts.json": &bintree{configStatusChainAccountsJson, map[string]*bintree{}},
		"status-chain-genesis.json":  &bintree{configStatusChainGenesisJson, map[string]*bintree{}},
		"test-data.json":             &bintree{configTestDataJson, map[string]*bintree{}},
	}},
	"keys": &bintree{nil, map[string]*bintree{
		"bootnode.key":                  &bintree{keysBootnodeKey, map[string]*bintree{}},
		"firebaseauthkey":               &bintree{keysFirebaseauthkey, map[string]*bintree{}},
		"test-account1-status-chain.pk": &bintree{keysTestAccount1StatusChainPk, map[string]*bintree{}},
		"test-account1.pk":              &bintree{keysTestAccount1Pk, map[string]*bintree{}},
		"test-account2-status-chain.pk": &bintree{keysTestAccount2StatusChainPk, map[string]*bintree{}},
		"test-account2.pk":              &bintree{keysTestAccount2Pk, map[string]*bintree{}},
		"test-account3-before-eip55.pk": &bintree{keysTestAccount3BeforeEip55Pk, map[string]*bintree{}},
		"wnodepassword":                 &bintree{keysWnodepassword, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
