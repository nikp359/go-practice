package maxvowels

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxVowels(t *testing.T) {
	testCases := []struct {
		input           string
		substringLength int
		wantMaxVowels   int
	}{
		{
			input:           "abciiidef",
			substringLength: 3,
			wantMaxVowels:   3,
		},
		{
			input:           "aeiou",
			substringLength: 2,
			wantMaxVowels:   2,
		},
		{
			input:           "leetcode",
			substringLength: 3,
			wantMaxVowels:   2,
		},
		{
			input:           "k",
			substringLength: 1,
			wantMaxVowels:   0,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case n: %d", i+1), func(t *testing.T) {
			got := maxVowels(tc.input, tc.substringLength)
			assert.Equal(t, tc.wantMaxVowels, got)
		})
	}
}
