package main

import (
	"gopkg.in/gcfg.v1"
	"log"
	"fmt"
)

func main() {
	config := struct {
		Section struct {
			Enabled bool
			Path    string
		}
	}{}

	err := gcfg.ReadFileInto(&config, "ch5/1217/1/config.ini")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config.Section.Path)
	fmt.Println(config.Section.Enabled)
}
