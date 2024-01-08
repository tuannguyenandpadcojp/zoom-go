package zoom

import (
	"fmt"
	"net/http"
)

type Client struct {
	conn   *http.Client
	key    string
	secret string
}

func DefaultHTTPClient() *http.Client {
	return http.DefaultClient
}

func NewClient(opts ...ClientOpt) (*Client, error) {
	cfg, err := loadConfig()
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}
	if cfg.SDKKey == "" || cfg.SDKSecret == "" {
		return nil, fmt.Errorf("missing sdk credentials")
	}
	c := &Client{
		conn:   DefaultHTTPClient(),
		key:    cfg.SDKKey,
		secret: cfg.SDKSecret,
	}
	for i := range opts {
		c = opts[i](c)
	}

	return c, nil
}
