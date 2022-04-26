package main

import (
	"fmt"
	"reflect"
)

type STest struct {
	A string `db:"a"`
	B int    `db:"b"`
}

// GetTagKeysFromStruct ...
func GetTagKeysFromStruct(obj interface{}, tag string) []string {
	r := reflect.Indirect(reflect.ValueOf(obj))
	n := reflect.TypeOf(obj)
	if n.Kind() == reflect.Ptr {
		n = n.Elem()
	}

	fieldNum := r.NumField()
	tags := make([]string, 0)
	for i := 0; i < fieldNum; i++ {
		tag := n.Field(i).Tag.Get(tag)
		if len(tag) > 0 && tag != "-" && r.Field(i).CanInterface() {
			//加个``用来兼容mysql保留字比如desc、time之类
			tags = append(tags, fmt.Sprintf("`%s`", tag))
		}
	}
	return tags
}

func main() {
	sTest := &STest{
		A: "aaa",
		B: 111,
	}
	res := GetTagKeysFromStruct(sTest, "db")
	fmt.Println(res)
}
