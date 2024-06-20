package concurency

import (
	"testing"
)

func TestPipeline(t *testing.T) {
	// received := consumerCounter(generator(33))
	// assert.Equal(t, 33, received)

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	// defer cancel()
	// received = consumerCounter(randomGenerator(ctx))
	// t.Log(received)

	funOut(10)
}
