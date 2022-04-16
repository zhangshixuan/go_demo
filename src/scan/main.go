package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Printf("你输入的姓名是：%v", name)
}

