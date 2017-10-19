package main

import (
	"fmt"
)

type Connecter interface {
	Connect()
}
type USB interface {
	//定义一个接口，这个接口下面包含两个方法
	//Name方法返回一个字符串
	Name() string
	Connecter
}

type PhoneConnecter struct {
	//用这个struct对这个USB的接口进行实现
	name string
}

func (pc PhoneConnecter) Name() string { //实现接口就是用这个结构，实现interface的方法
	return pc.name
}
func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}
func main() {
	pc := PhoneConnecter{"PhoneConnecter"}
	fmt.Printf("%T\n",pc)
	var a Connecter
	a = Connecter(pc)
	a.Connect()
	pc.name = "pc---"       //这里将对象赋值给接口时，会发生拷贝，接口内部存储的是指向这个复制品的指针，
	fmt.Println(pc.name) //既无法修改复制品的状态，也无法获取指针
	a.Connect()
}
func Disconnect(usb interface{}) { // 这里把USB替换成空接口
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected.", v.name)
	default:
		fmt.Println("Unknow device.")
	}

}
