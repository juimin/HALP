package boards

import (
	"fmt"
	"net/url"

	"gopkg.in/mgo.v2/bson"
)

// Board represents the broad category for user posts in the database
type Board struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Image       string        `json:"image"`
	Subscribers int           `json:"subscribers"`
	Posts       int           `json:"posts"`
}

// UpdateSubscriber represents the increment or decrement of subscribers
type UpdateSubscriber struct {
	Subscribers int `json:"subscribers"`
}

// UpdatePost represents the increment or decrement of posts
type UpdatePost struct {
	Posts int `json:"posts"`
}

// TempBoolStore represents the initial value given to the server from an HTTP request
type TempBoolStore struct {
	TempSubPost bool `json:"temp"`
}

// NewBoard represents a newly added board
type NewBoard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

// ChangeSubscriberCount takes in a boolean value to represent a change in subscribers
func (b *Board) ChangeSubscriberCount(update bool) {
	if b.Subscribers >= 0 {
		if update {
			b.Subscribers++
		} else {
			b.Subscribers--
		}
	} else {
		b.Subscribers = 0
	}
}

// ChangePostCount takes in a boolean value to represent a change in posts
func (b *Board) ChangePostCount(update bool) {
	if b.Posts >= 0 {
		if update {
			b.Posts++
		} else {
			b.Posts--
		}
	} else {
		b.Posts = 0
	}
}

//Validate checks that the new board is correctly formatted
func (nb *NewBoard) Validate() error {
	//Check for a title
	if len(nb.Title) == 0 {
		return fmt.Errorf("Please enter a title")
	}

	if len(nb.Description) == 0 {
		return fmt.Errorf("Please enter a description")
	}

	//Check for either caption or image
	if len(nb.Image) == 0 {
		return fmt.Errorf("Please use either an image or a caption")
	}
	//if image, verify url
	if len(nb.Image) > 0 {
		if _, err := url.ParseRequestURI(nb.Image); err != nil {
			return fmt.Errorf("Invalid Photo URL")
		}
	}
	return nil
}

//ToBoard converts the NewBoard to a Board
func (nb *NewBoard) ToBoard() (*Board, error) {
	err := nb.Validate()
	if err != nil {
		return nil, err
	}
	board := &Board{
		ID:          bson.NewObjectId(),
		Title:       nb.Title,
		Description: nb.Description,
		Image:       nb.Image,
		Subscribers: 0,
		Posts:       0,
	}
	return board, nil
}
