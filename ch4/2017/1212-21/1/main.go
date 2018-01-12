package main

import (
	"fmt"
	"os"
	"bufio"
	"bytes"
	"io"
)

func main() {
	var capx128 complex128 = complex(2, -2)
	var im64 = imag(capx128)
	var r64 = real(capx128)

	var m map[interface{}]bool

	m = make(map[interface{}]bool)
	fmt.Println(m["a"])

	fmt.Println(im64, r64)

	var c uint = 1

	switch c {
	case 1:
		fmt.Println("hello")
		fallthrough
	case 2:
		fmt.Println("world")

	}

	s := map[string]int{"你好": 1, "中国": 2, "hello": 3}
	var targetCount = make(map[string]int)
	for k, v := range s {
		var matched bool = true
	L:
		for _, r := range k {
			fmt.Printf("%U\n", r)
			if r < '\u4e00' || r > '\u9fbf' {
				matched = false
				break L
			}
		}
		if matched {
			targetCount[k] = v
		}
	}

	fmt.Println(targetCount)

L1:
	for v := range []int{1, 2, 3, 4, 5} {
		if v == 2 {
			break L1
		}
		fmt.Println(v)
	}

	ints := make([]int, 0)
	result := appendNumbers(ints)
	fmt.Println(result)

	openFIle(".gitignore")

	fetchDemo()

	def()

}

func appendNumbers(ints []int) (result []int) {
	result = append(ints, 1)
	defer func() {
		result = append(result, 2)
	}()
	result = append(result, 3)
	defer func() {
		result = append(result, 4)
	}()
	result = append(result, 5)
	defer func() {
		result = append(result, 6)
	}()
	return result
}

func openFIle(f string) {
	file, err := os.Open(f)
	if err != nil {
		if pe, ok := err.(*os.PathError); ok {
			fmt.Printf("Path error:%s (op=%s,path=%s)\n", pe.Err, pe.Op, pe.Path)
		} else {
			fmt.Printf("Uknown error:%s\n", err)
		}
	}
	r := bufio.NewReader(file)
	var buf bytes.Buffer
	for {
		byteArray, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("Read error:%s\n", err)
				break
			}
		} else {
			buf.Write(byteArray)
			buf.WriteByte('\n')
		}
	}
	fmt.Println(buf.String())
}

func fetchDemo() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("recovered a pain.[index=%d]\n", v)
		}
	}()
	ss := []string{"A", "B", "C"}
	fmt.Printf("fetch the elements in %v one by one...", ss)
	fetchElement(ss, 0)
	fmt.Println("The elements fetching is done.")
}

func fetchElement(s []string, index int) (element string) {
	if index >= len(s) {
		fmt.Printf("occur a pain! [index=%d]\n", index)
		panic(index)
	}
	fmt.Printf("fetching the element... [index=%d]\n", index)
	element = s[index]
	defer fmt.Printf("The element is %s. [index=%d]\n", element, index)
	fetchElement(s, index+1)
	return
}

func def() {
	var in int
	defer func() {
		fmt.Println(in)
	}()
	in = 1
}
