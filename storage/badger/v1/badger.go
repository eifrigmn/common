package v1

import (
	"github.com/dgraph-io/badger"
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
			// official doc suggests 0.5, example offer 0.7
			err := db.RunValueLogGC(0.7)
			if err == nil {
				goto again
			}
		}
	}()
	return &internalDB{db: db}, nil
}
