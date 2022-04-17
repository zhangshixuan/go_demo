package main

import "fmt"

func main() {
	// 闭包演示
	// 闭包 = 类变量， 匿名函数 = 类方法
	upper := AddUpper()
	fmt.Println(upper(1))
	fmt.Println(upper(2))

}

// 累加器
func AddUpper() func(int) int {
	n := 10
	return func(x int) int {
		n += x
		return n
	}
}
