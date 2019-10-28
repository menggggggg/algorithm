package sort

// BubbleSort ...
func BubbleSort(array []int) []int {
	if len(array) == 1 {
		return []int{}
	}

	dist := make([]int, len(array))
	copy(dist, array)

	for i := 0; i < len(dist)-1; i++ {
		for j := 0; j < len(dist)-i-1; j++ {
			if dist[j] > dist[j+1] {
				temp := dist[j]
				dist[j] = dist[j+1]
				dist[j+1] = temp
			}
		}
	}
	return dist
}
