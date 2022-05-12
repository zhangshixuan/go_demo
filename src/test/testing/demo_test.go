package main

import (
	"testing"
)

func TestAddUpper(t *testing.T) {
	res := AddUpper(4)
	if res != 3 {
		t.Fatalf("程序执行错误，期望值%v, 实际值%v", 3, res)
	}
	t.Logf("正确")
}
