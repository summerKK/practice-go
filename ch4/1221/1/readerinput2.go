package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.TrimSpace("hello world"))
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入:")
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Println(input)
	}
}
