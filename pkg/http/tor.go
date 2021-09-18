package http

import (
	"merryGoRound/pkg/proxy"
	"net"
	"net/http"
	"sync"
	"time"
)

// Tor represents an HTTP Client over the Tor network
type Tor struct {
	MaxTimeout time.Duration			// Max Timeout
	MaxIdleConnections int        	// Max Idle Connections
	once               sync.Once    // sync so we only set up 1 client
	netClient          *http.Client // http client
}

// New creates an instance of a Tor Client
func(t *Tor) New() *http.Client {
	// create one transport
	t.once.Do(func() {
		// transport configuration
		var netTransport = &http.Transport{
			Proxy:        proxy.TorProxy(),         // We can use Tor or Smart Proxy - rotating IP addresses - if nil no proxy is used
			MaxIdleConns: t.MaxIdleConnections, // max idle connections
			// TODO: Dial is deprecated
			Dial: (&net.Dialer{
				Timeout: t.MaxTimeout, // max dialer timeout
			}).Dial,
			TLSHandshakeTimeout: t.MaxTimeout, // transport layer security max timeout
		}
		// Client
		t.netClient = &http.Client{
			Timeout:   t.MaxTimeout, // round tripper timeout
			Transport: netTransport,     // how our HTTP requests are made
		}
	})

	return t.netClient
}


// NewIP swaps a client's transport with a new one
func (t *Tor) NewIP() {
	// transport configuration
	var newTransport = &http.Transport{
		Proxy:        proxy.TorProxy(),         // We can use Tor or Smart Proxy - rotating IP addresses - if nil no proxy is used
		MaxIdleConns: t.MaxIdleConnections, // max idle connections
		// TODO: Dial is deprecated
		Dial: (&net.Dialer{
			Timeout: t.MaxTimeout, // max dialer timeout
		}).Dial,
		TLSHandshakeTimeout: t.MaxTimeout, // transport layer security max timeout
	}
	t.netClient.Transport = newTransport
}
