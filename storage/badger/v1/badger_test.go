package v1_test

import (
	badger "github.com/eifrigmn/common/storage/badger/v1"
	"github.com/eifrigmn/common/storage/mock"
	"path"
	"testing"
)

func TestNewDatastore(t *testing.T) {
	filePath := path.Join(mock.BaseDataPath, "badger/v1")
	_, err := badger.NewDatastore(filePath)
	t.Log(err)
}
