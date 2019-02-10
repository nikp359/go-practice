package httpclient

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func HTTPBinMock(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{'uuid': 'ec12b7c7-02a8-4a8c-ac0f-ff322fce6628'}`)
}

func TestHTTPClient(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(HTTPBinMock))

	client := &Client{
		URL: ts.URL,
	}

	result, err := client.UUID()

	assert.Equal(t, result, "ec12b7c7-02a8-4a8c-ac0f-ff322fce6628", "they should be equal")
	assert.Nil(t, err)

}
