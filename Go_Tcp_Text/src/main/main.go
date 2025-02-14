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
	req, _ := http.NewRequest("GET", "http://localhost:1234/Transfer", bytes.NewBuffer(jsonStr))
	cookie := http.Cookie{Name: "gosessionid", Value: "9f7225a3942e82151cd45a5c3d550b3f", Expires: time.Now().Add(111 * time.Second)}
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
	req, _ := http.NewRequest("GET", "http://localhost:1234/Balance", nil)
	cookie := http.Cookie{Name: "gosessionid", Value: "9f7225a3942e82151cd45a5c3d550b3f", Expires: time.Now().Add(111 * time.Second)}
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
	req, _ := http.NewRequest("GET", "http://localhost:1234/DeleteAccount", bytes.NewBuffer(jsonStr))
	cookie := http.Cookie{Name: "gosessionid", Value: "9f7225a3942e82151cd45a5c3d550b3f", Expires: time.Now().Add(111 * time.Second)}
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
func Login(UserId, Passwd string) {
	args := &struct {
		UserId string
		Passwd string
	}{
		UserId,
		Passwd,
	}
	jsonStr, _ := json.Marshal(args)
	req, _ := http.NewRequest("GET", "http://localhost:1234/Login", bytes.NewBuffer(jsonStr))
	cookie := http.Cookie{Name: "gosessionid", Value: "9f7225a3942e82151cd45a5c3d550b3f", Expires: time.Now().Add(111 * time.Second)}
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
func OpenAccount(Passwd, Power string) {
	args := &struct {
		Passwd string
		Power  string
	}{
		Passwd,
		Power,
	}
	jsonStr, _ := json.Marshal(args)
	req, _ := http.NewRequest("GET", "http://localhost:1234/OpenAccount", bytes.NewBuffer(jsonStr))
	cookie := http.Cookie{Name: "gosessionid", Value: "9f7225a3942e82151cd45a5c3d550b3f", Expires: time.Now().Add(111 * time.Second)}
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
	req, _ := http.NewRequest("GET", "http://localhost:1234/Logout", nil)
	cookie := http.Cookie{Name: "gosessionid", Value: "9f7225a3942e82151cd45a5c3d550b3f", Expires: time.Now().Add(111 * time.Second)}
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
	// go OpenAccount("asas", "0")
	// go Login("9", "admin")
	// go DeleteAccount("7")
	// go Logout()
	go Balance()
	// go transfer("2", "20")
	// go transfer("6", "900")
	// fmt.Print("hello")
	for {

	}
	// conn, err := net.Dial("tcp", "localhost:777")
	// if err != nil {
	// 	fmt.Println("tcp erro")
	// }
	// conn.Write([]byte("hello tcp sockersajdkahsldjasldjalsdjlkasjdlkasjd;laks;dka;skda;s"))
	// fmt.Println("finish")
	// client, err := rpc.DialHTTP("tcp", "localhost"+":1234")
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
	// resp, _ := http.Get("http://localhost:12345/")
	// resp, _ = http.Get("http://localhost:12345/getSession")
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
	// req, _ := http.NewRequest("GET", "http://localhost:1234", bytes.NewBuffer(jsonStr))
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
