package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	ch1 := make(chan int, 1)

	wg.Add(1)
	go func(in chan int) {
		defer wg.Done()
		val := <-in
		fmt.Println("GO: get from chan", val)
		fmt.Println("GO: after read from chan 1")
	}(ch1)

	wg.Add(1)
	go func(in chan int) {
		defer wg.Done()
		val := <-in
		fmt.Println("GO: get from chan", val)
		fmt.Println("GO: after read from chan 2")
	}(ch1)
	ch1 <- 42
	ch1 <- 100500

	fmt.Println("MAIN: after put to chan")
	wg.Wait()
	//fmt.Scanln()
}
