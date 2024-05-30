package utils

import (
	"golang.org/x/exp/constraints"
)

// Number is a type that can be either a float or an integer
type Number interface {
	constraints.Float | constraints.Integer
}

// Min returns the minimum of two numbers
func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}
