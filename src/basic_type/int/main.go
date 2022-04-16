package main

import "fmt"
func main()  {
	// 有符号位：int8 int16 int32 int64
	// 无符号位：uint8 、、
	var a int8 = 127 //范围：-128～127
	fmt.Printf("%T\n", a)

	var b byte = 255
	fmt.Println(b)

	var r rune
	fmt.Printf("%T", r)
}
