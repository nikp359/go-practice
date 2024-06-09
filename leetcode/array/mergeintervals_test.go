package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeIntervals(t *testing.T) {
	testCases := []struct {
		input [][]int
		want  [][]int
	}{
		{
			input: [][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}},
			want:  [][]int{[]int{1, 6}, []int{8, 10}, []int{15, 18}},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case: %d", i), func(t *testing.T) {
			got := mergeIntervals(tc.input)

			assert.Equal(t, tc.want, got)
		})
	}
}
