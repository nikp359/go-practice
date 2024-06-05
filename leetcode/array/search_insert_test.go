package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	testCases := []struct {
		nums      []int
		target    int
		wantIndex int
	}{
		{
			nums:      []int{1, 2, 3, 5, 6},
			target:    5,
			wantIndex: 3,
		},
		{
			nums:      []int{1, 3, 5, 6},
			target:    2,
			wantIndex: 1,
		},
		{
			nums:      []int{1, 3, 5, 6},
			target:    7,
			wantIndex: 4,
		},
		{
			nums:      []int{1, 3, 5, 6},
			target:    0,
			wantIndex: 0,
		},
		{
			nums:      []int{1, 3, 5},
			target:    3,
			wantIndex: 1,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case number: %d", i+1), func(t *testing.T) {
			got := searchInsert(tc.nums, tc.target)
			assert.Equal(t, tc.wantIndex, got)
		})
	}
}
