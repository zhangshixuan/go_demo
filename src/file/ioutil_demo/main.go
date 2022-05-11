package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	files, err := ioutil.ReadFile("/Users/xuan/VPN.txt")
	if err != nil {
		fmt.Println("文件读取错误", err)
		return
	}
	fmt.Println(string(files))
}
