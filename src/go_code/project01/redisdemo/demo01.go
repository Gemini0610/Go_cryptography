package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	//通过go向redis写入读取数据
	//1、连接
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	//2、通过go向redis写入数据 string[key-val]
	_, err = conn.Do("Set", "name", "tom")
	if err != nil {
		fmt.Println("set err =", err)
		return
	}
	//3、读取
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("set err =", err)
		return
	}
	fmt.Println("OK", r)
}
