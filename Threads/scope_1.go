package main

import "fmt"
import "time"
import "sync"

func main(){
	var a string
	var wg sync.WaitGroup

	// a is declared out but still can be accessable inside the go routine. This is the closure property of go threads.

	wg.Add(1)
	go func(){
		a = "Hello World!"
		fmt.Println(a)
		wg.Done()
	}()
	wg.Wait()
	fmt.Print("Task 1 exited\n")

	// Note that all threads will point to the i variable in their closure, since i changes by the main go routine
	// all threads will print unexpected values, to avoid this pass i as an attribute as shown in task 3
	for i:= 0; i< 5; i++ {
		wg.Add(1)
		go func(){
			time.Sleep(1 * time.Second)
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Task 2 exited")

	for i:= 0; i< 5; i++ {
		wg.Add(1)
		go func(x int){
			time.Sleep(1 * time.Second)
			fmt.Println(x)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Task 3 exited")
}