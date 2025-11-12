package tech

import (
	"log"
	"time"
)

// This func is here just for demostration of what you can do with pulumi
// You can do HTTP requests while deploying your infrastructure, and make it more flexible
func FetchIpAllowList(url string) ([]string, error) {
	log.Printf("Fetching IP allow list for URL: %s\n", url)

	time.Sleep(time.Second)

	// Let's imagine that instead of sleeping, the program connects to your server,
	// receives data from there and converts it to the desired format.
	// Something like the example below

	/*
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		var ips []string
		json.Unmarshal(data, &ips)
		return ips, nil
	*/

	return []string{"192.168.1.2/32", "192.168.1.3/32"}, nil
}
