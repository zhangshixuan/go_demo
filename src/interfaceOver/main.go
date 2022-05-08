package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

// 给HeroSlice提供Len方法
func (hs HeroSlice) Len() int {
	return len(hs)
}

// 给HeroSlice提供Less方法
func (hs HeroSlice) Less(i int, j int) bool {
	return hs[i].Age < hs[j].Age
}

// 给HeroSlice提供Swap方法
func (hs HeroSlice) Swap(i int, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

/**
重写排序方法
*/
func main() {
	var heros HeroSlice

	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}
	for _, hero := range heros {
		fmt.Println(hero)
	}
	fmt.Println("---------按年龄排序后---------")
	sort.Sort(heros)
	for _, hero := range heros {
		fmt.Println(hero)
	}
}
