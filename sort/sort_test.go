package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BucketSort(t *testing.T) {
	src := []int{3, 4, 9, 7, 23, 14, 1, 33}
	dist := []int{1, 3, 4, 7, 9, 14, 23, 33}
	ret := BucketSort(src)
	assert.Equal(t, dist, ret)
}

func Test_BubbleSort(t *testing.T) {
	src := []int{3, 4, 9, 7, 23, 14, 1, 33}
	dist := []int{1, 3, 4, 7, 9, 14, 23, 33}
	ret := BubbleSort(src)
	assert.Equal(t, dist, ret)
}

func Test_QuickSort(t *testing.T) {
	src := []int{3, 4, 9, 7, 23, 14, 1, 33}
	dist := []int{1, 3, 4, 7, 9, 14, 23, 33}
	ret := QuickSort(src)
	assert.Equal(t, dist, ret)
}
