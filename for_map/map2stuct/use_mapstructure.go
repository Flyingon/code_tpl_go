package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type RspBookSvrCommon struct {
	Result     int                               `json:"result"`
	ReturnCode int                               `json:"returnCode"`
	ReturnMsg  string                            `json:"returnMsg"`
	Data       map[string]map[string]interface{} `json:"data"`
}

func main() {
	dataJson := `{"data":{"bookInfoList":[{"authorTxt":"nihao","bookId":636,"bookName":"可恋","coverPic":"https://cdn.story.qq.com/icon/20190726202611291.png","discount":{"isExpired":1,"name":"无","type":0},"friendsDesc":"","launchStatus":30,"launchTime":1565193600,"popularityTxt":"1506人气值","popularityValue":1506,"ratingDetail":{"avgRatingScore":"0.0","avgStarNum":0,"isShowStar":0,"ratingDesc":"评分人数较少"},"status":1,"tags":"互动的真实男/女"},{"authorTxt":"Jock","bookId":665,"bookName":"恋旅行","coverPic":"https://cdn.story.qq.com/icon/20190723194616169.jpg","discount":{},"friendsDesc":"","launchStatus":30,"launchTime":0,"popularityTxt":"970人气值","popularityValue":970,"ratingDetail":{"avgRatingScore":"6.0","avgStarNum":"3.00","isShowStar":0,"ratingDesc":"评分人数较少"},"status":1,"tags":"声控福利1 / 互动的真实男/女"}]},"result":0,"returnCode":0,"returnMsg":"成功"}`
	dataMap := make(map[string]interface{})
	json.Unmarshal([]byte(dataJson), &dataMap)
	fmt.Println(dataMap)
	rbi := RspBookSvrCommon{}

	mapDecoderConfig := &mapstructure.DecoderConfig{TagName: "json", Result: &rbi, WeaklyTypedInput: true, Metadata: nil}
	mapDecoder, _ := mapstructure.NewDecoder(mapDecoderConfig)
	err := mapDecoder.Decode(dataMap)

	fmt.Println("---------")
	fmt.Println(rbi, err)
}