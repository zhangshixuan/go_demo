package main

import (
	"fmt"
	"time"
)

func test() {
	for i := 1; i < 11; i++ {
		fmt.Printf("test()方法在执行：%v\n", i)
		time.Sleep(time.Second)
	}
}

/**
协程特点：
	有独立的栈空间
	共享堆空间
	手动调度
	轻量的线程
*/
func main() {
	go test()
	for i := 1; i < 11; i++ {
		fmt.Printf("main()方法在执行：%v\n", i)
		time.Sleep(time.Second)
	}
}
