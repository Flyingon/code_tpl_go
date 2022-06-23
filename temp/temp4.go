package main

import (
	"fmt"
	"github.com/siddontang/go-log/log"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

func main() {

	url := "http://127.0.0.1:8088/login/"
	method := "POST"

	payload := strings.NewReader("username=admin&password=Aa12345_")

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Errorf("Got error while creating cookie jar %s", err.Error())
	}

	client := &http.Client{
		//CheckRedirect: func(req *http.Request, via []*http.Request) error {
		//
		//},
		Jar: jar,
	}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	location, errLoc := res.Location()

	cookies := res.Cookies()
	fmt.Println("", res.StatusCode, location, errLoc, res.Header)
	for _, cookie := range cookies {
		fmt.Printf("cookie: %s\n", cookie)
	}
	fmt.Println(jar)
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(body))
}
