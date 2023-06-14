package bubblesort

// BubbleSort sort the array in place using the Bubble sort algorithm.
func BubbleSort(array []int) {
	var swapped bool

	for {
		swapped = false
		for i := 1; i < len(array); i++ {
			if array[i-1] > array[i] {
				array[i-1], array[i] = array[i], array[i-1]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}
