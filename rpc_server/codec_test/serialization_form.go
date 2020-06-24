package main

import (
	"fmt"
	"github.com/go-playground/form"
	"github.com/mitchellh/mapstructure"
	"net/url"
	"reflect"
)

func UnmarshalOld(in []byte, body interface{}) error {
	values, _ := url.ParseQuery(string(in))
	params := map[string]interface{}{}
	for k, v := range values {
		if len(v) == 1 {
			params[k] = v[0]
		} else {
			params[k] = v
		}
	}
	//fmt.Printf("in: %s\n params: %+v\n", in, params)
	config := &mapstructure.DecoderConfig{TagName: "json", Result: body, WeaklyTypedInput: true, Metadata: nil}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(params)
}

// Unmarshal 解包kv结构
func Unmarshal(in []byte, body interface{}) error {

	values, err := url.ParseQuery(string(in))
	if err != nil {
		return err
	}
	fmt.Println("values:", values)
	fmt.Println("body type: ", reflect.TypeOf(body))
	decoder := form.NewDecoder()
	decoder.SetTagName("json")
	err = decoder.Decode(&body, values)
	return err
}

type SizeInfo struct {
	Size int
	From int
}

func main() {
	//query := []byte("cGzip=1&cRand=1589459750079&cUserType=2&context=&rankId=1&subRankId=0&token=wq2fMa2f&userId=54523&cDeviceId=15F59DE1-4665-4BD4-A3ED-262665E349B8&cDeviceImei=&cDeviceMac=&cDevicePPI=&cDeviceScreenWidth=1125&cDeviceScreenHeight=2436&cDeviceModel=iPhone+XS+Max&cDeviceMem=3933552640&cDeviceCPU=ARM64&cDeviceNet=4G&cDeviceSP=%E4%B8%AD%E5%9B%BD%E8%81%94%E9%80%9A&cClientVersionCode=2102000000&cClientVersionName=2.0.0.0&cChannelId=0&cGameId=107&cSystem=ios&cSystemVersionCode=13.4.1&cSystemVersionName=iOS&cCurrentGameId=107")
	//form := map[string]interface{}{}
	//err := UnmarshalOld(query, &form)
	//fmt.Printf("1 form: %+v, err: %v\n", form, err)
	//
	//form2 := map[string]interface{}{}
	//err = Unmarshal(query, &form2)
	//fmt.Printf("2 form: %+v, err: %v\n", form2, err)
	sz := &SizeInfo{
		Size: 10,
		From: 20,
	}
	encoder := form.Encoder{}
	value, err := encoder.Encode(sz)
	fmt.Println(value.Encode(), err)
}
