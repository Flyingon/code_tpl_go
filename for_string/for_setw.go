package main

import (
	"fmt"
	"github.com/Flyingon/go-common/util"
	"strconv"
	"time"
)

func main() {
	sTimeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println(util.Setw(16, sTimeStamp, "0"))
}
