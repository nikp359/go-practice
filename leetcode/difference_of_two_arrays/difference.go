package differencearray

func findDifference(nums1 []int, nums2 []int) [][]int {
	m1 := unicMap(nums1)
	m2 := unicMap(nums2)

	diff1 := diffMap(m1, m2)
	diff2 := diffMap(m2, m1)
	return [][]int{diff1, diff2}
}

func unicMap(num []int) map[int]struct{} {
	um := make(map[int]struct{})

	for _, n := range num {
		if _, ok := um[n]; !ok {
			um[n] = struct{}{}
		}
	}

	return um
}

func diffMap(map1, map2 map[int]struct{}) []int {
	diff := []int{}

	for n := range map1 {
		if _, ok := map2[n]; !ok {
			diff = append(diff, n)
		}
	}

	return diff
}
