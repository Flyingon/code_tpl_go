package main

import (
	"errors"
	"fmt"
)

func deferSetErrV1() error {
	var err error
	defer func() {
		err = errors.New("test")
	}()
	return nil
}

func deferSetErrV2() (err error) {
	defer func() {
		err = errors.New("test")
	}()
	return nil
}

func main() {
	fmt.Printf("v1 retrun err: %v\n", deferSetErrV1())
	fmt.Printf("v2 retrun err: %v\n", deferSetErrV2())
}
