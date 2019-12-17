package main

import (
	_ "context"
	// "encoding/json"
	"fmt"
	"net"
	"sync"

	//	"strconv"
	//	"even"
	// "reflect"
	"time"
	//	"even"
)

type Token struct {
	//	id  int
	Name string
	Node
}
type Node struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type foo int
type F interface {
	//	goa(i int)
	Hell()
}

func (a *Node) Goa(i int) {
	a.Id = i
}
func (a Node) Hell() {
	fmt.Println("hello")
}

var ci chan int
var cp chan string

// var a int
func Handle(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 10)
	fmt.Printf("request IP = %v\n", conn.RemoteAddr())
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		if n == 0 {
			break
		}
		fmt.Println("messsage:", string(buf))
	}
}

type Lock struct {
	lock *sync.RWMutex
}

var m map[string]Lock
var a int

func addA() {
	defer func() {
		return
	}()
	defer func() {
		a++
	}()
	a++
}
func lockT(i int) {
	// var lock sync.RWMutex
	lock := m["a"]
	lock.lock.Lock()
	fmt.Printf("%p\n", lock.lock)
	fmt.Println("hello", i)
	for {

	}
	lock.lock.Unlock()
}
func main() {
	var a Node
	a.Id = 10
	a.Name = "ss"
	b := &a
	b.Id = 20
	fmt.Println(a)
	// node := Node{Id: 10, Name: "nihao"}
	// b, _ := json.Marshal(node)
	// fmt.Printf("%s\n", b)
	// c := make(map[string]interface{})
	// json.Unmarshal(b, &c)
	// for key, value := range c {
	// 	fmt.Println(key, value)
	// }
	// v := reflect.ValueOf(node)
	// t := reflect.TypeOf(&node)
	// fie := v.Field(0).Type()
	// tag := t.Elem().Field(0).Tag
	// fmt.Println(tag.Get("json"))
	// ctx := context.Background()
	// ctx = context.WithValue(ctx, "Name", "hel")
	// CtxText(ctx)
	//	h()
	//	fmt.Println(even.A)
	//	addA()
	//	println(a)
	// even.Say_Hello()
	// var a Token
	// a.Name = "as"
	// a.Node.Id = 10
	// a.Node.Name = "sda"
	// fmt.Println(a)
	// a = 1
	//	fmt.Println("a=", a)
	//	addA()
	//	fmt.Println("a= ", a)
	// m = make(map[string]Lock)
	// var lock sync.RWMutex
	// m["a"] = Lock{&lock}
	// go lockT(1)
	// go lockT(2)
	// for {

	// }
	fmt.Println(time.Now().Unix())
	// m := make(map[string]Lock)
	// if v, ok := m["a"]; ok {
	// 	fmt.Println("1", v)
	// }
	// if v, ok := m["a"]; ok {
	// 	fmt.Println("1", v)
	// }
	// fmt.Println("2", m["a"])
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }
	// fmt.Println(time.Now().Unix() + 20)
	// lock := Lock{}
	// fmt.Println(lock.lock)
	// fmt.Println(lock)
	// var lock sync.RWMutex
	// c := lock
	// lock.Lock()
	// c.Lock()
	// fmt.Println("sad")
	// c.Unlock()
	// lock.Unlock()
	// a := &Lock{lock}
	// lock.Lock()
	// fmt.Printf("%v\n", lock)
	// go a.lockT(1)
	// go a.lockT(2)
	// go a.lockT(3)
	// for {
	// }
	// lock.Unlock()
	//	ln, err := net.Listen("tcp", "0.0.0.0:777")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	for {
	//		fmt.Println("start listening ...")
	//		conn, err := ln.Accept()
	//		if err != nil {
	//			fmt.Println("erro")
	//		}
	//		go Handle(conn)
	//	}
	//	var a int; var b int
	//	var(
	//		a int
	//		b int
	//	)
	//	const(
	//		c = iota
	//		d = iota
	//	)
	//	var(
	//		c int = 10
	//		d = 20
	//	)
	//	var a = [...]int{1, 2, 3}
	//	a := [...]int{1, 2, 3}
	//	a = [...]int{1, 2, 3}
	//	b := a[0:2]
	//	s1 := append(b, 2, 3)
	//	s2 := append(b, 4)
	//	a[0] = 4
	//	s1[0] = 5
	//	b[1] = 4
	//	a[0] = 10
	//	 a := [...]int{1, 2, 3}
	//	 hello(a[:])
	//	dict := map[string]int{
	//		"hello": 1, "nihao": 2,
	//	}
	//	value, present := dict["asdk"]
	//	fmt.Println(value, present)
	//	for key, values := range(dict){
	//		fmt.Println("key=", key, "  values=", values)
	//	}
	//	 s := "利好"
	//	 s1 := []rune(s)
	//	 s1[0] = '我'
	//	 s = string(s1)
	//	 fmt.Printf("%T", s1[0])
	//	i := 1
	//
	//		fmt.Println(i)
	//		i++
	//		if i < 11 {
	//			goto H
	//		}
	//	const a int = 10
	//	a := make(map[string]int)
	//	a["hello"] = 1
	//	a["hi"] = 2
	//	a[0] = 1
	//	a = append(a, 2)
	//	for _, i := range a{
	//		println(i)
	//	}
	//	a := [...]int{1, 2, 3}
	//	println(throwsPanic(a[:]))
	//	b := a[:]
	//	change(a[:]...)
	//	print(b...)
	//	println(*pointer_text())
	//	var pointer *[3]int
	//	a := [...]int{1, 2, 3}
	//	pointer = &a
	//	fmt.Printf("%T", pointer)
	//	a := func(i int)int{
	//		var a = 1
	//		fmt.Println("hello")
	//		return a + i
	//	}
	//	even.Say_Hello()
	//	say()
	//	a := Token{id: 1, Name: "nihao"}
	//	println(a.id, a.Name)
	//	a := make([]int, 3)
	//	a := [3]int{1, 2, 3}
	////	a[0] = 1
	//	p := &a
	//	c := *p
	//	c[0] = 2
	//	fmt.Printf("%d", (*p)[0])
	//	throwsPanic(1)
	//	var a Token
	//	a := [...]int{1, 2, 3}
	//	a := new(Token)
	//	a.id = 2
	//	token_text(a)
	//	var a Token
	//	a.Id = 10
	//	fmt.Printf("%v\n", a)
	//	println(a)
	//	var a *Node = new(Node)
	//	a.Id = 10
	//	a.Name = new(string)
	//	*(*a).Name = "asd"
	//	var b F
	//	b = a
	//	b.goa()
	//	_, ok := b.(F)
	//	a.say()
	//	var b Node
	//	b.Id = 10
	//	b.Ia = 5
	//	b.Name = "sa"
	//	b.Node = &Node{11, "3asd2", 32, nil}
	//	var f F
	//	f = &b
	//	interface_test(f)
	//	fmt.Printf("%T", b)
	//	b := Token{Name: "adsad"}
	//	b.Id = 10
	//	b.Node.Name = "asd"
	//	b = new(Node)
	//	*a = &string("hello")
	//	t := reflect.TypeOf(f)
	//	v := reflect.ValueOf(f)
	//	f := v.Interface().(*Node)
	//	fmt.Printf("%v", t.NumMethod())
	//	param := make([]reflect.Value, 1)
	//	param[0] = reflect.ValueOf(13)
	//	method := v.MethodByName("Goa")
	//	method.Call(param)
	////
	//	fmt.Printf("%v", b)
	//	fmt.Println(t.Field(0).Type, v.Field(0).Interface())
	//	tag := t.Elem().Field(0).Name
	//	name := v.Elem().Field(1).String()
	//	str := "hello"
	//	v.Elem().Field(2).SetInt(10)
	//	b.goa(10)
	//	p := v.Interface().(*Node)
	//	fmt.Printf("%v", v.Kind())
	//	println(t)
	//	for index := 0; index < t.Elem().NumField(); index++ {
	//		field := t.Elem().Field(index)
	//		value := v.Elem().Field(index).Interface()
	//		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	//	}
	//	for index := 0; index < t.NumMethod(); index++ {
	//		m := t.Method(index)
	//		fmt.Printf("%s: %v\n", m.Name, m.Type)
	//	}
	//	ci = make(chan int, 10)
	//	cp = make(chan string, 10)
	//	go hello(1, 2, 3, 4, 6, 9)
	//	go change(5, 10, 42, 53, -1)
	//	i := 0
	//	var msg1 int
	//	var msg2 string
	//	msg1 := <- ci
	//	msg2 = <- cp
	//	i := 0
	//	go select_test()
	//	go change()
	// tick := time.NewTicker(1 * time.Second)
	// L:
	//	for i := 0; i < 10; i++ {
	//		select {
	//		case ci <- i:
	//			//				fmt.Println("first finish:", msg1)
	//		case cp <- strconv.Itoa(i):
	//			// fmt.Printf("%d: case <-tick.C\n", i)
	//			// default:
	//			// 	break L
	//		}
	//		//		fmt.Println("loop=", i)
	//	}
	//	time.Sleep(2 * time.Second)
	// if msg1, ok := <- ci; ok {
	// 	fmt.Println(msg1)
	// }
	// if msg2, ok := <- cp; ok {
	// 	fmt.Println(msg2)
	// }
	//	fmt.Println(msg1)
	//	fmt.Println(msg2)
	//	println("finish")
	//	var a Token = Token{"nihao", Node{10, "asd"}}
	//	t := reflect.TypeOf(a)
	//	v := reflect.ValueOf(a)
	//	fi := v.Field(1)
	////	v :=
	//	fmt.Printf("%v\n", fi.NumMethod())
	//	text_defer()
	//	println("sada")
}
func interface_test(any interface{}) {
	if t, ok := any.(F); ok {
		t.Hell()
	}
}
func select_test() {
	for {
		fmt.Println("val=", <-ci)
	}
}
func hello(args ...int) {
	for _, i := range args {
		println(i)
	}
	fmt.Printf("%v\n", args)
	cp <- "sas"
}

func token_text() {
	fmt.Println("hello")
}
func change() {
	for {
		fmt.Println("str=", <-cp)
	}
}
func text_defer() (i string) {
	//	i := 0
	defer func(x string) {
		i += x
		println(x)
	}("nihao")
	println("hello")
	return "2"
}
func pointer_text() *int {
	var i = 1
	defer func(a *int) {
		*a++
	}(&i)
	return &i
}
func throwsPanic(a []int) (b bool) {
	defer func() {
		if x := recover(); x == nil {
			fmt.Println("no panic", x)
			b = true
		} else {
			fmt.Println("have panic", x)
			b = false
		}
	}()
	//	fmt.Println(a[10])
	//	panic("out of range")
	//	println("hhh")
	return
}
