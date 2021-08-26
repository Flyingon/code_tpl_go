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

// RoundDown 浮点类型保留小数点后n位精度，舍弃后面的
func RoundDown(v float64, n int) (r float64) {
	pow10N := math.Pow10(n)
	r = math.Floor(v*pow10N) / pow10N
	return r
}

// RoundUp 浮点类型保留小数点后n位精度，向上进位
func RoundUp(v float64, n int) (r float64) {
	pow10N := math.Pow10(n)
	r = math.Ceil(v*pow10N) / pow10N
	return r
}

// RoundNormal 浮点类型保留小数点后n位精度, 四舍五入
func RoundNormal(v float64, n int) (r float64) {
	pow10N := math.Pow10(n)
	r = math.Round(v*pow10N) / pow10N
	return r
}
