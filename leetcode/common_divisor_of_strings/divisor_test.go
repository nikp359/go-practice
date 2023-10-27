package divisor

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGcdOfStrings(t *testing.T) {
	testCases := []struct {
		str1 string
		str2 string
		want string
	}{
		// {
		// 	str1: "A",
		// 	str2: "B",
		// 	want: "",
		// },
		// {
		// 	str1: "ABCABC",
		// 	str2: "ABC",
		// 	want: "ABC",
		// },
		{
			str1: "ABABAB",
			str2: "ABAB",
			want: "AB",
		},
		// {
		// 	str1: "ABCABCABC",
		// 	str2: "ABAB",
		// 	want: "",
		// },
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case number: %d", i+1), func(t *testing.T) {
			resp := gcdOfStrings(tc.str1, tc.str2)
			assert.Equal(t, tc.want, resp)
		})
	}
}

func TestGCD(t *testing.T) {
	testCases := []struct {
		num1 int
		num2 int
		want int
	}{
		{
			num1: 9,
			num2: 3,
			want: 3,
		},
		{
			num1: 4,
			num2: 14,
			want: 2,
		},
		{
			num1: 17,
			num2: 19,
			want: 1,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case number: %d", i+1), func(t *testing.T) {
			resp := gcd(tc.num1, tc.num2)
			assert.Equal(t, tc.want, resp)
		})
	}

}
