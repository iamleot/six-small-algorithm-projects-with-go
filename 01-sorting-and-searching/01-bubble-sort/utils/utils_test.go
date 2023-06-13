package utils

import "testing"

func TestMakeRandomArray(t *testing.T) {
	type testCase struct {
		numItems, max int
	}

	var tests = map[string]testCase{
		"zero elements array, 0 max": {
			numItems: 0,
			max:      0,
		},
		"five elements array, 1 max": {
			numItems: 5,
			max:      1,
		},
		"five elements array, 10 max": {
			numItems: 5,
			max:      10,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			a := MakeRandomArray(tc.numItems, tc.max)

			if len(a) != tc.numItems {
				t.Errorf("length expected: %q, got: %q",
					tc.numItems, len(a))
			}

			for _, v := range a {
				if v < 0 {
					t.Errorf("%q is negative, expected values only >= 0",
						v)
				}

				if v >= tc.max {
					t.Errorf("%q is bigger than %q, expected values only in [0, %q) interval",
						v, tc.max, tc.max)

				}
			}
		})
	}
}

func TestPrintArray(t *testing.T) {
	PrintArray([]int{0, 1, 2}, 0)
	// Output:
	// []

	PrintArray([]int{0, 1, 2}, 2)
	// Output:
	// [0 1]

	PrintArray([]int{0, 1, 2}, 3)
	// Output:
	// [0 1 2]

	PrintArray([]int{0, 1, 2}, 5)
	// Output:
	// [0 1 2]
}

func TestSprintArray(t *testing.T) {
	type testCase struct {
		array    []int
		numItems int
		expected string
	}

	var tests = map[string]testCase{
		"zero elements array, 0 items to print": {
			array:    []int{},
			numItems: 0,
			expected: "[]",
		},
		"zero elements array, 5 items to print": {
			array:    []int{},
			numItems: 5,
			expected: "[]",
		},
		"five elements array, 0 items to print": {
			array:    []int{0, 1, 2, 3, 4},
			numItems: 0,
			expected: "[]",
		},
		"five elements array, 3 items to print": {
			array:    []int{0, 1, 2, 3, 4},
			numItems: 3,
			expected: "[0 1 2]",
		},
		"five elements array, 5 items to print": {
			array:    []int{0, 1, 2, 3, 4},
			numItems: 5,
			expected: "[0 1 2 3 4]",
		},
		"five elements array, 10 items to print": {
			array:    []int{0, 1, 2, 3, 4},
			numItems: 10,
			expected: "[0 1 2 3 4]",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			printedArray := sprintArray(tc.array, tc.numItems)

			if printedArray != tc.expected {
				t.Errorf("expected: %q, got: %q",
					tc.expected, printedArray)
			}
		})
	}
}

func TestCheckSorted(t *testing.T) {
	CheckSorted([]int{})
	// Output:
	// The array is sorted

	CheckSorted([]int{0})
	// Output:
	// The array is sorted

	CheckSorted([]int{0, 1, 2})
	// Output:
	// The array is sorted

	CheckSorted([]int{0, 0, 0})
	// Output:
	// The array is sorted

	CheckSorted([]int{1, 2, 1})
	// Output:
	// The array is NOT sorted!
}

func TestSorted(t *testing.T) {
	type testCase struct {
		array  []int
		sorted bool
	}

	var tests = map[string]testCase{
		"zero elements array, sorted": {
			array:  []int{},
			sorted: true,
		},
		"one element array, sorted": {
			array:  []int{1},
			sorted: true,
		},
		"five elements array (all increasing), sorted": {
			array:  []int{1, 3, 5, 8, 10},
			sorted: true,
		},
		"five elements array (some increasing, some the same), sorted": {
			array:  []int{0, 0, 1, 2, 3},
			sorted: true,
		},
		"five elements array, not sorted": {
			array:  []int{0, 1, 2, 3, 0},
			sorted: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			sorted := isSorted(tc.array)

			if sorted != tc.sorted {
				t.Errorf("expected: %t, got: %t for array %q",
					tc.sorted, sorted, tc.array)
			}
		})
	}
}
