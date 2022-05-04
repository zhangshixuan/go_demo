package main

import "fmt"

type A struct {
	Num int64
}

func (a A) test() {
	fmt.Println(a.Num)
}

type Circle struct {
	Radios float64
}

/**
值传递
*/
func (c Circle) area() float64 {
	return 3.14 * c.Radios * c.Radios
}

/*
	指针传递
*/
func (c *Circle) area2() float64 {
	return 3.14 * (*c).Radios * (*c).Radios
}

func main() {
	var a A
	a.Num = 12
	a.test()

	// 计算面积type1
	//var circle = Circle{
	//	Radios: 4.0,
	//}
	//area := circle.area()
	//fmt.Println(area)

	// 地址传递type2
	var circle = Circle{
		Radios: 4.0,
	}
	var area = (&circle).area2()
	fmt.Println(area)

}
