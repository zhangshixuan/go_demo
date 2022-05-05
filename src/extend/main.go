package main

import "fmt"

//公共
type Student struct {
	Name string
	Age  uint8
}

func (s *Student) showInfo() {
	fmt.Printf("%v的年龄%v\n", s.Name, s.Age)
}

type Pupil struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Printf("%v正在进行小学考试\n", p.Name)
}

type Graduate struct {
	Student
}

func (g *Graduate) testing() {
	fmt.Printf("%v正在进行大学考试\n", g.Name)
}

/**
继承：通过匿名结构体实现（实际上没有继承）
	 可以直接使用匿名结构体中的方法和属性
*/
func main() {
	pupil := &Pupil{
		Student{
			Name: "zhangsx",
			Age:  12,
		},
	}
	(*pupil).showInfo()
	(*pupil).testing()

	graduate := &Graduate{
		Student{
			Name: "lisi",
			Age:  20,
		},
	}
	graduate.showInfo()
	graduate.testing()
}
