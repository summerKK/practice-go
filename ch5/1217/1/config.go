package main

import (
	"os"
	"encoding/json"
	"fmt"
)

type configuration struct {
	Enabled bool
	Path    string
}

func main() {
	file, _ := os.Open("ch5/1217/1/config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := configuration{}
	err := decoder.Decode(&config)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(config.Path)
}
