package twosum

// TwoSum array 必须为有序集合
func TwoSum(array []int, targe int) bool {
	if len(array) < 2 {
		return false
	}
	i, j := 0, len(array)-1
	for i < j {
		sum := array[i] + array[j]
		if sum == targe {
			return true
		}
		if sum < targe {
			i++
		} else {
			j--
		}
	}
	return false
}
