package main

//	"fmt"
import (
	// "even"
	"context"
	"fmt"
)

type bb struct {
	s int
}

func h() {
	// fmt.Println(even.A)
	// even.A.Name = "asd"
}
func CtxText(ctx context.Context) {
	fmt.Println(ctx.Value("Name"))
}
func (a *Node) say() {
	//	str := "aa"
	//	a.Name = &str
}
