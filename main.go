package main

import (
	"fmt"
	"io/ioutil"
	ht "merryGoRound/http"
	"net/http"
)

func main() {
	// create an instance of a new network client
	client := ht.NewNetClient()
	// check current requesting IP
	currentIP, err := checkIP(client)
	if err != nil {
		fmt.Errorf("Error: %v", err)

	}

	fmt.Printf("Current IP: %s", currentIP)

}

// checkIP will check what is the requesting IP address via an AWS server
func checkIP(client *http.Client) (string, error) {
	url := "https://checkip.amazonaws.com"
	res, err := client.Get(url)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	ip := string(body)
	return ip, nil
}
