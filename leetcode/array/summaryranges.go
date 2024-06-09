package array

import (
	"fmt"
	"strconv"
)

type Range struct {
	start  int
	finish int
}

func (r *Range) String() string {
	if r.start == r.finish {
		return strconv.Itoa(r.start)
	}

	return fmt.Sprintf("%d->%d", r.start, r.finish)
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	rangeList := make([]*Range, 0)

	lr := &Range{start: nums[0]}

	if len(nums) == 1 {
		lr.finish = lr.start
		return []string{lr.String()}
	}

	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i]-1 {
			lr.finish = nums[i-1]
			rangeList = append(rangeList, lr)
			lr = &Range{start: nums[i]}
		}
	}

	if lr.finish != nums[len(nums)-1] {
		lr.finish = nums[len(nums)-1]
		rangeList = append(rangeList, lr)
	}

	if lr.start > lr.finish {
		lr.finish = lr.start
		rangeList = append(rangeList, lr)
	}

	resp := make([]string, 0, len(rangeList))

	for _, r := range rangeList {
		resp = append(resp, r.String())
	}

	return resp
}
