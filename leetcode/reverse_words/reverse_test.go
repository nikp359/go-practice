package reversewords

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{
			input: "the sky is blue",
			want:  "blue is sky the",
		},
		{
			input: "  hello world  ",
			want:  "world hello",
		},
		{
			input: "  a good   example ",
			want:  "example good a",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case number: %d", i+1), func(t *testing.T) {
			got := reverseWords(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}
