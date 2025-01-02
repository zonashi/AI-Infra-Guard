// Package httpx http option
package httpx

import (
	"time"

	"github.com/projectdiscovery/fastdialer/fastdialer"
)

// HTTPOptions http options
type HTTPOptions struct {
	Timeout          time.Duration
	RetryMax         int
	FollowRedirects  bool
	HTTPProxy        string
	Unsafe           bool
	DefaultUserAgent string
	Dialer           *fastdialer.Dialer
}
