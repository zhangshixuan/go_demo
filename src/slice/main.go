package main

import "fmt"

/**
切片
*/
func main() {
	// 第一种
	arr := [...]int{1, 2, 3}
	slice := arr[1:3]
	fmt.Println("slice:", slice)
	fmt.Printf("arr[1]地址  %v\n", &arr[1])
	fmt.Printf("slice[0]地址%v\n", &slice[0])
	fmt.Println("--------------")

	// 第二种make
	var arrMake = make([]int, 2, 10)
	fmt.Println("arrMake", arrMake)
	fmt.Println("--------------")

	// 第三种
	var strArr = []string{"1", "2", "3"}
	fmt.Println(strArr)

	// 可以使用append扩容
	strArr = append(strArr, "4")
	fmt.Println("strings", strArr)
}
