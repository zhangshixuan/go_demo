package main

import "fmt"

func main() {
	// go中，所有的变量的值都会被分配一块内存，还有其代表的地址
	var i int = 10
	fmt.Printf("i的值是：%v 地址：%v\n", i, &i)

	// 1.声明了一个指针类型的变量
	// 2.赋值为一个地址&i，具体指向i的地址
	// 3.其par本身还有一个地址
	var par *int = &i
	fmt.Printf("par的值是：%v 地址：%v\n", par, &par)

	// *par可以解引用，将地址中所代表的值取出来
	fmt.Printf("par指针变量所代表的值：%v", *par)

	// 总结：
	// 每个变量都会指向一块空间的地址，空间存的是具体的值
	// 指针变量也会指向一块空间的地址，只不过具体存的值是另一块空间的地址
	// 像上面的par存的就是i变量空间的值
}
