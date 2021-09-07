package main

import "fmt"

//Queue Struct representing the object Queue
type Queue struct {
	items []int
}

//Enqueue method to add to queue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

//Dequeue to remove from the Queue
func (q *Queue) Dequeue() int{
	toDequeue := q.items[0] //returns a single value
	q.items = q.items[1:]  //returns a slice
	return toDequeue
}

func main(){
	myQueue := Queue{}
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	myQueue.Enqueue(500)
	myQueue.Enqueue(600)
	myQueue.Enqueue(700)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println("new",myQueue)
}