package main

import (
	"github.com/blang/semver"
	"fmt"
	"time"
	"io/ioutil"
	"os"
)

func main() {
	v, _ := semver.Make("0.0.1-alpha.preview+123.github")
	fmt.Printf("Major: %d\n", v.Major)
	fmt.Printf("Minor: %d\n", v.Minor)
	fmt.Printf("Patch: %d\n", v.Patch)
	fmt.Printf("Pre: %s\n", v.Pre)
	fmt.Printf("Build: %s\n", v.Build)

	// Prerelease versions array
	if len(v.Pre) > 0 {
		fmt.Println("Prerelease versions:")
		for i, pre := range v.Pre {
			fmt.Printf("%d: %q\n", i, pre)
		}
	}

	// Build meta data array
	if len(v.Build) > 0 {
		fmt.Println("Build meta data:")
		for i, build := range v.Build {
			fmt.Printf("%d: %q\n", i, build)
		}
	}

	i := 5
	fmt.Println(time.Duration(i))
	fmt.Printf("%T\n", time.Duration(i))
	fmt.Printf("%T\n", time.Second)

	ioutil.WriteFile("json_error.txt", []byte("hello world"), 0644)

	fmt.Println(os.Stat("hello.json"))

	fmt.Println("------------------")
	fmt.Println("hello")

	go test()

	time.Sleep(time.Second)

	var se semver.Version

	fmt.Println(se.String())

}

func test() {
	fmt.Println("world")
	fmt.Println("------------------")
}
