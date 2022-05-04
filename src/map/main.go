package main

import "fmt"

func main() {
	/**
	type 1
	map声明，需要分配空间make
	自动扩容
	*/
	var m map[string]string
	m = make(map[string]string, 2)
	m["key1"] = "jay"
	m["key2"] = "jj"
	m["key3"] = "zsx"
	fmt.Println(m)

	// type 2
	cities := make(map[string]string)
	cities["city1"] = "shanghai"
	fmt.Println(cities)

	// type 3
	hores := map[string]string{
		"1": "thor",
		"2": "spider man",
	}
	fmt.Println(hores)
	// 查询
	val, s := hores["1"]
	if s {
		fmt.Println(val)
	}
}
