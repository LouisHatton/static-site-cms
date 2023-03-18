package kv

import (
	"fmt"
)

// Forms and checks the connection to the database and returns a kvStore object
func Connect(tableName string) (KvStore, error) {
	connection := sqlDb()
	err := connection.Ping()
	if err != nil {
		return KvStore{}, fmt.Errorf("failed to ping database connection: " + err.Error())
	}

	store := KvStore{
		db:        connection,
		tableName: tableName,
	}

	if err := store.createTable(); err != nil {
		return KvStore{}, fmt.Errorf("failed to create table connection: " + err.Error())
	}
	return store, nil
}

// Closes the connection to the underlying database
func (kv *KvStore) Close() error {
	return kv.db.Close()
}

func (kv *KvStore) createTable() error {
	_, err := kv.db.Exec("CREATE TABLE IF NOT EXISTS " + kv.tableName + " (name TEXT NOT NULL PRIMARY KEY, value TEXT)")
	if err != nil {
		return err
	}

	kv.db.Exec("CREATE INDEX name_index ON " + kv.tableName + " (name)")
	return nil
}
