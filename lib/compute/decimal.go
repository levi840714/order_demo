package compute

import "github.com/shopspring/decimal"

func Add(left float64, right float64) float64 {
	l := decimal.NewFromFloat(left)
	r := decimal.NewFromFloat(right)
	result, _ := l.Add(r).Float64()
	return result
}

func Mult(left float64, right float64) float64 {
	l := decimal.NewFromFloat(left)
	r := decimal.NewFromFloat(right)
	result, _ := l.Mul(r).Float64()
	return result
}

func Sub(left float64, right float64) float64 {
	l := decimal.NewFromFloat(left)
	r := decimal.NewFromFloat(right)
	result, _ := l.Sub(r).Float64()
	return result
}

func Div(left float64, right float64) float64 {
	l := decimal.NewFromFloat(left)
	r := decimal.NewFromFloat(right)
	result, _ := l.Div(r).Float64()
	return result
}

func Mod(left float64, right float64) float64 {
	l := decimal.NewFromFloat(left)
	r := decimal.NewFromFloat(right)
	result, _ := l.Mod(r).Float64()
	return result
}
