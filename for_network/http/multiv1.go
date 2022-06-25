package main

import (
	"fmt"
	"github.com/siddontang/go-log/log"
	"net/http"
)

// https://medium.com/rungo/running-multiple-http-servers-in-go-d15300f4e59f

func main() {

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello: "+r.Host)
	})

	go func() {
		log.Fatal(http.ListenAndServe(":9000", nil))
	}()

	log.Fatal(http.ListenAndServe(":9001", nil))
}
