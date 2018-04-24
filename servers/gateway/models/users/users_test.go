package users

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// Add tests for the various functions in user.go, as described in the assignment.
// use `go test -cover` to ensure that you are covering all or nearly all of your code paths.

// Testing validation function
func TestValidate(t *testing.T) {

	cases := []struct {
		name           string
		input          *NewUser
		expectedOutput error
	}{
		{
			name:           "Empty",
			input:          &NewUser{},
			expectedOutput: fmt.Errorf("Email does not exist"),
		},
		{
			name: "Invalid email",
			input: &NewUser{
				Email:        "asdfasdf",
				Password:     "huptwothreefour",
				PasswordConf: "huptwothreefour",
				UserName:     "potato",
				FirstName:    "Neal",
				LastName:     "Jones",
			},
			expectedOutput: fmt.Errorf("Email invalid"),
		},
		{
			name: "Valid input",
			input: &NewUser{
				Email:        "asd.df@gmail.com",
				Password:     "huptwothreefour",
				PasswordConf: "huptwothreefour",
				UserName:     "potato",
				FirstName:    "Neal",
				LastName:     "Jones",
			},
			expectedOutput: nil,
		},
		{
			name: "too short password",
			input: &NewUser{
				Email:        "asd.df@gmail.com",
				Password:     "ddd",
				PasswordConf: "ddd",
				UserName:     "potato",
				FirstName:    "Neal",
				LastName:     "Jones",
			},
			expectedOutput: fmt.Errorf("Password must be at least 6 characters"),
		},
		{
			name: "too short password",
			input: &NewUser{
				Email:        "asd.df@gmail.com",
				Password:     "ddd",
				PasswordConf: "ddd",
				UserName:     "potato",
				FirstName:    "Neal",
				LastName:     "Jones",
			},
			expectedOutput: fmt.Errorf("Password must be at least 6 characters"),
		},
		{
			name: "confirmation password does not match original",
			input: &NewUser{
				Email:        "asd.df@gmail.com",
				Password:     "smokeemd",
				PasswordConf: "asdfasdf",
				UserName:     "potato",
				FirstName:    "Neal",
				LastName:     "Jones",
			},
			expectedOutput: fmt.Errorf("Password and password conf do not match"),
		},
		{
			name: "no username",
			input: &NewUser{
				Email:        "asd.df@gmail.com",
				Password:     "smokeemdf",
				PasswordConf: "smokeemdf",
				UserName:     "",
				FirstName:    "Neal",
				LastName:     "Jones",
			},
			expectedOutput: fmt.Errorf("Username must be greater than length zero"),
		},
	}

	for _, c := range cases {
		result := c.input.Validate()
		// The test should only error out if we have non nil inputs for both expected and result
		if result != nil && c.expectedOutput != nil {
			if result.Error() != c.expectedOutput.Error() {
				t.Errorf("%s: got %s but expected %s", c.name, result, c.expectedOutput)
			}
		}
		if (result == nil || c.expectedOutput == nil) && !(result == nil && c.expectedOutput == nil) {
			t.Errorf("%s: got %s but expected %s", c.name, result, c.expectedOutput)
		}
	}
}

// TestToUser tests the conversion of a new user to a user object
func TestToUser(t *testing.T) {
	cases := []struct {
		name           string
		input          *NewUser
		expectedOutput error
	}{
		{
			name: "Validate Check - Any error",
			input: &NewUser{
				Email:        "asd.df@gmail.com",
				Password:     "huptwothreefour",
				PasswordConf: "huptwothreefour",
				UserName:     "",
				FirstName:    "Neal",
				LastName:     "Jones",
			},
			expectedOutput: fmt.Errorf("Username must be greater than length zero"),
		},
		{
			name: "Valid Input",
			input: &NewUser{
				Email:        "asd.df@gmail.com",
				Password:     "huptwothreefour",
				PasswordConf: "huptwothreefour",
				UserName:     "potatoman",
				FirstName:    "Neal",
				LastName:     "Jones",
			},
			expectedOutput: nil,
		},
	}

	for _, c := range cases {
		_, output := c.input.ToUser()
		if output != nil && c.expectedOutput != nil {
			if output.Error() != c.expectedOutput.Error() {
				t.Errorf("%s: got %s but expected %s", c.name, output, c.expectedOutput)
			}
		}
		if (output == nil || c.expectedOutput == nil) && !(output == nil && c.expectedOutput == nil) {
			t.Errorf("%s: got %s but expected %s", c.name, output, c.expectedOutput)
		}
	}
}

