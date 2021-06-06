package proxy

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type ProxyService struct {
}

func (p *ProxyService) New() func(*http.Request) (*url.URL, error) {

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
