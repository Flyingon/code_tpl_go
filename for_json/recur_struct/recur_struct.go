package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type Module struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          int32     `protobuf:"varint,1,opt,name=type,proto3" json:"type"` // 0：案件; 1: 线索; 2:口供; 3: 推理
	Title         string    `protobuf:"bytes,2,opt,name=title,proto3" json:"title"`
	Icon          string    `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon"`
	SubModuleList []*Module `protobuf:"bytes,4,rep,name=sub_module_list,json=subModuleList,proto3" json:"sub_module_list,omitempty"` // 子模块
}

func main() {
	data := `[{"type": 0, "title": "案件", "icon": ""}, {"type": 1, "title": "线索", "icon": "", "sub_module_list": [{"type": 11, "title": "碎片", "icon": ""}, {"type": 12, "title": "时间线", "icon": ""}]}, {"type": 2, "title": "口供", "icon": ""}, {"type": 4, "title": "现场", "icon": ""}]`
	var modules []*Module
	err := jsoniter.UnmarshalFromString(data, &modules)
	fmt.Println(err)
	for _, module := range modules {
		fmt.Println(module)
		if len(module.SubModuleList) > 0 {
			for _, subModule := range module.SubModuleList {
				fmt.Println(subModule)
			}
		}
	}

}