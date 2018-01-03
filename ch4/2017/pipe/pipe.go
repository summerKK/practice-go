package main

import (
	"os/exec"
	"fmt"
	"bufio"
	"bytes"
)

//240
func main() {
	cmd1 := exec.Command("cmd", "/C", "netstat", "-a")
	cmd2 := exec.Command("findstr", "80")

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
	if err := cmd2.Wait(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outputBuf2)
}
