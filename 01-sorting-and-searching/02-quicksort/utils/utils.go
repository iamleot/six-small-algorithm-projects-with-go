package utils

import (
	"fmt"
	"math/rand"
)

// MakeRandomArray creates a random int array with numItems elements
// in the interval of [0, max).
func MakeRandomArray(numItems, max int) []int {
	a := make([]int, numItems)

	for i := range a {
		a[i] = rand.Intn(max)
	}

	return a
}

// PrintArray print at maximum numItems of the int array.
func PrintArray(array []int, numItems int) {
	fmt.Println(sprintArray(array, numItems))
}

// sprintArray returns a string containing at maximum numItems elements of the
// array.
//
// Only intended to make actual functionality PrintArray() testable.
func sprintArray(array []int, numItems int) string {
	if numItems >= len(array) {
		return fmt.Sprint(array)
	}

	return fmt.Sprint(array[:numItems])
}