// Test Set Password tests setting the password
func TestSetPassword(t *testing.T) {
	cases := []struct {
		name           string
		input          string
		expectedOutput error
	}{
		{
			name:           "Invalid password",
			input:          "ddd",
			expectedOutput: fmt.Errorf("Password length much be 6 or greater"),
		},
		{
			name:           "Valid password",
			input:          "hotpotato123",
			expectedOutput: nil,
		},
	}

	testUser := &User{}
	for _, c := range cases {
		output := testUser.SetPassword(c.input)
		if output != nil && c.expectedOutput != nil {
			if output.Error() != c.expectedOutput.Error() {
				t.Errorf("%s: got %s but expected %s", c.name, output, c.expectedOutput)
			}
		}
		if (output == nil || c.expectedOutput == nil) && !(output == nil && c.expectedOutput == nil) {
			t.Errorf("%s: got %s but expected %s", c.name, output, c.expectedOutput)
		}
	}
}

func TestAuthenticate(t *testing.T) {
	cases := []struct {
		name           string
		input          string
		cost           int
		expectedOutput error
	}{
		{
			name:           "Invalid password",
			input:          "",
			cost:           bcryptCost,
			expectedOutput: fmt.Errorf("Cannot authenticate no password"),
		},
		{
			name:           "Valid password",
			input:          "hotpotato123",
			cost:           bcryptCost,
			expectedOutput: nil,
		},
	}

	for _, c := range cases {
		ph, err := bcrypt.GenerateFromPassword([]byte(c.input), bcryptCost)
		if err != nil {
			if err != nil && c.expectedOutput != nil {
				if err.Error() != c.expectedOutput.Error() {
					t.Errorf("%s: got %s but expected %s", c.name, err, c.expectedOutput)
				}
			}
		} else {
			testUser := &User{
				PassHash: ph,
			}
			output := testUser.Authenticate(c.input)
			if output != nil && c.expectedOutput != nil {
				if output.Error() != c.expectedOutput.Error() {
					t.Errorf("%s: got %s but expected %s", c.name, output, c.expectedOutput)
				}
			}
			if (output == nil || c.expectedOutput == nil) && !(output == nil && c.expectedOutput == nil) {
				t.Errorf("%s: got %s but expected %s", c.name, output, c.expectedOutput)
			}
		}
	}
}
func TestApplyUpdates(t *testing.T) {
	cases := []struct {
		name           string
		input          *UserUpdate
		expectedOutput error
	}{
		{
			name: "First Name Invalid",
			input: &UserUpdate{
				FirstName:  "",
				LastName:   "Hello",
				Email:      "d95wang@gmail.com",
				Occupation: "Farmer",
			},
			expectedOutput: fmt.Errorf("Invalid input. First and last name must both have a non-zero length"),
		},
		{
			name: "Last Name Invalid",
			input: &UserUpdate{
				FirstName:  "Derek",
				LastName:   "",
				Email:      "d95wang@gmail.com",
				Occupation: "Farmer",
			},
			expectedOutput: fmt.Errorf("Invalid input. First and last name must both have a non-zero length"),
		},
		{
			name: "Invalid Email: NO input",
			input: &UserUpdate{
				FirstName:  "Derek",
				LastName:   "Hello",
				Email:      "",
				Occupation: "Farmer",
			},
			expectedOutput: fmt.Errorf("Invalid Input. Email cannot be empty"),
		},

		{
			name: "Invalid Email: bad input",
			input: &UserUpdate{
				FirstName:  "Derek",
				LastName:   "Hello",
				Email:      "d95wanglol",
				Occupation: "Farmer",
			},
			expectedOutput: fmt.Errorf("Invalid input. Email not a valid email"),
		},

		{
			name: "Valid Input",
			input: &UserUpdate{
				FirstName:  "Derek",
				LastName:   "Hello",
				Email:      "d95wang@gmail.com",
				Occupation: "Farmer",
			},
			expectedOutput: nil,
		},
	}
	testUser := &User{}
	for _, c := range cases {
		output := testUser.ApplyUpdates(c.input)
		if output != nil && c.expectedOutput != nil {
			if output.Error() != c.expectedOutput.Error() {
				t.Errorf("%s: got %s but expected %s", c.name, output, c.expectedOutput)
			}
		}
		if (output == nil || c.expectedOutput == nil) && !(output == nil && c.expectedOutput == nil) {
			t.Errorf("%s: got %s but expected %s", c.name, output, c.expectedOutput)
		}
	}
}

