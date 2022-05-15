package main

import (
	"fmt"
	"sync"
	"time"
)

/**
把数放入管道
*/
func putNum(intChan chan int) {
	for i := 1; i <= 10000; i++ {
		intChan <- i
	}
	close(intChan)
}

/**
在数管道里取数判断素数，如果是放入素数管道
并且判断如果取完了，将状态放入判断管道
*/
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for true {
		num, ok := <-intChan
		if !ok {
			exitChan <- true
			break
		}
		if IsPrime(num) {
			primeChan <- num
		}
	}
	//一个协程完成工作，不能关闭primeChan管道，可能还会有其他协程写入数据
}

// IsPrime 判断素数方法
func IsPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

/**
参考图示：
图示少画了匿名函数的协程
*/
func main() {
	intChan := make(chan int, 1000)   //所有数管道
	primeChan := make(chan int, 2000) //素数管道
	exitChan := make(chan bool, 4)    // 判断4个协程推出的管道
	start := time.Now().Unix()
	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	for true {
		_, ok := <-primeChan
		if !ok {
			break
		}
		//fmt.Println(res)
	}
	end := time.Now().Unix()
	fmt.Println(end - start)
}
