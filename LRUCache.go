package main

import "fmt"

type Node struct{
    key int
    val int
    next *Node
    prev *Node
}

func(this *Node)insert(temp *Node){
    temp.next = this.next
    this.next.prev = temp
    temp.prev = this
    this.next = temp
}

func (this *Node)del(m map[int]*Node){
    temp := this.prev.prev
    delete(m, this.prev.key)
    temp.next = this
    this.prev = temp
}

type LRUCache struct {
    head *Node
    tail *Node
    size int
    capacity int
    m map[int]*Node
}


func Constructor(capacity int) LRUCache {
    lrucache := LRUCache{}
    lrucache.capacity = capacity
    lrucache.size = 0
    lrucache.head = &Node{key:-1, val:-1}
    lrucache.tail = &Node{key:-1, val:-1}
    lrucache.head.next = lrucache.tail
    lrucache.tail.prev = lrucache.head
    lrucache.m = make(map[int]*Node)
    return lrucache
}


func (this *LRUCache) Get(key int) int {
    temp, is_present := this.m[key]
    if !is_present{
        return -1
    }
    value := temp.val
    temp.next.del(this.m)
    this.size -= 1 
    this.Put(key, value)
    return value
}

func(this *LRUCache)printer(){
    temp := this.head
    for temp!=nil{
        fmt.Print(temp.key, " ")
        temp = temp.next
    }
}

func (this *LRUCache) Put(key int, value int)  {
    temp := Node{key:key, val:value}
    temp1, is_present := this.m[key] // to update already existing node
    if is_present{
        temp1.next.del(this.m)
        this.size -= 1
    }
    this.m[key] = &temp
    if this.size<this.capacity{
        this.size += 1
        this.head.insert(&temp)
    }else{
        this.head.insert(&temp)
        this.tail.del(this.m)
    }
}

func main(){
    capacity := 2
    myCache := Constructor(capacity);
    myCache.Put(1,1)
    myCache.Put(1,2)
    myCache.Put(3,3)
    fmt.Print(myCache.Get(1), " ")
    myCache.Put(2,8)
    fmt.Print(myCache.Get(3), " ")
    fmt.Print(myCache.Get(2), " ")
}
