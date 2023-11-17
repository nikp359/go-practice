package productexceptself

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			nums: []int{3, 1, 2, 3, 4},
			want: []int{24, 72, 36, 24, 18},
		},
		{
			nums: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
		{
			nums: []int{-1, -1, 1, -1, -1, -1, 1, -1, 1, -1, 1, 1, -1, -1, -1, -1, -1, -1, 1, 1, 1, 1, 1, -1, -1, 1, 1, -1, 1, 1, 1, 1, 1, -1},
			want: []int{1, 1, -1, 1, 1, 1, -1, 1, -1, 1, -1, -1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, 1, 1, -1, -1, 1, -1, -1, -1, -1, -1, 1},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case number: %d", i+1), func(t *testing.T) {
			got := productExceptSelf(tc.nums)
			assert.Equal(t, tc.want, got)
		})
	}
}

func BenchmarkProductExceptSelf(b *testing.B) {
	nums := []int{1, 2, 3, 6, 2, 3, 1, 5, 6, 7 - 1, -1, 1, -1, -1, -1, 1, -1, 1, -1, 1, 1, -1, -1, -1, -1, -1, -1, 1, 1, 1, 1, 1, -1, -1, 1, 1, -1, 1, 1, 1, 1, 1, -1}

	for i := 0; i < b.N; i++ {
		productExceptSelf(nums)
	}
}
