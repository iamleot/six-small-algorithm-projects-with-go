package bubblesort

import (
	"bubblesort/utils"
	"testing"
)

func TestBubbleSort(t *testing.T) {
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
			BubbleSort(tc.array)

			if !utils.IsSorted(tc.array) {
				t.Errorf("got: %v, not sorted", tc.array)
			}
		})
	}
}
