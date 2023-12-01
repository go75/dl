package math

import (
	"dl/node"
	stmath "math"
)

func Pow[T node.Type](a, b T) T {
	return T(stmath.Pow(float64(a), float64(b)))
}

func Ln[T node.Type](a T) T {
	return T(stmath.Log(float64(a)))
}

func Tan[T node.Type](a T) T {
	return T(stmath.Tan(float64(a)))
}

func Sin[T node.Type](a T) T {
	return T(stmath.Sin(float64(a)))
}

func Cos[T node.Type](a T) T {
	return T(stmath.Cos(float64(a)))
}
