package main

import (
	socks5 "github.com/armon/go-socks5"
	"log"
	"os"
)

func startSvr() {
	// Create a socks server
	creds := socks5.StaticCredentials{
		"foo": "bar",
	}
	cator := socks5.UserPassAuthenticator{Credentials: creds}
	conf := &socks5.Config{
		AuthMethods: []socks5.Authenticator{cator},
		Logger:      log.New(os.Stdout, "", log.LstdFlags),
	}
	serv, err := socks5.New(conf)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	serv.ListenAndServe("tcp", "0.0.0.0:12365")
}

func main() {
	startSvr()
}
