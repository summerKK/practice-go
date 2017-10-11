// 获取一个对象的字段和方法
package main

import (
	"fmt"
	"reflect"
)

// 获取一个对象的字段和方法
func GetMembers(i interface{}) {
	// 获取 i 的类型信息
	t := reflect.TypeOf(i)

	for {
		// 进一步获取 i 的类别信息
		if t.Kind() == reflect.Struct {
			// 只有结构体可以获取其字段信息
			fmt.Printf("\n%-8v %v 个字段:\n", t, t.NumField())
			// 进一步获取 i 的字段信息
			for i := 0; i < t.NumField(); i++ {
				fmt.Println(t.Field(i).Name)
			}
		}
		// 任何类型都可以获取其方法信息
		fmt.Printf("\n%-8v %v 个方法:\n", t, t.NumMethod())
		// 进一步获取 i 的方法信息
		for i := 0; i < t.NumMethod(); i++ {
			fmt.Println(t.Method(i).Name)
		}
		if t.Kind() == reflect.Ptr {
			// 如果是指针，则获取其所指向的元素
			t = t.Elem()
		} else {
			// 否则上面已经处理过了，直接退出循环
			break
		}
	}
}

// 定义一个结构体用来进行测试
type sr struct {
	string
}

// 接收器为实际类型
func (s sr) Read() {
}

// 接收器为指针类型
func (s *sr) Write() {
}

func main() {
	// 测试
	s := &sr{}
	GetMembers(s)
	s.Read()
	s.Write()

	l := sr{}
	l.Read()
	l.Write()
}

/* 测试结果（可以读取私有字段）：
*main.sr 2 个方法:
Read
Write

main.sr  1 个字段:
string

main.sr  1 个方法:
Read
*/
