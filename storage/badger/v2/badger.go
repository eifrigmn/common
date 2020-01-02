package v2

import (
	"github.com/dgraph-io/badger/v2"
	"github.com/eifrigmn/common/storage"
	"time"
)

type internalDB struct {
	db *badger.DB
}

func NewDatastore(filePath string) (storage.DiskStorage, error) {
	db, err := badger.Open(badger.DefaultOptions(filePath))
	if err != nil {
		return nil, err
	}
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
		again:
			// 官方建议值为0.5，示例的值为0.7
			err := db.RunValueLogGC(0.7)
			if err == nil {
				goto again
			}
		}
	}()
	return &internalDB{db: db}, nil
}
