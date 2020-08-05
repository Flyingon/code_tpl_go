package main

import (
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/url"
)

func main() {
	u, _ := url.Parse("https://examplebucket-1250000000.cos.COS_REGION.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	// 1.永久密钥
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "COS_SECRETID",
			SecretKey: "COS_SECRETKEY",
		},
	})
}