package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/buger/jsonparser"
	"strings"
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
}

type OutPut struct {
	HbaseData map[string]interface{} `json:"hbase_data"`
}

type Rule struct {
	Input  map[string]interface{} `json:"input"`
	Rule   string                 `json:"rule"`
	Output map[string]interface{} `json:"output"`
}

type Result struct {
	EventID   uint64
	EventName string
	Data      map[string]interface{}
	Err       error
}

type RuleExecutor struct {
	Input  map[string]interface{}
	Rule   string
	Result interface{}
	Err    error
}

func (re *RuleExecutor) GetResult() error {
	var err error
	if re.Rule == "" {
		err = fmt.Errorf("rule conifg is nil")
		re.Err = err
		return err
	}
	err = re.ExecuteRule()
	if err != nil {
		err = fmt.Errorf("rule execute err: %v", err)
		re.Err = err
		return err
	}
	return nil
}

func (re *RuleExecutor) ExecuteRule() error {
	expression, err := govaluate.NewEvaluableExpressionWithFunctions(re.Rule, evaluableFunctions)
	if err != nil {
		fmt.Printf("[ERROR] rule[%s] init err: %v\n", err)
		return err
	}
	result, err := expression.Evaluate(re.Input)
	if err != nil {
		fmt.Printf("[ERROR] rule[%s] result: %v, err: %v\n", re.Rule, result, err)
		return err
	}
	re.Result = result
	re.Err = err
	return nil
}

// RuleReplace 根据规则获取字段结果
// 为了兼容字段需要根据规则生成
func RuleReplace(data string, dataMap map[string]interface{}) interface{} {
	fmt.Printf("result rule replace, data: %s\n", data)
	var ret interface{}
	if strings.HasPrefix(data, "rule.") {
		rule := data[5:]
		re := RuleExecutor{
			Input:  dataMap,
			Rule:   rule,
			Result: nil,
			Err:    nil,
		}
		err := re.GetResult()
		if err != nil {
			fmt.Printf("[ERROR] rule replace in format err: %v, data: %s\n", err, data)
			return data
		}
		ret = re.Result
	} else {
		ret = data
	}
	fmt.Printf("result rule replace, data: %s, ret: %+v\n", data, ret)
	return ret
}

func testFunc() {
	parameters := make(map[string]interface{}, 8)
	parameters["st_src"] = "23"
	parameters["qeh_flag_top"] = "1"
	parameters["qeh_minivideo_info"] = `{"Bgm":"","BgmId":"","BgmPicUrl":"","Character":2,"Excellent":"0","Feedid":"","FlagTop":0,"Gif":"http://puui.qpic.cn/vshare_gif/0/o0895ulyyqg_540_960.gif/0","Interactive_tag":"","Is_interactive":0,"Location":"","Recommand":"","Size":"720*1280","Style":100,"Topic":"","TopicId":"","WsDownloadurl":""}`

	re := RuleExecutor{
		Rule:  "jsonGetInt(qeh_minivideo_info, 'FlagTop') == 1 || qeh_flag_top >= '1'",
		Input: parameters,
	}
	err := re.GetResult()
	fmt.Println(re.Result, err)

	outPut := OutPut{HbaseData: map[string]interface{}{
		"is_flag_top": 1,
		"flag_top_detail": map[string]string{
			"qeh_minivideo_info.FlagTop": "Input.qeh_minivideo_info.FlagTop",
			"qeh_flag_top":               "Input.qeh_minivideo_info.FlagTop",
			"st_src":                     "Input.st_src",
		},
	}}
	fmt.Println(outPut)

	ruleRsp := RuleReplace("rule.jsonGetInt(qeh_minivideo_info, 'FlagTop')", parameters)
	fmt.Println(ruleRsp)
}


func main() {
	testFunc()
	//expression, err := govaluate.NewEvaluableExpression("foo > 0")
	//parameters := make(map[string]interface{}, 8)
	//parameters["foo"] = json.Number(1)
	//
	//result, err := expression.Evaluate(parameters);
	//fmt.Println(result, err)
}
