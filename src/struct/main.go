package main

import "fmt"

type Stu struct {
	Name string
	Age  uint8
}

func main() {
	// type 1
	var cat1 Stu
	cat1.Name = "Tom"
	cat1.Age = 255
	fmt.Println(cat1)

	// type2
	var cat2 Stu = Stu{
		"Cat",
		0,
	}
	fmt.Println(cat2)
}
