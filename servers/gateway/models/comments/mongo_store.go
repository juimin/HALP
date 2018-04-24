package comments

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// This file defines a mongo storage implementation of the store interface

// IDFilter contains a struct for filtering by an ID
type IDFilter struct {
	ID bson.ObjectId
}

// MongoStore outlines the storage struct for mongo db
type MongoStore struct {
	session *mgo.Session
	dbname  string
	colname string
}

// VoteInjector is a temporary update object that allows for updates to votes
type VoteInjector struct {
	Upvotes   int
	Downvotes int
}

//NewMongoStore constructs a new MongoStore
func NewMongoStore(sess *mgo.Session, dbName string, collectionName string) *MongoStore {
	if sess == nil {
		panic("nil pointer passed for session")
	}
	return &MongoStore{
		session: sess,
		dbname:  dbName,
		colname: collectionName,
	}
}

//GetByCommentID gets the user by the given id
func (s *MongoStore) GetByCommentID(id bson.ObjectId) (*Comment, error) {
	c := &Comment{}
	filter := &IDFilter{id}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.FindId(filter.ID).One(c); err != nil {
		return nil, fmt.Errorf("error getting comments by id: %v", err)
	}
	return c, nil
}

//GetBySecondaryID gets the user by the given id
func (s *MongoStore) GetBySecondaryID(id bson.ObjectId) (*SecondaryComment, error) {
	sc := &SecondaryComment{}
	filter := &IDFilter{id}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.FindId(filter.ID).One(sc); err != nil {
		return nil, fmt.Errorf("error getting secondary comments by id: %v", err)
	}
	return sc, nil
}

//GetCommentsByPostID gets all the comments associated with a given post
func (s *MongoStore) GetCommentsByPostID(id bson.ObjectId) (*[]Comment, error) {
	comments := &[]Comment{}
	filter := &IDFilter{id}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.Find(filter).All(comments); err != nil {
		return nil, fmt.Errorf("error getting comments by post: %v", err)
	}
	return comments, nil
}

//GetByParentID gets all the comments associated with a given parent comment
func (s *MongoStore) GetByParentID(id bson.ObjectId) (*[]SecondaryComment, error) {
	sc := &[]SecondaryComment{}
	filter := &IDFilter{id}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.Find(filter).All(sc); err != nil {
		return nil, fmt.Errorf("error getting comments by parent comment: %v", err)
	}
	return sc, nil
}

//InsertComment inserts a new comment into the store
func (s *MongoStore) InsertComment(newComment *NewComment) (*Comment, error) {
	comm, err := newComment.ToComment()
	if err != nil {
		return nil, err
	}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.Insert(comm); err != nil {
		return nil, fmt.Errorf("error inserting comment: %v", err)
	}
	return comm, nil
}

//InsertSecondaryComment inserts a new comment into the store
func (s *MongoStore) InsertSecondaryComment(newSecondary *NewSecondaryComment) (*SecondaryComment, error) {
	comm, err := newSecondary.ToSecondaryComment()
	if err != nil {
		return nil, err
	}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.Insert(comm); err != nil {
		return nil, fmt.Errorf("error inserting secondary comment: %v", err)
	}
	return comm, nil
}

// DeleteComment removes a comment from the database
func (s *MongoStore) DeleteComment(commentID bson.ObjectId) error {
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.RemoveId(commentID); err != nil {
		return err
	}
	// No error so delete must have been a success
	return nil
}

// UpdateComment updates the parent level comments
func (s *MongoStore) UpdateComment(commentID bson.ObjectId, updates *CommentUpdate) (*Comment, error) {
	change := mgo.Change{
		Update:    bson.M{"$set": updates}, //bson.M is map of string, to some value
		ReturnNew: true,
	}
	comment := &Comment{}
	col := s.session.DB(s.dbname).C(s.colname)
	if _, err := col.FindId(commentID).Apply(change, comment); err != nil {
		return nil, err
	}
	return comment, nil
}

// UpdateSecondaryComment updates a secondary level comment
func (s *MongoStore) UpdateSecondaryComment(secondaryID bson.ObjectId, updates *SecondaryCommentUpdate) (*SecondaryComment, error) {
	change := mgo.Change{
		Update:    bson.M{"$set": updates}, //bson.M is map of string, to some value
		ReturnNew: true,
	}
	comment := &SecondaryComment{}
	col := s.session.DB(s.dbname).C(s.colname)
	if _, err := col.FindId(secondaryID).Apply(change, comment); err != nil {
		return nil, err
	}
	return comment, nil
}

// CommentVote deals with comment votes
func (s *MongoStore) CommentVote(id bson.ObjectId, updates *CommentVote) (*Comment, error) {
	// Get the comment in question
	comment, err := s.GetByCommentID(id)
	if err != nil {
		return nil, err
	}

	// Change the votes
	comment.Vote(updates)

	change := mgo.Change{
		Update: bson.M{"$set": &VoteInjector{
			Upvotes:   comment.Upvotes,
			Downvotes: comment.Downvotes,
		}}, //bson.M is map of string, to some value
		ReturnNew: true,
	}

	col := s.session.DB(s.dbname).C(s.colname)
	if _, err := col.FindId(id).Apply(change, comment); err != nil {
		return nil, err
	}
	return comment, nil
}

// SecondaryCommentVote deals with comment votes
func (s *MongoStore) SecondaryCommentVote(id bson.ObjectId, updates *SecondaryCommentVote) (*SecondaryComment, error) {
	// Get the comment in question
	comment, err := s.GetBySecondaryID(id)
	if err != nil {
		return nil, err
	}

	// Change the votes
	comment.Vote(updates)

	change := mgo.Change{
		Update: bson.M{"$set": &VoteInjector{
			Upvotes:   comment.Upvotes,
			Downvotes: comment.Downvotes,
		}}, //bson.M is map of string, to some value
		ReturnNew: true,
	}

	col := s.session.DB(s.dbname).C(s.colname)
	if _, err := col.FindId(id).Apply(change, comment); err != nil {
		return nil, err
	}
	return comment, nil
}
