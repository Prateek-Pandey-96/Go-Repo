package main

type Queue struct {
	Q []string
}

func (q *Queue) Enqueue(val string) {
	q.Q = append(q.Q, val)
}

func (q *Queue) Dequeue() string {
	l := len(q.Q)
	if l == 0 {
		return "empty"
	}
	element := q.Q[0]
	if l == 1 {
		q.Q = []string{}
		return element
	}
	q.Q = q.Q[1:]
	return element
}
