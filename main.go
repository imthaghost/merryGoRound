package main

import (
	"fmt"
	"io/ioutil"
	"merryGoRound/http"
)

func main() {

	// check ip with aws
	url := "https://checkip.amazonaws.com"
	res, err := http.NewNetClient().Get(url)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	ip := string(body)
	fmt.Println(ip)
}
