package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func version4() {
	uV4 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", uV4)
}

func main() {
	version4()
}