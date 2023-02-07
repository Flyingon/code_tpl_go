package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) EchoName(name string) {
	fmt.Println("我的名字是：", name)
}

func main() {
	p := Person{Name: "无尘", Age: 18}

	v := reflect.ValueOf(p)

	// 获取方法控制权
	// 官方解释：返回v的名为name的方法的已绑定（到v的持有值的）状态的函数形式的Value封装
	mv := v.MethodByName("EchoName")
	// 拼凑参数
	args := []reflect.Value{reflect.ValueOf("wucs")}

	// 调用函数
	mv.Call(args)
}
