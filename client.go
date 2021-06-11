package merrygoround

import (
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

/*
If several goroutines are sending requests,
new connections will be created the pool has all
connections busy and will create new ones.
We limit the maximum number of connections per host.
*/
var (
	maxIdleConnections = 10         // Max Idle Connections
	once               sync.Once    // sync so we only setup 1 client - some important shit lmao
	netClient          *http.Client // http client
)

func NewClient(proxy func(*http.Request) (*url.URL, error)) *http.Client {
	once.Do(func() {
		// transport configuratin
		var netTransport = &http.Transport{
			Proxy:        proxy,              // We can use Tor or Smart Proxy - rotating IP addresses - if nil no proxy is used
			MaxIdleConns: maxIdleConnections, // max idle connections
			// Dialer
			Dial: (&net.Dialer{
				Timeout: 20 * time.Second, // max dialer timeout
			}).Dial,
			TLSHandshakeTimeout: 20 * time.Second, // transport layer security max timeout
		}
		// Client
		netClient = &http.Client{
			Timeout:   time.Second * 20, // roundtripper timeout
			Transport: netTransport,     // how our HTTP requests are made
		}
	})

	return netClient
}

func NewIP(client *http.Client, t *http.Transport) {
	client.Transport = t
}
