package concurency

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {
	received := consumerCounter(generator(33))
	assert.Equal(t, 33, received)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	received = consumerCounter(randomGenerator(ctx))
	t.Log(received)
}
