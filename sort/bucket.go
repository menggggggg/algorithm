package sort

// MaxNum ...
const MaxNum = 1000

// BucketSort ...
// 缺点空间复杂度比较大，比较的数值只能是数字
func BucketSort(array []int) []int {
	table := make([]int, MaxNum)
	for _, v := range array {
		table[v]++
	}

	dist := make([]int, 0, len(array))
	for k, v := range table {
		for j := 0; j < v; j++ {
			dist = append(dist, k)
		}
	}
	return dist
}
