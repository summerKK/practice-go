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
	"sync"
	"strconv"
)

var (
	db            *sql.DB
	chanLoopCount chan int
	chanImgUrl    chan string
	chs           chan int
	imgDir        string
	wg            *sync.WaitGroup
	mu            sync.Mutex
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
	//控制goruntine
	chs = make(chan int, 20)
	imgDir = "D:/goprojects/src/practice/ch3/vipstation/pic/"

	logFile, err := os.OpenFile("D:/goprojects/src/practice/ch3/vipstation/crawl.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	wg = &sync.WaitGroup{}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(logFile)
}

func main() {

	picDir := "D:/goprojects/src/practice/ch3/vipstation/pic/"
	_, err := os.Stat(picDir)

	if err != nil {
		err := os.Mkdir(picDir, os.ModePerm)
		if err != nil {
			fmt.Println("create Folder error:", err)
			os.Exit(1)
		}
	}

	rows, err := db.Query("select media_gallery,sku from lux_products")
	if err != nil {
		fmt.Println("fetech data failed:", err.Error())
		return
	}

	go downloadImage()

	defer rows.Close()
	for rows.Next() {
		var media_gallery, sku string
		rows.Scan(&media_gallery, &sku)
		imgPath := strings.Split(media_gallery, ";")
		pos := 1
		for _, url := range imgPath {
			chanImgUrl <- sku + "@" + strconv.Itoa(pos) + "@" + url
			pos++
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
		//控制并发
		chs <- 0
		wg.Add(1)
		go saveImages(imgUrl)
	}
	wg.Wait()
	close(chanLoopCount)
}

//下载图片
func saveImages(imgUrl string) {
	defer func() {
		//下载完后释放资源,保证20个goruntine
		<-chs
		wg.Done()
	}()
	//sku@position@img
	//vip124@1@http://erp.vipstation.com.hk/upload/image/20170601/20170601164454_0421.jpg
	pictureInfo := strings.Split(imgUrl, "@")
	sku := pictureInfo[0]
	pos := pictureInfo[1]
	imgUrl = pictureInfo[2]
	log.Println(imgUrl)
	fileDir := imgDir + sku + "/"

	mu.Lock()
	_, err := os.Stat(fileDir)
	mu.Unlock()
	fmt.Println(err)
	if err != nil {
		err := os.Mkdir(fileDir, os.ModePerm)
		if err != nil {
			fmt.Println("create Folder error:", err)
			log.Println("create Folder error:", err)
		}
	}

	extension := imgUrl[strings.LastIndex(imgUrl, "."):]
	filename := fileDir + sku + "_" + pos + extension

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
