package db

import (
	"github.com/boltdb/bolt"
)

func AddTask(task string) error {
	return DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TaskBucket))
		id, _ := b.NextSequence()
		return b.Put(itob(id), []byte(task))
	})
}

func GetAllTasks() ([]Task, error) {
	tasks := []Task{}
	err := DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TaskBucket))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{Id: btoi(k), Title: string(v)})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(key uint64) error {
	return DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TaskBucket))
		return b.Delete(itob(key))
	})
}
