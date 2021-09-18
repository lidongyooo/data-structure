package queue

type queueContract interface {
	push()
	pop()
}

type queue struct {
	length int
	items []string
}

func Init() *queue {
	return &queue{
		length: 0,
	}
}

func (q *queue) IsEmpty() bool  {
	return q.length <= 0
}

func (q *queue) Push(values ...string)  {
	q.length += len(values)
	q.items = append(q.items, values...)
}

func (q *queue) Pop() string {
	if q.length > 0 {
		val := q.items[0]
		q.items = q.items[1:]
		q.length--
		return val
	}

	return ""
}