package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	//布尔占一个字节
	var b bool = true
	fmt.Println(unsafe.Sizeof(b))
	//布尔只能是true活着false
	//b = 1
}
