package db

import (
	"encoding/binary"
)

const TaskBucket = "to_do_tasks"
const FinishedTaskBucket = "finished_tasks"

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

func btoi(byteNumber []byte) uint64 {
	return binary.BigEndian.Uint64(byteNumber)
}
