// 计算1-200的各个数的阶乘，并且把各个数的阶乘放入到map中，最后显示出来
// 思路:1、编写一个函数，来计算各个数的阶乘，并放到map中
//
//	2、我们启动的多个协程，统计的将结果放入到map中
//	3、map应该做成全局的
package main

import (
	"fmt"
	"sync" //synchronize同步
	"time"
)

var (
	myMap = make(map[int]int, 10) //定义全局map
	//声明一个全局的互斥锁(mutex)
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//将我们的res放入到myMap中
	//加锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}
func main() {
	//开启多个协程完成这个任务
	for i := 1; i <= 200; i++ {
		go test(i)
	} //启动了200个协程
	//休眠10秒钟
	time.Sleep(time.Second * 10)
	lock.Lock()
	//输出
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
