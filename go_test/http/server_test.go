package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGETPlayers(t *testing.T) {
	t.Run("Return Pepper's score", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/player/Pepper", nil)
		resp := httptest.NewRecorder()

		PlayerServer(resp, req)

		got := resp.Body.String()
		want := "20"

		assert.Equal(t, got, want)
	})
}
