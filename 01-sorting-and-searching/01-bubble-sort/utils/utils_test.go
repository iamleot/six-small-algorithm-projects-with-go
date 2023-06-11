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
