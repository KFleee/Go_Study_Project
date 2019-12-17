package main

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	r, err := conn.Do("HMSET", "gotest", "name", "h", "sex", "F")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
	r, err = conn.Do("HGETALL", "gotest")
	fmt.Println(r)
	hmap, _ := redis.StringMap(r, err)
	for key, value := range hmap {
		fmt.Println("key : ", key, "     value : ", value)
	}
}
