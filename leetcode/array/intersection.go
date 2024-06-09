package array

func intersection(nums1 []int, nums2 []int) []int {
	smaller := nums1
	bigger := nums2
	if len(nums2) < len(nums2) {
		smaller = nums2
		bigger = nums1
	}

	hash := make(map[int]bool, len(smaller))
	resp := make([]int, 0)

	for i := 0; i < len(smaller); i++ {
		if _, ok := hash[smaller[i]]; !ok {
			hash[smaller[i]] = false
		}
	}

	for _, n := range bigger {
		if appended, ok := hash[n]; ok && !appended {
			resp = append(resp, n)
			hash[n] = true
		}
	}

	return resp
}
