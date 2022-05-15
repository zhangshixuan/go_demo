package main

import (
	"fmt"
	"reflect"
)

func testReflect(inter interface{}) {
	// 1. 转换成reflect.type类型
	typeOf := reflect.TypeOf(inter)
	fmt.Println(typeOf)
	// 2. 转换成reflect.Value类型
	valueOf := reflect.ValueOf(inter)
	fmt.Println(valueOf)
	// 2.1 获取对应的kind
	kind := valueOf.Kind()
	k := typeOf.Kind()
	fmt.Println(kind)
	fmt.Println(k)
	// 3.将reflect.Value转成interface
	v2inter := valueOf.Interface()
	fmt.Println(v2inter)
	// 4.类型断言转成Stu
	stu, ok := v2inter.(Stu)
	if !ok {
		fmt.Printf("类型转换错误")
	}
	fmt.Println(stu.Name)
}

type Stu struct {
	Name string
	Age  int
}

func main() {
	stu := Stu{
		Name: "tom",
		Age:  18,
	}
	testReflect(stu)
}
