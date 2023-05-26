package util

import "strconv"

// ConvertToInt converts given string to int.
func ConvertToInt(number string) int {
	value, err := strconv.Atoi(number)
	if err != nil {
		return 0
	}
	return value
}

// ConvertToUint converts given string to uint.
func ConvertToUint(number string) uint {
	return uint(ConvertToInt(number))
}
