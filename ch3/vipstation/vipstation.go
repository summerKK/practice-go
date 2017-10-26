package main

import (
	"database/sql"
	"fmt"
	// 引入数据库驱动注册及初始化
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"log"
	"net/http"
	"io/ioutil"
	"os"
	"net/url"
)

var (
	db            *sql.DB
	chanLoopCount chan int
	chanImgUrl    chan string
	imgDir        string
)

func init() {
	connection, err := sql.Open("mysql", "root:177678012@tcp(127.0.0.1:3306)/vp53dev?charset=utf8")
	if err != nil {
		fmt.Println("failed to open database:", err.Error())
		return
	}
	db = connection

	chanImgUrl = make(chan string, 100)
	chanLoopCount = make(chan int, 1000)
	imgDir = "D:/goprojects/src/practice/ch3/vipstation/pic/"
}

func main() {

	rows, err := db.Query("select media_gallery from lux_products")
	if err != nil {
		fmt.Println("fetech data failed:", err.Error())
		return
	}

	go downloadImage()

	defer rows.Close()
	for rows.Next() {
		var media_gallery string
		rows.Scan(&media_gallery)
		imgPath := strings.Split(media_gallery, ";")
		for _, url := range imgPath {
			chanImgUrl <- url
		}
	}

	count := 0
	for num := range chanLoopCount {
		count += num
		fmt.Println("count:", count)
	}
	fmt.Println("downloaded")

}

//解析图片url
func downloadImage() {
	for imgUrl := range chanImgUrl {
		go func() {
			saveImages(imgUrl)
		}()
	}
}

//下载图片
func saveImages(imgUrl string) {
	log.Println(imgUrl)
	u, err := url.Parse(imgUrl)
	if err != nil {
		log.Println("parse url failed:", imgUrl, err)
		return
	}

	//去掉最左边的'/'
	tmp := strings.TrimLeft(u.Path, "/")
	filename := imgDir + strings.ToLower(strings.Replace(tmp, "/", "-", -1))

	exists := checkExists(filename)
	if exists {
		return
	}

	response, err := http.Get(imgUrl)
	if err != nil {
		log.Println("get img_url failed:", err)
		return
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("read data failed:", imgUrl, err)
		return
	}

	image, err := os.Create(filename)
	if err != nil {
		log.Println("create file failed:", filename, err)
		return
	}

	chanLoopCount <- 1
	defer image.Close()
	image.Write(data)
}

func checkExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
