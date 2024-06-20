package concurency

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"sync"
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

func randomGenerator1(ctx context.Context) <-chan int {
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

func worker(id int, jobs <-chan int, result chan<- int, wg *sync.WaitGroup) {
	slog.Info("Start worker:", slog.Int("ID:", id))

	for j := range jobs {
		result <- j * 2
	}

	wg.Done()
}

func merge(ctx context.Context, in <-chan int) []int {
	r := make([]int, 0, len(in))

	for {
		select {
		case i, ok := <-in:
			if ok {
				r = append(r, i)
			}
		case <-ctx.Done():
			return r
		}
	}
}

func funOut(numJobs int) {
	jobs := make(chan int, 0)
	result := make(chan int, 0)
	done := make(chan struct{}, 0)
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	for w := 0; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, result, wg)
	}

	go func(done chan<- struct{}) {
		r := merge(ctx, result)
		slog.Info(fmt.Sprintf("response: %+v", r))
		close(done)
	}(done)

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
	cancel()
	close(result)
	<-done
}
