package main

import (
	"log"
	"net/http"

	jwtauthentication "github.com/yigithanbalci/ldap-reverse-proxy/internal/jwt-authentication"
	"github.com/yigithanbalci/ldap-reverse-proxy/internal/ldap_proxy"
)

func main() {
	proxy, err := ldap_proxy.NewProxy("http://localhost:9092/superhero/hello")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/api", ldap_proxy.ProxyRequestHandler(proxy))
	http.HandleFunc("/signin", jwtauthentication.Signin)
	http.HandleFunc("/refresh", jwtauthentication.Refresh)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
