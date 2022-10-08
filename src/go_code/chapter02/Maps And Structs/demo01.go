package main

import (
	"fmt"
)

func main() {
	//statePopulations := make(map[string]int)//用make函数构造map
	statePopulations := map[string]int{ //map[key tyoe]value type
		//slice不能作为key  array可以作为key
		"Texas":    27862596,
		"New York": 19413303,
	}
	statePopulations["Georgia"] = 10310371 //map添加数据
	delete(statePopulations, "Georgia")    //map删除数据
	//map输出所要找的数据在map中没有时，返回0
	_, ok := statePopulations["Georgia"]
	fmt.Println(statePopulations)
	fmt.Println(ok) //ok进行判断map中是否含有该数据
}
