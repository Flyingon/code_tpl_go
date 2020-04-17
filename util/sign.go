package util

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func MakeSign(time uint64, appid string, skey string) string {
	timeString := strconv.FormatUint(time, 10)
	h := md5.New()
	h.Write([]byte(appid + skey + timeString)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	sign := hex.EncodeToString(cipherStr)
	return sign
}
