package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestGetRoot(t *testing.T) {
	s := prep()
	Convey("Empty GET to / should return 200 OK", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		s.Engine.ServeHTTP(w, req)
		So(w.Code, ShouldEqual, http.StatusOK)
		//assert.Equal(t, 200, w.Code)
	})
}

func TestGetColors(t *testing.T) {
	s := prep()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/color", nil)
	s.Engine.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	body, _ := ioutil.ReadAll(w.Body)
	colors := []Color{}
	json.Unmarshal(body, &colors)
	assert.Equal(t, 3, len(colors))

}

func prep() *Server {
	var server Server
	server.Store = NewColors()
	server.Engine = server.setupRouter()
	return &server
}
