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
	"sort"
)

var (
	db            *sql.DB
	chanLoopCount chan int
	chanImgUrl    chan string
	chanMysql     chan map[string]string
	chs           chan int
	imgDir        string
	linuxImgDir   string
	wg            *sync.WaitGroup
	mu            sync.Mutex
	//统计每个sku下载文件的个数
	record map[string]int
	//存放所有sku的imgPath
	newImgPaht map[string][]string
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
	chanMysql = make(chan map[string]string)
	//控制goruntine
	chs = make(chan int, 20)
	imgDir = "D:/goprojects/src/practice/ch3/vipstation/pic1/"
	//服务器上面保存图片的位置.主要是用来更新数据库的图片位置
	linuxImgDir = "/luxsens/robot/imgs/vipAPI/"

	logFile, err := os.OpenFile("D:/goprojects/src/practice/ch3/vipstation/crawl.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	wg = &sync.WaitGroup{}
	record = map[string]int{}
	newImgPaht = map[string][]string{}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(logFile)
}

func main() {

	_, err := os.Stat(imgDir)

	if err != nil {
		err := os.Mkdir(imgDir, os.ModePerm)
		if err != nil {
			fmt.Println("create Folder error:", err)
			os.Exit(1)
		}
	}

	rows, err := db.Query(`select media_gallery,sku from lux_products where media_gallery like "http://%"`)
	if err != nil {
		fmt.Println("fetech data failed:", err.Error())
		return
	}

	go downloadImage()
	go updateProduct()

	go func() {
		count := 0
		for num := range chanLoopCount {
			count += num
			fmt.Println("count:", count)
		}
	}()

	defer rows.Close()
	for rows.Next() {

		var media_gallery, sku string
		rows.Scan(&media_gallery, &sku)

		//判断图片路径是否正确
		if !strings.Contains(media_gallery, "http://") {
			continue
		}

		imgPath := strings.Split(media_gallery, ";")
		//把要下载的文件存放下来
		mu.Lock()
		record[sku] = len(imgPath)
		mu.Unlock()
		pos := 1
		//创建文件夹
		fileDir := imgDir + sku + "/"
		_, err := os.Stat(fileDir)
		if os.IsNotExist(err) {
			err := os.Mkdir(fileDir, os.ModePerm)
			if err != nil {
				fmt.Println("create Folder error:", err)
				log.Println("create Folder error:", err)
			}
		}
		for _, url := range imgPath {
			chanImgUrl <- sku + "@" + strconv.Itoa(pos) + "@" + url
			pos++
		}
	}

	fmt.Println("downloaded")
	log.Println("downloaded")
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
	linuxFileDir := linuxImgDir + sku + "/"

	extension := imgUrl[strings.LastIndex(imgUrl, "."):]
	filename := fileDir + sku + "_" + pos + extension
	//服务器上面图片放置的位置(图片在本地下载完成打包上传到服务器)
	linuxFileName := linuxFileDir + sku + "_" + pos + extension

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

	//把当前的路径存下来
	mu.Lock()
	//这里发送服务器文件的位置
	newImgPaht[sku] = append(newImgPaht[sku], linuxFileName)
	mu.Unlock()
	chanLoopCount <- 1
	defer image.Close()
	image.Write(data)

	//如果为1代表文件下载完成,通过mysqlChannel更新文件(这里==1就代表最后一个文件)
	if record[sku] == 1 {
		newPath := ""
		sort.Strings(newImgPaht[sku])
		for _, tmp := range newImgPaht[sku] {
			newPath += tmp + ";"
		}
		newPath = newPath[:strings.LastIndex(newPath, ";")]
		chanMysql <- map[string]string{"sku": sku, "img": newPath}
	} else {
		mu.Lock()
		//对应的文件总数减一
		record[sku] -= 1
		mu.Unlock()
	}
}

func checkExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func updateProduct() {
	for row := range chanMysql {
		fmt.Println(row)
		stmt, _ := db.Prepare(`update lux_products set media_gallery = ? where sku = ? `)
		_, err := stmt.Exec(row["img"], row["sku"])
		if err != nil {
			fmt.Println(err)
		}
	}
}
