package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

type DEF struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ABC struct {
	Id   int    `json:"id"`
	Data string `json:"data"`
	Def  *DEF   `json:"def"`
}

func genAbc(i int) *ABC {
	abc := ABC{
		Id:   i,
		Data: fmt.Sprint(i),
		Def: &DEF{
			Id:   i,
			Name: fmt.Sprint(i),
		},
	}
	return &abc
}

func main() {
	var al []*ABC
	var a *ABC
	for i := 0; i < 10; i++ {
		a = genAbc(i)
		fmt.Println(*a)
		al = append(al, &ABC{
			Id:   a.Id,
			Data: a.Data,
			Def: &DEF{
				Id:   a.Def.Id,
				Name: a.Def.Name,
			},
		})
	}
	alStr, _ := jsoniter.MarshalToString(al)
	fmt.Println(alStr)
}
