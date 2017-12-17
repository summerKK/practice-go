package main

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"fmt"
	"log"
)

func main() {
	config,err := yaml.ReadFile("ch5/1217/1/config.yaml")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(config.GetBool("enabled"))
	fmt.Println(config.Get("path"))
}
