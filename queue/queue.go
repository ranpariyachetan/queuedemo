package queue

func NewQueue() *Queue {
	return &Queue{}
}

type Queue struct {
	items []interface{}
}

func (queue *Queue) Enqueue(item interface{}) {
	queue.items = append(queue.items, item)
}

func (queue *Queue) Dequeue() (item interface{}, ok bool) {
	if len(queue.items) == 0 {
		return nil, false
	}

	item = queue.items[0]
	ok = true

	queue.items = queue.items[1:]

	return item, ok
}

func (queue *Queue) IsEmpty() bool {
	return len(queue.items) == 0
}

func (queue *Queue) Peek() interface{} {
	if len(queue.items) == 0 {
		return nil
	}

	return queue.items[0]
}
