package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("/Users/xuan/test.txt")
	if err != nil {
		fmt.Println("read file error", err)
	}
	err = ioutil.WriteFile("/Users/xuan/test1.txt", data, 0777)
	if err != nil {
		fmt.Println("write file error")
	}
}
