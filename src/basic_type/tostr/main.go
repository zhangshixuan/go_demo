package main

import (
	"fmt"
	"strconv"
)

func main()  {
	//其他类型转字符串
	//整型转字符串
	i := 12341
	itostr := fmt.Sprintf("%d", i)
	fmt.Printf("type:%T value:%v \n", itostr, itostr)
	itoa := strconv.Itoa(i)
	fmt.Printf("type:%T value:%v \n", itoa, itoa)

	//浮点型转字符串
	f := 3.1415926535
	ftostr := fmt.Sprintf("%f", f)
	fmt.Printf("type:%T value:%v \n", ftostr, ftostr)
	ftostr1 := strconv.FormatFloat(f, 'g', 10, 64)
	fmt.Printf("type:%T value:%v \n", ftostr1, ftostr1)

	//bool转字符串
	b := true
	btostr1 := fmt.Sprintf("%t", b)
	fmt.Printf("type:%T value:%v \n", btostr1, btostr1)
	formatBool := strconv.FormatBool(b)
	fmt.Printf("type:%T value:%v \n", formatBool, formatBool)


}
