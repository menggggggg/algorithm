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

func QuickSortV2(nums []int) {
	if len(nums) == 0 {
		return
	}
	var quickSort func(num []int, left, right int)
	quickSort = func(num []int, left, right int) {
		if left >= right {
			return
		}
		povit := num[(left+right)/2]
		i, j := left, right
		for {
			for num[i] < povit {
				i++
			}
			for num[j] > povit {
				j--
			}
			if i >= j {
				break
			}
			num[i], num[j] = num[j], num[i]
		}
		quickSort(num, left, i-1)
		quickSort(num, j+1, right)
	}
	quickSort(nums, 0, len(nums)-1)
}
