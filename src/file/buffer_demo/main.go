package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("D://test.txt")
	if err != nil {
		fmt.Println("open file error:", err)
	}
	defer file.Close()
	// 缓冲流读取数据
	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(readString)
	}
}