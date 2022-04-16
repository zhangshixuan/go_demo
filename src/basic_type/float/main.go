package main

import "fmt"

func main()  {
	//float由符号、指数、位数组成
	var sF float32 = 3.1415926535
	fmt.Println("精度丢失：", sF) //单精度位，精度丢失

	f := 3.1415926535
	fmt.Println(f)
	fmt.Printf("%T", f) //默认位float64双精度
}
