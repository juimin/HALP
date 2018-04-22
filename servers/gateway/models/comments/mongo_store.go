package comments

import (
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
