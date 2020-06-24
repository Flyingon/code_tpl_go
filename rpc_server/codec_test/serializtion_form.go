package main

import (
	"code_tpl_go/rpc_server/codec"
	"fmt"
)

var tag = "json"



func main() {
	formSerializer := codec.NewFormSerialization(tag)
	query := []byte("cGzip=1&cRand=1589459750079&cUserType=2&context=&rankId=1&subRankId=0&token=wq2fMa2f&userId=54523&cDeviceId=15F59DE1-4665-4BD4-A3ED-262665E349B8&cDeviceImei=&cDeviceMac=&cDevicePPI=&cDeviceScreenWidth=1125&cDeviceScreenHeight=2436&cDeviceModel=iPhone+XS+Max&cDeviceMem=3933552640&cDeviceCPU=ARM64&cDeviceNet=4G&cDeviceSP=%E4%B8%AD%E5%9B%BD%E8%81%94%E9%80%9A&cClientVersionCode=2102000000&cClientVersionName=2.0.0.0&cChannelId=0&cGameId=107&cSystem=ios&cSystemVersionCode=13.4.1&cSystemVersionName=iOS&cCurrentGameId=107")
	form := map[string]interface{}{}
	err := formSerializer.Unmarshal(query, &form)
	fmt.Printf("form: %+v, err: %v\n", form, err)
}
