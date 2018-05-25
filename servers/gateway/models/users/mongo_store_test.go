package users

import (
	"fmt"
	"testing"

	"gopkg.in/mgo.v2/bson"

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

func TestGetByID(t *testing.T) {

	testUser := &NewUser{
		Email:        "potato@gg.com",
		Password:     "something",
		PasswordConf: "something",
		UserName:     "user",
		FirstName:    "pot",
		LastName:     "tato",
		Occupation:   "Spud",
	}

	// Predefine a mongo store for all tests
	mongoSession, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("Error Connecting to MongoDB. Cannot perform Insertion Tests")
	}
	mongoStore := NewMongoStore(mongoSession, "test_users", "user")

	usr, err := mongoStore.Insert(testUser)

	if err != nil {
		t.Errorf("Error putting the user into the database")
	}

	cases := []struct {
		name           string
		id             bson.ObjectId
		expectedOutput error
	}{
		{
			name:           "Valid User Request - Known User",
			id:             usr.ID,
			expectedOutput: nil,
		},
		{
			name:           "Valid User Request - User Doesn't Exist",
			id:             bson.NewObjectId(),
			expectedOutput: fmt.Errorf("error getting users by id: %v", "not found"),
		},
	}

	for _, c := range cases {
		_, err = mongoStore.GetByID(c.id)
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

func TestDelete(t *testing.T) {

	testUser := &NewUser{
		Email:        "potato@gg.com",
		Password:     "something",
		PasswordConf: "something",
		UserName:     "user",
		FirstName:    "pot",
		LastName:     "tato",
		Occupation:   "Spud",
	}

	// Predefine a mongo store for all tests
	mongoSession, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("Error Connecting to MongoDB. Cannot perform Insertion Tests")
	}
	mongoStore := NewMongoStore(mongoSession, "test_users", "user")

	usr, err := mongoStore.Insert(testUser)

	if err != nil {
		t.Errorf("Error putting the user into the database")
	}

	cases := []struct {
		name           string
		id             bson.ObjectId
		expectedOutput error
	}{
		{
			name:           "Valid User Request - Known User",
			id:             usr.ID,
			expectedOutput: nil,
		},
		{
			name:           "Valid User Request - User Doesn't Exist",
			id:             bson.NewObjectId(),
			expectedOutput: fmt.Errorf("%v", "not found"),
		},
	}

	for _, c := range cases {
		err = mongoStore.Delete(c.id)
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

// TestUserUpdate tests the updating of user information
func TestUserUpdate(t *testing.T) {

	testUser := &NewUser{
		Email:        "potato@gg.com",
		Password:     "something",
		PasswordConf: "something",
		UserName:     "user",
		FirstName:    "pot",
		LastName:     "tato",
		Occupation:   "Spud",
	}

	// Predefine a mongo store for all tests
	mongoSession, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("Error Connecting to MongoDB. Cannot perform Insertion Tests")
	}
	mongoStore := NewMongoStore(mongoSession, "test_users", "user")

	usr, err := mongoStore.Insert(testUser)

	if err != nil {
		t.Error("There was an issue adding the test user to the database")
	}

	cases := []struct {
		name           string
		id             bson.ObjectId
		email          string
		firstname      string
		lastname       string
		occupation     string
		expectedOutput error
	}{
		{
			name:           "Test Updating User with Valid Update",
			id:             usr.ID,
			email:          "newemail@gmail.com",
			firstname:      "newname",
			lastname:       "newlastname",
			occupation:     "newjob",
			expectedOutput: nil,
		},
		{
			name:           "Test User Not found",
			id:             bson.NewObjectId(),
			email:          "got nothing",
			firstname:      "got nothing",
			lastname:       "got nothing",
			occupation:     "got nothing",
			expectedOutput: fmt.Errorf("not found"),
		},
	}

	for _, c := range cases {
		err = mongoStore.UserUpdate(c.id, &UserUpdate{
			Email:      c.email,
			FirstName:  c.firstname,
			LastName:   c.lastname,
			Occupation: c.occupation,
		})
		// Test Error
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
		// Test Results
		usr, err := mongoStore.GetByID(c.id)

		if err == nil && usr != nil {
			if usr.Occupation != c.occupation {
				t.Errorf("%s Failed: Occupation Expected to be %s but found %s", c.name, c.occupation, usr.Occupation)
			}
			if usr.Email != c.email {
				t.Errorf("%s Failed: Email Expected to be %s but found %s", c.name, c.email, usr.Email)
			}
			if usr.FirstName != c.firstname {
				t.Errorf("%s Failed: First Name Expected to be %s but found %s", c.name, c.firstname, usr.FirstName)
			}
			if usr.LastName != c.lastname {
				t.Errorf("%s Failed: Last Name Expected to be %s but found %s", c.name, c.lastname, usr.LastName)
			}
		}
	}
}

func TestPassUpdate(t *testing.T) {

	testUser := &NewUser{
		Email:        "potato@gg.com",
		Password:     "something",
		PasswordConf: "something",
		UserName:     "user",
		FirstName:    "pot",
		LastName:     "tato",
		Occupation:   "Spud",
	}

	// Predefine a mongo store for all tests
	mongoSession, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("Error Connecting to MongoDB. Cannot perform Insertion Tests")
	}
	mongoStore := NewMongoStore(mongoSession, "test_users", "user")

	usr, err := mongoStore.Insert(testUser)

	if err != nil {
		t.Error("There was an issue adding the test user to the database")
	}

	cases := []struct {
		name           string
		id             bson.ObjectId
		newPass        string
		newPassConf    string
		expectedOutput error
	}{
		{
			name:           "Test Updating User with Valid Update",
			id:             usr.ID,
			newPass:        "potato2",
			newPassConf:    "potato2",
			expectedOutput: nil,
		},
		{
			name:           "Test User Not found",
			id:             bson.NewObjectId(),
			newPass:        "tomato",
			newPassConf:    "tomato",
			expectedOutput: fmt.Errorf("not found"),
		},
		{
			name:           "New Password does not match conf",
			id:             bson.NewObjectId(),
			newPass:        "facebook",
			newPassConf:    "google",
			expectedOutput: fmt.Errorf("Password and password conf do not match"),
		},
		{
			name:           "New Password is not given",
			id:             bson.NewObjectId(),
			newPass:        "",
			newPassConf:    "",
			expectedOutput: fmt.Errorf("Invalid Input: New Password cannot be length 0"),
		},
		{
			name:           "New Password Conf is not given",
			id:             bson.NewObjectId(),
			newPass:        "google",
			newPassConf:    "",
			expectedOutput: fmt.Errorf("Invalid Input: New Password cannot be length 0"),
		},
	}

	for _, c := range cases {
		err = mongoStore.PassUpdate(c.id, &PasswordUpdate{
			NewPassword:     c.newPass,
			NewPasswordConf: c.newPassConf,
		})
		// Test Error
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

// TestFavoritesUpdate tests the mongo storage updating of the favorites for a user
func TestFavoritesUpdate(t *testing.T) {

	testUser := &NewUser{
		Email:        "potato@gg.com",
		Password:     "something",
		PasswordConf: "something",
		UserName:     "user",
		FirstName:    "pot",
		LastName:     "tato",
		Occupation:   "Spud",
	}

	// Predefine a mongo store for all tests
	mongoSession, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("Error Connecting to MongoDB. Cannot perform Insertion Tests")
	}
	mongoStore := NewMongoStore(mongoSession, "test_users", "user")

	usr, err := mongoStore.Insert(testUser)

	if err != nil {
		t.Error("There was an issue adding the test user to the database")
	}

	filler := bson.NewObjectId()

	cases := []struct {
		name           string
		id             bson.ObjectId
		addition       *FavoritesUpdate
		expectedOutput error
	}{
		{
			name: "Test Updating User with Valid Update Addition",
			id:   usr.ID,
			addition: &FavoritesUpdate{
				Adding:   true,
				UpdateID: filler,
			},
			expectedOutput: nil,
		},
		{
			name: "Test Updating User That doesn't exist",
			id:   bson.NewObjectId(),
			addition: &FavoritesUpdate{
				Adding:   true,
				UpdateID: bson.NewObjectId(),
			},
			expectedOutput: fmt.Errorf("error getting users by id: not found"),
		},
		{
			name: "Test Updating User with Valid Update Remove",
			id:   usr.ID,
			addition: &FavoritesUpdate{
				Adding:   false,
				UpdateID: filler,
			},
			expectedOutput: nil,
		},
	}

	for _, c := range cases {
		updatedUser, err := mongoStore.FavoritesUpdate(c.id, c.addition)
		if err != nil && c.expectedOutput != nil {
			if err.Error() != c.expectedOutput.Error() {
				t.Errorf("Error on %s: Expected %v but got %v", c.name, c.expectedOutput, err)
			}
		}
		if !(err == nil && c.expectedOutput == nil) && (err == nil || c.expectedOutput == nil) {
			t.Errorf("Error on %s: Expected %v but got %v", c.name, c.expectedOutput, err)
		}
		if updatedUser != nil {
			if c.addition.Adding {
				if len(updatedUser.Favorites) != 1 {
					t.Errorf("Error on %s: Expected %v but got %v", c.name, 1, len(updatedUser.Favorites))
				}
			} else {
				if len(updatedUser.Favorites) != 0 {
					t.Errorf("Error on %s: Expected %v but got %v", c.name, 0, len(updatedUser.Favorites))
				}
			}
		}
	}
}

// TestBookmarksUpdate tests the mongo updating of user's bookmarks
func TestBookmarksUpdate(t *testing.T) {

	testUser := &NewUser{
		Email:        "potato@gg.com",
		Password:     "something",
		PasswordConf: "something",
		UserName:     "user",
		FirstName:    "pot",
		LastName:     "tato",
		Occupation:   "Spud",
	}

	// Predefine a mongo store for all tests
	mongoSession, err := mgo.Dial("localhost")
	if err != nil {
		t.Errorf("Error Connecting to MongoDB. Cannot perform Insertion Tests")
	}
	mongoStore := NewMongoStore(mongoSession, "test_users", "user")

	usr, err := mongoStore.Insert(testUser)

	if err != nil {
		t.Error("There was an issue adding the test user to the database")
	}

	filler := bson.NewObjectId()

	cases := []struct {
		name           string
		id             bson.ObjectId
		addition       *BookmarksUpdate
		expectedOutput error
	}{
		{
			name: "Test Updating User with Valid Update Addition",
			id:   usr.ID,
			addition: &BookmarksUpdate{
				Adding:   true,
				UpdateID: filler,
			},
			expectedOutput: nil,
		},
		{
			name: "Test Updating User That doesn't exist",
			id:   bson.NewObjectId(),
			addition: &BookmarksUpdate{
				Adding:   true,
				UpdateID: bson.NewObjectId(),
			},
			expectedOutput: fmt.Errorf("error getting users by id: not found"),
		},
		{
			name: "Test Updating User with Valid Update Remove",
			id:   usr.ID,
			addition: &BookmarksUpdate{
				Adding:   false,
				UpdateID: filler,
			},
			expectedOutput: nil,
		},
	}

	for _, c := range cases {
		updatedUser, err := mongoStore.BookmarksUpdate(c.id, c.addition)
		if err != nil && c.expectedOutput != nil {
			if err.Error() != c.expectedOutput.Error() {
				t.Errorf("Error on %s: Expected %v but got %v", c.name, c.expectedOutput, err)
			}
		}
		if !(err == nil && c.expectedOutput == nil) && (err == nil || c.expectedOutput == nil) {
			t.Errorf("Error on %s: Expected %v but got %v", c.name, c.expectedOutput, err)
		}
		if updatedUser != nil {
			if c.addition.Adding {
				if len(updatedUser.Bookmarks) != 1 {
					t.Errorf("Error on %s: Expected %v but got %v", c.name, 1, len(updatedUser.Bookmarks))
				}
			} else {
				if len(updatedUser.Bookmarks) != 0 {
					t.Errorf("Error on %s: Expected %v but got %v", c.name, 0, len(updatedUser.Bookmarks))
				}
			}
		}
	}
}
