package main

import (
	"../util"
	"strconv"
	"time"
	"fmt"
)

func main() {
	sTimeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println(util.Setw(16, sTimeStamp, "0"))
}
