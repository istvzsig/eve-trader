package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

func Load() (EnvConfig, error) {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return EnvConfig{}, fmt.Errorf("loading .env: %w", err)
	}

	cfg := EnvConfig{
		ClientID:     os.Getenv("EVE_CLIENT_ID"),
		ClientSecret: os.Getenv("EVE_CLIENT_SECRET"),
		RedirectURI:  getenv("EVE_REDIRECT_URI", "http://localhost:8080/callback"),
	}

	if cfg.ClientID == "" {
		return EnvConfig{}, fmt.Errorf("EVE_CLIENT_ID is not set")
	}

	return cfg, nil
}

func getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}

const (
	BaseURL = "https://esi.evetech.net/latest"
	JitaID  = 10000002
	Timeout = 20 * time.Second
)
