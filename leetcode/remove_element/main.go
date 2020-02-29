package main

import "fmt"

func main() {
	//input := []int{2, 2, 3, 1}
	//input := []int{1}
	input := []int{1, 2, 3, 1, 97, 75, 1, 1, 1, 1}
	//input := []int{4, 1, 2, 4, 4, 0, 0}
	fmt.Println(input)
	l := removeElement(input, 1)
	fmt.Println(input)
	fmt.Println(l)
}

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		nums = removeLastValues(nums, val)
		if 1 >= len(nums) || len(nums) <= i {
			return len(nums)
		}
		if nums[i] == val {
			nums[i] = nums[len(nums)-1]
			nums[len(nums)-1] = 0
			nums = nums[:len(nums)-1]
		}
	}
	return len(nums)
}

func removeLastValues(nums []int, val int) []int {
	if len(nums) == 0 {
		return nums
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if val != nums[i] {
			return nums
		}
		nums = nums[:i]
	}
	return nums
}
