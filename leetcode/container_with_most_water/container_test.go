package containerwithmostwater

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		heights []int
		area    int
	}{
		{
			heights: []int{1, 8, 6, 2, 5, 4, 8, 3, 7, 3, 2},
			area:    49,
		},
		{
			heights: []int{1, 1},
			area:    1,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case n: %d", i), func(t *testing.T) {
			got := maxArea(tc.heights)
			assert.Equal(t, tc.area, got)
		})
	}

}
