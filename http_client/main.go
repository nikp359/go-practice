package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

const port = "localhost:8099"

func startWebserver() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 50)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	go http.ListenAndServe(port, nil)
}

func startLoadTest(client *http.Client) {
	count := 0
	for {
		resp, err := client.Get(fmt.Sprintf("http://%s/", port))
		if err != nil {
			panic(fmt.Sprintf("Got error: %v", err))
		}
		defer resp.Body.Close()
		io.Copy(io.Discard, resp.Body)
		log.Printf("Finished GET request #%v", count)
		count++
	}
}

func newCustomClient() *http.Client {
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).Dial,
		IdleConnTimeout:       90 * time.Second,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}
}

func newClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

func main() {
	fmt.Println("Start http client test")
	startWebserver()
	time.Sleep(time.Second * 1)

	http := newCustomClient()

	for i := 0; i < 100; i++ {
		go startLoadTest(http)
	}

	time.Sleep(time.Second * 2400)
}
