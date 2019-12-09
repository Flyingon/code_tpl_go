package main

import "fmt"

type Car struct {
	Weight int
	Name   string
}

type Bike struct {
	Car
	lunzi int
}

func main() {
	car1 := Car{
		Weight: 100,
		Name:   "car1",
	}
	bike := Bike{}
	bike.Car = car1
	fmt.Printf("bike: %+v\n", bike)
}