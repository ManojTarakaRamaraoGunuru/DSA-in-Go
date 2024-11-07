package main

import (
	"net/rpc"
	"fmt"
	"log"
)

type Args struct{
	A, B int
}

func main(){
	client, err := rpc.DialHTTP("tcp", "0.0.0.0:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := Args{4,2}
	var reply int

	err = client.Call("Arith.Divide", args, &reply)

	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println(reply)
}