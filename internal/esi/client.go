package esi

import (
	"net/http"
	"strings"
	"sync"
	"time"
)

type Client struct {
	clientID     string
	clientSecret string
	baseURL      string
	redirectURI  string

	httpClient *http.Client

	cache   map[int]string
	cacheMu sync.RWMutex
}

var (
	_ MarketClient = (*Client)(nil)
	_ ItemResolver = (*Client)(nil)
)

func NewClient(
	baseURL string,
	clientID string,
	clientSecret string,
	redirectURI string,
	timeout time.Duration,
) *Client {
	return &Client{
		baseURL:      strings.TrimRight(baseURL, "/"),
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURI:  redirectURI,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}
