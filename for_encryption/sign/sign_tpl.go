package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func RawUrlEncode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}

func ToString(key interface{}) string {
	if key == nil {
		return ""
	}

	var ret string
	switch key.(type) {
	case string:
		ret = key.(string)
	case int:
		ret = strconv.FormatInt(int64(key.(int)), 10)
	case int8:
		ret = strconv.FormatInt(int64(key.(int8)), 10)
	case int16:
		ret = strconv.FormatInt(int64(key.(int16)), 10)
	case int32:
		ret = strconv.FormatInt(int64(key.(int32)), 10)
	case int64:
		ret = strconv.FormatInt(key.(int64), 10)
	case uint:
		ret = strconv.FormatUint(uint64(key.(uint)), 10)
	case uint8:
		ret = strconv.FormatUint(uint64(key.(uint8)), 10)
	case uint16:
		ret = strconv.FormatUint(uint64(key.(uint16)), 10)
	case uint32:
		ret = strconv.FormatUint(uint64(key.(uint32)), 10)
	case uint64:
		ret = strconv.FormatUint(key.(uint64), 10)
	case float32:
		ret = strconv.FormatFloat(float64(key.(float32)), 'f', -1, 64)
	case float64:
		ret = strconv.FormatFloat(key.(float64), 'f', -1, 64)
	case bool:
		ret = strconv.FormatBool(key.(bool))
	case []byte:
		ret = string(key.([]byte))
	case json.Number:
		ret = key.(json.Number).String()
	default:
		retBytes, _ := json.Marshal(key)
		ret = string(retBytes)
	}
	return ret
}

func ToInt2(inputArg interface{}) int {
	var outputArg int = 0
	switch inputArg.(type) {
	case string:
		//outputArg, _ = strconv.Atoi(inputArg.(string))
		outputArg2, _ := strconv.ParseFloat(inputArg.(string), 64)
		outputArg = int(outputArg2)
		break
	case int32:
		outputArg = int(inputArg.(int32))
		break
	case int64:
		outputArg = int(inputArg.(int64))
		break
	case int: //todo 是否需要统一所有int int64
		outputArg = inputArg.(int)
		break
	case float32:
		outputArg = int(inputArg.(float32))
		break
	case float64:
		outputArg = int(inputArg.(float64))
		break
	default:
	}
	return outputArg
}

//生成签名
func CreateSig(method string, urlPath string, params map[string]interface{}, secret string) string {
	delete(params, "sig")

	strs := strings.ToUpper(method) + "&" + RawUrlEncode(urlPath) + "&"
	keys := make([]string, len(params))
	i := 0
	for k := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	queryString := ""
	for _, k := range keys {
		if queryString != "" {
			queryString += "&"
		}
		paramKtype := ""
		if params[k] != nil {
			paramKtype = reflect.TypeOf(params[k]).String()
		}
		if strings.HasPrefix(paramKtype, "int") || strings.HasPrefix(paramKtype, "float") {
			//go json decode 数字默认会处理成科学计数法，这里先转整形再转string
			queryString += (k + "=" + ToString(ToInt2(params[k])))
		} else {
			queryString += (k + "=" + ToString(params[k]))
		}
	}
	queryString = RawUrlEncode(queryString)
	queryString = strings.Replace(queryString, "~", "%7E", -1)
	mk := strs + queryString
	secret = strings.Replace(secret, "-", "+", -1)
	secret = strings.Replace(secret, "_", "/", -1)

	h := hmac.New(sha1.New, []byte(secret))
	h.Write([]byte(mk))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func main() {
	fmt.Println(CreateSig("POST", "/book/api/getAssets", map[string]interface{}{
		"userId":"54523", "osType": 2, "appid": "1000001", "token": "bXDaBN4C",
	}, "kRpKaAyrgDXBm7gOP3p4z6uE4eTFPZiz&"))
}
