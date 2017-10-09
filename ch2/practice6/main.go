package main

import (
	"os/exec"
	"fmt"
)

func main() {
	param := []string{`%s`,`-H`,`Host: www.vipstation.com.hk`,`-H`,`User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:49.0) Gecko/20100101 Firefox/49.0`,`-H`,`Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`,`-H`,`Accept-Language: en-US,en;q=0.8,pl;q=0.6,fr;q=0.4,zh-CN;q=0.2,zh;q=0.2,de;q=0.2" --compressed -H Referer: http://www.vipstation.com.hk/catalog/Handbag/`,`-H`,`Cookie: ASP.NET_SessionId=npcrpacpeke4c1qbdlmbsaur; Lang=en-us; vip_lang=en-us; _ga=GA1.3.1633574956.1468234378; POP800_REFERRER_URL=https"%"253A"%"252F"%"252Fwww.google.com.hk"%"252F; POP800_VISIT_TIMES=2; POP800_VISITOR_NEW_IF=1; PAGE_VIEW_TIMES=22; POP800_VISITOR_ID_L=40A36A76C32CBC06AA1DC8C176B71C16; vieweditem=0=AL3516TC&1=BK308FTGSS; _gat=1`,`-H`,`Connection: keep-alive`,`-H`,`Upgrade-Insecure-Requests: 1`,`-H`,`Cache-Control: max-age=0`}
	param[0] = fmt.Sprintf(param[0],"http://www.vipstation.com.hk/products/Detailed/1BG8382A4AF0K44")
	res,_ := exec.Command("curl",param...).Output()
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
