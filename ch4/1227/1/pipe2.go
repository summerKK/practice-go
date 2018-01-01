package main

import (
	"os/exec"
	"fmt"
	"bufio"
	"bytes"
)

func main() {
	cmd1 := exec.Command("dir")
	cmd2 := exec.Command("findstr","ch")

	stdout1, err := cmd1.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := cmd1.Start(); err != nil {
		fmt.Println(err)
		return
	}
	outputBuf1 := bufio.NewReader(stdout1)
	stdin2, err := cmd2.StdinPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	outputBuf1.WriteTo(stdin2)

	var outputBuf2 bytes.Buffer
	cmd2.Stdout = &outputBuf2
	if err := cmd2.Start(); err != nil {
		fmt.Println(err)
		return
	}
	err = stdin2.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := cmd2.Wait();err != nil{
		fmt.Println(err)
		return
	}
}
