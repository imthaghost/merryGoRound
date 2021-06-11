package main

import (
	"fmt"
	"io/ioutil"
	mgr "merryGoRound"

	"net/http"
)

func main() {
	// use Tor as our proxy
	tor := mgr.TorProxy()
	// create an instance of a new network client
	torClient := mgr.NewClient(tor)
	// check current requesting IP
	torIP, err := checkIP(torClient)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Tor IP: %s", torIP)

	// use stdlib http client
	standardClient := &http.Client{}
	// check current requesting IP
	standardIP, err := checkIP(standardClient)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Non proxy IP: %s", standardIP)

	if standardIP != torIP {
		fmt.Println("Success")
	}
}

// checkIP returns the IP address of the incoming HTTP request
func checkIP(client *http.Client) (string, error) {
	awsURL := "https://checkip.amazonaws.com"
	res, err := client.Get(awsURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	ip := string(body)

	return ip, nil
}
