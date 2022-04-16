package main

import "fmt"

func main()  {
	//Golang的字符编码集只是UTF-8，这样避免乱码问题
	//存：字符->数字->二进制
	//取：二进制->数字->字符
	var c1 byte = 'a'
	fmt.Println(c1)
	var c2 int = '中'
	fmt.Println(c2)
}
