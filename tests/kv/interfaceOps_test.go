package kv_test

import (
	"os"
	"testing"

	"github.com/LouisHatton/static-site-cms/pkg/kv"
)

func TestInterface(t *testing.T) {
	os.Setenv("KV_SQLITE_FILE_PATH", "./")
	connection, _ := kv.Connect("test")
	store = &connection

	type testingStruct struct {
		TestValue string `json:"testValue"`
		Other     bool   `json:"other"`
	}
	val := testingStruct{
		TestValue: "Louis",
		Other:     true,
	}

	store.Set("Key", val)

	var newVal testingStruct
	if err := store.Get("Key", &newVal); err != nil {
		t.Errorf("failed to get result: " + err.Error())
	}

	if newVal.TestValue != val.TestValue || newVal.Other != val.Other {
		t.Errorf("get restult does not match")
	}

	// os.Remove("database.sqlite")
}
