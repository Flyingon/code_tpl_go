package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
)

func main() {
	want := map[string]int{"a": 1, "b": 2}
	got := map[string]int{"a": 1, "b": 3}
	if diff := cmp.Diff(want, got); diff != "" {
		fmt.Printf("MakeGatewayInfo() mismatch (-want +got):\n%s", diff)
	}
}
