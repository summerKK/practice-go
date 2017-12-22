package main

import (
	"reflect"
	"fmt"
	"strings"
)

func main() {
	p := map[string]string{"summer": "Hello World"}

	kind := reflect.ValueOf(p).Kind()

	fmt.Println(kind)

	fmt.Println(reflect.Array)

	s := "hello;world"

	fmt.Println(strings.Split(s,";"))

	var mapping map[string]string = map[string]string{}

	mapping["summer"] = "hello"

	fmt.Println(mapping)

	format := "%scm"

	strs := []interface{}{"12"}

	fmt.Println(fmt.Sprintf(format,strs...))

}
