package main

import (
	"fmt"
	// "log"
	// "net/rpc"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type BB struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func transfer(destinationId, money string) {
	args := &struct {
		DestinationId string
		Money         string
	}{
		destinationId,
		money,
	}
	jsonStr, _ := json.Marshal(args)
	req, _ := http.NewRequest("GET", "http://10.13.74.212:1234/Transfer", bytes.NewBuffer(jsonStr))
	cookie := http.Cookie{Name: "gosessionid", Value: "27338caf308389e05c0fb57cfa0548c6", Expires: time.Now().Add(111 * time.Second)}
	req.AddCookie(&cookie)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64")
	resp, _ := (&http.Client{}).Do(req)
	for _, v := range resp.Cookies() {
		fmt.Printf("%v\t%v\n", v.Name, v.Value)
	}
	s, _ := ioutil.ReadAll(resp.Body)
	j := make(map[string]interface{})
	json.Unmarshal(jsonStr, &j)
	for i, j := range j {
		fmt.Println(i, j)
	}
	fmt.Println(string(s))
	fmt.Println("finish")
}
func Balance() {
	fmt.Println("balance start...")
	// jsonStr, _ := json.Marshal(args)
	req, _ := http.NewRequest("GET", "http://139.199.59.7:1234/Balance", nil)
	cookie := http.Cookie{Name: "gosessionid", Value: "0e2a7a421d1bc37caeb09e4df2a56a0c", Expires: time.Now().Add(111 * time.Second)}
	req.AddCookie(&cookie)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64")
	resp, _ := (&http.Client{}).Do(req)
	for _, v := range resp.Cookies() {
		fmt.Printf("%v\t%v\n", v.Name, v.Value)
	}
	s, _ := ioutil.ReadAll(resp.Body)
	// j := make(map[string]interface{})
	// json.Unmarshal(jsonStr, &j)
	// for i, j := range j {
	// 	fmt.Println(i, j)
	// }
	fmt.Println(string(s))
	fmt.Println("finish")
}
func DeleteAccount(DestinationId string) {
	args := &struct {
		DestinationId string
	}{
		DestinationId,
	}
	jsonStr, _ := json.Marshal(args)
	req, _ := http.NewRequest("GET", "http://139.199.59.7:1234/DeleteAccount", bytes.NewBuffer(jsonStr))
	cookie := http.Cookie{Name: "gosessionid", Value: "e7968e3bd26e0503ed2df8167cde224f", Expires: time.Now().Add(111 * time.Second)}
	req.AddCookie(&cookie)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64")
	resp, _ := (&http.Client{}).Do(req)
	for _, v := range resp.Cookies() {
		fmt.Printf("%v\t%v\n", v.Name, v.Value)
	}
	s, _ := ioutil.ReadAll(resp.Body)
	j := make(map[string]interface{})
	json.Unmarshal(jsonStr, &j)
	for i, j := range j {
		fmt.Println(i, j)
	}
	fmt.Println(string(s))
	fmt.Println("finish")
}
func Login(UserName, Passwd string) {
	args := &struct {
		UserName string
		Passwd   string
	}{
		UserName,
		Passwd,
	}
	jsonStr, _ := json.Marshal(args)
	req, _ := http.NewRequest("POST", "http://139.199.59.7:1234/Login", bytes.NewBuffer(jsonStr))
	cookie := http.Cookie{Name: "gosessionid", Value: "0e2a7a421d1bc37caeb09e4df2a56a0c", Expires: time.Now().Add(111 * time.Second)}
	req.AddCookie(&cookie)
	req.Header.Add("Accept", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64")
	resp, _ := (&http.Client{}).Do(req)
	for _, v := range resp.Cookies() {
		fmt.Printf("%v\t%v\n", v.Name, v.Value)
	}
	s, _ := ioutil.ReadAll(resp.Body)
	j := make(map[string]interface{})
	json.Unmarshal(jsonStr, &j)
	for i, j := range j {
		fmt.Println(i, j)
	}
	fmt.Println(string(s))
	fmt.Println("finish")
}
func OpenAccount(UserName, Passwd, Power string) {
	args := &struct {
		UserName string
		Passwd   string
		Power    string
	}{
		UserName,
		Passwd,
		Power,
	}
	jsonStr, _ := json.Marshal(args)
	req, _ := http.NewRequest("GET", "http://139.199.59.7:1234/OpenAccount", bytes.NewBuffer(jsonStr))
	cookie := http.Cookie{Name: "gosessionid", Value: "e7968e3bd26e05032df8167cde224f", Expires: time.Now().Add(111 * time.Second)}
	req.AddCookie(&cookie)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64")
	resp, _ := (&http.Client{}).Do(req)
	for _, v := range resp.Cookies() {
		fmt.Printf("%v\t%v\n", v.Name, v.Value)
	}
	s, _ := ioutil.ReadAll(resp.Body)
	j := make(map[string]interface{})
	json.Unmarshal(jsonStr, &j)
	for i, j := range j {
		fmt.Println(i, j)
	}
	fmt.Println(string(s))
	fmt.Println("finish")
}
func Logout() {
	fmt.Println("balance start...")
	// jsonStr, _ := json.Marshal(args)
	req, _ := http.NewRequest("GET", "http://139.199.59.7:1234/Logout", nil)
	cookie := http.Cookie{Name: "gosessionid", Value: "0e2a7a421d1bc37caeb09e4df2a56a0c", Expires: time.Now().Add(111 * time.Second)}
	req.AddCookie(&cookie)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64")
	resp, _ := (&http.Client{}).Do(req)
	for _, v := range resp.Cookies() {
		fmt.Printf("%v\t%v\n", v.Name, v.Value)
	}
	s, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(s))
}
func main() {
	// go OpenAccount("Helalso", "ssasas", "1")
	// go Login("admin", "admin")
	// go DeleteAccount("8")
	// go Logout()
	go Balance()
	// go transfer("4", "2")
	// go transfer("6", "10")
	// fmt.Print("hello")
	for {

	}
	// conn, err := net.Dial("tcp", "139.199.59.7:777")
	// if err != nil {
	// 	fmt.Println("tcp erro")
	// }
	// conn.Write([]byte("hello tcp sockersajdkahsldjasldjalsdjlkasjdlkasjd;laks;dka;skda;s"))
	// fmt.Println("finish")
	// client, err := rpc.DialHTTP("tcp", "139.199.59.7"+":1234")
	// if err != nil {
	// 	log.Fatal("dialing:", err)
	// }
	// args := BB{5, 10}
	// var reply int
	// err = client.Call("Arith.Multiply", args, &reply)
	// if err != nil {
	// 	log.Fatal("arith error:", err)
	// }
	// fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
	// resp, _ := http.Get("http://139.199.59.7:12345/")
	// resp, _ = http.Get("http://139.199.59.7:12345/getSession")
	// args := &struct {
	// 	UserName   string
	// 	Passwd     string
	// 	Power      string
	// 	OperatorId string
	// }{
	// 	"sadaasd",
	// 	"asdassadsd",
	// 	"1",
	// 	"1",
	// }
	// jsonStr, _ := json.Marshal(args)
	// req, _ := http.NewRequest("GET", "http://139.199.59.7:1234", bytes.NewBuffer(jsonStr))
	// cookie := http.Cookie{Name: "gosessionid", Value: "e343c9321a18018722e2b9c05163f817", Expires: time.Now().Add(111 * time.Second)}
	// req.AddCookie(&cookie)
	// req.Header.Add("Accept", "application/json")
	// req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	// req.Header.Add("Connection", "keep-alive")
	// req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64")
	// resp, _ := (&http.Client{}).Do(req)
	// for _, v := range resp.Cookies() {
	// 	fmt.Printf("%v\t%v\n", v.Name, v.Value)
	// }
	// s, _ := ioutil.ReadAll(resp.Body)
	// j := make(map[string]interface{})
	// json.Unmarshal(jsonStr, &j)
	// for i, j := range j {
	// 	fmt.Println(i, j)
	// }
	// fmt.Println(string(s))
	// fmt.Println("finish")
}
