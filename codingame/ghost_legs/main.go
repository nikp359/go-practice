package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//import "strings"
//import "strconv"

type GhostLegs struct {
	Width  int
	Height int
}

func main() {
	gl := readInput(os.Stdin)
	fmt.Println(gl)
}

func readInput(inpit io.Reader) *GhostLegs {
	scanner := bufio.NewScanner(inpit)
	scanner.Scan()

	gl := &GhostLegs{}
	fmt.Sscan(scanner.Text(), &gl.Width, &gl.Height)
	fmt.Printf("W: %v, H: %v\n", gl.Width, gl.Height)

	for i := 0; i < gl.Width; i++ {
		scanner.Scan()
		line := scanner.Text()
		fmt.Printf("%v\n", line)
	}
	return gl
}
