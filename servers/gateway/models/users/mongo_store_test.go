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
			} else {
				mongoStore = NewMongoStore(mongoSess, "users", "user") // Sess, database, collection
			}
		}
		if mongoStore == nil {
			t.Errorf("nil mongostore")
		}
	}
}

// Test Insertion of User into the database
func TestInsert(t *testing.T) {

	cases := []struct {
		name           string
		email          string
		password       string
		passwordConf   string
		username       string
		firstname      string
		lastname       string
		occupation     string
		expectedOutput error
	}{
		{
			name:           "Valid New User",
			email:          "test@test.com",
			password:       "potato123",
			passwordConf:   "potato123",
			username:       "test23",
			firstname:      "testname",
			lastname:       "testlastname",
			occupation:     "Welder",
			expectedOutput: nil,
		},
		{
			name:           "Invalid User Error",
			email:          "",
			password:       "potato123",
			passwordConf:   "potato123",
			username:       "test23",
			firstname:      "testname",
			lastname:       "testlastname",
			occupation:     "Welder",
			expectedOutput: fmt.Errorf("Email does not exist"),
		},
	}

	// Predefine a mongo store for all tests
	mongoSession, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("Error Connecting to MongoDB. Cannot perform Insertion Tests")
	}
	mongoStore := NewMongoStore(mongoSession, "test_users", "user")

	for _, c := range cases {
		// Define a new user by the input
		newUser := &NewUser{
			Email:        c.email,
			Password:     c.password,
			PasswordConf: c.passwordConf,
			UserName:     c.username,
			FirstName:    c.firstname,
			LastName:     c.lastname,
			Occupation:   c.occupation,
		}

		_, err := mongoStore.Insert(newUser)
		errText := ""
		expected := ""
		if err != nil {
			errText = err.Error()
		}
		if c.expectedOutput != nil {
			expected = c.expectedOutput.Error()
		}
		if expected != errText {
			t.Errorf("%s Failed: Expected %s but got %s", c.name, c.expectedOutput, err)
		}
	}
}

func TestGetByEmail(t *testing.T) {

	cases := []struct {
		name           string
		email          string
		expectedOutput error
	}{
		{
			name:           "Valid User Request - Known User",
			email:          "test@test.com",
			expectedOutput: nil,
		},
		{
			name:           "Valid User Request - User Doesn't Exist",
			email:          "test@gg.com",
			expectedOutput: fmt.Errorf("error getting users by email: %v", "not found"),
		},
		{
			name:           "Valid User Request - No Search  Input",
			email:          "",
			expectedOutput: fmt.Errorf("error getting users by email: %v", "not found"),
		},
	}

	// Predefine a mongo store for all tests
	mongoSession, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("Error Connecting to MongoDB. Cannot perform Insertion Tests")
	}
	mongoStore := NewMongoStore(mongoSession, "test_users", "user")

	for _, c := range cases {
		_, err := mongoStore.GetByEmail(c.email)
		errText := ""
		expected := ""
		if err != nil {
			errText = err.Error()
		}
		if c.expectedOutput != nil {
			expected = c.expectedOutput.Error()
		}
		if expected != errText {
			t.Errorf("%s Failed: Expected %v but got %v", c.name, expected, errText)
		}
	}
}

func TestGetByUserName(t *testing.T) {

	cases := []struct {
		name           string
		username       string
		expectedOutput error
	}{
		{
			name:           "Valid User Request - Known User",
			username:       "test23",
			expectedOutput: nil,
		},
		{
			name:           "Valid User Request - User Doesn't Exist",
			username:       "tast",
			expectedOutput: fmt.Errorf("error getting users by username: %v", "not found"),
		},
		{
			name:           "Valid User Request - No Search  Input",
			username:       "",
			expectedOutput: fmt.Errorf("error getting users by username: %v", "not found"),
		},
	}

	// Predefine a mongo store for all tests
	mongoSession, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("Error Connecting to MongoDB. Cannot perform Insertion Tests")
	}
	mongoStore := NewMongoStore(mongoSession, "test_users", "user")

	for _, c := range cases {
		_, err := mongoStore.GetByUserName(c.username)
		errText := ""
		expected := ""
		if err != nil {
			errText = err.Error()
		}
		if c.expectedOutput != nil {
			expected = c.expectedOutput.Error()
		}
		if expected != errText {
			t.Errorf("%s Failed: Expected %v but got %v", c.name, expected, errText)
		}
	}
}
