package storage

import (
	"github.com/eifrigmn/common/storage/common"
	"time"
)

type DiskStorage interface {
	Get(key []byte) ( []byte, error)
	Size() int64
	MGet(keys [][]byte) *common.Pair
	ExpireAt(key []byte) (int64, error)
	Set(key, val []byte) error
	SetWithTTL(key, val []byte, duration time.Duration) error
	MSet(data *common.Pair) error
	ListKeysByPrefix(prefix []byte) [][]byte
	Scan(pfx []byte, handler func(key, val []byte) error) error
	Del(keys [][]byte) error
	GetDB() interface{}
	Close() error
	BackData(filePath string) error
	LoadData(filePath, bkFilePath string) error
}
