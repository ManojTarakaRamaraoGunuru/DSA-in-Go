package main
import (
	"fmt"
	"math/rand"
)

// Implemented in a doubly linked list fashion, so all operations will be O(1)

type q_node struct{
	data int
	next *q_node
	pre *q_node
}

type queue struct{
	head *q_node
	tail *q_node
	size int
}

func(q *queue)enque(val int){
	q.size += 1
	temp:= q_node{val, nil, nil}
	if q.head == nil {
		q.head = &temp
		q.tail = &temp
		return
	}
	temp.next = q.head
	q.head.pre = &temp
	q.head = &temp
}

func (q *queue)deque()int{
	if q.size == 0 {
		fmt.Println("queue is empty")
		return -1;
	}
	discarderd_val := q.tail.data
	if q.size == 1{
		q.head = nil
		q.tail = nil
		q.size = 0
		return discarderd_val
	}
	q.size -=1
	temp := q.tail.pre
	temp.next = nil
	q.tail.pre = nil
	q.tail = temp
	return discarderd_val
}

func (q *queue)top() int{
	if q.size == 0 {
		fmt.Println("queue is empty")
		return -1;
	}
	return q.head.data
}


func main(){
	q:= queue{size:0}
	for i:= 0; i<10; i++ {
		t:= rand.Intn(100)
		q.enque(t)
	}
	for i:=0; i<10; i++ {
		fmt.Println(q.deque())
	}
	fmt.Println(q.top())
}