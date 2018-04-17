package posts

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

//ErrPostNotFound is returned when the user can't be found
var ErrPostNotFound = errors.New("post not found")

// AllPostLimit is a Limit for user queries
var AllPostLimit = 1000

//Store represents a store for Posts
type Store interface {
	// Insert converts the NewPost to a Post, inserts
	// it into the database, and returns it
	Insert(newPost *NewPost) (*Post, error)

	//GetByID returns the Post with the given ID
	GetByID(id bson.ObjectId) (*Post, error)

	//GetByAuthorID returns all Posts by a given author by ID
	GetByAuthorID(id bson.ObjectId) (*[]Post, error)

	//GetByBoardID returns all Posts for a given board by ID
	GetByBoardID(id bson.ObjectId) (*[]Post, error)

	//Delete deletes the post with the given ID
	Delete(id bson.ObjectId) error

	//PostUpdate applies a PostUpdate to a given post ID
	PostUpdate(id bson.ObjectId, update *PostUpdate) error
}
