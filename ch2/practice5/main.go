package main

import (
	"os"
	"fmt"
)

func main() {
	current, err := os.Getwd()
	if err != nil {
		fmt.Println("Getwd:", err)
	}
	separator := string(os.PathSeparator)
	path := current + separator + "ch2" + separator + "practice5" + separator + "main.go"
	checkFile, err := os.Stat(path)
	if err != nil {
		fmt.Println("Stat:", err)
	}
	fmt.Println(checkFile)
	err = os.Link(path, current+separator+"link.go")
	if err != nil{
		fmt.Println("Link:",err)
	}
}
