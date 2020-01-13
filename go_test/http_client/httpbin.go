package httpbin

import (
	"net/http"
	"net/url"
)

const defaultBaseURL = "https://httpbin.org/"

type Client struct {
	client  *http.Client
	baseURL string
}

func NewClient(cl *http.Client) *Client {
	return &Client{
		client:  cl,
		baseURL: defaultBaseURL,
	}
}

func (c *Client) Get(path string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, c.url(path), nil)
	if err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

func (c *Client) Post(path string) {

}

func (c *Client) url(path string) string {
	u, _ := url.Parse(c.baseURL)
	u.Path = path
	return u.String()
}
