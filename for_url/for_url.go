package main

import (
	"fmt"
)

func main() {
	esUrl := "http://elastic:XXXXXXX@1.10.10.10:8080/user_info/user_info/68188145/_update"
	u, err := url.Parse(esUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)
}
