package v1

import (
	"fmt"
	"github.com/dgraph-io/badger"
	"github.com/eifrigmn/common/storage/common"
	"io/ioutil"
	"os"
	"time"
)

// Get fetches the value of the specified k
func (s *internalDB) Get(key []byte) ([]byte, error) {
	var value []byte
	var err error
	// read only transactions
	err = s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		value, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return nil, nil
		}
		return nil, err
	}
	return value, nil
}

// Size returns the size of th database (LSM + ValueLog) in bytes
func (s *internalDB) Size() int64 {
	lsm, vlog := s.db.Size()
	return lsm + vlog
}

// Close close the badger db instance
func (s *internalDB) Close() error {
	return s.db.Close()
}

// MGet fetch multiple values of the specified keys
func (s *internalDB) MGet(keys [][]byte) *common.Pair {
	var result *common.Pair
	_ = s.db.View(func(txn *badger.Txn) error {
		for _, key := range keys {
			item, err := txn.Get(key)
			if err != nil {
				continue
			}

			val, err := item.ValueCopy(nil)
			if err != nil {
				continue
			}
			result.Store(key, val)
		}
		return nil
	})

	return result
}

// ExpiresAt returns a Unix time value indicating when the item will be
// considered expired. 0 indicates that the item will never expire.
func (s *internalDB) ExpireAt(key []byte) (int64, error) {
	var expires uint64
	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		expires = item.ExpiresAt()
		return nil
	})

	return int64(expires), err
}

// Set sets a key with the specified value
func (s *internalDB) Set(key, val []byte) error {
	return s.db.Update(func(txn *badger.Txn) error {
		err := txn.Set(key, val)
		if err == badger.ErrTxnTooBig {
			_ = txn.Commit()
			txn = s.db.NewTransaction(true)
			return txn.Set(key, val)
		}
		return err
	})
}

// SetWithTTL sets a key with the specified value and the ttl
// Once the TTL has elapsed, the key will no longer be retrievable and will be eligible for garbage collection
func (s *internalDB) SetWithTTL(key, val []byte, duration time.Duration) error {
	return s.db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry(key, val).WithTTL(duration)
		err := txn.SetEntry(e)
		return err
	})
}

// MSet sets multiple key-value pairs
func (s *internalDB) MSet(data *common.Pair) error {
	return s.db.Update(func(txn *badger.Txn) error {
		data.Range(func(key, val interface{}) bool {
			err := txn.Set(key.([]byte), val.([]byte))
			if err == badger.ErrTxnTooBig {
				_ = txn.Commit()
				txn = s.db.NewTransaction(true)
				_ = txn.Set(key.([]byte), val.([]byte))
			}
			return true
		})
		return nil
	})
}

// ListKeysByPrefix fetch all the matched keys with the specified prefix
// by iterator over the whole store.
func (s *internalDB) ListKeysByPrefix(prefix []byte) [][]byte {
	var keys [][]byte
	_ = s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.KeyCopy(nil)
			keys = append(keys, k)
		}
		return nil
	})
	return keys
}

// Scan iterate over the whole store using the handler function
func (s *internalDB) Scan(pfx []byte, handler func(key, val []byte) error) error {
	return s.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Seek(pfx); it.ValidForPrefix(pfx); it.Next() {
			item := it.Item()
			k := item.KeyCopy(nil)
			err := item.Value(func(v []byte) error {
				return handler(k, v)
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// Del removes key(s) from the store
func (s *internalDB) Del(keys [][]byte) error {
	return s.db.Update(func(txn *badger.Txn) error {
		for _, key := range keys {
			err := txn.Delete(key)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *internalDB) GetDB() interface{} {
	return s.db
}

func (s *internalDB) BackData(filePath string) error {
	// backup data of version 1
	bdbV1Instance := s.GetDB().(*badger.DB)
	bak, err := ioutil.TempFile(filePath, "badgerv1.backup")
	if err != nil {
		fmt.Println("open temp file path fail")
		return err
	}
	defer bak.Close()
	_, err = bdbV1Instance.Backup(bak, 0)
	if err != nil {
		fmt.Println("back v1 data fail", err)
		return err
	}

	return nil
}

func (s *internalDB) LoadData(filePath, bkFilePath string) error {
	bak, err := os.Open(bkFilePath)
	if err != nil {
		fmt.Println("open backed file error", err)
		return err
	}
	defer bak.Close()
	err = s.db.Load(bak, 16)
	if err != nil {
		fmt.Println("load backed file error", err)
		return err
	}
	//s.db.DropPrefix()
	return os.Remove(bak.Name())
}
