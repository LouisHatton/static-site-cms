package kv_test

import (
	"os"
	"testing"

	"github.com/LouisHatton/static-site-cms/pkg/kv"
)

var store *kv.KvStore

func TestConnect(t *testing.T) {
	os.Setenv("KV_SQLITE_FILE_PATH", "./")
	connection, err := kv.Connect("test")
	if err != nil {
		t.Errorf("failed to connect to kv store: " + err.Error())
	}
	store = &connection
	err = store.Close()
	if err != nil {
		t.Errorf("failed to close connection: " + err.Error())
	}
}
