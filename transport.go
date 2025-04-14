package basicauth

import (
	"maps"
	"net/http"
)

// Transport adds HTTP Basic Authentication to each HTTP request
type Transport struct {
	// Username is the username for HTTP Basic Authentication
	Username string
	// Password is the password for HTTP Basic Authentication
	Password string
	// Base is the base RoundTripper used to make HTTP requests.
	// If nil, http.DefaultTransport is used.
	Base http.RoundTripper
}

// RoundTrip implements the http.RoundTripper interface
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Per the http.RoundTripper interface, we need to clone the request to avoid modifying it
	req2 := new(http.Request)
	*req2 = *req
	req2.Header = maps.Clone(req.Header)
	// Add HTTP Basic Authentication to the cloned request
	req2.SetBasicAuth(t.Username, t.Password)
	// And finally call the base http.RoundTripper
	if t.Base == nil {
		return http.DefaultTransport.RoundTrip(req2)
	}
	return t.Base.RoundTrip(req2)
}
