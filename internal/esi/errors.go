package esi

import "errors"

// Sentinel errors let callers use errors.Is to make retry decisions
// based on error *kind*, not by parsing status text.
var (
	ErrNotFound     = errors.New("esi: resource not found")
	ErrUnauthorized = errors.New("esi: unauthorized (token missing or expired)")
	ErrForbidden    = errors.New("esi: forbidden (missing scope)")
	ErrRateLimited  = errors.New("esi: rate limited")
	ErrServer       = errors.New("esi: server error")
	ErrUnexpected   = errors.New("esi: unexpected status")
)

// classifyStatus maps an HTTP status code to a sentinel error.
// Kept in one place so every endpoint classifies consistently.
func classifyStatus(status int) error {
	switch {
	case status >= 200 && status < 300:
		return nil
	case status == 401:
		return ErrUnauthorized
	case status == 403:
		return ErrForbidden
	case status == 404:
		return ErrNotFound
	case status == 429:
		return ErrRateLimited
	case status >= 500:
		return ErrServer
	default:
		return ErrUnexpected
	}
}
