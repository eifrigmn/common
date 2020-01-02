package badger_test

import (
	"fmt"
	"github.com/eifrigmn/common/storage"
	"path/filepath"
	//v1 "storage/badger/v1"
	v2 "github.com/eifrigmn/common/storage/badger/v2"
	"github.com/eifrigmn/common/storage/bench/common"
	"github.com/eifrigmn/common/storage/mock"
	"testing"
)

var badgerDB storage.DiskStorage

func init() {
	var err error
	badgerDB, err = v2.NewDatastore(filepath.Join(mock.BaseDataPath, mock.BadgerV1DataPath))
	if err != nil {
		fmt.Println("got error", err)
	}
}
func InitBadgerDBData() {
	for n := 0; n < 10000; n++ {
		key := common.GetKey(n)
		val := common.GeyValue64B()
		badgerDB.Set(key, val)
	}
}

func BenchmarkBadgerDBPutValue64B(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	fmt.Println()
	for n := 0; n < b.N; n++ {
		key := common.GetKey(n)
		val := common.GeyValue64B()
		if err := badgerDB.Set(key, val); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBadgerDBPutValue128B(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := common.GetKey(n)
		val := common.GeyValue128B()

		if err := badgerDB.Set(key, val); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBadgerDBPutValue256B(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := common.GetKey(n)
		val := common.GeyValue256B()

		if err := badgerDB.Set(key, val); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBadgerDBPutValue512B(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		key := common.GetKey(n)
		val := common.GeyValue512B()
		if err := badgerDB.Set(key, val); err != nil {
			b.Fatal(err)
		}
	}
}
func BenchmarkBadgerDBGet(b *testing.B) {
	InitBadgerDBData()

	b.ReportAllocs()
	b.ResetTimer()
	key := []byte("key_" + fmt.Sprintf("%07d", 99))
	for n := 0; n < b.N; n++ {
		if _, err := badgerDB.Get(key); err != nil {
			b.Fatal(err)
		}
	}
}
