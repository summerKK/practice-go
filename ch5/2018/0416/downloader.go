package main

import (
	"io"
	"net/http"
	"os"
	"golang.org/x/net/proxy"
	"net"
	"time"
	"log"
	"flag"
)

var (
	generatorHttpClient *http.Client
	url                 string
	fileName            string
)

func main() {

	flag.StringVar(&url, "u", "", "下载地址")
	flag.StringVar(&fileName, "n", "", "文件名")

	flag.Parse()

	if url == "" || fileName == "" {
		log.Fatal("参数不能为空")
	}

	httpClient := genHttpClient()
	res, err := httpClient.Get(url)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
}

func genHttpClient() *http.Client {

	if generatorHttpClient == nil {
		dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:1080",
			nil,
			&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			},
		)
		if err != nil {
			log.Fatalln("get dialer error", dialer)
		}
		httpTransport := &http.Transport{Dial: dialer.Dial}
		generatorHttpClient = &http.Client{Transport: httpTransport}
		return generatorHttpClient
	}

	return generatorHttpClient
}
