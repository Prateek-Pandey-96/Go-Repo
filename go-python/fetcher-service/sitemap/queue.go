package sitemap

type Queue struct {
	arr []string
}

func (q *Queue) Enqueue(val string) {
	q.arr = append(q.arr, val)
}

func (q *Queue) Dequeue() string {
	l := len(q.arr)
	if l == 0 {
		return "empty"
	}
	element := q.arr[0]
	if l == 1 {
		q.arr = []string{}
		return element
	}
	q.arr = q.arr[1:]
	return element
}
