package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"practice/ch2/scann_/source"
)

func main() {
	current, _ := os.Getwd()
	SEP := string(os.PathSeparator)
	path := current + SEP + "ch2" + SEP + "scann_" + SEP + "main.go"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	var test1 *source.HELLO = &source.HELLO{}
	test1.Hello = "summer"
	var test2 *source.HELLO = &source.HELLO{}
	fmt.Println(test1, test2)
	test3 := test1
	test3.Hello = "ella"

	fmt.Printf("%#v\n", test1)

	var i int
	//defer func() {
	//	fmt.Println(i)
	//}()

	defer fmt.Println(i)

	for i = 0; i < 10; i++ {

	}
	fmt.Println(i)
}
