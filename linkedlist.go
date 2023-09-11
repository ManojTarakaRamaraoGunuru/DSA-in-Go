package main

import "fmt"

type node struct{
	data int
	next *node
}

type linkedlist struct{
	head *node
}

// Note that the receive over here exactly points to the instance of the class.

func (ll *linkedlist)insert(val int){
	temp := node{val, nil}
	if(ll.head == nil){
		ll.head = &temp
		return
	}
	temp.next = ll.head
	ll.head = &temp
}

func (ll *linkedlist)insert_at_end(val int){
	temp:= node{val,nil}

	if(ll.head == nil){				// Note that condition is not ll == nil but ll.head == nil
		ll.head = &temp
		return
	}
	cur_last := ll.head
	for cur_last.next!= nil{
		cur_last = cur_last.next
	}
	cur_last.next = &temp
}

func (ll *linkedlist)printer(){
	t := ll.head
	for t!=nil{
		fmt.Print(t.data)
		fmt.Print(" ")
		t = t.next
	}
}

func main(){
	fmt.Println("ok")
	ll := linkedlist{}

	for i:=0; i<5; i++{
		ll.insert_at_end(i)
		// ll.printer()
	}
	ll.printer()
}