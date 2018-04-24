package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/JuiMin/HALP/servers/gateway/models/boards"

	"github.com/JuiMin/HALP/servers/gateway/handlers"
	"github.com/JuiMin/HALP/servers/gateway/models/sessions"
	"github.com/JuiMin/HALP/servers/gateway/models/users"
	"github.com/go-redis/redis"
	mgo "gopkg.in/mgo.v2"
)

// getEnvVariable takes in an environment variable as a string
// and checks if the variable is set, if it is not set, return the defaul
// If the error message is set, display and exit since these are vital
func getEnvVariable(name string, defaultValue string, errorMessage string) (string, error) {
	envVariable := os.Getenv(name)
	if len(envVariable) == 0 {
		// Check if a default is set
		if len(defaultValue) != 0 {
			return defaultValue, nil
		}
		return "", fmt.Errorf(errorMessage)
	}
	return envVariable, nil
}

func main() {
	// Check if the port is set
	port, err := getEnvVariable("ADDR", ":443", "Port Variable Not Set")

	// If it is not set, default the port to be the 443 Https ENABLED port
	if err != nil {
		fmt.Printf("Problem Encountered getting Environment Variable %s =: %v", "ADDR", err)
		os.Exit(1)
	}

	// Get the TLS Cert and TLS Key from the environment variables
	tlskey, err := getEnvVariable("TLSKEY", "", "TLS Key not Set")

	if err != nil {
		fmt.Printf("Problem Encountered getting Environment Variable %s =: %v", "TLSKEY", err)
		os.Exit(1)
	}

	tlscert, err := getEnvVariable("TLSCERT", "", "TLS Cert Not Set")

	if err != nil {
		fmt.Printf("Problem Encountered getting Environment Variable %s =: %v", "TLSCERT", err)
		os.Exit(1)
	}

	// Connection to the Session Store
	redisAddr, err := getEnvVariable("REDISADDR", "localhost:6379", "Redis Address Not Set")

	if err != nil {
		fmt.Printf("Problem Encountered getting Environment Variable %s =: %v", "REDISADDR", err)
		os.Exit(1)
	}

	// Connection to the Session Store
	mongoAddr, err := getEnvVariable("DBADDR", "localhost:27017", "Mongo Address Not Set")

	if err != nil {
		fmt.Printf("Problem Encountered getting Environment Variable %s =: %v", "DBADDR", err)
		os.Exit(1)
	}

	// Ge tthe variable for the session key
	sessionKey, err := getEnvVariable("SESSIONKEY", "", "Session Key Not Set")

	if err != nil {
		fmt.Printf("Problem Encountered getting Environment Variable %s =: %v", "SESSIONKEY", err)
		os.Exit(1)
	}

	// Prepare the redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", //no password set
		DB:       0,  //use default DB
	})

	redisStore := sessions.NewRedisStore(redisClient, time.Minute*30)

	fmt.Printf("Redis Online...\n")

	// Dial the mongo Server
	mongoSession, err := mgo.Dial(mongoAddr)
	// Check if there was an error dialing the mongo server
	if err != nil {
		fmt.Println("Mongo " + err.Error())
		os.Exit(1)
	}

	mongoStore := users.NewMongoStore(mongoSession, "users", "user")

	boardStore := boards.NewMongoStore(mongoSession, "boards", "board")

	fmt.Printf("Mongodb Online...\n")

	cr, err := handlers.NewContextReceiver(sessionKey, mongoStore, redisStore, boardStore)

	// Create a new mux to start the server
	mux := http.NewServeMux()

	// TODO: DEFINE HANDLERS

	// Default Root handling
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/users", cr.UsersHandler)
	mux.HandleFunc("/users/me", cr.UsersMeHandler)
	mux.HandleFunc("/sessions", cr.SessionsHandler)
	mux.HandleFunc("/sessions/mine", cr.SessionsMineHandler)
	mux.HandleFunc("/boards", cr.BoardsAllHandler)
	mux.HandleFunc("/boards/single", cr.SingleBoardHandler)
	mux.HandleFunc("/boards/updatepost", cr.UpdatePostHandler)
	mux.HandleFunc("/boards/updatesubscriber", cr.UpdateSubscriberHandler)
	mux.HandleFunc("/bookmarks", cr.BookmarksHandler)
	mux.HandleFunc("/favorites", cr.FavoritesHandler)

	// CORS Handling
	// This takes over for the mux after it has done everything the server needs
	corsHandler := handlers.NewCORSHandler(mux)

	fmt.Println("CORS Mounted Successfully...")

	// Notify that the server is started
	fmt.Printf("Server started on port %s\n", port)

	// Start the listener with TLS, logging when errors occur
	log.Fatal(http.ListenAndServeTLS(port, tlscert, tlskey, corsHandler))
}
