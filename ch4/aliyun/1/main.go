package main

import (
	"os"
	"log"
	"bufio"
	"io/ioutil"
	"regexp"
	"strconv"
	"sort"
	"fmt"
)

var (
	logPath    string         = "ch4/aliyun/1/log/"
	tmpFile    string         = "ch4/aliyun/1/tmp/tmp.log"
	statisData map[string]int = map[string]int{}
)

type res []string

func main() {

	os.Remove(tmpFile)

	combineFiles()
	readLine()
	r := res(make([]string, 0, len(statisData)))
	for i, v := range statisData {
		r = append(r, i+"->"+strconv.Itoa(v))
	}
	sort.Sort(r)

	fmt.Println(r)
}

func (r res) Len() int {
	return len(r)
}

func (r res) Less(i, j int) bool {
	iCount := regexp.MustCompile(`->(\d+)`).FindStringSubmatch(r[i])[1]
	jCount := regexp.MustCompile(`->(\d+)`).FindStringSubmatch(r[j])[1]
	iInt, _ := strconv.Atoi(iCount)
	jInt, _ := strconv.Atoi(jCount)
	return iInt > jInt
}

func (r res) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLine() {
	f, err := os.Open(tmpFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			return
		}
		statistics(splitInfo(string(line)))
	}
}

func splitInfo(str string) string {
	res := regexp.MustCompile(`\[(?P<time>.+)\] (?P<res_ip>\S+) (?P<origin_ip>\S+) (?P<res_time>[0-9]+) "(?P<referer>[^"]+)" "(?P<req_url>[^"]+)" (?P<http_code>[0-9]+) (?P<req_size>\S+) (?P<res_size>[0-9]+) (?P<cache_status>\S+) "(?P<ua>.*)" "(?P<content_type>.*)"`).FindStringSubmatch(str)
	return res[2]

}

func statistics(key string) {
	if _, ok := statisData[key]; ok {
		statisData[key] += 1
	} else {
		statisData[key] = 1
	}
}

func combineFiles() {
	files, err := ioutil.ReadDir(logPath)
	if err != nil {
		check(err)
	}
	for _, v := range files {
		writeTo(logPath + v.Name())
	}
}

func writeTo(path string) {

	tmpFile, err := os.OpenFile(logPath+"../"+"tmp/"+"tmp.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
	defer tmpFile.Close()
	if err != nil {
		check(err)
	}
	data, err := ioutil.ReadFile(path)

	_, err = tmpFile.Write(data)
	check(err)

}
