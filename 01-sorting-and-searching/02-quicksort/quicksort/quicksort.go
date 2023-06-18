package quicksort

// QuickSort sort the array in place using the Quicksort algorithm using the
// Lomuto partition scheme.
func QuickSort(array []int) {
	if len(array) <= 1 {
		return
	}

	p := partition(array)

	QuickSort(array[:p])
	QuickSort(array[p:])
}

// partition partitions the array using the Lomuto partition scheme and returns
// the index of the pivot.
//
// The array should have at least one element.
func partition(array []int) int {
	lo, hi := 0, len(array)-1
	pivot := array[hi]

	i := lo - 1
	for j := 0; j < hi; j++ {
		if array[j] <= pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}

	i++
	array[i], array[hi] = array[hi], array[i]
	return i
}
