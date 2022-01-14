package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	var project = flag.String("ams-project", "", "AMS project to be used (required)")
	var sub = flag.String("ams-sub", "", "AMS subscription to pull messages from (required)")
	var host = flag.String("ams-host", "", "AMS host and port (host:port) (required)")
	var token = flag.String("ams-token", "", "AMS access token (required)")
	var pollRate = flag.Int64("pull-interval", 300, "Interval in milliseconds between pulling the next message")
	var remoteEndpoint = flag.String("remote-endpoint", "", "Remote endpoint url to push messages to (required)")
	var remoteEndpointAuthHeader = flag.String("auth-header", "", "Remote endpoint expected authorization header")

	flag.Parse()

	requiredParams := []string{"ams-host", "ams-project", "ams-sub", "ams-token", "remote-endpoint"}

	for _, item := range requiredParams {
		if flag.Lookup(item).Value.String() == "" {
			fmt.Printf("required flag not set: -%v \n", item)
			flag.Usage()
			os.Exit(1)
		}
	}

	// build the client and execute the request
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transCfg, Timeout: time.Duration(30 * time.Second)}
	ams := AMSClient{Endpoint: *host, Project: *project, Token: *token, Client: client}

	for {
		err := ams.Push(*sub, *remoteEndpoint, *remoteEndpointAuthHeader)

		if err != nil {
			log.Printf("ERROR: %v", err)
		}

		time.Sleep(time.Duration(*pollRate) * time.Millisecond)
	}

}
