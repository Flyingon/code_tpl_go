package main

import (
	"encoding/json"
	"fmt"
)

type LowClassifierDetail struct {
	LowClassifierList []LowClassifierOne `json:"low"`
}

type LowClassifierOne struct {
	Confidence *float64 `json:"confidence"`
	Level      *int     `json:"level"`
	Score      *int     `json:"score"`
	Type       *int     `json:"type"`
}

type ABC struct {
	A string `json:"a,omitempty"`
	B string `json:"b"`
}

func main() {
	//contentStr := `{"a": 1, "b": "2", "c": {"d": 3, "e": "4"}}`
	//contentMap := make(map[string]interface{})
	//err := util.JSONUnMarshal(util.StringToBytesFast(contentStr), &contentMap)
	//fmt.Println(err, contentMap).Marshal(&abc)
	//	//fmt.Printf("abc new json: %s\n", abc
	//for k, v := range contentMap {
	//	fmt.Println(k, reflect.TypeOf(v))
	//}
	data := `{"low":[{"confidence":0.37699482,"level":0,"score":0,"type":12},{"confidence":0,"level":0,"score":0,"type":8},{"confidence":0.01519581,"level":0,"score":0,"type":13},{"confidence":0.01,"level":0,"score":0,"type":17},{"confidence":0.25547466,"level":0,"score":0,"type":0},{"level":0,"score":0,"type":5},{"confidence":0.26630497,"level":0,"score":0,"type":10},{"confidence":0,"level":0,"score":0,"type":21},{"level":0,"score":0,"type":22},{"level":0,"score":0,"sub_lowinfo":[{"sub_info":"","sub_level":0,"sub_type":0},{"sub_info":"","sub_level":0,"sub_type":1},{"sub_info":"","sub_level":0,"sub_type":2},{"sub_info":"","sub_level":0,"sub_type":3},{"sub_info":"","sub_level":0,"sub_type":4},{"sub_info":"","sub_level":0,"sub_type":5}],"type":23},{"confidence":1,"level":2,"score":0,"sub_lowinfo":[{"sub_info":"","sub_level":0,"sub_type":0},{"sub_info":"","sub_level":2,"sub_type":1}],"type":20},{"level":0,"score":0,"type":26},{"level":0,"score":0,"type":27},{"level":0,"score":0,"sub_lowinfo":[{"sub_info":"","sub_level":0,"sub_type":0},{"sub_info":"","sub_level":0,"sub_type":1},{"sub_info":"","sub_level":0,"sub_type":2},{"sub_info":"","sub_level":0,"sub_type":3},{"sub_info":"","sub_level":0,"sub_type":4},{"sub_info":"","sub_level":0,"sub_type":5},{"sub_info":"","sub_level":0,"sub_type":6},{"sub_info":"","sub_level":0,"sub_type":7},{"sub_info":"","sub_level":0,"sub_type":8},{"sub_info":"","sub_level":0,"sub_type":9}],"type":24},{"confidence":0,"level":0,"score":0,"type":14},{"level":0,"score":0,"type":15},{"level":0,"score":0,"type":16},{"level":0,"score":0,"type":19},{"level":0,"score":0,"type":18},{"confidence":0.25547466,"level":0,"score":0,"sub_lowinfo":[{"sub_confidence":0.25547466,"sub_level":0,"sub_type":1},{"sub_confidence":0,"sub_level":0,"sub_type":2},{"sub_confidence":0,"sub_level":0,"sub_type":3}],"type":31},{"confidence":0.0019000173,"level":0,"score":0,"type":32},{"confidence":0.0038,"level":0,"score":0,"type":33}]}`

	lowClassifierDetail := LowClassifierDetail{}
	err := json.Unmarshal([]byte(data), &lowClassifierDetail)
	fmt.Println(err)

	abcJson := `{"a":"1", "b":"2"}`
	abc := ABC{}
	errABC := json.Unmarshal([]byte(abcJson), &abc)
	fmt.Printf("err: %v, abc: %+v\n", errABC, abc)
	abcdJson := `{"a":"a", "b":"b", "d": "d"}`
	errABC = json.Unmarshal([]byte(abcdJson), &abc)
	fmt.Printf("err: %v, abc: %+v\n", errABC, abc)
	//abcJsonNew, _ := jsonJsonNew)
}
