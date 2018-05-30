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
	AuthorID bson.ObjectId `json:"author_id"`
}

//BoardIDFilter filters by board id
type BoardIDFilter struct {
	BoardID bson.ObjectId `json:"board_id"`
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
	// filter := &AuthorIDFilter{id}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.Find(bson.M{"authorid": id}).All(&posts); err != nil {
		return nil, fmt.Errorf("error getting posts by author id: %v", err)
	}
	return posts, nil
}

//GetByBoardID gets all posts from a given board (via id)
func (s *MongoStore) GetByBoardID(id bson.ObjectId) ([]*Post, error) {
	posts := []*Post{}
	// filter := &BoardIDFilter{id}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.Find(bson.M{"boardid": id}).All(&posts); err != nil {
		return nil, fmt.Errorf("error getting posts by board id: %v", err)
	}
	return posts, nil
}

//Delete removes a post from the database
func (s *MongoStore) Delete(id bson.ObjectId) error {
	col := s.session.DB(s.dbname).C(s.colname)
	return col.RemoveId(id)
}

// GetLastN gets the last N posts inserted into the store
func (s *MongoStore) GetLastN(n int) ([]*Post, error) {
	var results []*Post
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.Find(bson.M{}).Sort("-time_created").Limit(n).All(&results); err != nil {
		return nil, fmt.Errorf("Error getting the last N posts %v", err)
	}
	return results, nil
}

// GetAll gets every post from the store
func (s *MongoStore) GetAll() ([]*Post, error) {
	var result []*Post
	col := s.session.DB(s.dbname).C(s.colname)
	err := col.Find(bson.M{}).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
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

// VoteInjector is used by the change in mgo to update the votes in a post object
type VoteInjector struct {
	Upvotes   int
	Downvotes int
}

// PostVotes takes care of updating the votes in the post with the given input
func (s *MongoStore) PostVotes(id bson.ObjectId, update *PostVote) (*Post, error) {
	post, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Apply the votes
	err = post.ApplyVotes(update)

	if err != nil {
		return nil, err
	}
	change := mgo.Change{
		Update: bson.M{"$set": &VoteInjector{
			Upvotes:   post.Upvotes,
			Downvotes: post.Downvotes,
		}},
		ReturnNew: true,
	}
	output := &Post{}
	col := s.session.DB(s.dbname).C(s.colname)
	if _, err := col.FindId(id).Apply(change, output); err != nil {
		return nil, err
	}
	return output, nil
}
