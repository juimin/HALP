package boards

import "gopkg.in/mgo.v2/bson"

// Board represents the broad category for user posts in the database
type Board struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Icon        string        `json:"icon"`
	Image       string        `json:"image"`
	Subcribers  int           `json:"subscribers"`
	Posts       int           `json:"posts"`
}
