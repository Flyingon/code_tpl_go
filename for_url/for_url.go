package main

import (
	"fmt"
)

var ENV string

func setEnv(env string){
	ENV = env
	fmt.Println("!!!!", ENV)
}

func main() {
	//esUrl := "http://elastic:XXXXXXX@1.10.10.10:8080/user_info/user_info/68188145/_update"
	//u, err := url.Parse(esUrl)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(u)

	setEnv("!11232131")
	fmt.Println("~~~~~~", ENV)
}
