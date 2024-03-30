package pivotindex

func pivotIndex(nums []int) int {
	var ls int
	rs := sum(nums)

	for i := range nums {
		if i > 0 {
			ls += nums[i-1]
		}

		rs -= nums[i]

		if ls == rs {
			return i
		}
	}

	return -1
}

func sum(nums []int) int {
	var r int
	for _, n := range nums {
		r += n
	}

	return r
}
