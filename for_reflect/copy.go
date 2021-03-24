package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	a string
	b *string
}

func InterfaceCopy(d interface{}) interface{} {
	fmt.Println(reflect.TypeOf(d).Kind())
	n := reflect.New(reflect.TypeOf(d))
	n.Elem().Set(reflect.ValueOf(d))
	return n.Interface()
}

func main() {
	a := "a"
	d := &Data{
		a: "a",
		b: &a,
	}
	c := InterfaceCopy(d)
	d.a = "c"
	e := c.(**Data)
	fmt.Println(c, *e)
}
