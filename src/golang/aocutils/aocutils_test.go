package aocutils

import (
	"testing"
)

func Test_IntegerOrderOfmagnitude(t *testing.T) {
	for i, number := range []int{2, 54, 538, 9764, 97646} {
		expected := i
		actual := OrderOfMagnitude(number)

		if expected != actual {
			t.Errorf("Order of magnitute not calculated correctly for %d.\n\tExpected: %d\n\tGot: %d", number, expected, actual)
		}
	}
}
