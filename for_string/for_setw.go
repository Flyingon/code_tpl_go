package main

import (
	"code_tpl_go/util"
	"fmt"
	"strconv"
	"time"
)

func main() {
	sTimeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println(util.Setw(16, sTimeStamp, "0"))
}
