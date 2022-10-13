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

func (q *Qimiao) Song(name string) (restr string) {
	restr = "真是的"
	fmt.Printf("%v唱了一手%v,观众觉得%v", q.Name, name, restr)
	return restr

}
func main() {
	// var qm Qimiao
	// qm.Age = 18
	// qm.Name = "奇妙"
	// qm.Sex = false
	// qm.Hobbys = []string{"玩游戏", "唱歌"}

	// fmt.Println(qm)
	qm := Qimiao{
		Name:   "唱歌个",
		Age:    18,
		Sex:    false,
		Hobbys: nil,
	}
	re := qm.Song("惊雷")
	fmt.Println("\n", re)
}
