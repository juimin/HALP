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
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Email      string        `json:"email"`
	PassHash   []byte        `json:"-"` // Stored, but not encoded to clients
	UserName   string        `json:"userName"`
	FirstName  string        `json:"firstName"`
	LastName   string        `json:"lastName"`
	PhotoURL   string        `json:"photoURL"`
	Occupation string        `json:"occupation"`
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
	NewPassword     string `json:"newPassword"`
	NewPasswordConf string `json:"newPasswordConf"`
}

// PassUpdate holds the hashed password for the new pass
type PassUpdate struct {
	PassHash []byte
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
	if nu.Password != nu.PasswordConf {
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
		ID:         bson.NewObjectId(),               // Generate a new bson object ID
		PhotoURL:   gravatarBasePhotoURL + emailHash, // Gravatar for the given email
		Occupation: nu.Occupation,
	}

	// Set the password using the given hash from the password generator
	user.SetPassword(nu.Password)
	// Return the user and no error
	return user, nil
}

// FullName outputs the full name of the given user
// Empty string is returned if no name is seen
func (u *User) FullName() string {
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

// SetPassword hashes the password and stores it in the PassHash field
func (u *User) SetPassword(password string) error {
	// use the bcrypt package to generate a new hash of the password
	// https://godoc.org/golang.org/x/crypto/bcrypt
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
	err := bcrypt.CompareHashAndPassword(u.PassHash, []byte(password))
	if err != nil {
		return fmt.Errorf("Bcrypt hash error")
	}
	return nil
}

//ApplyUpdates applies the updates to the user. An error
//is returned if the updates are invalid
func (u *User) ApplyUpdates(updates *UserUpdate) error {
	if len(updates.FirstName) == 0 || len(updates.LastName) == 0 {
		return fmt.Errorf("Invalid input. First and last name must both have a non-zero length")
	}

	// We can't deal with empty emails either because this is not optional
	if len(updates.Email) == 0 {
		return fmt.Errorf("Invalid Input. Email cannot be empty")
	}

	// Check Email valid
	if _, err := mail.ParseAddress(updates.Email); err != nil {
		return fmt.Errorf("Invalid input. Email not a valid email")
	}

	// We aren't dealing with occupation because it is optional
	u.FirstName = updates.FirstName
	u.LastName = updates.LastName
	u.Email = updates.Email
	u.Occupation = updates.Occupation

	return nil
}

// PassUpdate allows for the changing of a password given knowledge of the old password
func (u *User) PassUpdate(updates *PasswordUpdate) error {
	// The user should be authenticated already
	// Check password and password conf
	if len(updates.NewPassword) == 0 || len(updates.NewPasswordConf) == 0 {
		return fmt.Errorf("Invalid Input: New Password cannot be length 0")
	}

	if subtle.ConstantTimeCompare([]byte(updates.NewPassword), []byte(updates.NewPasswordConf)) != 1 {
		return fmt.Errorf("Password and password conf do not match")
	}

	// Set Password since we confirmed
	u.SetPassword(updates.NewPassword)

	// No problems setting the new password so we can return no error
	return nil
}
