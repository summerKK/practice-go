package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(j)
			if j == 2{
				break
			}
		}
	}

	fmt.Println(time.Now())


	fmt.Println(strings.TrimSpace("hello world"))
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入:")
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Println(input)
	}
}
