package kv_test

import (
	"os"
	"testing"
	"time"

	"github.com/LouisHatton/static-site-cms/pkg/kv"
)

func TestInterface(t *testing.T) {
	os.Setenv("KV_SQLITE_FILE_PATH", "./")
	connection, _ := kv.Connect("test")
	store = &connection

	type testingStruct struct {
		RoutePath   string    `json:"routePath"`
		IsPublic    bool      `json:"isPublic"`
		LastUpdated time.Time `json:"lastUpdated"`
	}
	val := testingStruct{
		RoutePath:   "/app/login",
		IsPublic:    true,
		LastUpdated: time.Now(),
	}

	store.Set("Key", val)

	var newVal testingStruct
	if err := store.Get("Key", &newVal); err != nil {
		t.Errorf("failed to get result: " + err.Error())
	}

	if newVal.RoutePath != val.RoutePath {
		t.Errorf("get result 'string' does not match")
	}
	if newVal.IsPublic != val.IsPublic {
		t.Errorf("get result 'bool' does not match")
	}
	// This checks if two time values match down to the nanosecond .
	//
	// This is done because Time values can contain a monotonic field which is not saved during marshalling and
	// using the .Equal() function would cause the test to fail.
	if val.LastUpdated.Sub(newVal.LastUpdated).Abs().Nanoseconds() != 0 {
		t.Errorf("get result 'time' does not match, value:%s, newValue: %s", val.LastUpdated, newVal.LastUpdated)
	}
}
