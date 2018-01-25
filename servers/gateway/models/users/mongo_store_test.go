package users

import (
	"fmt"
	"testing"

	mgo "gopkg.in/mgo.v2"
)

//TODO: add tests for the various functions in user.go, as described in the assignment.
//use `go test -cover` to ensure that you are covering all or nearly all of your code paths.

func TestNewMongoStore(t *testing.T) {

	cases := []struct {
		name           string
		host           string
		dbname         string
		colname        string
		expectedOutput error
	}{
		{
			name:           "Test New User Mem Store Constructor",
			host:           "localhost",
			dbname:         "users",
			colname:        "user",
			expectedOutput: nil,
		},
		{
			name:           "Null session",
			host:           "",
			dbname:         "",
			colname:        "",
			expectedOutput: fmt.Errorf("session null"),
		},
	}

	for _, c := range cases {
		mongoStore := &MongoStore{}
		if len(c.host) != 0 {
			mongoSess, err := mgo.Dial("localhost")
			if err != nil {
				t.Errorf("Error on %s: Expected %s but got %s", c.name, c.expectedOutput, err.Error())
			}
			mongoStore = NewMongoStore(mongoSess, "users", "user")
		}
		if mongoStore == nil {
			t.Errorf("nil mongostore")
		}
	}
}
