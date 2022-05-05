package main

import (
	"encapsulation/model"
	"fmt"
)

/**
go封装：隐藏一些私有变量或者方法（实际上没有封装，只是一种形式）
实现模式：工厂模式
*/
func main() {
	person := model.NewPerson("Jay")
	fmt.Printf("%v没开工资之前:%v\n", person.Name, person.GetSalary())
	person.SetSalary(20000.0)
	fmt.Printf("%v开工资之后:%v\n", person.Name, person.GetSalary())
}
