package kv

import (
	"database/sql"
	"encoding/json"
)

type KvStore struct {
	db        *sql.DB
	tableName string
}

type KvItem struct {
	Key   string
	Value string
}

// Internal get function returning a KvItem of key and value
func (kv *KvStore) get(key string) (*KvItem, error) {
	item := new(KvItem)
	query := "SELECT name, value FROM " + kv.tableName + " WHERE name = ?"
	err := kv.db.QueryRow(query, key).Scan(&item.Key, &item.Value)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return item, nil
}

// Internal set function which takes KvItem
func (kv *KvStore) set(item KvItem) error {
	row, err := kv.get(item.Key)
	var query string

	if row == nil && err == nil {
		query = "INSERT INTO " + kv.tableName + " (value, name) VALUES( ? , ? )"
	} else {
		query = "UPDATE " + kv.tableName + " SET value = ? WHERE name = ? "
	}

	_, err = kv.db.Exec(query, item.Value, item.Key)
	return err
}

// Takes a key and a reference to an interface, fetches the key and unmarshal the JSON
// back into the interface provided.
func (kv *KvStore) Get(key string, value interface{}) error {
	item, err := kv.get(key)
	if err != nil {
		return err
	}

	if err = json.Unmarshal([]byte(item.Value), &value); err != nil {
		return err
	}
	return nil
}

// Takes an interface and stores it in a JSON format in the db.
//
// Interfaces must have json tags for exported values or these will be lost on save.
func (kv *KvStore) Set(key string, value interface{}) error {
	encoded, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return kv.set(KvItem{Key: key, Value: string(encoded)})
}
