package users

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Define Filters for mongo search
type emailFilter struct {
	Email string
}

type idFilter struct {
	ID bson.ObjectId
}

type userNameFilter struct {
	UserName string
}

// Define Updates for update statements
type updateUser struct {
	FirstName string
	LastName  string
	Email     string
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

//Insert inserts a new user into the store
func (s *MongoStore) Insert(nu *NewUser) (*User, error) {
	user, err := nu.ToUser()
	if err != nil {
		return nil, err
	}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.Insert(user); err != nil {
		return nil, fmt.Errorf("error inserting user: %v", err)
	}
	return user, nil
}

//GetByEmail gets all the users by the given email
func (s *MongoStore) GetByEmail(email string) (*User, error) {
	user := &User{}
	filter := &emailFilter{email}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.Find(filter).One(user); err != nil {
		return nil, fmt.Errorf("error getting users by email: %v", err)
	}
	return user, nil
}

//GetByUserName gets all the users by the given username
func (s *MongoStore) GetByUserName(userName string) (*User, error) {
	user := &User{}
	filter := &userNameFilter{userName}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.Find(filter).One(user); err != nil {
		return nil, fmt.Errorf("error getting users by username: %v", err)
	}
	return user, nil
}

//GetByID gets all the users by the given id
func (s *MongoStore) GetByID(id bson.ObjectId) (*User, error) {
	user := &User{}
	filter := &idFilter{id}
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.Find(filter).One(user); err != nil {
		return nil, fmt.Errorf("error getting users by id: %v", err)
	}
	return user, nil
}

// GetByIDs takes a slice of user ids and returns the user objects
func (s *MongoStore) GetByIDs(ids []*bson.ObjectId) ([]*User, error) {
	col := s.session.DB(s.dbname).C(s.colname)
	output := []*User{}
	for _, id := range ids {
		var user User
		// filter := &idFilter{id}
		if err := col.Find(bson.M{"_id": id}).One(&user); err != nil {
			fmt.Printf("...error getting users by id: %v\n", err)
		} else {
			output = append(output, &user)
		}
	}
	return output, nil
}

// Delete removes a user from the database
func (s *MongoStore) Delete(id bson.ObjectId) error {
	col := s.session.DB(s.dbname).C(s.colname)
	if err := col.RemoveId(id); err != nil {
		return err
	}
	// No error so delete must have been a success
	return nil
}
