package pivotindex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPivotIndex(t *testing.T) {
	testCases := []struct {
		input     []int
		wantIndex int
	}{
		{
			input:     []int{1, 7, 3, 6, 5, 6},
			wantIndex: 3,
		},
		{
			input:     []int{1, 2, 3},
			wantIndex: -1,
		},
		{
			input:     []int{2, 1, -1},
			wantIndex: 0,
		},
		{
			input:     []int{-1, -1, -1, -1, -1, 0},
			wantIndex: 2,
		},
		{
			input:     []int{-1, -1, 0, 0, -1, -1},
			wantIndex: 2,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case n: %d", i+1), func(t *testing.T) {
			got := pivotIndex(tc.input)
			assert.Equal(t, tc.wantIndex, got)
		})
	}
}
