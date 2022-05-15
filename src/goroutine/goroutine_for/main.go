package main

import "fmt"

/**
在无限循环的时候，如果不关闭管道，就会出现deadlock

*/
func main() {
	boolChan := make(chan bool, 3)
	boolChan <- true
	boolChan <- false
	boolChan <- true
	close(boolChan)

	for true {
		b, ok := <-boolChan
		if !ok {
			break
		}
		fmt.Println(b)
	}
}
