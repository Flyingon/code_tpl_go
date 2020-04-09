package main

import (
	"code_tpl_go/investigate/govaluate/ruleExecutor"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/buger/jsonparser"
)

var evaluableFunctionsV1 = map[string]govaluate.ExpressionFunction{
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
}



func testFunc() {
	parameters := make(map[string]interface{}, 8)
	parameters["st_src"] = "23"
	parameters["qeh_flag_top"] = "1"
	parameters["qeh_minivideo_info"] = `{"Bgm":"","BgmId":"","BgmPicUrl":"","Character":2,"Excellent":"0","Feedid":"","FlagTop":0,"Gif":"http://puui.qpic.cn/vshare_gif/0/o0895ulyyqg_540_960.gif/0","Interactive_tag":"","Is_interactive":0,"Location":"","Recommand":"","Size":"720*1280","Style":100,"Topic":"","TopicId":"","WsDownloadurl":""}`

	re := ruleExecutor.RuleExecutor{
		Rule:  "jsonGetInt(qeh_minivideo_info, 'FlagTop') == 1 || qeh_flag_top >= '1'",
		Input: parameters,
		FuncMap: evaluableFunctionsV1,
	}
	err := re.GetResult()
	fmt.Println(re.Result, err)

	outPut := ruleExecutor.OutPut{HbaseData: map[string]interface{}{
		"is_flag_top": 1,
		"flag_top_detail": map[string]string{
			"qeh_minivideo_info.FlagTop": "Input.qeh_minivideo_info.FlagTop",
			"qeh_flag_top":               "Input.qeh_minivideo_info.FlagTop",
			"st_src":                     "Input.st_src",
		},
	}}
	fmt.Println(outPut)

	ruleRsp := ruleExecutor.RuleReplace("rule.jsonGetInt(qeh_minivideo_info, 'FlagTop')", parameters)
	fmt.Println(ruleRsp)
}

func main() {
	//testFunc()
	expression, err := govaluate.NewEvaluableExpression("! (foo > 4)")
	parameters := make(map[string]interface{}, 8)
	parameters["foo"] = 3

	result, err := expression.Evaluate(parameters)
	fmt.Println(result, err)
}
