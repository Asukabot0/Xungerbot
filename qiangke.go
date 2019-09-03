package main

import (
	"net/http"
	"io/ioutil"
	"regexp"
	"net/url"
	"net/http/cookiejar"
	"time"
	"fmt"
	"strings"
)

func main() {
	start := time.Now()
	var digitsRegexp = regexp.MustCompile(`on" value="(.+?)"`)
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: cookieJar,
	}
	resp, err := client.Get("https://cas.sustech.edu.cn/cas/login?service=http%3A%2F%2Fjwxt.sustech.edu.cn%2Fjsxsd%2F")
	if err != nil {
		println("CAS连接建立失败")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println("CAS连接建立失败")
	}
	println("CAS连接成功")

	exeRaw := digitsRegexp.FindSubmatch(body)[0]
	exe := string(exeRaw[11:len(exeRaw)-1])

	resp, err = client.PostForm("https://cas.sustech.edu.cn/cas/login?service=http%3A%2F%2Fjwxt.sustech.edu.cn%2Fjsxsd%2F",
		url.Values{"username": {"11811214"}, "password": {"biubiubiu23314"}, "execution": {exe}, "_eventId": {"submit"}, "geolocation": {""},})

	if err != nil {
		println("CAS认证失败")
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		println("CAS认证失败")
	}
	println("弱智科技教务系统启动")
	//fmt.Println(string(body))	
	i:=1
	println("少女祈祷中(" , i , "/ ? )")
	var isStarted = regexp.MustCompile(`href="(.+)" target="blank">进入选课`)
	resp, err = client.Get("http://jwxt.sustech.edu.cn/jsxsd/xsxk/xklc_list?Ves632DSdyV=NEW_XSD_PYGL")
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	for !isStarted.Match(body) {
		time.Sleep(time.Millisecond*200)
		println("少女祈祷中(" , i , "/ ? )")
		i++
		resp, err = client.Get("http://jwxt.sustech.edu.cn/jsxsd/xsxk/xklc_list?Ves632DSdyV=NEW_XSD_PYGL")
		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		//fmt.Println(string(body))	
	}
	keyRaw := isStarted.FindSubmatch(body)[0]
	key := string(keyRaw[6:len(keyRaw)-29])
	println("弱智科技抢课系统已开启")
	resp, err = client.Get("http://jwxt.sustech.edu.cn" + key)
	defer resp.Body.Close()
	println("开始抢课")

	flag1 := false

	for (!flag1){

		if !flag1 {
			resp, err = client.Get("http://jwxt.sustech.edu.cn/jsxsd/xsxkkc/ggxxkxkOper?jx0404id=201920201001174&xkzy=&trjf=")
			defer resp.Body.Close()
			body, err = ioutil.ReadAll(resp.Body)
			fmt.Println(string(body))	
		}
		if strings.Contains(string(body), "true"){
			flag1 = true
		}
		time.Sleep(time.Millisecond*200)
	}

	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)

}