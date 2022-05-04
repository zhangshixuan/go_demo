package main

import (
	"fmt"
	"struct/factory_pattern/model"
)

func main() {
	//student := model.Student{
	//	Name: "zhangsx",
	//	Age:  24.0,
	//}
	//fmt.Println(student)

	student := model.GetStudent("zhangsx", 24.0)
	fmt.Println(*student)
}
