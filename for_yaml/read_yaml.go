package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
	"reflect"
)

type BPtr struct {
	B  string                 `yaml:"b"`
	BC map[string]interface{} `yaml:"bc"`
	BE []string               `yaml:"be"`
}

type Data struct {
	A string                 `yaml:"a"`
	B *BPtr                  `yaml:"b"`
	C map[string]interface{} `yaml:"c"`
	E []string               `yaml:"e"`
}

func main() {
	//j1 := `{"A":"a","B":{"B":"b"},"C":{"c":1,"d":2}, "E": ["1", "2"]}`
	//j2 := `{"A":"a2","B":{"B":"b2"},"C":{"c":11,"d":22}, "E": ["3", "4"]}`
	dd1 := Data{
		A: "a",
		B: &BPtr{
			B: "b",
			BC: map[string]interface{}{
				"111": 111,
				"222": 222,
			},
			BE: []string{"111", "222"},
		},
		C: map[string]interface{}{
			"c": 1,
			"d": 2,
			"pp": &BPtr{
				B: "qqqqqq",
				BC: map[string]interface{}{
					"pb1111": 1111,
					"pb2222": 2222,
				},
				BE: []string{"1111", "2222"},
			},
		},
		E: []string{"1", "2"},
	}
	dd2 := Data{
		A: "aa",
		B: &BPtr{
			B: "bb",
			BC: map[string]interface{}{
				"3333": 3333,
				"4444": 4444,
			},
			BE: []string{"1111", "2222"},
		},
		C: map[string]interface{}{
			"c": 11,
			"d": 21,
			"pp": &BPtr{
				B: "pppbbb",
				BC: map[string]interface{}{
					"pb3333": 5555,
					"pb4444": 6666,
				},
				BE: []string{"3333", "4444"},
			},
		},
		E: []string{"33", "44"},
	}
	//j1, _ := jsoniter.MarshalToString(d)
	y1, _ := yaml.Marshal(dd1)
	y2, _ := yaml.Marshal(dd2)
	//fmt.Println("yy: \n", string(yy))

	var d1 interface{}
	d0 := &Data{}

	d1 = d0
	yaml.Unmarshal(y1, d1)
	fmt.Println(d1)
	spew.Dump(d1)
	fmt.Println("--------------------------------------------------------")

	newType := reflect.TypeOf(d1).Elem()
	d2 := reflect.New(newType).Interface()

	yaml.Unmarshal(y2, d2)
	fmt.Println(d2, reflect.TypeOf(d2))
	spew.Dump(d2)
	fmt.Println("--------------------------------------------------------")

	d1Val := reflect.ValueOf(d1).Elem()
	d1Val.Set(reflect.ValueOf(d2).Elem())

	fmt.Println(d1)
	spew.Dump(d1)
	fmt.Println("--------------------------------------------------------")
}
