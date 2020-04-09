package main

import (
	"code_tpl_go/investigate/govaluate/ruleExecutor"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/buger/jsonparser"
)

const (
	maxListLen = 5000
)

var evaluableFunctions = map[string]govaluate.ExpressionFunction{
	"jsonGetInt": func(args ...interface{}) (interface{}, error) {
		data, ok := args[0].(string)
		if !ok {
			return nil, fmt.Errorf("args.0 type is not string")
		}
		var keys []string
		for _, arg := range args[1:] {
			if key, ok := arg.(string); ok {
				keys = append(keys, key)
			} else {
				return nil, fmt.Errorf("key[%v] is not valid", arg)
			}

		}
		val, err := jsonparser.GetInt([]byte(data), keys...)
		if err != nil {
			return nil, err
		}
		return (float64)(val), nil
	},
	"jsonGetStr": func(args ...interface{}) (interface{}, error) {
		data, ok := args[0].(string)
		if !ok {
			return nil, fmt.Errorf("args.0 type is not string")
		}
		var keys []string
		for _, arg := range args[1:] {
			if key, ok := arg.(string); ok {
				keys = append(keys, key)
			} else {
				return nil, fmt.Errorf("key[%v] is not valid", arg)
			}

		}
		val, err := jsonparser.GetString([]byte(data), keys...)
		if err != nil {
			return nil, err
		}
		return (string)(val), nil
	},
	"mapGetInt": func(args ...interface{}) (interface{}, error) {
		data, ok := args[0].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("args.0 type is not map[string]interface{}")
		}
		key := args[1].(string)
		if val, fond := data[key]; fond {
			valInt, ok := val.(int)
			if ok {
				return (float64)(valInt), nil
			} else {
				return nil, fmt.Errorf("key[%s] type is not int", key)
			}
		} else {
			return nil, fmt.Errorf("key[%s] path not found", key)
		}
	},
	"mapGetStr": func(args ...interface{}) (interface{}, error) {
		data, ok := args[0].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("args.0 type is not map[string]interface{}")
		}
		key := args[1].(string)
		if val, fond := data[key]; fond {
			valStr, ok := val.(string)
			if ok {
				return (string)(valStr), nil
			} else {
				return nil, fmt.Errorf("key[%s] type is not string", key)
			}
		} else {
			return nil, fmt.Errorf("key[%s] path not found", key)
		}
	},
	"isListContain":     IsListContain,
	"isListContainList": IsListContainList,
}

// IsListContainList 判断参数args[0](interface{}) 是否在 args[1]([]interface{}) 里面
func IsListContain(args ...interface{}) (interface{}, error) {
	for i, arg := range args {
		fmt.Printf("isListContain params, index: %d, arg: %+v\n", i, arg)
	}
	if len(args) != 2 {
		return nil, fmt.Errorf("args num is not matched")
	}
	key, ok := args[0].(interface{})
	if !ok {
		return nil, fmt.Errorf("args.0 type is not interface{}")
	}
	data, ok := args[1].([]interface{})
	if !ok {
		return nil, fmt.Errorf("args.1 type is not []interface{}")
	}
	if len(data) > maxListLen {
		return nil, fmt.Errorf("list len is exceed max[%d]", maxListLen)
	}
	for _, v := range data {
		if key == v {
			return true, nil
		}
	}
	return false, nil
}

// IsListContainList 判断参数args[1]([]interface{}) 是否包含 args[2]([]interface{})
// args[0] 不好支持传入数组，具体原因再分析下
func IsListContainList(args ...interface{}) (interface{}, error) {
	for i, arg := range args {
		fmt.Printf("IsListContainList params, index: %d, arg: %+v\n", i, arg)
	}
	if len(args) != 3 {
		return nil, fmt.Errorf("args num is not matched")
	}
	fmt.Printf("listA: %v\n", args[1])
	fmt.Printf("listB: %v\n", args[2])
	listA, ok := args[1].([]interface{})
	if !ok {
		return nil, fmt.Errorf("IsListContainList params, args.1 type is not []interface{}")
	}
	listB, ok := args[2].([]interface{})
	if !ok {
		return nil, fmt.Errorf("IsListContainList params, args.2 type is not []interface{}")
	}
	if len(listB) > maxListLen || len(listA) > maxListLen {
		return nil, fmt.Errorf("list len is exceed max[%d]", maxListLen)
	}
	if len(listB) > len(listA) {
		return false, nil
	}
	mapA := make(map[interface{}]bool)
	for _, k := range listA {
		mapA[k] = true
	}
	ret := true
	for _, e := range listB {
		if !mapA[e] {
			ret = false
		}
	}
	return ret, nil
}

func testListContain() {
	parameters := make(map[string]interface{}, 8)
	//parameters["dataA"] = []interface{}{[]interface{}{"b", "1", 1, 3}, []interface{}{"b", "1", 1, 3}}
	parameters["dataA"] = []interface{}{"b", "1", 1, 3}

	re := ruleExecutor.RuleExecutor{
		Rule: "isListContain('b', dataA)",
		//Rule: "",
		Input:   parameters,
		FuncMap: evaluableFunctions,
	}
	err := re.GetResult()
	fmt.Println(re.Result, err)
}

func testListContainList() {
	parameters := make(map[string]interface{}, 8)
	//parameters["dataA"] = []interface{}{[]interface{}{"b", "1", 1, 3}, []interface{}{"b", "1", 1, 3}}
	parameters["dataA"] = []interface{}{"b", "1", 1, 3}

	re := ruleExecutor.RuleExecutor{
		Rule: "isListContainList('', dataA, ('a', 'b'))",
		//Rule: "",
		Input:   parameters,
		FuncMap: evaluableFunctions,
	}
	err := re.GetResult()
	fmt.Println(re.Result, err)
}

func main() {
	testListContain()
	//testListContainList()
}
