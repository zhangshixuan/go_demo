package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("D://test.txt")
	if err != nil {
		fmt.Println("open file error:", err)
	}
	fmt.Println(&file)
	if file.Close() != nil {
		fmt.Println("close file error:", file.Close())
	}
}
