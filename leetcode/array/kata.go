package array

import (
	"fmt"
	"log/slog"
)

func Kata() {
	a := []int{5, 6, 7, 8, 9, 10, 11, 12}

	for i := 0; i < len(a); i++ {
		slog.Info("F", slog.Int("I:", a[i]))
	}

	reverse := make([]int, 0, len(a))

	for j := len(a) - 1; j >= 0; j-- {
		reverse = append(reverse, a[j])
	}

	slog.Info(fmt.Sprintf("%+v", reverse))
}

func reverse(list []int) []int {
	r := make([]int, 0, len(list))

	for j := len(list) - 1; j >= 0; j-- {
		r = append(r, list[j])
	}

	return r
}
