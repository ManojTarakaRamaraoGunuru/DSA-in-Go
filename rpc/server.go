package main

import (
	"net"
	"net/http"
	"net/rpc"
	"log"
)

type Args struct{
	A, B int
}
type Arith int


func(this *Arith) Add (args Args, reply *int)error{
	*reply = args.A + args.B
	return nil
}

func(this *Arith) Subtract (args Args, reply *int) error{
	*reply = args.A - args.B
	return nil
}

func(this *Arith) Multiply (args Args, reply *int) error{
	*reply = args.A * args.B
	return nil
}

func(this *Arith) Divide (args Args, reply *int) error{
	*reply = args.A / args.B
	return nil
}

func main(){
	myArith := new(Arith)
	
	rpc.Register(myArith)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(l, nil)

}

