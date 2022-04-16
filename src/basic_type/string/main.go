package main

import "fmt"

func main()  {
	//字符串由多个字符组成
	//字符串不可变性
	var str = "不可变性"
	//str[0] = '1'
	fmt.Println(str)
	//反引号，es6模版字符串
	str1 := `\n\t\r`
	fmt.Println(str1)

}
