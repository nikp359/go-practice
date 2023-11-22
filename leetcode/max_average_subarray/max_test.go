package maxaveragesubarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxAverage(t *testing.T) {
	testCases := []struct {
		nums           []int
		subarrayLength int
		wantAverage    float64
	}{
		{
			nums:           []int{1, 12, -5, -6, 50, 3},
			subarrayLength: 4,
			wantAverage:    12.75,
		},
		{
			nums:           []int{1, 12, -5, 0, 50, 7},
			subarrayLength: 3,
			wantAverage:    19,
		},
		{
			nums:           []int{7},
			subarrayLength: 1,
			wantAverage:    7,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case n: %d", i+1), func(t *testing.T) {
			got := findMaxAverage(tc.nums, tc.subarrayLength)
			assert.InDelta(t, tc.wantAverage, got, 0.01)
		})
	}
}
