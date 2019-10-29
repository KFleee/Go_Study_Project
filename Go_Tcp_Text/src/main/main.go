package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type BB struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	// conn, err := net.Dial("tcp", "localhost:777")
	// if err != nil {
	// 	fmt.Println("tcp erro")
	// }
	// conn.Write([]byte("hello tcp sockersajdkahsldjasldjalsdjlkasjdlkasjd;laks;dka;skda;s"))
	// fmt.Println("finish")
	client, err := rpc.DialHTTP("tcp", "localhost"+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := BB{5, 10}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}
