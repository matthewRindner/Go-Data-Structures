package main

import "fmt"

//Queue represetn a queue that holds a slice
type Queue struct {
	items []int
}

//Enqueue adds value at the end
func (q *Queue) enqueue(value int) {
	q.items = append(q.items, value)
}

//Enqueue removes value at the beginning
//RETURNS the removed value
func (q *Queue) dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}
func main() {
	myQueue := Queue{}
	myQueue.enqueue(1)
	myQueue.enqueue(2)
	myQueue.enqueue(3)
	myQueue.enqueue(4)
	fmt.Println(myQueue)
	res := myQueue.dequeue()
	fmt.Println(myQueue, "the dequeued value is: ", res)

}
