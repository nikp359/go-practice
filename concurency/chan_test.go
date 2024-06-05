package concurency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChan(t *testing.T) {
	ch := make(chan int, 1)

	ch <- 1

	close(ch)

	got, ok := <-ch
	assert.True(t, ok)
	assert.Equal(t, 1, got)

	// read from closed chanel
	got, ok = <-ch
	assert.False(t, ok)
	assert.Equal(t, 0, got)

	// write to closed chanel
	assert.Panics(t, func() {
		ch <- 5
	})

}
