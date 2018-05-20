package comments

import (
	"testing"

	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

// Test the creation of a mongostore
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

// Test the getting of a comment by it's object id in the database
func TestCommentCRUD(t *testing.T) {
	conn, err := mgo.Dial("localhost:27017")

	if err != nil {
		t.Errorf("Dialing Mongo Failed: %v", err)
	}

	ms := NewMongoStore(conn, "test_db", "test_col")

	newComment := &NewComment{
		AuthorID: bson.NewObjectId(),
		Content:  "Hello",
		PostID:   bson.NewObjectId(),
		ImageURL: "https://google.com",
	}

	comment, err := ms.InsertComment(newComment)

	if err != nil {
		t.Errorf("Testing InsertComment Failed: %v", err)
	}

	// Try get by comment id
	_, err = ms.GetByCommentID(comment.ID)

	if err != nil {
		t.Errorf("Testing GetByCommentID Failed: %v", err)
	}

	// Try get by post
	byPost, err := ms.GetCommentsByPostID(comment.PostID)

	if len(*byPost) == 0 {
		t.Errorf("Get by post should not be empty")
	}

	if err != nil {
		t.Errorf("Testing Get all by post id failed: %v", err)
	}

	newSecondarybadComment := &NewSecondaryComment{
		AuthorID: bson.NewObjectId(),
		Content:  "Hello",
		PostID:   bson.NewObjectId(),
		ImageURL: "https://google.com",
		Parent:   bson.NewObjectId(),
	}

	_, err = ms.InsertSecondaryComment(newSecondarybadComment)

	if err == nil {
		t.Errorf("Shouldn't be able to add orphan secondary comment")
	}

	newSecondaryComment := &NewSecondaryComment{
		AuthorID: bson.NewObjectId(),
		Content:  "Hello",
		PostID:   bson.NewObjectId(),
		ImageURL: "https://google.com",
		Parent:   comment.ID,
	}

	secondaryComment, err := ms.InsertSecondaryComment(newSecondaryComment)

	_, err = ms.GetBySecondaryID(secondaryComment.ID)

	if err != nil {
		t.Errorf("Testing get by secondary comment failed: %v", err)
	}

	byParent, err := ms.GetByParentID(secondaryComment.Parent)

	if err != nil {
		t.Errorf("Getting by parent id failed: %v", err)
	}

	if len(*byParent) == 0 {
		t.Errorf("Get by post should not be empty")
	}

	// Test Update

	updateComment := &CommentUpdate{
		ImageURL: "http://potato.com",
		Content:  "Tomato",
	}

	_, err = ms.UpdateComment(comment.ID, updateComment)

	if err != nil {
		t.Errorf("update comment failed :%v", err)
	}

	_, err = ms.UpdateComment(secondaryComment.ID, updateComment)

	if err == nil {
		t.Errorf("bad update comment failed :%v", err)
	}

	_, err = ms.UpdateComment("123j12;3j12lk;j1d=2=3d", updateComment)

	if err == nil {
		t.Errorf("bad update failed :%v", err)
	}
	// Test Delete

	err = ms.DeleteComment(secondaryComment.ID)

	if err != nil {
		t.Errorf("Deleting id failed: %v", err)
	}
}

func TestCommentCRUDDeadDB(t *testing.T) {
	conn, err := mgo.Dial("localhost:27017")

	if err != nil {
		t.Errorf("Dialing Mongo Failed: %v", err)
		t.Errorf("%v is a thing", conn)
	}

	ms := NewMongoStore(conn, "test_db", "test_col")

	newComment := &NewComment{
		AuthorID: bson.NewObjectId(),
		Content:  "Hello",
		PostID:   bson.NewObjectId(),
		ImageURL: "https://google.com",
	}

	comment, err := ms.InsertComment(newComment)

	if err != nil {
		t.Errorf("Testing InsertComment Failed: %v", err)
	}

	badNewComment := &NewComment{
		AuthorID: "1231232`-`-`-`-`13",
		Content:  "",
		PostID:   bson.NewObjectId(),
		ImageURL: "",
	}

	_, err = ms.InsertComment(badNewComment)

	if err == nil {
		t.Errorf("Testing InsertComment Failed: %v", err)
	}

	// Try get by comment id
	_, err = ms.GetByCommentID(comment.ID)

	if err != nil {
		t.Errorf("Testing GetByCommentID Failed: %v", err)
	}

	// Try get by comment id
	_, err = ms.GetByCommentID("asdfasdfasdas122-2-121-12")

	if err == nil {
		t.Errorf("Testing GetByCommentID Failed: %v", err)
	}

	// Try get by post
	_, err = ms.GetCommentsByPostID(comment.PostID)

	if err != nil {
		t.Errorf("Testing Get all by post id failed: %v", err)
	}

	newSecondaryComment := &NewSecondaryComment{
		AuthorID: bson.NewObjectId(),
		Content:  "Hello",
		PostID:   bson.NewObjectId(),
		ImageURL: "https://google.com",
		Parent:   comment.ID,
	}

	secondaryComment, err := ms.InsertSecondaryComment(newSecondaryComment)

	if err != nil {
		t.Errorf("Testing insert secondary comment failed: %v", err)
	}

	badNewSecondaryComment := &NewSecondaryComment{
		AuthorID: bson.NewObjectId(),
		Content:  "",
		PostID:   bson.NewObjectId(),
		ImageURL: "",
		Parent:   comment.ID,
	}

	_, err = ms.InsertSecondaryComment(badNewSecondaryComment)

	if err == nil {
		t.Errorf("Testing insert secondary comment failed: %v", err)
	}

	_, err = ms.GetBySecondaryID(secondaryComment.ID)

	if err != nil {
		t.Errorf("Testing get by secondary comment failed: %v", err)
	}

	_, err = ms.GetBySecondaryID("123412h;3kj123n213jkn123==``")

	if err == nil {
		t.Errorf("Testing get by secondary comment failed: %v", err)
	}

	// Update
	secondaryUpdate := &SecondaryCommentUpdate{
		ImageURL: "https://htoeld.com",
		Content:  "Davin",
	}

	_, err = ms.UpdateSecondaryComment(secondaryComment.ID, secondaryUpdate)

	if err != nil {
		t.Errorf("Update failed for secondary : %v ", err)
	}

	_, err = ms.UpdateSecondaryComment(comment.ID, secondaryUpdate)

	if err == nil {
		t.Errorf("Update failed for secondary : %v ", err)
	}

	_, err = ms.UpdateSecondaryComment("12kl;3123h1hi``p`", secondaryUpdate)

	if err == nil {
		t.Errorf("bad Update failed for secondary : %v ", err)
	}

	_, err = ms.GetByParentID(comment.ID)

	if err != nil {
		t.Errorf("Getting by parent id failed: %v", err)
	}

	err = ms.DeleteComment(bson.NewObjectId())

	if err == nil {
		t.Errorf("Deleting id failed: %v", err)
	}
}

func TestErrors(t *testing.T) {
	conn, err := mgo.Dial("localhost:27017")

	if err != nil {
		t.Errorf("Dialing Mongo Failed: %v", err)
		t.Errorf("%v is a thing", conn)
	}

	ms := NewMongoStore(conn, "test_db", "test_col")

	_, err = ms.GetCommentsByPostID("k;l12j3lkj3-3`=31jpj2lk12")

	if err == nil {
		t.Errorf("Get Comments by Post ID Failed:%v", err)
	}

	_, err = ms.GetByParentID("k;l12j3lkj3-3`=31jpj2lk12")

	if err == nil {
		t.Errorf("Get Comments by Post ID Failed:%v", err)
	}
}

// Testing Voiting
func TestVoting(t *testing.T) {

	conn, err := mgo.Dial("localhost:27017")

	if err != nil {
		t.Errorf("Dialing Mongo Failed: %v", err)
		t.Errorf("%v is a thing", conn)
	}

	ms := NewMongoStore(conn, "test_db", "test_col")

	newComment := &NewComment{
		AuthorID: bson.NewObjectId(),
		Content:  "Hello",
		PostID:   bson.NewObjectId(),
		ImageURL: "https://google.com",
	}

	comment, err := ms.InsertComment(newComment)

	newSecondaryComment := &NewSecondaryComment{
		AuthorID: bson.NewObjectId(),
		Content:  "Hello",
		PostID:   bson.NewObjectId(),
		ImageURL: "https://google.com",
		Parent:   comment.ID,
	}

	secondaryComment, err := ms.InsertSecondaryComment(newSecondaryComment)

	commentVote := &CommentVote{
		Upvote:   1,
		Downvote: -1,
	}

	_, err = ms.CommentVote(comment.ID, commentVote)

	if err != nil {
		t.Errorf("Comment Vote Failed %v", err)
	}

	badCommentVote := &CommentVote{
		Upvote:   1,
		Downvote: -1,
	}

	_, err = ms.CommentVote(comment.ID, badCommentVote)

	if err == nil {
		t.Errorf("Comment Vote Failed %v", err)
	}

	_, err = ms.CommentVote("1@#!@#~~~@##@?@?#?#", commentVote)

	if err == nil {
		t.Errorf("bad Comment Vote Failed %v", err)
	}

	secondaryCommentVote := &SecondaryCommentVote{
		Upvote:   0,
		Downvote: 1,
	}

	_, err = ms.SecondaryCommentVote(secondaryComment.ID, secondaryCommentVote)

	if err != nil {
		t.Errorf("Comment Vote Failed %v", err)
	}

	_, err = ms.SecondaryCommentVote("1@#!@#~~~@##@?@?#?#", secondaryCommentVote)

	if err == nil {
		t.Errorf("bad secondary Comment Vote Failed %v", err)
	}

	badsecondaryCommentVote := &SecondaryCommentVote{
		Upvote:   201,
		Downvote: 1,
	}

	_, err = ms.SecondaryCommentVote(secondaryComment.ID, badsecondaryCommentVote)

	if err == nil {
		t.Errorf("Comment Vote Failed %v", err)
	}
}
