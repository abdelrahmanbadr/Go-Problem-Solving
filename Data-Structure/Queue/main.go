package main

import "fmt"

type Queue struct {
	QueueArray []int
	Rear       int
	//Front      int
}

func (q *Queue) Empty() bool {
	return q.Rear == 0
}

func (q *Queue) Enqueue(v int) {
	q.QueueArray = append(q.QueueArray, v)
	q.Rear++
}

func (q *Queue) Dequeue() int {
	v := q.QueueArray[0]
	//remove first element in the queue
	q.QueueArray = q.QueueArray[1:q.Rear]
	q.Rear--
	return v
}

func main() {

	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Empty())
	// Output:
	// 1
	// 2
	// true
}
