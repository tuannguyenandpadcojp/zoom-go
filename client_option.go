package zoom

import (
	"net/http"
)

type ClientOpt func(c *Client) *Client

func WithHTTPClient(conn *http.Client) ClientOpt {
	return func(c *Client) *Client {
		if conn != nil {
			c.conn = conn
		}
		return c
	}
}
