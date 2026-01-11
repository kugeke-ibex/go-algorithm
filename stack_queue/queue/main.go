package main

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-queue/queue"
	"github.com/golang-queue/queue/job"
)

func main() {
	q := Queue{}
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	fmt.Println(q.queue)
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	q.enqueue(5)
	q.enqueue(6)
	q.enqueue(7)
	q.enqueue(8)
	q.enqueue(9)
	q.enqueue(10)
	fmt.Println(q.queue)

	q.reverse()
	fmt.Println(q.queue)

	taskN := 100
	resultChan := make(chan int, taskN)
	queue := queue.NewPool(10)
	defer queue.Release()

	for index := 0; index < taskN; index++ {
		idx := index
		var taskFunc job.TaskFunc = func(ctx context.Context) error {
			time.Sleep(time.Second * 1)
			resultChan <- idx
			return nil
		}
		if err := queue.QueueTask(taskFunc); err != nil {
			fmt.Println("Error queuing task:", err)
		}
	}

	for i := 0; i < taskN; i++ {
		fmt.Println("Processing result:", <-resultChan)
	}
}

type Queue struct {
	queue []any
}

func (q *Queue) enqueue(value any) {
	q.queue = append(q.queue, value)
}

func (q *Queue) dequeue() any {
	if len(q.queue) != 0 {
		value := q.queue[0]
		q.queue = q.queue[1:]
		return value
	}

	return nil
}

func (q *Queue) reverse() {
	newQueue := make([]any, 0, len(q.queue))
	for index := len(q.queue) - 1; index >= 0; index-- {
		newQueue = append(newQueue, q.queue[index])
	}

	q.queue = newQueue
}
