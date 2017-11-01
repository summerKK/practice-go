package main

import (
	"os/exec"
	"fmt"
)

func main() {
	cmd0 := exec.Command("echo", "-n", "My first command from golang.")
	stdOut0, err := cmd0.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	outPut0 := make([]byte, 30)
	n, err := stdOut0.Read(outPut0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
}
