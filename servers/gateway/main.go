package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/JuiMin/HALP/servers/gateway/handlers"
	"github.com/go-redis/redis"
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

	// Connection to the Session Store
	redisAddr := os.Getenv("REDISADDR")

	// Check redis addr being set in the terminal
	if len(redisAddr) == 0 {
		// Default to localhost
		redisAddr = "localhost:6379"
	}

	// Prepare the redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", //no password set
		DB:       0,  //use default DB
	})

	// temp usage of redisClient so the build doesn't die
	if redisClient != nil {
		fmt.Println("The client is there lol")
	}

	// Create a new redis store
	// Set the session store to be a redis store with the given time duration
	// redisSess := sessions.NewRedisStore(redisClient, time.Duration(time.Second)*time.Second)

	// Create a new mux to start the server
	mux := http.NewServeMux()

	// TODO: DEFINE HANDLERS

	// Default Root handling
	mux.HandleFunc("/", handlers.RootHandler)

	// CORS Handling
	// This takes over for the mux after it has done everything the server needs
	corsHandler := handlers.NewCORSHandler(mux)
	fmt.Println("CORS Mounted Successfully")

	// Notify that the server is started
	fmt.Printf("Server started on port %s\n", port)

	// Start the listener with TLS, logging when errors occur
	log.Fatal(http.ListenAndServeTLS(port, tlscert, tlskey, corsHandler))
}
