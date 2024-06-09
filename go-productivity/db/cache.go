package db

import (
	"fmt"
	"os"

	"github.com/boltdb/bolt"
)

func Put(key string, val string) error {
	return DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(CacheBucket))
		return b.Put([]byte(key), []byte(val))
	})
}

func Get(key string) (string, error) {
	var val string
	err := DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(CacheBucket))
		val = string(b.Get([]byte(key)))
		return nil
	})
	if err != nil {
		fmt.Println("Error occured while getting key: ", key)
		os.Exit(1)
	}
	return val, nil
}

func Del(key string) error {
	return DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(CacheBucket))
		return b.Delete([]byte(key))
	})
}
