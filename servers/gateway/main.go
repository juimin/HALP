package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/JuiMin/HALP/servers/gateway/indexes"
	"github.com/JuiMin/HALP/servers/gateway/models/boards"
	"github.com/JuiMin/HALP/servers/gateway/models/posts"

	"github.com/JuiMin/HALP/servers/gateway/handlers"
	"github.com/JuiMin/HALP/servers/gateway/models/comments"
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

func generateContextHandler() (*handlers.ContextReceiver, string, string, string, error) {
	// Check if the port is set
	port, err := getEnvVariable("ADDR", ":443", "Port Variable Not Set")

	// Get the TLS Cert and TLS Key from the environment variables
	tlskey, err := getEnvVariable("TLSKEY", "", "TLS Key not Set")

	if err != nil {
		fmt.Printf("Problem Encountered getting Environment Variable %s =: %v", "TLSKEY", err)
		return nil, "", "", "", err
	}

	tlscert, err := getEnvVariable("TLSCERT", "", "TLS Cert Not Set")

	if err != nil {
		fmt.Printf("Problem Encountered getting Environment Variable %s =: %v", "TLSCERT", err)
		return nil, "", "", "", err
	}

	// Connection to the Session Store
	redisAddr, err := getEnvVariable("REDISADDR", "localhost:6379", "Redis Address Not Set")

	// Connection to the Session Store
	mongoAddr, err := getEnvVariable("DBADDR", "localhost:27017", "Mongo Address Not Set")

	// Ge tthe variable for the session key
	sessionKey, err := getEnvVariable("SESSIONKEY", "", "Session Key Not Set")

	if err != nil {
		fmt.Printf("Problem Encountered getting Environment Variable %s =: %v", "SESSIONKEY", err)
		return nil, "", "", "", err
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

	// Data Store
	userStore := users.NewMongoStore(mongoSession, "users", "user")
	commentStore := comments.NewMongoStore(mongoSession, "comments", "comment")
	boardStore := boards.NewMongoStore(mongoSession, "boards", "board")
	postStore := posts.NewMongoStore(mongoSession, "posts", "post")

	fmt.Printf("Mongodb Online...\n")

	// Search Tries
	userTrie := indexes.NewSearchTrie()
	commentTrie := indexes.NewSearchTrie()
	boardTrie := indexes.NewSearchTrie()
	postTrie := indexes.NewSearchTrie()

	// IMPORT DATA FROM THE DATABASE FOR EACH TRIE
	users, err := userStore.GetAll()
	if err == nil {
		for _, u := range users {
			// Insert the keys into the trie by building+room
			err := userTrie.Insert(u.UserName, u.ID, 0)
			if err != nil {
				log.Printf("Failed to insert %s for ID: %v, Error: %v", u.UserName, u.ID, err)
			}
		}
	} else {
		fmt.Printf("%v", err == nil)
	}

	boards, err := boardStore.GetAllBoards()
	if err == nil {
		for _, b := range boards {
			// Insert the keys into the trie by building+room
			err := boardTrie.Insert(b.Title, b.ID, 0)
			if err != nil {
				log.Printf("Failed to insert %s for ID: %v, Error: %v", b.Title, b.ID, err)
			}
		}
	} else {
		fmt.Printf("%v\n", err)
	}

	posts, err := postStore.GetAll()
	if err == nil {
		for _, p := range posts {
			// Insert the keys into the trie by building+room
			err := postTrie.Insert(p.Title, p.ID, 0)
			if err != nil {
				log.Printf("Failed to insert %s for ID: %v, Error: %v", p.Title, p.ID, err)
			}
		}
	} else {
		fmt.Printf("%v\n", err)
	}

	// Build the CR
	cr, err := handlers.NewContextReceiver(
		sessionKey,
		userStore,
		redisStore,
		commentStore,
		postStore,
		boardStore,
		userTrie,
		commentTrie,
		boardTrie,
		postTrie)

	return cr, port, tlscert, tlskey, err
}

func generateMux(cr *handlers.ContextReceiver) *handlers.CORSHandler {
	// Create a new mux to start the server
	mux := http.NewServeMux()

	// Default Root handling
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/users", cr.UsersHandler)
	mux.HandleFunc("/users/me", cr.UsersMeHandler)
	mux.HandleFunc("/sessions", cr.SessionsHandler)
	mux.HandleFunc("/sessions/mine", cr.SessionsMineHandler)
	mux.HandleFunc("/posts/new", cr.NewPostHandler)
	mux.HandleFunc("/posts/update", cr.UpdatePostHandler)
	mux.HandleFunc("/posts/get", cr.GetPostHandler)
	mux.HandleFunc("/posts/get/board", cr.GetPostByBoardHandler)
	mux.HandleFunc("/posts/get/author", cr.GetPostByAuthorHandler)
	mux.HandleFunc("/posts/get/recent", cr.GetLastNHandler)
	mux.HandleFunc("/boards", cr.BoardsAllHandler)
	mux.HandleFunc("/boards/single", cr.SingleBoardHandler)
	mux.HandleFunc("/boards/updatepost", cr.UpdatePostCountHandler)
	mux.HandleFunc("/boards/updatesubscriber", cr.UpdateSubscriberCountHandler)
	mux.HandleFunc("/boards/createboard", cr.CreateBoardHandler)
	mux.HandleFunc("/bookmarks", cr.BookmarksHandler)
	mux.HandleFunc("/favorites", cr.FavoritesHandler)
	mux.HandleFunc("/search", cr.SearchHandler)
	// CORS Handling
	// This takes over for the mux after it has done everything the server needs
	corsHandler := handlers.NewCORSHandler(mux)
	return corsHandler
}

func main() {
	cr, port, tlscert, tlskey, err := generateContextHandler()

	if err != nil {
		fmt.Printf("Could not generate the Context Handler: %v", err)
		os.Exit(1)
	}

	corsHandler := generateMux(cr)
	fmt.Println("CORS Mounted Successfully...")

	// Notify that the server is started
	fmt.Printf("Server started on port %s\n", port)

	// Start the listener with TLS, logging when errors occur
	log.Fatal(http.ListenAndServeTLS(port, tlscert, tlskey, corsHandler))
}
