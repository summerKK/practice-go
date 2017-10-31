package threadPool

import (
	"testing"
	"fmt"
	"strings"
	"os"
	"net/http"
	"io"
	"time"
)

func TestThreadPool(t *testing.T) {
	urls := []string{
		"http://dlsw.baidu.com/sw-search-sp/soft/44/17448/Baidusd_Setup_4.2.0.7666.1436769697.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/3a/12350/QQ_V7.4.15197.0_setup.1436951158.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/9d/14744/ChromeStandalone_V43.0.2357.134_Setup.1436927123.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/6f/15752/iTunes_V12.2.1.16_Setup.1436855012.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/70/17456/BaiduAn_Setup_5.0.0.6747.1435912002.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/40/12856/QIYImedia_1_06_v4.0.0.32.1437470004.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/42/37473/BaiduSoftMgr_Setup_7.0.0.1274.1436770136.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/49/16988/YoudaoNote_V4.1.0.300_setup.1429669613.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/55/11339/bdbrowserSetup-7.6.100.2089-1212_11000003.1437029629.exe",
		"http://dlsw.baidu.com/sw-search-sp/soft/53/21734/91zhushoupc_Windows_V5.7.0.1633.1436844901.exe",
	}

	pool := new(GorountinePool)
	//开三个goruntine下载文件
	pool.Init(3, len(urls))

	//添加任务
	for i := range urls {
		url := urls[i]
		pool.AddTask(func() error {
			return download(url)
		})
	}

	//设置任务状态
	isFinish := false
	pool.SetFinishCallback(func() {
		func(isFinish *bool) {
			//执行完毕回调函数把任务状态设为true
			//*isFinish指针引用可以改变地址对应变量的值
			*isFinish = true
			//这里&isFinish传指针(作用域)
		}(&isFinish)
	})

	//任务开始
	pool.Start()

	//轮询查看任务是否结束
	for !isFinish {
		//优化代码,当任务状态还没有改变的时候查询一次休眠100毫秒
		time.Sleep(time.Microsecond * 100)
	}

	//结束任务
	pool.Stop()
	fmt.Println("所有任务执行完成")

}

func download(s string) error {
	fmt.Println("开始下载:", s)
	sp := strings.Split(s, "/")
	fileName := sp[len(sp)-1]

	saveDIr := "D:/goprojects/src/practice/ch3/threadPool/down/"
	if _, err := os.Stat(saveDIr); os.IsNotExist(err) {
		err := os.Mkdir(saveDIr, os.ModePerm)
		if err != nil {
			fmt.Println("创建文件夹失败,", saveDIr, err)
			os.Exit(1)
		}
	}

	file, err := os.Create("D:/goprojects/src/practice/ch3/threadPool/down/" + fileName)
	if err != nil {
		return err
	}

	res, err := http.Get(s)
	if err != nil {
		return err
	}

	length, err := io.Copy(file, res.Body)
	if err != nil {
		return err
	}

	fmt.Println("##下载完成!", s, "文件长度:", length)

	return nil

}
