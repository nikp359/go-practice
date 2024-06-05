package array

func searchInsert(nums []int, target int) int {
	var low, mid, index int

	hight := len(nums) - 1

	if nums[hight] < target {
		return len(nums)
	}

	for low <= hight {
		mid = (low + hight) / 2
		if nums[mid] < target {
			low = mid + 1
			index = mid + 1
		} else if nums[mid] == target {
			return index
		} else {
			hight = mid - 1
		}
	}

	return index
}
