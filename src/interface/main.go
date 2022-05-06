package main

import "fmt"

//定义一个usb接口
type Usb interface {
	Start()
	Stop()
}

//定义一个相机结构体
type Camera struct {
	Name string
}

//让相机结构体方法，实现usb接口的开始方法
func (c Camera) Start() {
	fmt.Printf("%v相机开始工作\n", c.Name)
}

//让相机结构体方法，实现usb接口的结束方法
func (c Camera) Stop() {
	fmt.Printf("%v相机结束工作\n", c.Name)
}

type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println("手机结束工作")
}

// 计算机结构体
type Computer struct {
}

//编写一个方法Working方法，接收一个Usb接口类型变量
//只要是实现了Usb接口（所谓实现Usb接口，就是指实现了Usb接口声明所有方法）
func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

/**
接口
*/
func main() {
	camera := Camera{
		Name: "三星手机",
	}
	phone := Phone{}
	computer := Computer{}
	computer.Working(camera)
	computer.Working(phone)
}
