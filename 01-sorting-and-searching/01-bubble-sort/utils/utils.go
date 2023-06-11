package utils

import "math/rand"

// MakeRandomArray creates a random int array with numItems elements
// in the interval of [0, max).
func MakeRandomArray(numItems, max int) []int {
	a := make([]int, numItems)

	for i := range a {
		a[i] = rand.Intn(max)
	}

	return a
}
