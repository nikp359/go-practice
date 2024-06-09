package array

func majorityElement(nums []int) int {
	hash := make(map[int]int, len(nums)/2)

	border := (len(nums)) / 2

	for i := 0; i < len(nums); i++ {
		hash[nums[i]]++

		if ch := hash[nums[i]]; ch > border {
			return nums[i]
		}
	}

	return 0
}
