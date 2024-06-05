package concurency

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
)

func generator(ln int) <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < ln; i++ {
			out <- i
		}

		close(out)
	}()

	return out
}

func randomGenerator(ctx context.Context) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for {
			select {
			case out <- rand.Int():
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

func consumer(in <-chan int) {
	for v := range in {
		fmt.Printf("Recieved: %d\n", v)
	}

	fmt.Println("Done")
}

func consumerCounter(in <-chan int) int {
	var counter int

	for i := range in {
		slog.Info(fmt.Sprintf("Recived: %d", i))
		counter++
	}

	return counter
}
