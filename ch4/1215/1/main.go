package main

import (
	"net/http"
	"log"
	"fmt"
	"os"
)

func main() {
	getHtml("https://www.1stdibs.com/fashion/handbags-purses-bags/top-handle-bags/hermes-birkin-30cm-gold-camel-tan-togo-gold-hardware-bag-stamp/id-v_3643261/")
}

func getHtml(url string) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == http.StatusOK {
		fmt.Println(resp.StatusCode)
		buf := make([]byte, 1024)
		f, _ := os.OpenFile("1stdibs.html", os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModePerm)
		defer f.Close()
		for {
			n, _ := resp.Body.Read(buf)
			if 0 == n {
				break
			}
			f.WriteString(string(buf[:n]))
		}
	}
}
