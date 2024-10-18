package main

import (
	"errors"
	"fmt"
)

// function that always returns an error
func alwaysErrors() error {
	return errors.New("this is a basic error")
}

// function that can return an error
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return a / b, nil
}

func main() {

	fmt.Println("basic error handling")
	err := alwaysErrors()
	fmt.Println("The error: ", err)
	fmt.Printf("The error is of the type: %T\n", err)

	// standard pattern:
	result, err := divide(2, 2)
	if err != nil {
		// handle the error
		fmt.Println("divide failed:", err)
	} else {
		fmt.Println("divide result: ", result)
	}
	result, err = divide(2, 0)
	if err != nil {
		// handle the error
		fmt.Println("divide failed:", err)
	} else {
		fmt.Println("divide result: ", result)
	}

}
