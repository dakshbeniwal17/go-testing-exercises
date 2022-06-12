package exercise1

import "errors"

// without using any external package create a test case for each end case of this method
// only testing package should be used for this exercise
func Factorial(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("number should be greater than zero")
	} else if n > 10 {
		return 0, errors.New("number should be less than ten")
	}

	result := 1
	for i := n; i > 0; i-- {
		result *= i
	}
	return result, nil
}
