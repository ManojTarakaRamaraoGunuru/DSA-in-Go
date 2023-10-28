package main

import (
    "fmt";
    "sync";
    )


// Note that instead of using (done)sync.WaitGroup, we can use another channel that can send a message to main go routine to finish.
// It is a simple hack

var ch chan int 
// var fin chan int
var done sync.WaitGroup

func producer(){
    for i:=0; i<10; i++ {
        ch <- i
    }
    close(ch)
    done.Done()
}

func consumer(){
    for i := range(ch){
        fmt.Println(i)
    }
    // fin <- 1
    // close(fin)
    done.Done()
}

func main(){
    ch = make(chan int)
    // fin = make(chan int)
    done.Add(1)
    go producer()
    done.Add(1)
    go consumer()
    done.Wait()
    // for 1 != <- fin{}
}