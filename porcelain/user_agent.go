package porcelain

import (
	"net/http"
)

const (
	DefaultUserAgent = "go-databricks"
)

// UserAgentTransport is a http.RoundTripper that adds a User-Agent header
type UserAgentTransport struct {
	agent string
	rt    http.RoundTripper
}

// NewUserAgentTransport initializes a new UserAgentTransport
func NewUserAgentTransport(rt http.RoundTripper, agent string) http.RoundTripper {
	if agent == "" {
		agent = DefaultUserAgent
	}

	return &UserAgentTransport{
		agent: agent,
		rt:    rt,
	}
}

// RoundTrip implements the http.RoundTripper, setting a User-Agent header
func (ua *UserAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", ua.agent)
	return ua.rt.RoundTrip(req)
}
