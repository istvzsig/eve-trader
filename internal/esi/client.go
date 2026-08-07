package esi

import (
	"net/http"
	"sync"
	"time"
)

type Client struct {
	baseURL string

	httpClient *http.Client

	cache   map[int]string
	cacheMu sync.RWMutex
}

var (
	_ MarketClient = (*Client)(nil)
	_ ItemResolver = (*Client)(nil)
)

func NewClient(baseURL string, timeout time.Duration) *Client {
	return &Client{
		baseURL: baseURL,

		httpClient: &http.Client{
			Timeout: timeout,
		},

		cache: make(map[int]string),
	}
}
