package main

import "fmt"

func goodMerge() {
	s1 := []string{"a.txt", "b.txt", "c.txt"}
	s2 := []string{"b.txt", "c.txt", "d.txt", "e.txt"}

	s1 = append(s1, s2...)
	seen := make(map[string]struct{}, len(s1))
	j := 0
	for _, s := range s1 {
		if _, ok := seen[s]; ok {
			continue
		}
		seen[s] = struct{}{}
		s1[j] = s
		j++
	}

	fmt.Printf("s1: %#v\n", s1)
	mergedSlice := s1[:j]
	fmt.Printf("mergedSlice: %#v\n", mergedSlice)
}
