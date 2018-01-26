package users

import (
	"crypto/subtle"
	"fmt"

	"golang.org/x/crypto/bcrypt"
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

type passUpdate struct {
	PassHash []byte
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
	if err := col.FindId(filter.ID).One(user); err != nil {
		return nil, fmt.Errorf("error getting users by id: %v", err)
	}
	return user, nil
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

// UserUpdate updates the user with the general information
func (s *MongoStore) UserUpdate(userID bson.ObjectId, updates *UserUpdate) error {
	change := mgo.Change{
		Update:    bson.M{"$set": updates}, //bson.M is map of string, to some value
		ReturnNew: true,
	}
	user := &User{}
	col := s.session.DB(s.dbname).C(s.colname)
	if _, err := col.FindId(userID).Apply(change, user); err != nil {
		return err
	}
	return nil
}

// PassUpdate updates the password of the given user
func (s *MongoStore) PassUpdate(userID bson.ObjectId, updates *PasswordUpdate) error {
	// The user should be authenticated already
	// Check password and password conf
	if len(updates.NewPassword) == 0 || len(updates.NewPasswordConf) == 0 {
		return fmt.Errorf("Invalid Input: New Password cannot be length 0")
	}

	if subtle.ConstantTimeCompare([]byte(updates.NewPassword), []byte(updates.NewPasswordConf)) != 1 {
		return fmt.Errorf("Password and password conf do not match")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(updates.NewPassword), bcryptCost)
	if err != nil {
		return fmt.Errorf("Bcrypt error")
	}
	update := &passUpdate{
		PassHash: pass,
	}
	change := mgo.Change{
		Update:    bson.M{"$set": update}, //bson.M is map of string, to some value
		ReturnNew: true,
	}
	user := &User{}
	col := s.session.DB(s.dbname).C(s.colname)
	if _, err := col.FindId(userID).Apply(change, user); err != nil {
		return err
	}
	return nil
}
