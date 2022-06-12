package exercise3

import "errors"

// Using mockgen, create a test case for each end case of following function
// Hint: you need to mock Expo function
// You may also edit CalculatePower and Expo functions but no existing should be removed
func CalculatePower(num int, pow int) (int, error) {
	result := Expo(num+1, pow)
	if result == 0 {
		return 0, errors.New("invalid number")
	} else if result == -1 {
		return 0, errors.New("invalid power")
	}

	return result, nil
}
