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
	fmt.Println(valueOf) //此时不是真正的int类型，是reflect的value
	//做运算的时候调用Value结构体的Int()方法，转成int
	i := valueOf.Int()
	fmt.Println(i + 1)
	// 3.将reflect.Value转成interface
	v2inter := valueOf.Interface()
	fmt.Println(v2inter)
	// 4.类型断言转成int
	intValue := v2inter.(int)
	fmt.Println(intValue)
}

/**
int -> interface{} -> reflect.value
reflect.value -> interface{} -> int
*/
func main() {
	i := 10
	testReflect(i)
}
