package main

import (
	"os"
	"fmt"
	"time"
	"io"
)

func main() {
	pipeBasedFile()
	fmt.Println("----------------------------------")
	pipeBasedMemory()
}

//内存形式的管道
func pipeBasedMemory() {
	reader, writer := io.Pipe()
	go func() {
		output := make([]byte, 100)
		n, err := reader.Read(output)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Read %d byte(s) [memory-pipe]\n", n)
		fmt.Println(string(output))
	}()

	input := make([]byte, 26)
	for i := 65; i <= 90; i++ {
		input[i-65] = byte(i)
	}
	n, err := writer.Write(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Write %d byte(s) [memory-pipe]\n", n)
	time.Sleep(time.Microsecond * 200)
}

//文件形式的管道
func pipeBasedFile() {
	reader, writer, err := os.Pipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	//读和写操作要同时进行,因为会阻塞
	go func() {
		output := make([]byte, 100)
		n, err := reader.Read(output)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Read %d byte(s) .[file-based pipe]\n", n)
		fmt.Println(string(output))
	}()

	input := make([]byte, 26)
	for i := 65; i <= 90; i++ {
		input[i-65] = byte(i)
	}
	n, err := writer.Write(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Write %d byte(s) .[file-based pipe]\n", n)
	time.Sleep(time.Microsecond * 200)
}
