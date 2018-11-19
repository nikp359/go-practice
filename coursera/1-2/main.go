package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("MAIN")

	freeFlowJobs := []job{
		job(func(in, out chan interface{}) {
			out <- int(0)
			out <- int(1)
			out <- int(3)
		}),
		job(SingleHash),
		job(MultiHash),
		job(CombineResults),
		job(func(in, out chan interface{}) {
			for val := range in {
				fmt.Println("collected", val)
			}
		}),
	}
	start := time.Now()
	ExecutePipeline(freeFlowJobs...)
	fmt.Println("Execute time:", time.Since(start))
}
