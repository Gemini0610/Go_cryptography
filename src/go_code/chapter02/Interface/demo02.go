package main

import (
	"fmt"
)

type Qimiao struct {
	Name   string
	Age    int
	Sex    bool
	Hobbys []string
}

func main() {
	// var qm Qimiao
	// qm.Age = 18
	// qm.Name = "奇妙"
	// qm.Sex = false
	// qm.Hobbys = []string{"玩游戏", "唱歌"}

	// fmt.Println(qm)
	qm := Qimiao{
		"奇妙", 18, false, []string{"玩游戏", "唱歌"},
	}
	fmt.Println(qm)
}
