package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
创建文件，写入文字
*/
func main() {
	file, err := os.OpenFile("/Users/xuan/test.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}
	defer file.Close()
	str := "Hello, xuan\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()

}
