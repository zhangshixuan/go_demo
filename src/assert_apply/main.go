package main

import "fmt"

func TypeJudge(items ...interface{}) {
	for i, item := range items {
		switch item.(type) {
		case int, int32, int64:
			fmt.Printf("第%v是int:%v\n", i+1, item)
		case string:
			fmt.Printf("第%v是string:%v\n", i+1, item)
		default:
			fmt.Printf("第%v是不确定:%v\n", i+1, item)
		}
	}
}

/**
断言实例2
*/
func main() {
	v1 := "hahaha"
	v2 := 12

	TypeJudge(v1, v2)
}
