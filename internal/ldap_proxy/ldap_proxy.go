package ldap_proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	jwtauthentication "github.com/yigithanbalci/ldap-reverse-proxy/internal/jwt-authentication"
)

// NewProxy takes target host and creates a reverse proxy
func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	return httputil.NewSingleHostReverseProxy(url), nil
}

// ProxyRequestHandler handles the http request using proxy
func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		successful, err := jwtauthentication.ValidateToken(w, r)
		if err == nil && successful == true {
			proxy.ServeHTTP(w, r)
		}
	}
}
