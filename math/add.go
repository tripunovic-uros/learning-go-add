// Package math provides methods for math operations.
package math

import "golang.org/x/exp/constraints"

// A Number is any integer- or float-point type.
type Number interface {
  constraints.Integer | constraints.Float
}

// Add takes in two int paramters and returns the sum
//
// More information on addition can be found at [mathisfun]
// 
// [mathisfun]: https://mathisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
  return a + b
}

