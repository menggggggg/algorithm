package hash

import (
	"sort"
)

// 两数之和
// https://leetcode.cn/problems/two-sum/description/
func TwoSum(nums []int, target int) []int {
	// sumMap := make(map[int]int, 0)
	// for idx, value := range nums {
	// 	v, ok := sumMap[target-value]
	// 	if ok {
	// 		data := []int{idx, v}
	// 		return data
	// 	}
	// 	sumMap[value] = idx
	// }
	// return []int{}
	numsIdxMap := make(map[int]int, 0)
	for k, v := range nums {
		numsIdxMap[v] = k
	}
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{numsIdxMap[nums[left]], numsIdxMap[nums[right]]}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}

// 三数之和
// https://leetcode.cn/problems/3sum/
func ThreeSum(nums []int) [][]int {
	datas := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > 0 {
			break
		}
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			// 去重
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				datas = append(datas, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				continue
			} else if sum < 0 {
				left++
				continue
			}
			right--
		}
	}
	return datas
}
