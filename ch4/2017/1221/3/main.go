package main

import (
	"fmt"
	"golang.org/x/net/proxy"
	"net/http"
	"log"
	"io/ioutil"
	"net"
	"time"
)

func main() {
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:1080",
		nil,
		&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
		},
	)
	if err != nil {
		log.Fatalln("get dialer error", dialer)
	}
	httpTransport := &http.Transport{Dial: dialer.Dial}
	httpClient := &http.Client{Transport: httpTransport}
	resp, err := httpClient.Get("https://www.google.com.hk/")
	if err != nil {
		log.Fatalln(err)
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s\n", body)
	}
}
