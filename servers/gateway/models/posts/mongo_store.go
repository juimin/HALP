package posts

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Filters for mongo search

//IDFilter filters by post id
type IDFilter struct {
	ID bson.ObjectId
}

//AuthorIDFilter filters by author id
type AuthorIDFilter struct {
	AuthorID bson.ObjectId
}

//BoardIDFilter filters by board id
type BoardIDFilter struct {
	BoardID bson.ObjectId
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

//Insert inserts a new post into the store
func (s *MongoStore) Insert(np *NewPost) (*Post, error) {
	post, err := np.ToPost()
	if err != nil {
		return nil, err
	}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.Insert(post); err != nil {
		return nil, fmt.Errorf("error inserting post: %v", err)
	}
	return post, nil
}

//GetByID gets all posts by the given ID
func (s *MongoStore) GetByID(id bson.ObjectId) (*Post, error) {
	post := &Post{}
	filter := &IDFilter{id}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.FindId(filter.ID).One(post); err != nil {
		return nil, fmt.Errorf("error getting posts by id: %v", err)
	}
	return post, nil
}

//GetByAuthorID gets all posts by the given author (via id)
func (s *MongoStore) GetByAuthorID(id bson.ObjectId) ([]*Post, error) {
	posts := []*Post{}
	filter := &AuthorIDFilter{id}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.FindId(filter.AuthorID).All(&posts); err != nil {
		return nil, fmt.Errorf("error getting posts by author id: %v", err)
	}
	return posts, nil
}

//GetByBoardID gets all posts from a given board (via id)
func (s *MongoStore) GetByBoardID(id bson.ObjectId) ([]*Post, error) {
	posts := []*Post{}
	filter := &BoardIDFilter{id}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.FindId(filter.BoardID).All(&posts); err != nil {
		return nil, fmt.Errorf("error getting posts by board id: %v", err)
	}
	return posts, nil
}

//Delete removes a post from the database
func (s *MongoStore) Delete(id bson.ObjectId) error {
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.RemoveId(id); err != nil {
		return err
	}
	return nil
}

//PostUpdate updates the post with general information
func (s *MongoStore) PostUpdate(id bson.ObjectId, update *PostUpdate) error {
	change := mgo.Change{
		Update:    bson.M{"$set": update},
		ReturnNew: true,
	}
	post := &Post{}
	col := s.session.DB(s.dbname).C(s.colname)
	if _, err := col.FindId(id).Apply(change, post); err != nil {
		return err
	}
	return nil
}
