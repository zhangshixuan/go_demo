package main

import (
	"fmt"
	"time"
)

func write(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Println("写入的数据是：", i)
		//time.Sleep(time.Second)
	}
	close(intChan)
}

func read(intChan chan int, exitChan chan bool) {
	for true {
		v, ok := <-intChan
		if !ok {
			break
		}
		//time.Sleep(time.Second)
		fmt.Println("读取的数据是：", v)
	}
	exitChan <- true
	close(exitChan)
}

/**
往管道里放入50
一边放一边取50
在主线程中运行等待
*/
func main() {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go write(intChan)
	go read(intChan, exitChan)
	for true {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
