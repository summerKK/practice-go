package main

import (
	"os/exec"
	"fmt"
	"encoding/json"
	"strings"
)

func main() {
	param := 		[]string{
		`http://www.italystation.com/search/`,
		`-H`,
		`Cookie: __cfduid=d78d764de2ef360773b8f685b1c5bbdca1468564432; PHPSESSID=48te5gsg3gkva37uoflgfn0eq7; app_version=2.2.3; Hm_lvt_d6c1c469611f4d18a7613d940ea5f4f8=1468564431; Hm_lpvt_d6c1c469611f4d18a7613d940ea5f4f8=1468564431; _dc_gtm_UA-11318687-1=1; pageReferrInSession=http%3A//www.italystation.com/welcome.html%3Fparam%3D%26sex%3D%26start%3D0%26searchtime%3D1; firstEnterUrlInSession=http%3A//www.italystation.com/search/%3Fcurrency%3DHKD%26lang%3Den_US; VisitorCapacity=1; _ga=GA1.2.778989578.1468564431; PHPSESSID=48te5gsg3gkva37uoflgfn0eq7; currencyId=HKD; langId=en_US`,
		`-H`,
		`Origin: http://www.italystation.com`,
		`-H`,
		`Accept-Encoding: gzip, deflate`,
		`-H`,
		`Accept-Language: zh-TW,zh;q=0.8,en-US;q=0.6,en;q=0.4`,
		`-H`,
		`User-Agent: Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36`,
		`-H`,
		`Content-Type: application/x-www-form-urlencoded; charset=UTF-8`,
		`-H`,
		`Accept: application/json, text/javascript, */*; q=0.01`,
		`-H`,
		`Referer: http://www.italystation.com/search/?stylegroup%5B%5D=Bags&start=0&searchtime=1`,
		`-H`,
		`X-Requested-With: XMLHttpRequest`,
		`-H`,
		`Connection: keep-alive`,
		`--data`,
		`stylegroup%5B%5D=Bags&start=0&searchtime=1`,
		`--compressed`}
	res,_ := exec.Command("curl",param...).Output()

	var data map[string]interface{}

	json.Unmarshal([]byte(strings.Split(string(res),"MyPhpRuntimeException")[0]),&data)

	fmt.Println(data)

	fmt.Println(string(res))

	hello := make([]string,10)

	fmt.Printf("%T\n",hello)

	set(hello)

	fmt.Println(hello)

}

func set(hello []string) {
	hello[1] = "hello"
	hello[2] = "world"
}
