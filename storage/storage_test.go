package storage_test

import (
	"encoding/binary"
	"flag"
	"github.com/eifrigmn/common/storage"
	v1 "github.com/eifrigmn/common/storage/badger/v1"
	"github.com/eifrigmn/common/storage/mock"
	"path"
)

var count = flag.Int("count", 1000, "item count for test")

var stores = []struct {
	Name    string
	Path    string
	Factory func(filepath string) (storage.DiskStorage, error)
}{
	{"badger", path.Join(mock.BaseDataPath, "badger/v1"), v1.NewDatastore},
}

func prefixKey(i int) []byte {
	r := make([]byte, 8)
	binary.BigEndian.PutUint64(r, uint64(i))
	return r
}
