package array

import (
	"cmp"
	"slices"
)

// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Input: intervals = [[1,4],[4,5]]
// Output: [[1,5]]
func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	slices.SortFunc(intervals, func(a []int, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	resp := make([][]int, 0, len(intervals)/2)
	head := intervals[0]

	for _, interval := range intervals[1:] {
		if head[1] < interval[0] {
			resp = append(resp, head)
			head = interval
		} else if interval[1] > head[1] {
			head[1] = interval[1]
		}
	}

	resp = append(resp, head)

	return resp
}
