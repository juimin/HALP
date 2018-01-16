package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the port addressd that we want the gateway to run on
	// This should be set to the envirnoment variable in the container we run this on
	port := os.Getenv("ADDR")

	// Check if the port is set
	// If it is not set, default the port to be the 443 Https ENABLED port
	if len(port) == 0 {
		port = ":443"
	}

	// Get the TLS Cert and TLS Key from the environment variables
	tlskey := os.Getenv("TLSKEY")
	tlscert := os.Getenv("TLSCERT")

	// Check that both the TLS Cert and the TLS Key are available
	if len(tlskey) == 0 || len(tlscert) == 0 {
		fmt.Println("TLS Key or TLS Cert not set")
		// Exit with a non zero exit code
		os.Exit(1)
	}
}
