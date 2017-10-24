package main

import (
	"os"
	"fmt"
	"log"
	"runtime"
	"flag"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"strconv"
	"net/url"
	"net/http"
	"io/ioutil"
)

const HOST string = "http://www.aitaotu.com"
const DOC_URL string = "https://www.aitaotu.com/search/%E6%B8%85%E7%BA%AF/"

var (
	ch1    chan string
	ch2    chan string
	ch3    chan int
	imgDir string
	suffix string = ".html"
)

func init() {
	ch1 = make(chan string, 20)
	ch2 = make(chan string, 1000)
	ch3 = make(chan int, 1000)

	logFile, err := os.OpenFile("D:/goprojects/src/practice/ch3/crawl-pic-practice/crawl.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(logFile)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	imgPath := flag.String("imgPath", "D:/goprojects/src/practice/ch3/crawl-pic-practice/pic/", "where is image to save")

	flag.Parse()

	imgDir = *imgPath

	//check folder exitst
	file, err := os.Stat(imgDir)

	if err != nil || !file.IsDir() {
		err := os.Mkdir(imgDir, os.ModePerm)
		if err != nil {
			fmt.Println("create dir failed")
			os.Exit(1)
		}
	}

	go getListUrl()
	go parseListUrl()
	go downloadImg()

	count := 0
	for num := range ch3 {
		count = count + num
		fmt.Println("count:", count)
	}

	fmt.Println("crawl end")
}

//获取列表的每个item的URL
func getListUrl() {
	doc, err := goquery.NewDocument(DOC_URL)
	if err != nil {
		fmt.Println("getListUrl:", err)
		os.Exit(1)
	}

	doc.Find(".picbox a	").Each(func(i int, selection *goquery.Selection) {
		text, _ := selection.Attr("href")
		listUrl := HOST + text
		ch1 <- listUrl
	})
}

//通过列表的URL统计每个item的图片总页数
func parseListUrl() {
	for listUrl := range ch1 {
		pageCount := getPageCount(listUrl)
		preFix := strings.TrimRight(listUrl, suffix)
		for i := 1; i <= pageCount; i++ {
			imgListUrl := preFix + "_" + strconv.Itoa(i) + suffix
			ch2 <- imgListUrl
		}
	}
}

// 解析图片URL
func downloadImg() {
	for imgUrl := range ch2 {
		doc, _ := goquery.NewDocument(imgUrl)
		doc.Find("#big-pic a img").Each(func(i int, selection *goquery.Selection) {
			imgSrc, _ := selection.Attr("src")
			go func() {
				saveImages(imgSrc)
			}()
		})
	}
}

//下载图片
func saveImages(s string) {
	log.Println(s)
	u, err := url.Parse(s)
	if err != nil {
		log.Println("parse url failed:", s, err)
		return
	}

	//去掉左边的"/"
	tmp := strings.TrimLeft(u.Path, "/")
	fileName := imgDir + strings.ToLower(strings.Replace(tmp, "/", "-", -1))

	exists := checkExists(fileName)

	if exists {
		return
	}

	resp, err := http.Get(s)

	if err != nil {
		log.Println("get imgUrl failed:", err)
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("read data failed:", s, err)
		return
	}

	image, err := os.Create(fileName)

	if err != nil {
		log.Println("create file failed:", fileName, err)
		return
	}

	ch3 <- 1
	defer image.Close()
	image.Write(data)
}

func getPageCount(url string) (count int) {
	count = 0
	doc, _ := goquery.NewDocument(url)
	doc.Find(".pages li a").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		if text == "末页" {
			lastPageUrl, _ := selection.Attr("href")
			preFix := strings.Trim(lastPageUrl, suffix)
			index := strings.Index(preFix, "_")
			lastPageNum := preFix[index+1:]
			count, _ = strconv.Atoi(lastPageNum)
		}
	})

	return count
}

func checkExists(s string) bool {
	_, err := os.Stat(s)
	return err == nil
}
