package handlers

import (
	"fmt"
	"testing"

	"github.com/JuiMin/HALP/servers/gateway/models/users"
	mgo "gopkg.in/mgo.v2"
)

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
		_, err = NewContextReceiver(c.key, mongoStore)

		if err != nil {
			actualErr = err.Error()
		}

		// Check if the error is the same
		if actualErr != expectedErr {
			t.Errorf("%s Failed: Expected %s but got %s", c.name, expectedErr, actualErr)
		}
	}
}
