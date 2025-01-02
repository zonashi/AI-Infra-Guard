// Package httpx provides HTTP utility functions and types for handling user agents
package httpx

import (
	"math/rand"
)

// UserAgent represents a slice of user agent strings
type UserAgent []string

// userAgents contains a collection of common browser user agent strings
var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML like Gecko) Chrome/44.0.2403.155 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML like Gecko) Chrome/46.0.2486.0 Safari/537.36 Edge/13.10586",
}

// GetRandomUserAgent returns a random user agent string from the predefined list.
// It uses math/rand to generate a random index.
func GetRandomUserAgent() string {
	return userAgents[rand.Intn(len(userAgents))]
}
