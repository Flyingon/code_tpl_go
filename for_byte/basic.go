package main

import (
	"encoding/json"
	"fmt"
)

type HotSpotStruct struct {
	HotStatus *int32          `json:"hot_status,omitempty"`
	HotSource []int           `json:"hot_source,omitempty"`
	IsHotCp   *int32          `json:"is_hot_cp,omitempty"`
	HotDetail json.RawMessage `json:"hot_detail,omitempty"`
	HotEvent  json.RawMessage `json:"hot_event,omitempty"`
}

func main() {
	data := []byte{123, 34, 104}
	fmt.Printf("%s\n", string(data))
	hotSpot := HotSpotStruct{}
	err := json.Unmarshal(data, &hotSpot)
	fmt.Printf("%+v, %v\n", hotSpot, err)
}
