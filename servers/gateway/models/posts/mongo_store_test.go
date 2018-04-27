package posts

import (
	"testing"

	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

func TestNewMongoStore(t *testing.T) {

	ms, err := mgo.Dial("localhost:27017")

	if err != nil {
		t.Errorf("Dialing Mongo Failed: %v", err)
	}

	cases := []struct {
		name           string
		session        *mgo.Session
		db             string
		collectionName string
	}{
		{
			name:           "session",
			session:        ms,
			db:             "test_db",
			collectionName: "test_col",
		},
	}

	for _, c := range cases {
		ms := NewMongoStore(c.session, c.db, c.collectionName)
		if ms == nil {
			t.Errorf("%s Failed: MongoStore is null somehow", c.name)
		}
	}
}

func TestCRUDValid(t *testing.T) {
	author := bson.NewObjectId()

	board := bson.NewObjectId()

	np := &NewPost{
		Title:    "Test Post",
		ImageURL: "https://github.com",
		Caption:  "Alex I'm doing them lol",
		AuthorID: author,
		BoardID:  board,
	}

	// Generate a mongo store
	conn, err := mgo.Dial("localhost:27017")

	if err != nil {
		t.Errorf("Dialing Mongo Failed: %v", err)
		t.Errorf("%v is a thing", conn)
	}
	ms := NewMongoStore(conn, "test_db", "test_col")

	// Insert the post
	post, err := ms.Insert(np)

	if err != nil {
		t.Errorf("Insertion of Post Failed: %v", err)
	}

	getIDPost, err := ms.GetByID(post.ID)

	if err != nil {
		t.Errorf("Get By ID Failed for inserted board: %v", err)
	}

	// Check match
	if getIDPost.ID != post.ID {
		t.Errorf("Got the wrong post out of the database, expected %v but got %v", post.ID, getIDPost.ID)
	}

	getAuthorPosts, err := ms.GetByAuthorID(author)

	if err != nil {
		t.Errorf("Get By ID Failed for inserted board: %v", err)
	}

	// Check matches
	for _, p := range getAuthorPosts {
		if p.AuthorID != post.AuthorID {
			t.Errorf("Got the wrong post out of the database, expected %v but got %v", post.AuthorID, p.AuthorID)
		}
	}

	getBoardPosts, err := ms.GetByBoardID(board)

	if err != nil {
		t.Errorf("Get By ID Failed for inserted board: %v", err)
	}

	// Check matches
	for _, p := range getBoardPosts {
		if p.BoardID != post.BoardID {
			t.Errorf("Got the wrong post out of the database, expected %v but got %v", post.BoardID, p.BoardID)
		}
	}

	// Update

	pu := &PostUpdate{
		Title:    "LOL",
		Caption:  "Tomato",
		ImageURL: "https://update.com",
	}

	err = ms.PostUpdate(post.ID, pu)

	if err != nil {
		t.Errorf("Error updating the post: %v", err)
	}

	// Delete

	err = ms.Delete(post.ID)

	if err != nil {
		t.Errorf("Error deleting the post: %v", err)
	}

	// Check deletion
	post, err = ms.GetByID(post.ID)

	if err == nil {
		t.Errorf("Deletion was bad %v", err)
	}
}

func TestInsert(t *testing.T) {
	author := bson.NewObjectId()

	board := bson.NewObjectId()

	np := &NewPost{
		Title:    "",
		ImageURL: "https://github.com",
		Caption:  "Alex I'm doing them lol",
		AuthorID: author,
		BoardID:  board,
	}

	// Generate a mongo store
	conn, err := mgo.Dial("localhost:27017")

	if err != nil {
		t.Errorf("Dialing Mongo Failed: %v", err)
		t.Errorf("%v is a thing", conn)
	}
	ms := NewMongoStore(conn, "test_db", "test_col")

	// Insert the post
	_, err = ms.Insert(np)

	if err == nil {
		t.Errorf("Expected Error here for invalid new post")
	}
}

func TestGetBadIDs(t *testing.T) {
	author := bson.NewObjectId()

	board := bson.NewObjectId()

	np := &NewPost{
		Title:    "adsfadsf",
		ImageURL: "https://github.com",
		Caption:  "Alex I'm doing them lol",
		AuthorID: author,
		BoardID:  board,
	}

	// Generate a mongo store
	conn, err := mgo.Dial("localhost:27017")

	if err != nil {
		t.Errorf("Dialing Mongo Failed: %v", err)
		t.Errorf("%v is a thing", conn)
	}
	ms := NewMongoStore(conn, "test_db", "test_col")

	// Insert the post
	_, err = ms.Insert(np)

	// Test the get methods for bad inputs

	_, err = ms.GetByAuthorID("!@#!@#!@C!@#!@C!~~~?~?~?!!!")

	if err == nil {
		t.Errorf("Expected Error here for Getting author by bad id")
	}

	_, err = ms.GetByBoardID("!@#!@#!@C!@#!@C!~~~?~?~?!!!")

	if err == nil {
		t.Errorf("Expected Error here for Getting author by bad id")
	}
}

func TestUpdateBadID(t *testing.T) {
	author := bson.NewObjectId()

	board := bson.NewObjectId()

	np := &NewPost{
		Title:    "adsfadsf",
		ImageURL: "https://github.com",
		Caption:  "Alex I'm doing them lol",
		AuthorID: author,
		BoardID:  board,
	}

	// Generate a mongo store
	conn, err := mgo.Dial("localhost:27017")

	if err != nil {
		t.Errorf("Dialing Mongo Failed: %v", err)
		t.Errorf("%v is a thing", conn)
	}
	ms := NewMongoStore(conn, "test_db", "test_col")

	// Insert the post
	_, err = ms.Insert(np)

	pu := &PostUpdate{
		Title:    "LOL",
		Caption:  "Tomato",
		ImageURL: "https://update.com",
	}

	err = ms.PostUpdate("1231`2312312312321d-d-fd-fdfa", pu)

	if err == nil {
		t.Errorf("Expected error updating post with bad id")
	}
}
