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

func (ll * linkedlist) delete(val int){
	if ll.head == nil{
		return
	}

	if ll.head.data == val{
		ll.head = ll.head.next
		return
	}

	temp:= ll.head

	for temp.next!=nil && temp.next.data!=val{
		temp = temp.next
	}
	
	if temp.next != nil{
		temp.next = temp.next.next
	}

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
	ll := linkedlist{}

	for i:=0; i<5; i++{
		ll.insert_at_end(i)
	}
	ll.printer()

	for i:=5; i>=3;i--{
		ll.delete(i)
	}
	fmt.Println(" ")
	ll.printer()
}