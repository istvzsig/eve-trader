package esi

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
)

func (c *Client) GetAccessToken(ctx context.Context) (string, error) {
	verifier := randomString(32)

	hash := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(hash[:])

	state := randomString(32)

	callback := make(chan string, 1)
	errors := make(chan error, 1)

	server := &http.Server{
		Addr: ":8080",
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != state {
			http.Error(w, "invalid state", http.StatusBadRequest)
			errors <- fmt.Errorf("invalid OAuth state")
			return
		}

		code := r.URL.Query().Get("code")
		if code == "" {
			errors <- fmt.Errorf("missing authorization code")
			return
		}

		fmt.Fprintln(w, "EVE authorization successful. You can close this window.")

		callback <- code
	})

	server.Handler = mux

	go func() {
		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			errors <- err
		}
	}()

	params := url.Values{}
	params.Set("response_type", "code")
	params.Set("redirect_uri", c.redirectURI)
	params.Set("client_id", c.clientID)
	params.Set("scope", "esi-skills.read_skills.v1")
	params.Set("state", state)
	params.Set("code_challenge", challenge)
	params.Set("code_challenge_method", "S256")

	authURL := "https://login.eveonline.com/v2/oauth/authorize?" + params.Encode()

	fmt.Println("Opening EVE Online login...")
	fmt.Println(authURL)

	if err := openBrowser(authURL); err != nil {
		fmt.Println("Could not open browser automatically.")
		fmt.Println("Open the URL above manually.")
	}

	var code string

	select {
	case code = <-callback:
	case err := <-errors:
		_ = server.Shutdown(ctx)
		return "", err
	case <-ctx.Done():
		_ = server.Shutdown(ctx)
		return "", ctx.Err()
	}

	_ = server.Shutdown(ctx)

	return c.exchangeCode(ctx, code, verifier)
}

func (c *Client) exchangeCode(
	ctx context.Context,
	code string,
	verifier string,
) (string, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("client_id", c.clientID)
	data.Set("redirect_uri", c.redirectURI)
	data.Set("code_verifier", verifier)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"https://login.eveonline.com/v2/oauth/token",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set(
		"Content-Type",
		"application/x-www-form-urlencoded",
	)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("oauth token request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf(
			"oauth token request failed: %s",
			resp.Status,
		)
	}

	var token struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return "", err
	}

	if token.AccessToken == "" {
		return "", fmt.Errorf("OAuth response contained no access token")
	}

	return token.AccessToken, nil
}

func randomString(n int) string {
	b := make([]byte, n)

	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	return base64.RawURLEncoding.EncodeToString(b)
}

func openBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}

	return cmd.Start()
}
