package main

import (
	"fmt"
	"strconv"
)

func main()  {
	//字符串转其他基本类型
	//转bool
	str1 := "true"
	parseBool, _ := strconv.ParseBool(str1)
	fmt.Printf("type:%T value:%v \n", parseBool, parseBool)
	//转int
	str2 := "123120"
	parseInt, _ := strconv.ParseInt(str2, 10, 64)
	fmt.Printf("type:%T value:%v \n", parseInt, parseInt)
	//转float
	str3 := "1.21212312"
	parseFloat, _ := strconv.ParseFloat(str3, 64)
	fmt.Printf("type:%T value:%v \n", parseFloat, parseFloat)
}
