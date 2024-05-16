package array

import "math"

// 二分查找
// https://leetcode.cn/problems/binary-search/
func BinarySearch(nums []int, target int) bool {
	if len(nums) <= 0 {
		return false
	}
	mid := len(nums) / 2
	if target == nums[mid] {
		return true
	} else if target > nums[mid] {
		return BinarySearch(nums[mid+1:], target)
	} else {
		return BinarySearch(nums[:mid], target)
	}
}

// 长度最小的子数组
// https://leetcode.cn/problems/minimum-size-subarray-sum/description/
func MinSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 1
		}
		return 0
	}
	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	minLen := math.MaxInt32
	sum := 0
	i := 0
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		if sum >= target {
			minLen = min(minLen, j-i+1)
			sum -= nums[i]
			i++
		}
	}
	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}
