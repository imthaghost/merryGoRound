package http

import (
	"merryGoRound/proxy"
	"net"
	"net/http"
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
	once               sync.Once    // sync so we only setup 1 client
	netClient          *http.Client // http client
)

func NewNetClient() *http.Client {
	once.Do(func() {
		// transport configuratin
		var netTransport = &http.Transport{
			Proxy:        proxy.TorProxy(),   // default - rotating IP addresses
			MaxIdleConns: maxIdleConnections, // max idle connections
			Dial: (&net.Dialer{ // Dialer
				Timeout: 20 * time.Second, // max dialer timeout
			}).Dial,
			TLSHandshakeTimeout: 20 * time.Second, // transport layer security max timeout
		}
		netClient = &http.Client{
			Timeout:   time.Second * 20, // roundtripper timeout
			Transport: netTransport,     // how our HTTP requests are made
		}
	})

	return netClient
}
