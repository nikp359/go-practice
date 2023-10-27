package divisor

import "strings"

func gcdOfStrings(str1 string, str2 string) string {
	if str1 == str2 {
		return str1
	}

	if len(str2) > len(str1) {
		str1, str2 = str2, str1
	}

	rem := gcd(len(str1), len(str2))
	if rem == 1 {
		return ""
	}

	if strings.HasPrefix(str1, str2[:rem]) {
	}

	if str1 == str2[:rem-1] {
		return str1
	}

	return gcdOfStrings(str1[rem:], str2)
}

func gcd(num1, num2 int) int {
	if num1 < num2 {
		num1, num2 = num2, num1
	}

	remainder := num1 % num2
	if remainder == 0 {
		return num2
	}

	return gcd(remainder, num2)
}
