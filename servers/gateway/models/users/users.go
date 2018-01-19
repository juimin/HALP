package users

import (
	"crypto/md5"
	"crypto/subtle"
	"fmt"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// Constant URL for obtaining a gravatar avatar img if you have one
// This will default to some blue thing if one does not exist for the given user
const gravatarBasePhotoURL = "https://www.gravatar.com/avatar/"

// Cost of Bcrypt
var bcryptCost = 13

// User represents a user account in the database
type User struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Email       string        `json:"email"`
	PassHash    []byte        `json:"-"` // Stored, but not encoded to clients
	UserName    string        `json:"userName"`
	FirstName   string        `json:"firstName"`
	LastName    string        `json:"lastName"`
	PhotoURL    string        `json:"photoURL"`
	Occupation  string        `json:"occupation"`
	Specialties []string      `json:"specialties"` // List representation of expert level categories
}

// Credentials represents user sign-in credentials
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// NewUser represents a new user signing up for an account
type NewUser struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordConf string `json:"passwordConf"`
	UserName     string `json:"userName"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Occupation   string `json:"occupation"`
}

// UserUpdate represents allowed updates to a user profile
// Updatable Elements:
// - Name Elements (First and Last)
// - Email
type UserUpdate struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
}

// PasswordUpdate represents requirements for changing the user's password
type PasswordUpdate struct {
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	NewPasswordConf string `json:"newPasswordConf"`
}

// Validate confirms that a new user contains information that we
// can work with
// If everything is good input, we can return nil, else return an
// error describing the problem
func (nu *NewUser) Validate() error {
	// Check email is not empty string
	if len(nu.Email) > 0 {
		// Parses an email address
		if _, err := mail.ParseAddress(nu.Email); err != nil {
			return fmt.Errorf("Email invalid")
		}
	} else {
		// An empty or no email newuser will hit this error first
		return fmt.Errorf("Email does not exist")
	}

	// Check password length
	if len(nu.Password) < 6 {
		return fmt.Errorf("Password must be at least 6 characters")
	}

	// Check password and password conf
	if subtle.ConstantTimeCompare([]byte(nu.Password), []byte(nu.PasswordConf)) != 1 {
		return fmt.Errorf("Password and password conf do not match")
	}

	// Check username length
	if len(nu.UserName) == 0 {
		return fmt.Errorf("Username must be greater than length zero")
	}

	// Everything checks out
	return nil
}

//ToUser converts the NewUser to a User, setting the
//PhotoURL and PassHash fields appropriately
func (nu *NewUser) ToUser() (*User, error) {
	//TODO: set the PhotoURL field of the new User to
	//the Gravatar PhotoURL for the user's email address.
	//see https://en.gravatar.com/site/implement/hash/
	//and https://en.gravatar.com/site/implement/images/

	//TODO: also set the ID field of the new User
	//to a new bson ObjectId
	//http://godoc.org/labix.org/v2/mgo/bson

	//TODO: also call .SetPassword() to set the PassHash
	//field of the User to a hash of the NewUser.Password

	// Validate the new user object to confirm that we can convert it
	// into a valid user
	err := nu.Validate()
	if err != nil {
		// Something went wrong
		return nil, err
	}

	// MD5 hasher
	hash := md5.New()
	// Hash the email using md5
	emailHash := string(hash.Sum([]byte(strings.ToLower(strings.Trim(nu.Email, " ")))))

	// We have a valid new user so we can generate a user object
	user := &User{
		Email:      nu.Email,
		UserName:   nu.UserName,
		FirstName:  nu.FirstName,
		LastName:   nu.LastName,
		ID:         bson.NewObjectId(),
		PhotoURL:   gravatarBasePhotoURL + emailHash,
		Occupation: nu.Occupation,
	}

	// Set the password using the given hash from the password generator
	user.SetPassword(nu.Password)
	// Return the user and no error
	return user, nil
}

//FullName returns the user's full name, in the form:
// "<FirstName> <LastName>"
//If either first or last name is an empty string, no
//space is put betweeen the names
func (u *User) FullName() string {
	//TODO: implement according to comment above
	if len(u.FirstName) > 0 {
		if len(u.LastName) > 0 {
			return u.FirstName + " " + u.LastName
		}
		return u.FirstName
	}
	if len(u.LastName) > 0 {
		return u.LastName
	}
	return ""
}

//SetPassword hashes the password and stores it in the PassHash field
func (u *User) SetPassword(password string) error {
	//TODO: use the bcrypt package to generate a new hash of the password
	//https://godoc.org/golang.org/x/crypto/bcrypt

	if len(password) < 6 {
		return fmt.Errorf("Password length much be 6 or greater")
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return fmt.Errorf("Bcrypt error")
	}
	u.PassHash = pass
	return nil
}

//Authenticate compares the plaintext password against the stored hash
//and returns an error if they don't match, or nil if they do
func (u *User) Authenticate(password string) error {
	//TODO: use the bcrypt package to compare the supplied
	//password with the stored PassHash
	//https://godoc.org/golang.org/x/crypto/bcrypt
	if len(password) == 0 {
		return fmt.Errorf("Cannot authenticate no password")
	}
	err := bcrypt.CompareHashAndPassword(u.PassHash, []byte(password))
	if err != nil {
		return fmt.Errorf("Bcrypt hash error")
	}
	return nil
}

//ApplyUpdates applies the updates to the user. An error
//is returned if the updates are invalid
func (u *User) ApplyUpdates(updates *Updates) error {
	//TODO: set the fields of `u` to the values of the related
	//field in the `updates` struct, enforcing the following rules:
	//- the FirstName must be non-zero-length
	//- the LastName must be non-zero-length
	if len(updates.FirstName) == 0 || len(updates.LastName) == 0 {
		return fmt.Errorf("Invalid updates. First and last name must both have a non-zero length")
	}
	u.FirstName = updates.FirstName
	u.LastName = updates.LastName
	return nil
}
