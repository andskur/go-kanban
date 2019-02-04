package utils

import "strings"

// ContainSlice check if providing slice have given value
func ContainSlice(slice []string, value string) bool {
	for _, a := range slice {
		if a == value {
			return true
		}
	}
	return false
}

// divide divides given string
func Divide(str string) []string {
	return strings.Split(str, "|")
}
