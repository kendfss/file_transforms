package main

import (
	"fmt"
	"reflect"

	"github.com/kendfss/iters/slices"
	"github.com/kendfss/oprs"
	"github.com/kendfss/rules"
)

// Polynomial represents a polynomial with coefficients of type T.
type Polynomial[T rules.Number] []T

// NewPolynomial creates a new Polynomial from a variadic sequence of bytes.
func NewPolynomial[T rules.Number](bytes ...byte) Polynomial[T] {
	coefficients := make([]T, 256) // 256 possible byte values
	for _, b := range bytes {
		coefficients[b]++
	}

	return Polynomial[T](coefficients)
}

func rto[T any](arg T) reflect.Type {
	return reflect.TypeOf(arg)
}

func convert[O, I any](in I) (out *O) {
	ti := rto(in)

	numtypes := []reflect.Type{
		rto(*new(complex128)),
		rto(*new(complex64)),
		rto(*new(float64)),
		rto(*new(float32)),
		rto(*new(int)),
		rto(*new(int64)),
		rto(*new(int32)),
		rto(*new(int16)),
		rto(*new(int8)),
		rto(*new(uint)),
		rto(*new(uint64)),
		rto(*new(uint32)),
		rto(*new(uint16)),
		rto(*new(uint8)),
	}

	if slices.Contains(numtypes[:2], ti) {
		if slices.Contains(numtypes[:2], to) {
			return O(in)
		}
		return O(complex(in, 0))
	}

	switch ty {
	case rto(oprs.New[complex128]), rto(oprs.New[complex64]):
		return T(ar)

	case rto(oprs.New[float32]):
	case rto(oprs.New[float64]):
	case rto(oprs.New[int]):
	case rto(oprs.New[int8]):
	case rto(oprs.New[int16]):
	case rto(oprs.New[int32]):
	case rto(oprs.New[int64]):
	case rto(oprs.New[uint]):
	case rto(oprs.New[uint8]):
	case rto(oprs.New[uint16]):
	case rto(oprs.New[uint32]):
	case rto(oprs.New[uint64]):
	}

	return out
}

// Call evaluates the Polynomial at the given value x.
func (p Polynomial[T]) Call(x T) T {
	result := T(0)
	power := T(1)
	for _, coefficient := range p {
		result += coefficient * power
		power *= x
	}
	return result
}

func main() {
	// Example usage
	bytes := []byte{1, 2, 3, 2, 3, 3} // Example byte sequence
	polynomial := NewPolynomial[int](bytes...)
	x := 2 // Example value to evaluate the polynomial at
	result := polynomial.Call(x)
	fmt.Println(result)
}
