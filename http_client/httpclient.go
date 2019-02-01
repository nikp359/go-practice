package httpclient

type Client struct {
	URL string
}

func (c *Client) UUID() (string, error) {
	return "123", nil
}
