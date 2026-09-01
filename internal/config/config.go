package config

import (
	"os"
	"time"
)

type env struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

// TODO: add error handling
func getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}

var Env = env{
	ClientID:     getenv(os.Getenv("EVE_CLIENT_ID"), ""),
	ClientSecret: getenv(os.Getenv("EVE_CLIENT_SECRET"), ""),
	RedirectURI:  getenv("EVE_REDIRECT_URI", "http://localhost:4916/callback"),
}

const (
	BaseURL = "https://esi.evetech.net/latest"
	JitaID  = 10000002
	Timeout = 20 * time.Second
)
