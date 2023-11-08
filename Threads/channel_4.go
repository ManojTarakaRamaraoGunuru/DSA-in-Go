package main

import (
	"fmt";
	"sync";
	"time";
)

func checker(ch chan int, i int){
	time.Sleep(10 * time.Second)
	ch <- i
}

func main(){
	nThreads := 5
	var wg sync.WaitGroup
	
	for i:=0; i<nThreads; i++{
		wg.Add(1)
		go func(i int){
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Task 1 executed")

	ch := make(chan int)
	
	for i:= 0; i<nThreads; i++{
		go checker(ch, i)
	}
	
	for i:=0; i<nThreads; i++{
		// We are actually blocking exactly like waitGroup till some sender places a variable in the queue
		// This is equivalent to task1, 
		// Using condition variables and mutexes are more recommended than using channels, buffered channels are more prone to code errors
		fmt.Println(<-ch)
	}
	fmt.Println("Task 2 executed")
}