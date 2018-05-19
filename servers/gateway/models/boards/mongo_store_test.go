package boards

import (
	"testing"

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

func TestCRUDValidInput(t *testing.T) {

	conn, err := mgo.Dial("localhost:27017")

	if err != nil {
		t.Errorf("Dialing Mongo Failed: %v", err)
	}

	ms := NewMongoStore(conn, "test_db_board", "test_col_board")

	nb := &NewBoard{
		Title:       "test tile",
		Description: "Testing more stuff",
		Image:       "https://gpdfasdf.goc",
	}

	board, err := ms.CreateBoard(nb)

	if err != nil {
		t.Errorf("PRoblems creating a board")
	}

	allboards, err := ms.GetAllBoards()

	if err != nil {
		t.Errorf("Problem getting all boards")
	}

	if len(allboards) == 0 {
		t.Errorf("There should be more than one board")
	}

	if allboards[0].ID != board.ID {
		t.Errorf("This should be the only id in the baords")
	}

	byID, err := ms.GetByID(board.ID)

	if err != nil {
		t.Errorf("Problem getting board by id")
	}

	if byID.ID != board.ID {
		t.Errorf("Get By ID got the wrong object")
	}

	byTitle, err := ms.GetByBoardName(board.Title)

	if err != nil {
		t.Errorf("Problem getting board by title")
	}

	if byTitle.ID != board.ID {
		t.Errorf("Get By title got the wrong object")
	}

	// Update the stuff
	subUpdate := &UpdateSubscriber{
		Subscribers: 1,
	}

	postUpdate := &UpdatePost{
		Posts: 1,
	}

	_, err = ms.UpdateSubscriberCount(board.ID, subUpdate)

	if err != nil {
		t.Errorf("Something went wrong updating the bsub")
	}

	_, err = ms.UpdatePostCount(board.ID, postUpdate)

	if err != nil {
		t.Errorf("Something went wrong updating the bsub")
	}

}

// Test for the errors when getting stuff in crud
func TestCRUDERrors(t *testing.T) {

	conn, err := mgo.Dial("localhost:27017")

	if err != nil {
		t.Errorf("Dialing Mongo Failed: %v", err)
	}

	ms := NewMongoStore(conn, "test_db_board", "test_col_board")

	nb := &NewBoard{
		Title:       "",
		Description: "Testing more stuff",
		Image:       "https://gpdfasdf.goc",
	}

	_, err = nb.ToBoard()

	if err == nil {
		t.Errorf("This should fail because no title")
	}

	_, err = ms.GetByBoardName("")

	if err == nil {
		t.Errorf("Failed Error on get by title, got %v but expected nil", err)
	}

	// Bad ID
	_, err = ms.GetByID("!#!@#??@?@>@>!@#>!@#")

	if err == nil {
		t.Errorf("Failed Error on get by title, got %v but expected nil", err)
	}

}
