package reverse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseVowels(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{
			input: "hello",
			want:  "holle",
		},
		{
			input: "healthcare",
			want:  "healthcare",
		},
		{
			input: "leetcode",
			want:  "leotcede",
		},
		{
			input: "a",
			want:  "a",
		},
		{
			input: "aabbe",
			want:  "eabba",
		},
		{
			input: "a.",
			want:  "a.",
		},
		{
			input: "aA",
			want:  "Aa",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case number: %d", i+1), func(t *testing.T) {
			got := reverseVowels(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}