func TestUpdateFavorites(t *testing.T) {

	filler := []bson.ObjectId{
		bson.NewObjectId(),
	}

	cases := []struct {
		name           string
		user           *NewUser
		expectedUpdate *FavoritesUpdate
		filler         bool
	}{
		{
			name: "Testing Update",
			user: &NewUser{
				Email:        "asd.df@gmail.com",
				Password:     "huptwothreefour",
				PasswordConf: "huptwothreefour",
				UserName:     "potatoman",
				FirstName:    "Neal",
				LastName:     "Jones",
			},
			expectedUpdate: &FavoritesUpdate{
				Adding:   true,
				UpdateID: bson.NewObjectId(),
			},
			filler: false,
		},
		{
			name: "Testing Failed Removal",
			user: &NewUser{
				Email:        "asd.df@gmail.com",
				Password:     "huptwothreefour",
				PasswordConf: "huptwothreefour",
				UserName:     "potatoman",
				FirstName:    "Neal",
				LastName:     "Jones",
			},
			expectedUpdate: &FavoritesUpdate{
				Adding:   false,
				UpdateID: bson.NewObjectId(),
			},
			filler: false,
		},
		{
			name: "Testing Successful Removal",
			user: &NewUser{
				Email:        "asd.df@gmail.com",
				Password:     "huptwothreefour",
				PasswordConf: "huptwothreefour",
				UserName:     "potatoman",
				FirstName:    "Neal",
				LastName:     "Jones",
			},
			expectedUpdate: &FavoritesUpdate{
				Adding:   false,
				UpdateID: filler[0],
			},
			filler: true,
		},
	}

	for _, c := range cases {
		u, err := c.user.ToUser()
		if err != nil {
			t.Errorf("Error on %s: %v", c.name, err)
		}
		if c.filler {
			u.Favorites = filler
		}
		u.UpdateFavorites(c.expectedUpdate)
		if c.expectedUpdate.Adding {
			if u.Favorites[0] != c.expectedUpdate.UpdateID {
				t.Errorf("Error on %s: Epected %s at %d but got %s", c.name,
					c.expectedUpdate.UpdateID, 0, u.Favorites[0])
			}
		} else {
			if len(u.Favorites) != len(filler)-1 {
				t.Errorf("Error on %s: Epected %s at %d but got %s", c.name,
					c.expectedUpdate.UpdateID, 0, u.Favorites[0])
			}
		}
	}
}

// TestUpdateBookmarks tests the updating of bookmarks
func TestUpdateBookmarks(t *testing.T) {

	filler := []bson.ObjectId{
		bson.NewObjectId(),
	}

	cases := []struct {
		name           string
		user           *NewUser
		expectedUpdate *BookmarksUpdate
		filler         bool
	}{
		{
			name: "Testing Update",
			user: &NewUser{
				Email:        "asd.df@gmail.com",
				Password:     "huptwothreefour",
				PasswordConf: "huptwothreefour",
				UserName:     "potatoman",
				FirstName:    "Neal",
				LastName:     "Jones",
			},
			expectedUpdate: &BookmarksUpdate{
				Adding:   true,
				UpdateID: bson.NewObjectId(),
			},
			filler: false,
		},
		{
			name: "Testing Failed Removal",
			user: &NewUser{
				Email:        "asd.df@gmail.com",
				Password:     "huptwothreefour",
				PasswordConf: "huptwothreefour",
				UserName:     "potatoman",
				FirstName:    "Neal",
				LastName:     "Jones",
			},
			expectedUpdate: &BookmarksUpdate{
				Adding:   false,
				UpdateID: bson.NewObjectId(),
			},
			filler: false,
		},
		{
			name: "Testing Success Removal",
			user: &NewUser{
				Email:        "asd.df@gmail.com",
				Password:     "huptwothreefour",
				PasswordConf: "huptwothreefour",
				UserName:     "potatoman",
				FirstName:    "Neal",
				LastName:     "Jones",
			},
			expectedUpdate: &BookmarksUpdate{
				Adding:   false,
				UpdateID: filler[0],
			},
			filler: true,
		},
	}

	for _, c := range cases {
		u, err := c.user.ToUser()
		if err != nil {
			t.Errorf("Error on %s: %v", c.name, err)
		}
		if c.filler {
			u.Bookmarks = filler
		}
		u.UpdateBookmarks(c.expectedUpdate)
		if c.expectedUpdate.Adding {
			if u.Bookmarks[0] != c.expectedUpdate.UpdateID {
				t.Errorf("Error on %s: Epected %s at %d but got %s", c.name,
					c.expectedUpdate.UpdateID, 0, u.Bookmarks[0])
			}
		} else {
			if len(u.Bookmarks) != 0 {
				t.Errorf("Error on %s: Epected %s at %d but got %s", c.name,
					c.expectedUpdate.UpdateID, 0, u.Bookmarks[0])
			}

		}
	}
}
