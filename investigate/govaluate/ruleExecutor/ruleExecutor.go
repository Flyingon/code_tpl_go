package ruleExecutor

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"strings"
)

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
	FuncMap map[string]govaluate.ExpressionFunction
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
	expression, err := govaluate.NewEvaluableExpressionWithFunctions(re.Rule, re.FuncMap)
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

