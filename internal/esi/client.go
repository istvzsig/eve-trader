package esi

import (
	"net/http"
	"sync"
	"time"
)

type Client struct {
	clientID    string
	baseURL     string
	redirectURI string

	httpClient *http.Client

	cache   map[int]string
	cacheMu sync.RWMutex
}

var (
	_ MarketClient = (*Client)(nil)
	_ ItemResolver = (*Client)(nil)
)

func NewClient(clientID, baseURL, redirectURI string, timeout time.Duration) *Client {
	return &Client{
		clientID:    clientID,
		baseURL:     baseURL,
		redirectURI: redirectURI,

		httpClient: &http.Client{
			Timeout: timeout,
		},

		cache: make(map[int]string),
	}
}
