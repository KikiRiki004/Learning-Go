package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) { // By convention, errors are the last return value and have type error
	if arg == 42 {

		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil // Nil value in the error position indicates that there was no error
}

var ErrOutOfTea = fmt.Errorf("no more tea available") // A sentinel error is a predeclared variable
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {

		return fmt.Errorf("making tea: %w", ErrPower) // Error wrapping with %w
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {

		if r, e := f(i); e != nil { // Commonly checked in the if statement
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {

			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}

// Error wrapping is the practice of adding context or additional information to an error while preserving the original error,
// making it easier to trace and debug the root cause of a problem.
