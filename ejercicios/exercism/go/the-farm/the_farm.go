package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fodercalculator FodderCalculator, nCows int) (float64, error) {
	foodPerCow, err := fodercalculator.FodderAmount(nCows)
	if err != nil {
		return 0.0, err
	}
	fatteringFactor, err := fodercalculator.FatteningFactor()
	if err != nil {
		return 0.0, err
	}
	return (fatteringFactor * foodPerCow) / float64(nCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodercalculator FodderCalculator, nCows int) (float64, error) {
	if nCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fodercalculator, nCows)
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
	nCows       int
	customError string
}

func (i *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", i.nCows, i.customError)
}
func ValidateNumberOfCows(nCows int) error {
	if nCows < 0 {
		return &InvalidCowsError{nCows: nCows, customError: "there are no negative cows"}
	} else if nCows == 0 {
		return &InvalidCowsError{nCows: nCows, customError: "no cows don't need food"}
	} else {
		return nil
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
