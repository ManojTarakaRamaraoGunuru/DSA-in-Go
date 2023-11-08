package main

import "fmt"
import "math/rand"
import "sync"
import "time"


// We are conducting election

func RequestVote()bool{
	x := rand.Intn(100)
	fmt.Println(x)
	return x%2 == 0
}

func main(){
	cnt, finished := 0, 0
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	rand.Seed(time.Now().UnixNano())

	for i:=0; i< 10; i++ {
		go func(){
			vote := RequestVote()

			mu.Lock()
			defer mu.Unlock()

			if (vote){
				cnt++;
			}
			finished++;
			cond.Broadcast()   // We do broadcast when we change shared variables and we broadcst by holding the lock and we release the lock
		}()
	}
	mu.Lock()
	
	for cnt < 5 && finished != 10{
		cond.Wait()  // when thread sleeps here it releases lock atomically so the other threads can avoid busy waiting
					 // when threads got awake again it will acquire the lock and checks the condition
	}
	if (cnt >= 5){
		fmt.Println("Won Election")
	}else{
		fmt.Println("Lost")
	}

	mu.Unlock()
	
}