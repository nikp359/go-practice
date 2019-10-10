package main

import "fmt"

func main() {
	y := NbYear(1500, 5, 100, 5000)
	fmt.Println(y)
}

//NbYear How many years does the town need to see its population greater or equal to responce inhabitants
func NbYear(p0 int, percent float64, aug int, p int) int {
	var year int

	for {
		p0 = growth(p0, percent, aug)
		year++
		if p0 >= p || p0 == 0 {
			return year
		}
	}
}

func growth(p0 int, percent float64, aug int) int {
	return p0 + int(float64(p0)*(percent/100.00)) + aug
}
