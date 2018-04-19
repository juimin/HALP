package boards

import (
	"fmt"
	"testing"

	mgo "gopkg.in/mgo.v2"
)

func TestNewMongoStore(t *testing.T) {

	cases := []struct {
		name           string
		host           string
		dbname         string
		colname        string
		expectedOutput error
	}{
		{
			name:           "Testing New Board",
			host:           "localhost",
			dbname:         "boards",
			colname:        "board",
			expectedOutput: nil,
		},
		{
			name:           "Null Session",
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
			} else {
				mongoStore = NewMongoStore(mongoSess, "boards", "board") // Sess, database, collection
			}
		}
		if mongoStore == nil {
			t.Errorf("nil mongostore")
		}
	}
}

func TestGetByID(t *testing.T) {

}

func TestGetByBoardName(t *testing.T) {

}
