package main

import "fmt"

/**
在无限循环的时候，如果不关闭管道，就会出现deadlock

*/
func main() {
	boolChan := make(chan bool, 3)

	go func() {
		boolChan <- true
		boolChan <- false
		boolChan <- true
		// 只有关闭管道，才能出现!ok跳出的状态
		close(boolChan)
	}()

	//for i := 0; i < 3; i++ {
	//	b := <-boolChan
	//	fmt.Println(b)
	//}

	for true {
		b, ok := <-boolChan
		if !ok {
			break
		}
		fmt.Println(b)
	}
}
