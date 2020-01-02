package v1_test

import (
	v1 "github.com/eifrigmn/common/storage/badger/v1"
	"github.com/eifrigmn/common/storage/mock"
	"path"
	"testing"
)

var db, _ = v1.NewDatastore(path.Join(mock.BaseDataPath, "badger/v1"))

func TestInternalDB_Size(t *testing.T) {
	t.Log(db.Size())
}
