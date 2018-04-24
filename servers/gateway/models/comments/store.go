package comments

import (
	"gopkg.in/mgo.v2/bson"
)

// Store defines the comments store interface, which can be used when constructing
// storage handlers through different database systems (Mongo and SQL)
type Store interface {

	// GetCommentsByPostId retrieves all the comments associated with the given post id
	GetCommentsByPostID(postID bson.ObjectId) (*[]Comment, error)

	// GetByParentId takes a comment id and returns all secondary comments assoicated with it
	GetByParentID(commentID bson.ObjectId) (*[]SecondaryComment, error)

	// GetByCommentId retreives a primary comment by its Id
	GetByCommentID(commentID bson.ObjectId) (*Comment, error)

	// GetBySecondaryId retreives a secondary comment by its Id
	GetBySecondaryID(secondaryID bson.ObjectId) (*SecondaryComment, error)

	// InsertComment inserts a new comment into the store
	InsertComment(newComment *NewComment) (*Comment, error)

	// InsertSecondaryComment inserts a new secondary comment into the store
	InsertSecondaryComment(newSecondary *NewSecondaryComment) (*SecondaryComment, error)

	// DeleteComment removes a primary comment from a store
	DeleteComment(commentID bson.ObjectId) error

	// DeleteSecondaryComment removes a secondary comment from the store
	DeleteSecondaryComment(secondaryID bson.ObjectId) error

	// UpdateComment updates the parent level comments
	UpdateComment(commentID bson.ObjectId, updates *CommentUpdate) (*Comment, error)

	// UpdateSecondaryComment updates a secondary level comment
	UpdateSecondaryComment(secondaryID bson.ObjectId, updates *SecondaryCommentUpdate) (*SecondaryComment, error)
}
