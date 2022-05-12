package main

func AddUpper(param int) int {
	var sum int
	for i := 1; i < param+1; i++ {
		sum += i
	}
	return sum
}
