package quicksort

import (
	"sort"
	"testing"
)

func TestPartition(t *testing.T) {
	type testCase struct {
		array         []int
		expectedPivot int
	}

	var tests = map[string]testCase{
		"one element array": {
			array:         []int{1},
			expectedPivot: 0,
		},
		"two elements array, sorted": {
			array:         []int{1, 10},
			expectedPivot: 1,
		},
		"two elements array, not sorted": {
			array:         []int{10, 1},
			expectedPivot: 0,
		},
		"five elements array, not sorted": {
			array:         []int{3, 1, 4, 5, 10},
			expectedPivot: 4,
		},
		"five elements array, already sorted": {
			array:         []int{0, 0, 1, 2, 3},
			expectedPivot: 4,
		},
		"six elements array, not sorted": {
			array:         []int{0, 1, 2, 0, 3, 0},
			expectedPivot: 2,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			pivot := partition(tc.array)

			if pivot != tc.expectedPivot {
				t.Errorf("got: %v, expected: %v", pivot, tc.expectedPivot)
			}
		})
	}
}
func TestQuickSort(t *testing.T) {
	type testCase struct {
		array []int
	}

	var tests = map[string]testCase{
		"zero elements array": {
			array: []int{},
		},
		"one element array": {
			array: []int{1},
		},
		"five elements array, not sorted": {
			array: []int{3, 1, 4, 5, 10},
		},
		"five elements array, already sorted": {
			array: []int{0, 0, 1, 2, 3},
		},
		"six elements array, not sorted": {
			array: []int{0, 1, 2, 0, 3, 0},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			QuickSort(tc.array)

			if !sort.IntsAreSorted(tc.array) {
				t.Errorf("got: %v, not sorted", tc.array)
			}
		})
	}
}
