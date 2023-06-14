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

// CheckSorted print if an array is sorted or not.
func CheckSorted(array []int) {
	if IsSorted(array) {
		fmt.Sprintln("The array is sorted")
	} else {
		fmt.Sprintln("The array is NOT sorted!")
	}
}

// IsSorted returns true if the array is sorted, false otherwise.
func IsSorted(array []int) bool {
	for i := 1; i < len(array); i++ {
		if array[i-1] > array[i] {
			return false
		}
	}
	return true
}
