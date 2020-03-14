package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (m *MyApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/user/profile":
		m.handlerProfile(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (m *OtherApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Addr:", "123", "URL:", r.URL.String())
}

type APIResponce struct {
	Error    string      `json:"error"`
	Responce interface{} `json:"responce"`
}

func (m *MyApi) handlerProfile(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query().Get("login")
	ctx := context.Background()
	user, err := m.Profile(ctx, ProfileParams{Login: login})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responce := APIResponce{
		Responce: user,
	}

	js, err := json.Marshal(responce)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
