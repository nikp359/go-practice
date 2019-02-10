package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestReadInput(t *testing.T) {
	// want := &GhostLegs{
	// 	Width:  7,
	// 	Height: 7,
	// }
	Input := `7 7
A  B  C
|  |  |
|--|  |
|  |--|
|  |--|
|  |  |
1  2  3`

	//scanner := bufio.NewScanner(strings.NewReader(Input))
	//scanner.Scan()
	//fmt.Println(scanner.Text())

	got := readInput(strings.NewReader(Input))
	fmt.Printf("GOT: %+v", got)
}
