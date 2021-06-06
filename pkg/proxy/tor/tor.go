package proxy

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
)

type ProxyService struct {
}

func (p *ProxyService) New() func(*http.Request) (*url.URL, error) {

	num := rand.Intn(0x7fffffff-10000) + 10000
	// base url
	base := "socks5://%d:x@:9050"
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
