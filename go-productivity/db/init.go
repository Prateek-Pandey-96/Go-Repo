package db

import (
	"time"

	"github.com/boltdb/bolt"
)

var DB *bolt.DB

func Init(path string) error {
	db, err := bolt.Open(path+"tasks_db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	DB = db

	return db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(TaskBucket)); err != nil {
			return err
		}
		if _, err := tx.CreateBucketIfNotExists([]byte(FinishedTaskBucket)); err != nil {
			return err
		}
		return nil
	})
}
