package posts

import (
	"fmt"
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

	if len(getAuthorPosts) == 0 {
		t.Errorf("There should be a post with authors in it")
	}
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
	if len(getBoardPosts) == 0 {
		t.Errorf("There should be boards in the database")
	}

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

func TestFaultyPost(t *testing.T) {
	// Generate a mongo store
	conn, err := mgo.Dial("localhost:27017")

	if err != nil {
		t.Errorf("Dialing Mongo Failed: %v", err)
		t.Errorf("%v is a thing", conn)
	}
	ms := NewMongoStore(conn, "test_db", "test_col")
	_, err = ms.PostVotes("!234123123~?~?", &PostVote{
		Upvote:   -231231231,
		Downvote: 1231231,
	})

	if err == nil {
		t.Errorf("Bad Get By ID Failed for nonexistent ID")
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

// TestVoting tests the vote function for the mongo store and makes sure that we
// Get valid votes
func TestVoting(t *testing.T) {
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
	post, err := ms.Insert(np)

	if err != nil {
		t.Errorf("Inserting the post failed")
	}

	cases := []struct {
		name              string
		votes             *PostVote
		destination       bson.ObjectId
		expectedError     error
		expectedUpvotes   int
		expectedDownvotes int
	}{
		{
			name: "Valid Vote 1 0",
			votes: &PostVote{
				Upvote:   1,
				Downvote: 0,
			},
			destination:       post.ID,
			expectedError:     nil,
			expectedUpvotes:   1,
			expectedDownvotes: 0,
		},
		{
			name: "Valid Vote 0 1",
			votes: &PostVote{
				Upvote:   0,
				Downvote: 1,
			},
			destination:       post.ID,
			expectedError:     nil,
			expectedUpvotes:   1,
			expectedDownvotes: 1,
		},
		{
			name: "Bad Vote 1 1",
			votes: &PostVote{
				Upvote:   1,
				Downvote: 1,
			},
			destination:       post.ID,
			expectedError:     fmt.Errorf("The updates to both upvotes and downvotes cannot be the same"),
			expectedUpvotes:   1,
			expectedDownvotes: 1,
		},
		{
			name: "Bad Vote -1 -1",
			votes: &PostVote{
				Upvote:   -1,
				Downvote: -1,
			},
			destination:       post.ID,
			expectedError:     fmt.Errorf("The updates to both upvotes and downvotes cannot be the same"),
			expectedUpvotes:   1,
			expectedDownvotes: 1,
		},
		{
			name: "Bad Vote 0 0",
			votes: &PostVote{
				Upvote:   0,
				Downvote: 0,
			},
			destination:       post.ID,
			expectedError:     fmt.Errorf("The updates to both upvotes and downvotes cannot be the same"),
			expectedUpvotes:   1,
			expectedDownvotes: 1,
		},
		{
			name: "Valid Vote +1 -1",
			votes: &PostVote{
				Upvote:   1,
				Downvote: -1,
			},
			destination:       post.ID,
			expectedError:     nil,
			expectedUpvotes:   2,
			expectedDownvotes: 0,
		},
		{
			name: "Valid Vote -1 +1",
			votes: &PostVote{
				Upvote:   -1,
				Downvote: 1,
			},
			destination:       post.ID,
			expectedError:     nil,
			expectedUpvotes:   1,
			expectedDownvotes: 1,
		},
		{
			name: "Bad Vote - Value out of bounds for upvote",
			votes: &PostVote{
				Upvote:   -231231231,
				Downvote: 0,
			},
			destination:       post.ID,
			expectedError:     fmt.Errorf("Upvotes are out of bounds"),
			expectedUpvotes:   1,
			expectedDownvotes: 1,
		},
		{
			name: "Bad Vote - Value out of bounds for downvote",
			votes: &PostVote{
				Upvote:   0,
				Downvote: 1231231,
			},
			destination:       post.ID,
			expectedError:     fmt.Errorf("Downvotes are out of bounds"),
			expectedUpvotes:   1,
			expectedDownvotes: 1,
		},
	}

	// Test all the vote test cases
	for _, c := range cases {
		p, err := ms.PostVotes(c.destination, c.votes)
		if err != nil {
			if c.expectedError != nil {
				if err.Error() != c.expectedError.Error() {
					t.Errorf("%s Failed. Incorrect Error: expected %v but got %v", c.name, c.expectedError, err)
				}
			} else {
				t.Errorf("%s failed: expected nil but got %v", c.name, err)
			}
		} else {
			if nil != c.expectedError {
				t.Errorf("%s Failed: Expected Error %v but got nil", c.name, c.expectedError)
			} else {
				// Both are nil which means that we should have succeeded
				if p.Downvotes != c.expectedDownvotes {
					t.Errorf("%s Failed: Updated downvotes is wrong: expected %d but got %d", c.name, c.expectedDownvotes, p.Downvotes)
				}
				if p.Upvotes != c.expectedUpvotes {
					t.Errorf("%s Failed: Updated upvotes is wrong: expected %d but got %d", c.name, c.expectedUpvotes, p.Upvotes)
				}
			}
		}
	}

}
