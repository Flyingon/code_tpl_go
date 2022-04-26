package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func MulFloat(x, y float32) float32 {
	score2 := decimal.NewFromFloat32(x).Mul(decimal.NewFromFloat32(y))
	ret, _ := score2.Float64()
	return float32(ret)
}

func main() {
	srcScore := 700
	buf := float32(1.3)
	score := int64(float32(srcScore) * buf)
	fmt.Println(score)

	score2 := MulFloat(float32(srcScore), buf)
	fmt.Println(score2)
}
