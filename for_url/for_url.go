package main

import (
	"fmt"
	"net/url"
)

func main() {
	addr := "https://www.baidu.com"
	u, err := url.Parse(addr)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Host)
}
