package kv

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func connect() *sql.DB {
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

	return sqlite
}

func sqlDb() *sql.DB {
	return connect()
}
