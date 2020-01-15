package httpbin

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	cl := NewStubClient(t)

	c := NewClient(cl)
	resp, err := c.Get("/get")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

type StubTransport struct {
	t *testing.T
}

func (s *StubTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	s.t.Log(req)
	return &http.Response{
		StatusCode: http.StatusOK,
	}, nil
}

func NewStubClient(t *testing.T) *http.Client {
	return &http.Client{
		Transport: &StubTransport{
			t: t,
		},
	}
}
