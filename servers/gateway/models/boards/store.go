package boards

import (
	"gopkg.in/mgo.v2/bson"
)

// Store represents a store for Boards
type Store interface {
	// GetByID returns the User with the given ID
	GetByID(id bson.ObjectId) (*Board, error)

	//GetByUserName returns the User with the given Username
	GetByBoardName(title string) (*Board, error)

	//GetAllBoards returns every board that exists in the app
	GetAllBoards() ([]*Board, error)

	//UpdateSubscriberCount returns the board with a new subscriber count
	UpdateSubscriberCount(BoardID bson.ObjectId, subs *UpdateSubscriber) (*Board, error)

	//UpdatePostCount returns the board with a new post count
	UpdatePostCount(BoardID bson.ObjectId, posts *UpdatePost) (*Board, error)
}
