package main

import (
	"fmt"
	"net/url"
)

func FormatQueryParams(params map[string]string) string {
	if params == nil {
		return ""
	}
	reqBodyVal := url.Values{}
	for k, v := range params {
		reqBodyVal.Add(k, v)
	}
	return reqBodyVal.Encode()
}

func main() {
	fmt.Println(FormatQueryParams(map[string]string{
		"size": "10",
		"from": "220",
	}))
}
