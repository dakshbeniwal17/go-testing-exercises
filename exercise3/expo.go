package exercise3

func Expo(number int, power int) int {
	if number < 1 || number > 10 {
		return 0
	}
	if power < 1 || power > 5 {
		return -1
	}

	result := 1
	for i := power; i > 0; i-- {
		result *= number
	}
	return result
}
