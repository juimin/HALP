package handlers

// Construct a test CR - These should only affect the travis
import (
	"fmt"
	"os"
	"time"

	"github.com/JuiMin/HALP/servers/gateway/models/posts"

	"github.com/JuiMin/HALP/servers/gateway/models/boards"
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

// Generates a Test CR for testing
func prepTestCR() *ContextReceiver {
	// Connection to the Session Store
	redisAddr, err := getEnvVariable("REDISADDR", "localhost:6379", "Redis Address Not Set")

	if err != nil {
		fmt.Printf("Problem Encountered getting Environment Variable %s =: %v", "TLSCERT", err)
		os.Exit(1)
	}

	// Connection to the Session Store
	mongoAddr, err := getEnvVariable("DBADDR", "localhost:27017", "Mongo Address Not Set")

	if err != nil {
		fmt.Printf("Problem Encountered getting Environment Variable %s =: %v", "TLSCERT", err)
		os.Exit(1)
	}

	// Ge tthe variable for the session key
	sessionKey, err := getEnvVariable("SESSIONKEY", "Potato", "Session Key Not Set")

	if err != nil {
		fmt.Printf("Problem Encountered getting Environment Variable %s =: %v", "SESSIONKEY", err)
		os.Exit(1)
	}

	// Prepare the redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       3,  // Use the non default database
	})

	redisStore := sessions.NewRedisStore(redisClient, time.Minute*30)

	// Dial the mongo Server
	mongoSession, err := mgo.Dial(mongoAddr)
	// Check if there was an error dialing the mongo server
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	mongoStore := users.NewMongoStore(mongoSession, "test", "users")

	commentStore := comments.NewMongoStore(mongoSession, "test", "comments")
	boardStore := boards.NewMongoStore(mongoSession, "test", "board")
	postStore := posts.NewMongoStore(mongoSession, "test", "post")

	cr, err := NewContextReceiver(sessionKey, mongoStore, redisStore, commentStore, postStore, boardStore, nil, nil, nil, nil)

	return cr
}

// Broken Mongo Returns a CR with a bad mongo store connection
func BrokenMongoCR() *ContextReceiver {
	// Connection to the Session Store
	redisAddr, err := getEnvVariable("REDISADDR", "localhost:6379", "Redis Address Not Set")

	if err != nil {
		fmt.Printf("Problem Encountered getting Environment Variable %s =: %v", "TLSCERT", err)
		os.Exit(1)
	}

	// Ge tthe variable for the session key
	sessionKey, err := getEnvVariable("SESSIONKEY", "Potato", "Session Key Not Set")

	if err != nil {
		fmt.Printf("Problem Encountered getting Environment Variable %s =: %v", "SESSIONKEY", err)
		os.Exit(1)
	}

	// Prepare the redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       3,  // Use the non default database
	})

	redisStore := sessions.NewRedisStore(redisClient, time.Minute*30)

	// Dial the mongo Server
	mongoSession, err := mgo.Dial("localhost:2222")
	// Check if there was an error dialing the mongo server
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	mongoStore := users.NewMongoStore(mongoSession, "test", "users")

	commentStore := comments.NewMongoStore(mongoSession, "test", "comments")
	boardStore := boards.NewMongoStore(mongoSession, "test", "board")
	postStore := posts.NewMongoStore(mongoSession, "test", "post")

	cr, err := NewContextReceiver(sessionKey, mongoStore, redisStore, commentStore, postStore, boardStore, nil, nil, nil, nil)

	return cr
}

// Generates a CR with broken Redis Connection
func BrokenRedisCR() *ContextReceiver {
	// Connection to the Session Store
	mongoAddr, err := getEnvVariable("DBADDR", "localhost:27017", "Mongo Address Not Set")

	if err != nil {
		fmt.Printf("Problem Encountered getting Environment Variable %s =: %v", "TLSCERT", err)
		os.Exit(1)
	}

	// Ge tthe variable for the session key
	sessionKey, err := getEnvVariable("SESSIONKEY", "Potato", "Session Key Not Set")

	if err != nil {
		fmt.Printf("Problem Encountered getting Environment Variable %s =: %v", "SESSIONKEY", err)
		os.Exit(1)
	}

	// Prepare the redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:2222",
		Password: "", // no password set
		DB:       3,  // Use the non default database
	})

	redisStore := sessions.NewRedisStore(redisClient, time.Minute*30)

	// Dial the mongo Server
	mongoSession, err := mgo.Dial(mongoAddr)
	// Check if there was an error dialing the mongo server
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	mongoStore := users.NewMongoStore(mongoSession, "test", "users")

	commentStore := comments.NewMongoStore(mongoSession, "test", "comments")
	boardStore := boards.NewMongoStore(mongoSession, "test", "board")
	postStore := posts.NewMongoStore(mongoSession, "test", "post")

	cr, err := NewContextReceiver(sessionKey, mongoStore, redisStore, commentStore, postStore, boardStore, nil, nil, nil, nil)

	return cr
}
