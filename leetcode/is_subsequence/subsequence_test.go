package subsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	testCases := []struct {
		input  string
		target string
		want   bool
	}{
		{
			input:  "abc",
			target: "ahbgdc",
			want:   true,
		},
		{
			input:  "axc",
			target: "ahbgdc",
			want:   false,
		},
		{
			input:  "aaaaaa",
			target: "bbaaaa",
			want:   false,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case number: %d", i+1), func(t *testing.T) {
			got := isSubsequence(tc.input, tc.target)
			assert.Equal(t, tc.want, got)
		})
	}
}
