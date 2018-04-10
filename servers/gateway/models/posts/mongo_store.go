package posts

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Filters for mongo search
type IdFilter struct {
	ID bson.ObjectId
}

type AuthorIdFilter struct {
	AuthorID bson.ObjectId
}

type BoardIdFilter struct {
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
	filter := &IdFilter{id}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.FindId(filter.ID).One(post); err != nil {
		return nil, fmt.Errorf("error getting posts by id: %v", err)
	}
	return post, nil
}

//GetByAuthorID gets all posts by the given author (via id)
func (s *MongoStore) GetByAuthorID(id bson.ObjectId) ([]*Post, error) {
	posts := []*Post{}
	filter := &AuthorIdFilter{id}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.FindId(filter.AuthorID).All(&posts); err != nil {
		return nil, fmt.Errorf("error getting posts by author id: %v", err)
	}
	return posts, nil
}

//GetByBoardID gets all posts from a given board (via id)
func (s *MongoStore) GetByBoardID(id bson.ObjectId) ([]*Post, error) {
	posts := []*Post{}
	filter := &BoardIdFilter{id}
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
//fix postupdate in posts.go to apply up/down/total votes
