package main

import (
	"flag"
	"fmt"
)

func main() {
	csvPath := flag.String("f", "problems.csv", "test file name")
	flag.Parse()
	fmt.Printf("File name: %s\n", *csvPath)
}
