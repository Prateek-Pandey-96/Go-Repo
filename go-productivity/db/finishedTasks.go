package db

import (
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
)

func FinishTask(task string) error {
	return DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(FinishedTaskBucket))
		id, _ := b.NextSequence()

		entry := FinishedTask{
			Timestamp: uint64(time.Now().Unix()),
			Title:     task,
		}
		value, err := json.Marshal(entry)
		if err != nil {
			return err
		}

		return b.Put(itob(id), value)
	})
}

func GetAllFinishedTasks() ([]Task, error) {
	tasks := []Task{}
	err := DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(FinishedTaskBucket))
		c := bucket.Cursor()
		cutoff := uint64(time.Now().Add(-24 * time.Hour).Unix())

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var finishedTask FinishedTask
			err := json.Unmarshal(v, &finishedTask)
			if err != nil {
				return err
			}
			if finishedTask.Timestamp >= cutoff {
				tasks = append(tasks, Task{Id: btoi(k), Title: string(finishedTask.Title)})
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
