package main

import (
	"os/exec"
	"fmt"
	"github.com/hunterhug/GoSpider/query"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func main() {

	feature := map[string]string{"Reference Number": "summer"}

	if _, ok := feature["Reference Number"]; ok {
		delete(feature, "Reference Number")
	}
	fmt.Println(feature)


	curlParameter := []string{`https://www.1stdibs.com/fashion/handbags-purses-bags/luggage-travel-bags/fendi-like-new-snakeskin-leather-weekender-carryall-top-handle-tote-bag/id-v_3611303/`, ` -H "cookie: ab=as^^97:at^^57:br^^52:bu^^25:co^^58:cr^^03:de^^84:df^^19:hp^^51:lo^^93:ne^^85:pd^^55:re^^76:sa^^51:se^^39:sw^^45; __qca=P0-812246926-1505109318839; __gads=ID=f0286a151961c2cb:T=1505109370:S=ALNI_MaJmbj0lCVo37EdbWU8Gci7BlzB4Q; 3060738.3440491=706de5e4-2c9b-456e-ba44-ad030cb81003; SSID=CABWqx0cAAAAAACV995ZLJJABpX33lkBAAAAAADBhQZalffeWQANSEN3AAGzGg4AlffeWQEAW3gAA2U4DgCV995ZAQA; tracker_device=db70e544-8f4a-4df4-bf66-31f84b973db4; SSRT=9_feWQADAQ; guestId=40f32494-e97e-63d5-97f0-91c5fea58829; newUserPool=non-test-pool-nonpurchase; __insp_wid=759407264; __insp_nv=true; __insp_targlpu=aHR0cHM6Ly93d3cuMXN0ZGlicy5jb20vZmFzaGlvbi9oYW5kYmFncy1wdXJzZXMtYmFncy8^%^2FcGFnZT0x; __insp_targlpt=VmludGFnZSBhbmQgRGVzaWduZXIgQmFncyAtIDEzLDcwMiBGb3IgU2FsZSBhdCAxc3RkaWJz; __insp_norec_sess=true; _dc_gtm_UA-28839796-1=1; cp_session_in_progress=true; __utma=78698571.519788163.1505109314.1513182164.1513267552.8; __utmb=78698571.0.10.1513267552; __utmc=78698571; __utmz=78698571.1506327022.4.3.utmcsr=baidu^|utmccn=(organic)^|utmcmd=organic; __utmv=78698571.^|1=Guest^%^20Id=1ff19801-3a12-0411-ab84-e4dd0fc78928=1; lcv=^{^%^221^%^22:^{^%^22slot^%^22:1^%^2C^%^22name^%^22:^%^22Guest^%^20Id^%^22^%^2C^%^22value^%^22:^%^221ff19801-3a12-0411-ab84-e4dd0fc78928^%^22^%^2C^%^22scope^%^22:1^}^%^2C^%^223^%^22:^{^%^22slot^%^22:3^%^2C^%^22name^%^22:^%^22Not^%^20Registered^%^22^%^2C^%^22value^%^22:^%^22USD^%^20Default^%^22^%^2C^%^22scope^%^22:2^}^%^2C^%^224^%^22:^{^%^22slot^%^22:4^%^2C^%^22name^%^22:^%^22PDP-Available-Price^%^22^%^2C^%^22value^%^22:^%^22log=n^|pur=y^|neg=n^|ship=n^|ret=n^|pr=21450^|twl=y^|sale=n^|pd=12-13-2017^|sd=none^|id=v_3643261^|dd=v_1302^%^22^%^2C^%^22check_login^%^22:true^%^2C^%^22scope^%^22:3^}^%^2C^%^225^%^22:^{^%^22slot^%^22:5^%^2C^%^22name^%^22:^%^22Chicjoy^%^20-^%^20v_1302^%^22^%^2C^%^22value^%^22:^%^22pdp-available-price^%^22^%^2C^%^22scope^%^22:3^}^}; _ga=GA1.2.519788163.1505109314; _gid=GA1.2.1623245082.1513178957; _gaexp=GAX1.2.s6Wutwp7ReSzK0VVl7Q8Hw.17531.1; mp_1stdibs_mixpanel=^%^7B^%^22distinct_id^%^22^%^3A^%^20^%^2215e6f818cfb3c4-0705d15df8012f-5848211c-1fa400-15e6f818cfc2ca^%^22^%^7D; _uetsid=_uetc8503a13; __insp_slim=1513268337780; intercom-id-fnj80w14=f6f79a69-5572-440e-a358-8de5d1c45ad4" -H "accept-encoding: gzip, deflate, br" -H "accept-language: zh-CN,zh;q=0.9" -H "upgrade-insecure-requests: 1" -H "user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36" -H "accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8" -H "cache-control: max-age=0" -H "authority: www.1stdibs.com" -H "referer: https://www.1stdibs.com/fashion/handbags-purses-bags/" --compressed`}

	str := cur(curlParameter)

	fmt.Println(str)

	doc, _ := query.QueryString(str)
	val := map[string]string{}
	doc.Find(".pdp-details-entry-content-container > div > span").Each(func(i int, selection *goquery.Selection) {
		key := doc.Find(".pdp-details-entry-title").Eq(i).Text()
		val[key] = selection.Text()
	})
	fmt.Println(val)

	val = map[string]string{}
	doc.Find(".pdp-details-entry-content-container").Each(func(i int, selection *goquery.Selection) {
		key := doc.Find(".pdp-details-entry-title").Eq(i).Text()
		if key == "Dimensions" {
			val[key] = strings.TrimSpace(selection.Find("div").Eq(1).Find("span").Text())
		} else {
			val[key] = strings.TrimSpace(selection.Text())
		}
	})

	fmt.Println(val)
	//fmt.Println(str)

}

func cur(s []string) string {
	out, err := exec.Command("curl", s...).Output()
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}
