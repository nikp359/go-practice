package move_zeroes

func moveZeroes(nums []int) {
	var j int

	for i := range nums {
		if nums[i] == 0 {
			continue
		}

		nums[j] = nums[i]
		j++
	}

	for ; j < len(nums); j++ {
		nums[j] = 0
	}
}
