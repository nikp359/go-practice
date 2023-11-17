package productexceptself

func productExceptSelf(nums []int) []int {
	ln := len(nums)
	left := make([]int, ln)
	left[0], left[ln-1] = 1, 1

	for i := 1; i < ln; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right := 1
	for i := ln - 2; i >= 0; i-- {
		right *= nums[i+1]
		left[i] *= right
	}

	return left
}
