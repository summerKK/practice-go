package main

import (
	"fmt"
	"strings"
	"strconv"
	"net/http"
	"net/url"
	"io/ioutil"
	"os"
	"log"
	"runtime"
	"flag"
	"github.com/PuerkitoBio/goquery"
)

const HOST string = "http://www.aitaotu.com"
const DOC_URL string = "https://www.aitaotu.com/search/%E6%B8%85%E7%BA%AF/"

var (
	ch1     chan string
	ch2     chan string
	ch3     chan int
	img_dir string
)

//初始化变量
func init() {
	ch1 = make(chan string, 20)
	ch2 = make(chan string, 1000)
	ch3 = make(chan int, 1000)

	logfile, err := os.OpenFile("D:/goprojects/src/practice/ch3/crawl-pic/crawl.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(logfile)
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	img_path := flag.String("img_path", "D:/goprojects/src/practice/ch3/crawl-pic/pic/", "where  is image to save")
	flag.Parse()

	img_dir = *img_path
	//检查目录是否存在
	file, err := os.Stat(img_dir)
	if err != nil || !file.IsDir() {
		dir_err := os.Mkdir(img_dir, os.ModePerm)
		if dir_err != nil {
			fmt.Println("create dir failed")
			os.Exit(1)
		}
	}

	go getListUrl()
	go parseListUrl()
	go downloadImage()

	count := 0
	for num := range ch3 {
		count = count + num
		fmt.Println("count:", count)
	}
	fmt.Println("crawl end")
}

func getListUrl() {
	doc, err := goquery.NewDocument(DOC_URL)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	doc.Find(".picbox").Each(func(i int, s *goquery.Selection) {
		text, _ := s.Find("a").Attr("href")
		list_url := HOST + text
		ch1 <- list_url
	})
}

//根据模块和总数据列出所有的图片页面
func parseListUrl() {
	suffix := ".html"
	for list_url := range ch1 {
		page_count := getPageCount(list_url)
		prefix := strings.TrimRight(list_url, suffix)
		for i := 1; i <= page_count; i++ {
			img_list_url := prefix + "_" + strconv.Itoa(i) + suffix
			ch2 <- img_list_url
		}
	}
}

//获取总页数
func getPageCount(list_url string) (count int) {
	count = 0
	doc, _ := goquery.NewDocument(list_url)
	doc.Find(".pages ul li").Each(func(i int, s *goquery.Selection) {
		text := s.Find("a").Text()
		if text == "末页" {
			last_page_url, _ := s.Find("a").Attr("href")
			prefix := strings.Trim(last_page_url, ".html")
			index := strings.Index(prefix, "_")
			last_page_num := prefix[index+1:]
			page_num, _ := strconv.Atoi(last_page_num)
			count = page_num
		}
	})
	return count
}

//解析图片url
func downloadImage() {
	for img_list_url := range ch2 {
		doc, _ := goquery.NewDocument(img_list_url)
		doc.Find("#big-pic p a").Each(func(i int, s *goquery.Selection) {
			img_url, _ := s.Find("img").Attr("src")
			go func() {
				saveImages(img_url)
			}()
		})
	}
}

//下载图片
func saveImages(img_url string) {
	log.Println(img_url)
	u, err := url.Parse(img_url)
	if err != nil {
		log.Println("parse url failed:", img_url, err)
		return
	}

	//去掉最左边的'/'
	tmp := strings.TrimLeft(u.Path, "/")
	filename := img_dir + strings.ToLower(strings.Replace(tmp, "/", "-", -1))

	exists := checkExists(filename)
	if exists {
		return
	}

	response, err := http.Get(img_url)
	if err != nil {
		log.Println("get img_url failed:", err)
		return
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("read data failed:", img_url, err)
		return
	}

	image, err := os.Create(filename)
	if err != nil {
		log.Println("create file failed:", filename, err)
		return
	}

	ch3 <- 1
	defer image.Close()
	image.Write(data)
}

func checkExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
