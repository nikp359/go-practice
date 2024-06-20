package array

import (
	"fmt"
	"log/slog"
)

func Kata() {
	a := []int{1, 3, 5, 6, 8, 11, 17}
	b := []int{1, 2, 4, 9, 14, 19, 21}

	got := merge(a, b)

	fmt.Printf("Got: %+v\n", got)

	as, bs := separation(got)

	slog.Info(fmt.Sprintf("Sep A: %+v, B: %+v\n", as, bs))
}

func separation(input []int) ([]int, []int) {
	a := make([]int, 0, len(input)/2)
	b := make([]int, 0, len(input)/2)

	for _, v := range input {
		if v%2 == 0 {
			a = append(a, v)
		} else {
			b = append(b, v)
		}
	}

	return a, b
}

func merge(a, b []int) []int {
	var i, j int
	resp := make([]int, 0, len(a)+len(b))

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			resp = append(resp, a[i])
			i++
		} else {
			resp = append(resp, b[j])
			j++
		}
	}

	if i < j {
		resp = append(resp, a[i:]...)
	} else {
		resp = append(resp, b[j:]...)
	}

	return resp
}

func kataPlusOne(list []int) {
	for i := 0; i < len(list); i++ {
		list[i]++
	}
}

func reverse(list []int) []int {
	for i, j := 0, len(list)-1; i <= j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}

	return list
}
