package main

import (
	"fmt"
	"utils"
)

// 引包
func main() {
	fmt.Println("start package")
	i := utils.Par()
	fmt.Printf("%v", i)
}
