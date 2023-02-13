package calculator

import (
	"errors"
	"math"
)

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 || b == math.NaN() || a == math.NaN() {
		return 0, errors.New("division by 0 is not allowed")
	}
	return a / b, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 || a == math.NaN() {
		return 0, errors.New("square root of negative numbers is not allowed")
	}
	return math.Sqrt(a), nil
}
