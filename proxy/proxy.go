package proxy

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"
)

// smartProxy initializes and returns a proxy function for use in a Transport
func SmartProxy() func(*http.Request) (*url.URL, error) {
	// base url
	base := "http://%s:%s@%s"
	// fill credentials into url
	rawUrl := fmt.Sprintf(base, os.Getenv("SMARTPROXY_USERNAME"), os.Getenv("SMARTPROXY_PASSWORD"), os.Getenv("SMARTPROXY_ADDRESS"))
	// parse proxy url
	proxyUrl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println("invalid url to parse when creating proxy transport. err: ", err)
	}
	// setup proxy transport
	proxy := http.ProxyURL(proxyUrl)

	return proxy
}

// TODO: dockerize and start TOR
// torProxy initializes and returns a proxy function for use in a Transport
func TorProxy() func(*http.Request) (*url.URL, error) {
	// a source of uniformly-distributed pseudo-random
	rand.Seed(time.Now().UnixNano())
	// random int
	num := rand.Intn(0x7fffffff-10000) + 10000
	// base url localhost for now
	base := "socks5://%d:x@127.0.0.1:9050"
	// proxy url with random credentials
	rawUrl := fmt.Sprintf(base, num)
	// parse proxy url
	proxyUrl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println("invalid url to parse when creating proxy transport. err: ", err)
	}
	// setup proxy transport
	proxy := http.ProxyURL(proxyUrl)

	return proxy
}
