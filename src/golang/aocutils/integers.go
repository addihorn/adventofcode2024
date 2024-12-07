package aocutils

import (
	"math"
	"sort"
)

func Min(intValues []int) int {
	sort.Ints(intValues)
	return intValues[0]
}

func Max(intValues []int) int {
	sort.Ints(intValues)
	return intValues[len(intValues)-1]
}

func Abs(value int) int {
	if value < 0 {
		return value * -1
	}
	return value
}

func OrderOfMagnitude(value int) int {
	return int(math.Floor(math.Log10(float64(value))))
}
