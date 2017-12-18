package main

import (
	"regexp"
	"fmt"
	"os"
	"net/http"
	"time"
	"io/ioutil"
	"io"
	"bytes"
	"sync"
	"strconv"
)

func main() {
	var wg sync.WaitGroup
	urls := []string{
		"https://a.1stdibscdn.com/archivesE/upload/1121189/v_36432611513079272365/3643261_master.jpg?width=240",
		"https://a.1stdibscdn.com/archivesE/upload/1121189/v_35765811512477005519/3576581_master.jpg?width=240",
		"https://a.1stdibscdn.com/archivesE/upload/v_336/v_36084831512408362811/2947713_master.jpg?width=240",
		"https://a.1stdibscdn.com/archivesE/upload/v_842/v_36434211513016100807/3488741_master.jpg?width=240",
		"https://a.1stdibscdn.com/archivesE/upload/1121189/v_36412511513078517857/3641251_master.jpg?width=240",
		"https://a.1stdibscdn.com/archivesE/upload/1121189/v_35271211511263112351/3527121_master.jpg?width=240",
	}
	img := make(chan string, 5)
	imgStr := ""
	var mu sync.Mutex
	for i, url := range urls {
		wg.Add(1)
		go func(url string, i int) {
			s := DownloadImg(url, "ch4/1218/1/img/", i, 5)
			mu.Lock()
			imgStr += s + ";"
			mu.Unlock()
			img <- s
			wg.Done()
		}(url, i)
	}

	go func() {
		wg.Wait()
		close(img)
	}()

	for s := range img {
		fmt.Println(s)
	}

	fmt.Println(imgStr)

}

func DownloadImg(url string, path string, order int, retry int) string {
	regexFile := regexp.MustCompile(`(\w|\d|_)*\.(jpg|jpeg|png|gif)`)
	fileName := regexFile.FindString(url)
	if len(fileName) == 0 {
		fmt.Println("文件解析失败!")
		return ""
	}
	fileName = strconv.Itoa(order) + "_" + fileName
	//判断文件是否已经下载过
	if fileInfo, err := os.Stat(path + fileName); os.IsNotExist(err) || fileInfo.Size() < 5120 {

		//文件小于5120b
		if err == nil {
			os.Remove(path + fileName)
		}

		resp, err := http.Get(url)
		defer resp.Body.Close()

		if err != nil || resp.StatusCode != 200 {
			time.Sleep(time.Second)
			if retry > 0 {
				return DownloadImg(url, path, order, retry-1)
			} else {
				fmt.Println("下载图片失败!重试5次")
				return ""
			}
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			out, _ := os.Create(path + fileName)
			io.Copy(out, bytes.NewReader(body))
			out.Close()
			os.Chmod(path+fileName, os.ModePerm)
			return fileName
		}

	} else {
		return fileName
	}
}
