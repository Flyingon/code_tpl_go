package main

import (
	"fmt"
	"github.com/buger/jsonparser"
	jsoniter "github.com/json-iterator/go"
	"log"
	stdhttp "net/http"
	"time"
)

type NewResponse struct {
	StdResponse stdhttp.ResponseWriter
}

func (nr NewResponse) Header() stdhttp.Header {
	return nr.StdResponse.Header()
}

func (nr NewResponse) Write(in []byte) (int, error) {
	if nIn, e := jsonparser.Set(in, []byte(fmt.Sprintf("%d", time.Now().Unix())), "timestamp"); e == nil {
		return nr.StdResponse.Write(nIn)
	}
	return nr.StdResponse.Write(in)
}

func (nr NewResponse) WriteHeader(statusCode int) {
	nr.StdResponse.WriteHeader(statusCode)
}

func main() {
	mux := stdhttp.NewServeMux()
	mux.HandleFunc("/", func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		nw := NewResponse{StdResponse: w}
		setRspTimeStamp(nw, r)
	})
	log.Fatal(stdhttp.ListenAndServe(":80", mux))
}

func setRspTimeStamp(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	rspMap := map[string]interface{}{
		"return_code": 0,
		"return_msg":  "",
	}
	rspJson, _ := jsoniter.Marshal(rspMap)
	num, e := w.Write(rspJson)
	fmt.Printf("AAAA: %+v, err: %+v\n", num, e)
	fmt.Printf("BBBB: %+v, err: %+v\n", r.Response, e)

}
