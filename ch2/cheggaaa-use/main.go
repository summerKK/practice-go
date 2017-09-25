package main

import (
	"github.com/blang/semver"
	"fmt"
	"time"
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

	ticker := time.NewTicker(4 * time.Second / time.Duration(i))
	for t := range ticker.C {
		fmt.Println(t)
		fmt.Printf("%T\n",t)
	}
}
