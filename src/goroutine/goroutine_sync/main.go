package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	reMap = make(map[int]float64)
	lock  = sync.Mutex{}
)

func factorial(n int) {
	var res = 1.0
	for i := 1; i < n+1; i++ {
		res *= float64(i)
	}
	lock.Lock()
	reMap[n] = res
	lock.Unlock()
}

func main() {
	for i := 1; i < 21; i++ {
		go factorial(i)
	}
	time.Sleep(time.Second * 10)
	for i, f := range reMap {
		fmt.Printf("map[%v]=%v\n", i, f)
	}
}
