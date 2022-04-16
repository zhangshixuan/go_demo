package main

import "fmt"

func main() {
	/**
	1.函数可以像变量一样当参数。比如isOdd
	2.filter方法事先需要自定义函数类型参数
	*/
	slice := []int{1, 2, 3, 4}
	ints := filter(slice, isOdd)
	fmt.Println(ints)

	/**
	函数还可以返回多个值，支持命名
	*/
	sum, sub := getSumAndSub(2, 1)
	fmt.Printf("sum:%v, sub:%v", sum, sub)
}

type myFunctype func(num int) bool

func filter(slice []int, f myFunctype) []int {
	var res []int
	// type1
	for i := 0; i < len(slice); i++ {
		if f(slice[i]) {
			res = append(res, slice[i])
		}
	}
	// type2
	//for _, value := range slice {
	//	if f(value) {
	//		res = append(res, value)
	//	}
	//}
	return res
}

func isOdd(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

func getSumAndSub(num1 int, num2 int) (sum int, sub int) {
	return num1 + num2, num1 - num2
}
