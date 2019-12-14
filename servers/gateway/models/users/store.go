package users

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

//ErrUserNotFound is returned when the user can't be found
var ErrUserNotFound = errors.New("user not found")

// AllUserLimit is a Limit for user queries
var AllUserLimit = 1000

// Store represents a store for Users
type Store interface {
	// GetByEmail returns the User with the given email
	GetByEmail(email string) (*User, error)

	// GetByID returns the User with the given ID
	GetByID(id bson.ObjectId) (*User, error)

	//GetByUserName returns the User with the given Username
	GetByUserName(username string) (*User, error)

	// Insert converts the NewUser to a User, inserts
	// it into the database, and returns it
	Insert(newUser *NewUser) (*User, error)

	// Update applies UserUpdates to the given user ID
	UserUpdate(userID bson.ObjectId, updates *UserUpdate) error

	// Update applies password changes
	PassUpdate(userID bson.ObjectId, updates *PasswordUpdate) error

	// Delete deletes the user with the given ID
	Delete(userID bson.ObjectId) error

	// UpdateFavorites facilitates the updating of the users favorite boards
	FavoritesUpdate(userID bson.ObjectId, updates *FavoritesUpdate) (*User, error)

	// UpdateBookmarks allows for updating the list of bookmarks the user has
	BookmarksUpdate(userID bson.ObjectId, updates *BookmarksUpdate) (*User, error)

	// PostVoteUpdate
	PostVoteUpdate(userID bson.ObjectId, updates *PostVoting) (*User, error)

	// CommentVoteUpdate
	CommentVoteUpdate(userID bson.ObjectId, updates *CommentVoting) (*User, error)

	// GetAll gets every user in the store
	GetAll() ([]*User, error)
}
