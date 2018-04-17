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

	//GetById returns the Post with the given ID
	GetById(id bson.ObjectId) (*Post, error)

	//GetByAuthorId returns all Posts by a given author by ID
	GetByAuthorId(id bson.ObjectId) (*[]Post, error)

	//GetByBoardId returns all Posts for a given board by ID
	GetByBoardId(id bson.ObjectId) (*[]Post, error)

	//Delete deletes the post with the given ID
	Delete(id bson.ObjectId) error

	//PostUpdate applies a PostUpdate to a given post ID
	PostUpdate(id bson.ObjectId, update *PostUpdate) error
}
