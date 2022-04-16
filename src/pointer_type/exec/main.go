package main

import "fmt"

func main() {
	// 1.写一个程序，获取int变量num的地址，并输出到终端
	var num int = 10
	fmt.Printf("num变量地址：%v\n", &num)
	// 2.将num的地址赋给ptr，并通过ptr去修改num的值
	var par *int = &num
	*par = 20
	fmt.Println(num)

}
