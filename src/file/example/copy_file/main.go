package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
文件copy
*/
func main() {

	src, err := os.Open("/Users/xuan/Downloads/jay.jpeg")
	if err != nil {
		fmt.Println("open file error")
	}
	defer src.Close()
	reader := bufio.NewReader(src)

	dest, err := os.OpenFile("/Users/xuan/jay.jpeg", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}
	defer dest.Close()
	writer := bufio.NewWriter(dest)

	io.Copy(writer, reader)
}
