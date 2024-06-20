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
			input: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			input: [][]int{{1, 4}, {2, 3}},
			want:  [][]int{{1, 4}},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case: %d", i), func(t *testing.T) {
			got := mergeIntervals(tc.input)

			assert.Equal(t, tc.want, got)
		})
	}
}
