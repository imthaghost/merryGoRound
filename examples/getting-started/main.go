package main

import (
	"fmt"
	"io/ioutil"
	mgr "merryGoRound"

	"net/http"
)

func main() {
	// create an instance of a new network client
	client := mgr.NewClient()
	// check current requesting IP
	currentIP, err := checkIP(client)
	if err != nil {
		fmt.Println(err)

	}

	fmt.Printf("Current IP: %s", currentIP)

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
