package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	url := "https://baidu.com"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		content, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(content))
	}
}
