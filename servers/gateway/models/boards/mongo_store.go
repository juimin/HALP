package boards

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Define Filters for mongo search
type idFilter struct {
	ID bson.ObjectId
}

type titleFilter struct {
	Title string
}

//MongoStore outlines the storage struct for mongo db
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

//GetByID gets and returns all Boards from the given ID
func (s *MongoStore) GetByID(id bson.ObjectId) (*Board, error) {
	board := &Board{}
	filter := &idFilter{id}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.FindId(filter.ID).One(board); err != nil {
		return nil, fmt.Errorf("error getting boards by Board id: %v", err)
	}
	return board, nil
}

//GetByBoardName gets and returns all Boards from the given Name
func (s *MongoStore) GetByBoardName(title string) (*Board, error) {
	board := &Board{}
	filter := &titleFilter{title}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.Find(filter.Title).One(board); err != nil {
		return nil, fmt.Errorf("error getting boards by Board name: %v", err)
	}
	return board, nil
}

//GetAllBoards returns all boards in the application
func (s *MongoStore) GetAllBoards() ([]*Board, error) {
	boards := []*Board{}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.Find(nil).All(&boards); err != nil {
		return nil, fmt.Errorf("error getting all boards because: %v", err)
	}
	return boards, nil
}

//UpdateSubscriberCount updates the subscriber count in MongoDB
func (s *MongoStore) UpdateSubscriberCount(BoardID bson.ObjectId, subs *UpdateSubscriber) (*Board, error) {
	change := mgo.Change{
		Update:    bson.M{"$set": subs},
		ReturnNew: true,
	}
	board := &Board{}
	col := s.session.DB(s.dbname).C(s.colname)
	if _, err := col.FindId(BoardID).Apply(change, board); err != nil {
		return nil, fmt.Errorf("error adding subs to Board by id: %v", err)
	}
	return board, nil
}

//UpdatePostCount updates the subscriber count in MongoDB
func (s *MongoStore) UpdatePostCount(BoardID bson.ObjectId, posts *UpdatePost) (*Board, error) {
	change := mgo.Change{
		Update:    bson.M{"$set": posts},
		ReturnNew: true,
	}
	board := &Board{}
	col := s.session.DB(s.dbname).C(s.colname)
	if _, err := col.FindId(BoardID).Apply(change, board); err != nil {
		return nil, fmt.Errorf("error adding posts to Board by id: %v", err)
	}
	return board, nil
}

//CreateBoard adds a new board to the database to use for testing
func (s *MongoStore) CreateBoard(NewBoard *NewBoard) (*Board, error) {
	board, err := NewBoard.ToBoard()
	if err != nil {
		return nil, err
	}
	col := s.session.DB(s.dbname).C(s.colname)

	if err := col.Insert(board); err != nil {
		return nil, fmt.Errorf("error adding board: %v", err)
	}
	return board, nil
}
