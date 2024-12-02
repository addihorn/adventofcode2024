package aocutils

import "strconv"

func CBool2Int(value bool) int {
	if value {
		return 1
	}
	return 0
}

func CString2Int(value string) int {
	intValue, _ := strconv.Atoi(value)
	return intValue
}
