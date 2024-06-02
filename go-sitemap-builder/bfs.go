package main

type empty struct{}

func BFS(urlFlag string, depth int) []string {
	queue := Queue{Q: []string{}}
	visited := make(map[string]empty)

	queue.Enqueue(urlFlag)
	visited[urlFlag] = empty{}

	level := -1

	for len(queue.Q) != 0 {
		size := len(queue.Q)
		level++
		if level == depth {
			break
		}
		for i := 0; i < size; i++ {
			from_link := queue.Q[0]
			queue.Dequeue()

			to_links := GetLinks(from_link)

			for _, to_link := range to_links {
				if _, ok := visited[to_link]; ok {
					continue
				}
				visited[to_link] = empty{}
				queue.Enqueue(to_link)
			}

		}
	}

	pages := make([]string, len(visited))
	i := 0
	for page := range visited {
		pages[i] = page
		i++
	}
	return pages
}
