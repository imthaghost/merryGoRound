package main

import (
	"fmt"
	"io/ioutil"
	"log"

	ht "merryGoRound/pkg/http"
	"net/http"
	"time"
)

func main() {
	// configuration for tor client
	tor := ht.Tor {
		MaxTimeout: 20 * time.Second,
		MaxIdleConnections: 10,
	}

	// new instance of tor client
	torClient := tor.New()

	// check current requesting IP
	torIP1, err := getIP(torClient)
	if err != nil {
		fmt.Println(err)
	}

	// give the tor Client a new IP
	tor.NewIP()
	torIP2, err := getIP(torClient)
	if err != nil {
		fmt.Println(err)
	}

	// check
	if torIP1 != torIP2 {
		fmt.Println("Success")
	}

	// use stdlib http client
	standardClient := &http.Client{}

	// check current requesting IP
	_, err = getIP(standardClient)
	if err != nil {
		fmt.Println(err)
	}
}

// getIP returns the IP address of the incoming HTTP request
func getIP(client *http.Client) (string, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		"https://checkip.amazonaws.com",
		nil,
	)
	if err != nil {
		return "", err
	}
	// set a normal/non-hackerman user agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return  "", err
	}
	ip := string(body)

	log.Printf("IP: %s", ip)
	return ip, nil
}