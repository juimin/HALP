package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/JuiMin/HALP/servers/gateway/models/sessions"
	"github.com/JuiMin/HALP/servers/gateway/models/users"
	"github.com/go-redis/redis"
	mgo "gopkg.in/mgo.v2"
)

func TestRootHandler(t *testing.T) {

	cases := []struct {
		name                string
		expectedContentType string
		expectedStatus      int
		expectedOutput      string
	}{
		{
			name:                "Passing Test",
			expectedContentType: contentTypeText,
			expectedStatus:      http.StatusOK,
			expectedOutput:      "Welcome to the gateway! There is no resource here right now!",
		},
	}

	for _, c := range cases {
		// Create http request for the root path
		req, err := http.NewRequest("GET", "/", nil)
		// Fatal error report for test
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()

		// Serve the handler
		handler := http.HandlerFunc(RootHandler)
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("%s Failed: Testing Status Code: expected %v but got %v",
				c.name, c.expectedStatus, status)
		}

		if header := rr.Header().Get(headerContentType); header != c.expectedContentType {
			t.Errorf("%s Failed: Testing header, expected %s but got %s",
				c.name, c.expectedContentType, header)
		}

		// Check the response body is what we expect.
		if body := rr.Body.String(); body != c.expectedOutput {
			t.Errorf("%s Failed: Testing response body: expected %s but got %s",
				c.name, c.expectedOutput, body)
		}
	}
}

func TestCORSHandler(t *testing.T) {
	cases := []struct {
		name           string
		expHeaders     []string
		expHeaderVals  []string
		origin         string
		expectedStatus int
		reqType        string
		expectedOutput string
	}{
		{
			name: "Evil Origin Test",
			expHeaders: []string{
				accessControlAllowOrigin,
				accessControlAllowMethods,
				exposeHeaders,
				accessControlAllowAge,
				allowHeaders,
			},
			expHeaderVals: []string{
				accessControlValue,
				accessControlMethods,
				exposedHeaders,
				accessControlAge,
				allowedHeaders,
			},
			origin:         "http://evil.com",
			reqType:        "GET",
			expectedStatus: http.StatusUnauthorized,
			expectedOutput: "Sorry, bad request blocked\n",
		},
		{
			name: "Option Header Test",
			expHeaders: []string{
				accessControlAllowOrigin,
				accessControlAllowMethods,
				exposeHeaders,
				accessControlAllowAge,
				allowHeaders,
			},
			expHeaderVals: []string{
				accessControlValue,
				accessControlMethods,
				exposedHeaders,
				accessControlAge,
				allowedHeaders,
			},
			origin:         "",
			reqType:        "OPTIONS",
			expectedStatus: http.StatusOK,
			expectedOutput: "",
		},
		{
			name: "Nothing Special Get Request Test",
			expHeaders: []string{
				accessControlAllowOrigin,
				accessControlAllowMethods,
				exposeHeaders,
				accessControlAllowAge,
				allowHeaders,
			},
			expHeaderVals: []string{
				accessControlValue,
				accessControlMethods,
				exposedHeaders,
				accessControlAge,
				allowedHeaders,
			},
			origin:         "",
			reqType:        "GET",
			expectedStatus: http.StatusNotFound,
			expectedOutput: "404 page not found\n",
		},
	}

	for _, c := range cases {
		// Create http request for the root path
		req, err := http.NewRequest(c.reqType, "/", nil)
		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()

		// Fatal error report for test if we could not gen a req
		if err != nil {
			t.Fatal(err)
		}

		// Set origin if exist
		if len(c.origin) != 0 {
			req.Header.Add("Origin", c.origin)
		}

		// Serve the handler
		cors := NewCORSHandler(http.NewServeMux())
		cors.ServeHTTP(rr, req)

		// Test Header Results
		if rr.Code == c.expectedStatus {
			// We only care about these headers if we set them
			if c.expectedStatus == http.StatusOK {
				// Check all the headers we set
				for i, header := range c.expHeaders {
					if headerVal := rr.Header().Get(header); headerVal != c.expHeaderVals[i] {
						t.Errorf("%s Failed Header Test: Expected Header %s to be %s but got %s",
							c.name, header, c.expHeaderVals[i], headerVal)
					}
				}
			}

			// Check Content Body
			if body := rr.Body.String(); c.expectedOutput != body {
				t.Errorf("%s Failed Response Body Test: Expected %s but got %s",
					c.name, c.expectedOutput, body)
			}
		} else {
			t.Errorf("%s Failed Response Body Test: Expected %v but got %v",
				c.name, c.expectedStatus, rr.Code)
		}

	}
}

func TestContextHandler(t *testing.T) {

	cases := []struct {
		name           string
		key            string
		expectedOutput error
	}{
		{
			name:           "Passing Test",
			key:            "potato",
			expectedOutput: nil,
		},
		{
			name:           "No Key Test",
			key:            "",
			expectedOutput: fmt.Errorf("No key set for signing key"),
		},
	}

	for _, c := range cases {
		expectedErr := ""
		actualErr := ""
		if c.expectedOutput != nil {
			expectedErr = c.expectedOutput.Error()
		}

		// Predefine a mongo store for all tests
		mongoSession, err := mgo.Dial("localhost")
		if err != nil {
			t.Errorf("Error Connecting to MongoDB. Cannot perform Insertion Tests")
		}

		mongoStore := users.NewMongoStore(mongoSession, "test_users", "user")

		// Prepare the redis client
		redisClient := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", //no password set
			DB:       1,  //use default DB
		})

		redisStore := sessions.NewRedisStore(redisClient, time.Minute*30)

		_, err = NewContextReceiver(c.key, mongoStore, redisStore)

		if err != nil {
			actualErr = err.Error()
		}

		// Check if the error is the same
		if actualErr != expectedErr {
			t.Errorf("%s Failed: Expected %s but got %s", c.name, expectedErr, actualErr)
		}
	}
}
