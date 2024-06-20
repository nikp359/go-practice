package array

import "slices"

func arrayPairSum(nums []int) int {
	slices.Sort(nums)
	var sum int

	for i := 0; i < len(nums); i = i + 2 {
		sum += nums[i]
	}

	return sum
}
