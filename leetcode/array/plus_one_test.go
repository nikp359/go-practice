package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	testCases := []struct {
		digits []int
		want   []int
	}{
		{
			digits: []int{6, 9, 9, 9},
			want:   []int{7, 0, 0, 0},
		},
		{
			digits: []int{9, 9, 9, 9},
			want:   []int{1, 0, 0, 0, 0},
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case number: %d", i+1), func(t *testing.T) {
			got := plusOne(tc.digits)
			assert.Equal(t, tc.want, got)
		})
	}
}
