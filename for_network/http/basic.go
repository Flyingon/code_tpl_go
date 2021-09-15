package main

import (
	"fmt"
	"github.com/buger/jsonparser"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
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
	log.Fatal(stdhttp.ListenAndServe(":8080", mux))
}

func setRspTimeStamp(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	data, err := ioutil.ReadAll(r.Body)
	fmt.Printf("req-addr: %s, req-body: %s, err: %+v\n", r.RemoteAddr, data, err)
	rspMap := map[string]interface{}{
		"return_code": 0,
		"return_msg":  "",
	}
	rspJson, _ := jsoniter.Marshal(rspMap)
	_, _ = w.Write(rspJson)
	//num, e := w.Write(rspJson)
	//fmt.Printf("rsp-write: %+v, err: %+v\n", num, e)
	//fmt.Printf("rsp: %+v, err: %+v\n", r.Response, e)

}
