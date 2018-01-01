package main

import (
	"os/exec"
	"fmt"
	"bufio"
)

func main() {
	command0 := exec.Command("echo", "-n", "my first command from golang -------------------------------------------------")
	stdout0, err := command0.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := command0.Start(); err != nil {
		fmt.Println(err)
		return
	}
	//读取管道数据
	//方法一
	/*
	var outputBufo bytes.Buffer
	for {
		tempOut := make([]byte, 5)
		n, err := stdout0.Read(tempOut)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		if n > 0 {
			outputBufo.Write(tempOut[:n])
		}
	}
	fmt.Println(outputBufo.String())
	*/
	//方法二
	outputBufo := bufio.NewReader(stdout0)
	output0, _, err := outputBufo.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output0))

}
