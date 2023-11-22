package maxaveragesubarray

func findMaxAverage(nums []int, k int) float64 {
	maxSum := sum(nums[0:k])
	s := maxSum

	for left, right := 0, k; right < len(nums); left, right = left+1, right+1 {
		s = s - nums[left] + nums[right]
		if s > maxSum {
			maxSum = s
		}
	}

	return float64(maxSum) / float64(k)
}

func sum(nums []int) (s int) {
	for i := range nums {
		s += nums[i]
	}

	return s
}
