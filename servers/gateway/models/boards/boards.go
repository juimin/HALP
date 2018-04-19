package boards

import "gopkg.in/mgo.v2/bson"

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
	Sub int `json:"sub"`
}

// UpdatePost represents the increment or decrement of posts
type UpdatePost struct {
	Post int `json:"post"`
}

// TempBoolStore represents the initial value given to the server from an HTTP request
type TempBoolStore struct {
	TempSubPost bool `json:"temp"`
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
