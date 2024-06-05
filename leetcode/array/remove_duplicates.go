package array

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	var j int

	for i := 1; i < len(nums); i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}
