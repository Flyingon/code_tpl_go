package main

import (
	"bytes"
	test "code_tpl_go/for_proto/proto"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
)

// Marshaler jsonpb序列化结构体，可自己更改参数
var Marshaler = jsonpb.Marshaler{EmitDefaults: true, OrigName: true, EnumsAsInts: true}

// Unmarshaler jsonpb反序列化结构体，可自己更改参数
var Unmarshaler = jsonpb.Unmarshaler{AllowUnknownFields: true}

func main() {
	res := test.Result{}
	data := `{"JSON": {"a": 1, "b": "2", "c": {"c": "c"}}, "JSON2": {"a": 111111111122222222, "b": "", "c": {"c": "c"}}}`
	e := Unmarshaler.Unmarshal(bytes.NewReader([]byte(data)), &res)
	//e := json.Unmarshal([]byte(data), &res)
	fmt.Println(e, res.JSON2)
	fmt.Println(uint64(res.JSON2["a"].GetNumberValue()))
}
