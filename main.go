package main

import (
	"fmt"

	"github.com/ranpariyachetan/simplequeue/queue"
)

func main() {
	q := queue.NewQueue()

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Enqueue(50)
	q.Enqueue(60)

	for q.Peek() != nil {
		i, ok := q.Dequeue()
		fmt.Printf("%d : %v\n", i, ok)
	}
}
