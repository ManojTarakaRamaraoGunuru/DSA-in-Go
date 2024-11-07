package main

import ("fmt"
"math/rand"
)

type node struct{
	data int;
	next *node;
}

type stack struct{
	head *node
	size int
}

func (sta *stack)push(data int){
	sta.size += 1
	temp := node{data,nil}
	if sta.head == nil{
		sta.head = &temp
		return
	}
	temp.next = sta.head
	sta.head = &temp
}

func (sta *stack)pop(){
	if sta.size == 0 {
		return
	}
	sta.size -= 1
	sta.head = sta.head.next
}

func (sta *stack)top() int{
	if sta.size ==0 {
		fmt.Println("Stack has no elements")
		return -1;
	}
	return sta.head.data
}


func main(){

	sta := stack{size:0}
	for i:=0; i<10; i++{
		sta.push(rand.Intn(100));
	}
	for i := 0; i<11; i++{
		fmt.Println(sta.top())
		sta.pop();
	}
}