package main

import (
	"fmt"
	"reflect"
)

/**
通过refelct 改变str
重点是Value.Elem()
*/
func main() {
	str := "123"
	of := reflect.ValueOf(&str)
	of.Elem().SetString("456")
	fmt.Println(str)
}
