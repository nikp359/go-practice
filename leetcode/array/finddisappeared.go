package array

// Find All Numbers Disappeared in an Array
// Input: nums = [4,3,2,7,8,2,3,1]
// Output: [5,6]
//
// Input: nums = [1,1]
// Output: [2]
func findDisappearedNumbers(nums []int) []int {
	result := make([]int, len(nums))

	for _, n := range nums {
		result[n-1] = n
	}

	resp := make([]int, 0, len(result))

	for i := 0; i < len(result); i++ {
		if result[i] == 0 {
			resp = append(resp, i+1)
		}
	}

	return resp
}
