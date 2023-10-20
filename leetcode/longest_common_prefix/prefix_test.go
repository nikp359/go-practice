package prefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		input []string
		want  string
	}{
		{
			input: nil,
			want:  "",
		},
		{
			input: []string{"flower", "flow", "flight"},
			want:  "fl",
		},
		{
			input: []string{"dog", "racecar", "car"},
			want:  "",
		},
	}

	for _, tc := range testCases {
		resp := longestCommonPrefix(tc.input)
		assert.Equal(t, tc.want, resp)
	}
}
