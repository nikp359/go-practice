package differencearray

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPivotIndex(t *testing.T) {
	testCases := []struct {
		nums1     []int
		nums2     []int
		wantDiff1 []int
		wantDiff2 []int
	}{
		{
			nums1:     []int{1, 2, 3, 3, 4},
			nums2:     []int{1, 1, 2, 2, 6},
			wantDiff1: []int{3, 4},
			wantDiff2: []int{6},
		},
		{
			nums1:     []int{1, 2},
			nums2:     []int{1},
			wantDiff1: []int{2},
			wantDiff2: []int{},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case n: %d", i+1), func(t *testing.T) {
			got := findDifference(tc.nums1, tc.nums2)
			slices.Sort(got[0])
			slices.Sort(got[1])

			assert.Equal(t, tc.wantDiff1, got[0])
			assert.Equal(t, tc.wantDiff2, got[1])
		})
	}
}
