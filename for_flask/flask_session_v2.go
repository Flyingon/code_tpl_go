package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/kcorlidy/dangerous"
)

func decodeFlaskCookie(secretKey, salt, cookieStr string) (interface{}, error) {
	ser := dangerous.Serializer{Secret: secretKey, Salt: salt}
	ser.Signerkwargs = map[string]interface{}{"KeyDerivation": "hmac", "DigestMethod": sha1.New}
	maxAge := int64(99999999999)
	return ser.URLSafeTimedLoads(cookieStr, maxAge)
}

func encodeFlaskCookie(secretKey, salt string, cookieDict interface{}) (interface{}, error) {
	ser := dangerous.Serializer{Secret: secretKey, Salt: salt}
	ser.Signerkwargs = map[string]interface{}{"KeyDerivation": "hmac", "DigestMethod": sha1.New}
	return ser.URLSafeTimedDumps(cookieDict)
}

func main() {
	secretKey := "CHANGE_ME_TO_A_COMPLEX_RANDOM_SECRET"
	salt := "cookie-session"
	data, err := decodeFlaskCookie(secretKey, salt,
		".eJwlz0FqAzEMheG7eD0LSZYsKZcZbI9EQ0IDM8mq9O41dP8-eP9P2fOM66vc3ucntrLfj3IrJmipHjDZjQaLJcLRq0Gi80QS6FgDhxOqmYS7Z28eEsMAhqp2TrURySBM0iZQjgrKPKeLd2ZgwdpSvFbv08wU4wDpk0XLVuZ15v5-PeJ7_QEm6qDVI0N9rbBRw0TuB44WdIhFAuJyz9fsz1hmwa18rjj_k6j8_gH2mEDE.YqgmMg.6zaZu8oS3PoaBCjppn4CATsO4uA")
	fmt.Println(data, err)

	enCodeData, err := encodeFlaskCookie(secretKey, salt, map[string]interface{}{
		"_fresh":     true,
		"_id":        "8518f79e0c4982b458f10da380f194c1250a13e1b9217885e999fa69e5eb800b777a4f78bef4054256c02fb30744cc959a44045136f59339ac88871ed05ac457",
		"csrf_token": "0422a0739efe79d0516261f14ad1b6e2d58ef011",
		"locale":     "en",
		"user_id":    "2",
	})
	fmt.Printf("%s %v\n", enCodeData, err)
}
