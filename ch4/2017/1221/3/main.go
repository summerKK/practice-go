package main

import (
	"reflect"
	"fmt"
	"strings"
	"regexp"
)

func main() {
	p := map[string]string{"summer": "Hello World"}

	kind := reflect.ValueOf(p).Kind()

	fmt.Println(kind)

	fmt.Println(reflect.Array)

	s := "hello;world"

	fmt.Println(strings.Split(s,";"))

	var mapping map[string]string = map[string]string{}

	mapping["summer"] = "hello"

	fmt.Println(mapping)

	format := "%scm"

	strs := []interface{}{"12"}

	fmt.Println(fmt.Sprintf(format,strs...))

	defer func() {
		fmt.Println("---hello world")
	}()

	str := `[25/Dec/2017:10:48:07 +0800] 123.125.67.208 - 187 "-" "GET https://chs.luxsens.com/m/zhcn/index.php/view/product/list.html/+productOrder/price,desc/+attr/item_style/18560/+price/0-500.9/+category/1556" 200 324 62451 MISS "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.93 Safari/537.36" "text/html; charset=UTF-8"`

	reg := regexp.MustCompile(`\[(?P<time>.+)\] (?P<res_ip>\S+) (?P<origin_ip>\S+) (?P<res_time>[0-9]+) "(?P<referer>[^"]+)" "(?P<req_url>[^"]+)" (?P<http_code>[0-9]+) (?P<req_size>\S+) (?P<res_size>[0-9]+) (?P<cache_status>\S+) "(?P<ua>.*)" "(?P<content_type>.*)"`)

	result := reg.FindStringSubmatch(str)


	fmt.Println(result)


	panic(1)

}
