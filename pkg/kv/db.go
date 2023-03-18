package kv

import (
	"database/sql"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var once sync.Once
var connection *sql.DB

func connect() {
	fileName := os.Getenv("KV_SQLITE_FILE_NAME")
	if fileName == "" {
		fileName = "database.sqlite"
	}
	filePath := os.Getenv("KV_SQLITE_FILE_PATH")
	if filePath == "" {
		filePath = "/tmp/"
	}
	sqlite, err := sql.Open("sqlite3", filePath+fileName)
	if err != nil {
		panic(err)
	}

	connection = sqlite
}

func sqlDb() *sql.DB {
	once.Do(connect)
	return connection
}
