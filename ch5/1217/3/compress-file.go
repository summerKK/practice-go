package main

import (
	"os"
	"compress/gzip"
	"io"
	"sync"
	"fmt"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var files []string = []string{
		"ch5/1217/1/config.go",
		"ch5/1217/1/config.ini",
		"ch5/1217/1/config.yaml",
	}
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			compress(file)
			wg.Done()
		}(file)

	}
	wg.Wait()

	test()

	time.Sleep(time.Second * 3)

}

func test() {
	for i, inter := range []string{"1", "2"} {
		fmt.Println(&inter) //内存地址都是一样的
		fmt.Println(&i)     //内存地址都是一样的
		go func() {
			time.Sleep(time.Second)
			fmt.Println(inter) // 2 2
			fmt.Println(i)     // 1 1
		}()
	}
}

func compress(file string) error {
	in, err := os.Open(file)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(file + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()
	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	defer gzout.Close()
	return err
}
