package main

func twoSum(nums []int, target int) []int {
	for i, n := range nums {
		diff := target - n
		foundIdx, ok := searcNum(nums[i+1:len(nums)], diff)
		if ok {
			return []int{i, i + foundIdx + 1}
		}
	}

	return []int{0, 0}
}

func twoSumHash(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, n := range nums {
		diff := target - n
		foundIdx, ok := hash[diff]
		if ok {
			return []int{foundIdx, i}
		}

		hash[n] = i
	}

	return []int{0, 0}
}

func searcNum(nums []int, search int) (int, bool) {
	for i, n := range nums {
		if n == search {
			return i, true
		}
	}
	return 0, false
}
