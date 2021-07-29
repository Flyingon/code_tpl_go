package util

import (
	"github.com/shopspring/decimal"
	"math"
)

func DivideFloat(d1, d2 float64) float64 {
	decimalD1 := decimal.NewFromFloat(d1)
	decimalD2 := decimal.NewFromFloat(d2)
	decimalResult := decimalD1.Div(decimalD2)
	float64Result, _ := decimalResult.Float64()
	return float64Result
}

// RoundFloat64 浮点类型保留小数点后n位精度
func RoundFloat64(v float64, n int) (r float64) {
	pow10N := math.Pow10(n)
	r = math.Trunc((v+0.5/pow10N)*pow10N) / pow10N
	return r
}
