package main

import (
	"net/rpc"
	"errors"
//	"net"
	"net/http"
	"fmt"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	fmt.Println("hsadjsad")
	*reply = args.A * args.B
	return nil
}
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	e := http.ListenAndServe(":1234", nil)
	if e != nil {
		fmt.Println("listen error:", e)
	}
//	go http.Serve(l, nil)
	fmt.Println("finish")
}