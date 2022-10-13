// writeData协程 向管道intChan写入50个整数
// readData协程  从管道读取writeData写入的数据
// 上述两个协程都用一个管道
// 主线程需要等待writeData和readData协程都完成后才能退出
package main

import (
	"fmt"
	"time"
)

func writeData(intChan chan int) {
	for i := 0; i <= 50; i++ {
		intChan <- i
		fmt.Println("writeData=", i)
		time.Sleep(time.Second)
	} //for循环写入
	close(intChan)
}
func readData(intChan chan int, exitChan chan bool) {
	for {
		_, ok := <-intChan
		if !ok {
			break
		}
		time.Sleep(time.Second)
		fmt.Printf("readData =%v\n", v)
	}
	exitChan <- true
	close(exitChan)
}
func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

}
