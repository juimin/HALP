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
