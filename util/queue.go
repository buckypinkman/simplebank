package util

type Queue struct {
	items []string
}

func (q *Queue) Enqueue(item string) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (string, bool) {
	if len(q.items) == 0 {
		return "", false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// func main() {
// 	q := Queue{}
// 	q.Enqueue("Task Data")
// 	task, _ := q.Dequeue()
// 	fmt.Println("Processing", task)
// }
