package mymath

import (
	"math"
)

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func Ceil(x float64) int {
	intPart := int(x)
	if x-float64(intPart) > 0 {
		return intPart + 1
	}
	return intPart
}

func Floor(x float64) int {
	return int(x)
}

func Pow(x, y float64) float64 {
	var total float64 = 1
	for i := 0; i < int(y); i++ {
		total *= x
	}
	return total
}

func Max(x, y float64) float64 {
	if x > y {
		return x
	} else {
		return y
	}
}

func Min(x, y float64) float64 {
	if x < y {
		return x
	} else {
		return y
	}
}

func Abs(x float64) float64 {
	return math.Abs(x)
}

func Yn(x int, y float64) float64 {
	return math.Yn(x, y)
}
