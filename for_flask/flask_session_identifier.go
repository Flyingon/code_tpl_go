package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

// CreateIdentifier ...
func CreateIdentifier(remoteAddr, userAgent string) string {
	base := fmt.Sprintf("b'%s'|b'%s'", remoteAddr, userAgent)
	fmt.Println(base)
	data := sha512.Sum512([]byte(base))

	return hex.EncodeToString(data[:])
}

func main() {
	fmt.Println(CreateIdentifier("127.0.0.1",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36"))
}
