package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		nums     []int
		wantNums []int
		wantLen  int
	}{
		{
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			wantNums: []int{0, 1, 2, 3, 4},
			wantLen:  5,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case number: %d", i+1), func(t *testing.T) {
			ln := removeDuplicates(tc.nums)
			assert.Equal(t, tc.wantLen, ln)
			assert.Equal(t, tc.wantNums, tc.nums[:ln])
		})
	}
}
