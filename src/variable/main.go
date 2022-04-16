package main

import (
	"fmt"
)

var name, age = "zhangsan", "nan" //定义全局变量
var (
	animal = "dog"
	like = "meet"
)//第二种

func main() {
	var i int = 10
	fmt.Println("i:", i)
	var name = "zhangsx"
	fmt.Println("name:", name)
	sex := 1
	fmt.Println("sex:", sex)

	fmt.Println("name:", name, "age:", age)

	fmt.Println("animal:", animal, "like:", like)
}
