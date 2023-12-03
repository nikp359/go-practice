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
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case n: %d", i+1), func(t *testing.T) {
			got := maxVowels(tc.input, tc.substringLength)
			assert.Equal(t, tc.wantMaxVowels, got)
		})
	}
}
