package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Count struct {
	ZhC     int
	SpaceC  int
	NumberC int
	Other   int
}

/**
统计字符
*/
func main() {
	file, err := os.Open("/Users/xuan/count.txt")
	if err != nil {
		fmt.Println("open file error")
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var count Count
	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, value := range readString {
			switch {
			case value > 'a' && value < 'z':
				fallthrough
			case value > 'A' && value < 'Z':
				count.ZhC++
			case value >= '0' && value <= '9':
				count.NumberC++
			case value == ' ' || value == '\t':
				count.SpaceC++
			default:
				count.Other++
			}
		}
	}
	fmt.Printf("字母：%v, 数字：%v, 空格：%v, 其他：%v", count.ZhC, count.NumberC, count.SpaceC, count.Other)
}
