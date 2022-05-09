package main

import "fmt"

type Point struct {
	x int
	y int
}

/**
	接口断言：将已知类型的接口复制给某一个已知变量类型，如将i赋值给f就是进行了类型断言
 */
func main() {
	fmt.Println("----------------数据类型断言---------------")
	var i interface{}
	var f float32 = 18.8
	i = f
	fmt.Println(i)
	f = i.(float32)
	fmt.Println(f)
	fmt.Println("----------------结构体类型断言---------------")
	var is interface{}
	point := Point{
		x: 2,
		y: 3,
	}
	is = point
	fmt.Println(is)
	point = is.(Point)
	fmt.Println(point)
}
