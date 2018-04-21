package comments

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// SecondaryComment stores the secondary comments, which are responses to comments
// allowing further discussion
type SecondaryComment struct {
	ID          bson.ObjectId   `json:"id" bson:"_id"`
	ImageURL    string          `json:"image_url"`
	Content     string          `json:"caption"`
	Parent      bson.ObjectId   `json:"parent"`
	AuthorID    bson.ObjectId   `json:"author_id"`
	PostID      bson.ObjectId   `json:"post_id"`
	Upvotes     map[string]bool `json:"upvotes"`
	Downvotes   map[string]bool `json:"downvotes"`
	TotalVotes  int             `json:"total_votes"`
	TimeCreated time.Time       `json:"time_created"`
	TimeEdited  time.Time       `json:"time_edited"`
}

// NewSecondaryComment contains the information required for a secondary comment
type NewSecondaryComment struct {
	ImageURL string        `json:"image_url"`
	Content  string        `json:"caption"`
	PostID   bson.ObjectId `json:"post_id"`
	Parent   bson.ObjectId `json:"parent"`
	AuthorID bson.ObjectId `json:"author_id"`
}

// SecondaryCommentUpdate contains all the information that could be updated in a comment
type SecondaryCommentUpdate struct {
	ImageURL string `json:"image_url"`
	Content  string `json:"caption"`
}

// SecondaryCommentVote contains an integer that represents the vote of this user
type SecondaryCommentVote struct {
	Vote int `json:"vote"`
}

// Validate should validate the new comment object to confirm that we have a proper comment
func (nc *NewSecondaryComment) Validate() error {
	// Check that the comment contains meaningful information
	if len(nc.ImageURL) == 0 && len(nc.Content) == 0 {
		return fmt.Errorf("Cannot have no image or content in comment")
	}
	// Passed the validation
	return nil
}

// ToComment takes a new comment and converts it to a comment object
func (nc *NewSecondaryComment) ToComment() (*Comment, error) {

	// Validate the new comment structure is admissable
	if err := nc.Validate(); err != nil {
		return nil, err
	}

	// Construct the new comment
	comment := &Comment{
		ID:          bson.NewObjectId(),
		ImageURL:    nc.ImageURL,
		Content:     nc.Content,
		AuthorID:    nc.AuthorID,
		Comments:    []bson.ObjectId{},
		PostID:      nc.PostID,
		Upvotes:     map[string]bool{},
		Downvotes:   map[string]bool{},
		TotalVotes:  0,
		TimeCreated: time.Now(),
		TimeEdited:  time.Now(),
	}

	// Return the created comment
	return comment, nil
}
