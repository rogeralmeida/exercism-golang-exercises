package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

type SillyNephewError struct {
	cowsCount int
}

func (sillyNephewError SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", sillyNephewError.cowsCount)
}

var DivisionByZeroError = errors.New("division by zero")
var NegativeFodderError = errors.New("negative fodder")

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
	switch {
	case fodder < 0 && err == ErrScaleMalfunction:
		return 0, NegativeFodderError
	case fodder > 0 && err == ErrScaleMalfunction:
		fodder *= 2
	case err != nil:
		return 0, err
	case fodder < 0:
		return 0, NegativeFodderError
	case cows == 0:
		return 0, DivisionByZeroError
	case cows < 0:
		return 0, &SillyNephewError{cowsCount: cows}
	}
	return fodder / float64(cows), nil
}
