package main

import (
	"fmt";
	"time";
	"sync";
)

// Alice and bob keeps doing transactions

func main(){

	var mu sync.Mutex
	alice, bob := 10000, 10000

	for i:=0; i<1000; i++{
		go func(){
			time.Sleep(20 * time.Millisecond)
			mu.Lock()
			defer mu.Unlock()
			alice -= 1
			bob += 1
		}()
	}

	for i:= 0; i < 1000; i++{
		go func(){
			time.Sleep(20 * time.Millisecond)
			mu.Lock()
			defer mu.Unlock()
			alice += 1
			bob -= 1
		}()
	}

	start := time.Now()
	for time.Since(start)< 1*time.Second {
		go func(){
			mu.Lock()
			defer mu.Unlock()
			if alice+bob != 20000{
				fmt.Println("Issue in transactions")
			}
		}()
	}
}