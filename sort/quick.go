package sort

// QuickSort ...
func QuickSort(array []int) []int {
	if len(array) == 1 {
		return []int{}
	}

	dist := make([]int, len(array))
	copy(dist, array)
	quick(dist, 0, len(dist)-1)
	return dist
}

func quick(array []int, left, right int) {
	if left < right {
		index := adjustArray(array, left, right)
		quick(array, left, index-1)
		quick(array, index+1, right)
	}
}
func adjustArray(array []int, left, right int) int {
	i, j := left, right
	x := array[i]
	for i < j {
		for i < j && array[j] >= x {
			j--
		}
		if i < j {
			array[i] = array[j]
			i++
		}
		for i < j && array[i] < x {
			i++
		}
		if i < j {
			array[j] = array[i]
		}
	}
	array[i] = x
	return i
}
