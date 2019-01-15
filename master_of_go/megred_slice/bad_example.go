package main

import (
	"fmt"
)

func badExample() {
	s1 := []string{"a.txt", "b.txt", "c.txt"}
	s2 := []string{"b.txt", "c.txt", "d.txt"}

	tempMap := make(map[string]struct{}, len(s1)+len(s2))
	for _, f := range s1 {
		tempMap[f] = struct{}{}
	}

	fmt.Println(tempMap)

	for _, f := range s2 {
		tempMap[f] = struct{}{}
	}

	MergedSlice := make([]string, len(tempMap))
	i := 0
	for f, _ := range tempMap {
		MergedSlice[i] = f
		i++
	}

	fmt.Println(MergedSlice)
}
