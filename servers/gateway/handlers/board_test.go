package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JuiMin/HALP/servers/gateway/models/boards"
)

func TestBoardsAllHandler(t *testing.T) {

	cases := []struct {
		name                string
		reqType             string
		expectedContentType io.Reader
		expectedStatus      int
		expectedOutput      string
		newBoard            io.Reader
	}{
		{
			name:                "Passing Test",
			reqType:             "GET",
			expectedContentType: nil,
			expectedStatus:      http.StatusOK,
			expectedOutput:      "",
			newBoard: bytes.NewBuffer([]byte(
				`{
					"title": "Good test AllBoardsHandler",
					"description": "describe AllBoardsHandler",
					"image": "https://storage.googleapis.com/gweb-uniblog-publish-prod/images/Chrome__logo.max-2800x2800.png"
				}`)),
		},
		{
			name:                "Bad Method",
			reqType:             "POST",
			expectedContentType: nil,
			expectedStatus:      http.StatusMethodNotAllowed,
			expectedOutput:      "",
			newBoard: bytes.NewBuffer([]byte(
				`{
					"title": "Bad Method test AllBoardsHandler",
					"description": "describe AllBoardsHandler",
					"image": "https://storage.googleapis.com/gweb-uniblog-publish-prod/images/Chrome__logo.max-2800x2800.png"
				}`)),
		},
	}

	for _, c := range cases {
		req, err := http.NewRequest(c.reqType, "/boards", c.expectedContentType)
		rr := httptest.NewRecorder()
		cr := prepTestCR()
		httpBoardsAllHandler := http.HandlerFunc(cr.BoardsAllHandler)

		//Create test instance of a board
		newBoard := &boards.NewBoard{}
		decoder := json.NewDecoder(c.newBoard)
		decoder.Decode(newBoard)
		cr.BoardStore.CreateBoard(newBoard)
		if err != nil {
			t.Errorf("AllBoards Test Board not created: %v", err)
		}

		// Error if request could not be made
		if err != nil {
			t.Errorf("%s Failed: Error %v", c.name, err)
		} else {
			httpBoardsAllHandler.ServeHTTP(rr, req)
			if rr.Code != c.expectedStatus {
				t.Errorf("%s Failed. Expected %d but got %d.", c.name, c.expectedStatus, rr.Code)
			}
		}
	}
}

func TestSingleBoardHandler(t *testing.T) {

	//Create test instance of a board
	cr := prepTestCR()
	board, err := cr.BoardStore.GetAllBoards()
	if err != nil {
		t.Errorf("Board not created: %v", err)
	}

	cases := []struct {
		name                string
		reqType             string
		expectedContentType io.Reader
		destinationURL      string
		expectedStatus      int
		testBoard           io.Reader
	}{
		{
			name:                "Passing Test",
			reqType:             "GET",
			expectedContentType: nil,
			destinationURL:      "/boards/single?id=" + board[0].ID.Hex(),
			expectedStatus:      http.StatusOK,
		},
		{
			name:                "Bad Request / Bad ID",
			reqType:             "GET",
			expectedContentType: nil,
			destinationURL:      "/boards/single/?id=ffdddfsfwefwecewcwec",
			expectedStatus:      http.StatusBadRequest,
		},
		{
			name:                "Bad Request Method",
			reqType:             "DELETE",
			expectedContentType: nil,
			destinationURL:      "/boards/single?id=" + board[0].ID.Hex(),
			expectedStatus:      http.StatusMethodNotAllowed,
		},
	}

	for _, c := range cases {
		req, err := http.NewRequest(c.reqType, c.destinationURL, c.expectedContentType)
		rr := httptest.NewRecorder()
		httpSingleBoardHandler := http.HandlerFunc(cr.SingleBoardHandler)

		// Error if request could not be made
		if err != nil {
			t.Errorf("%s Failed: Error %v", c.name, err)
		} else {
			httpSingleBoardHandler.ServeHTTP(rr, req)
			if rr.Code != c.expectedStatus {
				t.Errorf("%s Failed. Expected %d but got %d.", c.name, c.expectedStatus, rr.Code)
			}
		}
	}
}

func TestUpdatePostCountHandler(t *testing.T) {

	//Create test instance of a board
	cr := prepTestCR()

	// Insert a board into the cr
	_, err := cr.BoardStore.CreateBoard(&boards.NewBoard{
		Title:       "Davin",
		Description: "Test Desc",
		Image:       "https://potato.com",
	})

	if err != nil {
		t.Errorf("Boards not added")
	}

	board, err := cr.BoardStore.GetAllBoards()
	if err != nil {
		t.Errorf("Board not created: %v", err)
	}

	cases := []struct {
		name           string
		body           io.Reader
		reqType        string
		destinationURL string
		expectedStatus int
	}{
		{
			name:           "Passing Test",
			body:           bytes.NewBuffer([]byte(`{"temp": true}`)),
			reqType:        "PATCH",
			destinationURL: "/boards/updatepost?id=" + board[0].ID.Hex(),
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Passing Test",
			body:           bytes.NewBuffer([]byte(`{"temp": true}`)),
			reqType:        "GET",
			destinationURL: "/boards/updatepost?id=" + board[0].ID.Hex(),
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Passing Test",
			body:           nil,
			reqType:        "PATCH",
			destinationURL: "/boards/updatepost?id=" + board[0].ID.Hex(),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid JSON Test",
			body:           bytes.NewBuffer([]byte(`{"temp" true}`)),
			reqType:        "PATCH",
			destinationURL: "/boards/updatepost?id=" + board[0].ID.Hex(),
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, c := range cases {
		req, err := http.NewRequest(c.reqType, c.destinationURL, c.body)
		rr := httptest.NewRecorder()
		httpUpdatePostHandler := http.HandlerFunc(cr.UpdatePostCountHandler)
		// Error if request could not be made
		if err != nil {
			t.Errorf("%s Failed: Error %v", c.name, err)
		} else {
			httpUpdatePostHandler.ServeHTTP(rr, req)
			if rr.Code != c.expectedStatus {
				t.Errorf("%s Failed. Expected %d but got %d.", c.name, c.expectedStatus, rr.Code)
			}
		}
	}
}

func TestUpdateSubscriberCountHandler(t *testing.T) {
	//Create test instance of a board
	cr := prepTestCR()
	board, err := cr.BoardStore.GetAllBoards()
	if err != nil {
		t.Errorf("Board not created: %v", err)
	}

	cases := []struct {
		name           string
		body           io.Reader
		reqType        string
		destinationURL string
		expectedStatus int
	}{
		{
			name:           "Passing Test",
			body:           bytes.NewBuffer([]byte(`{"temp": false}`)),
			reqType:        "PATCH",
			destinationURL: "/boards/updatesubscriber?id=" + board[0].ID.Hex(),
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Passing Test",
			body:           bytes.NewBuffer([]byte(`{"temp": true}`)),
			reqType:        "GET",
			destinationURL: "/boards/updatepost?id=" + board[0].ID.Hex(),
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Passing Test",
			body:           nil,
			reqType:        "PATCH",
			destinationURL: "/boards/updatepost?id=" + board[0].ID.Hex(),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid JSON Test",
			body:           bytes.NewBuffer([]byte(`{"temp" false}`)),
			reqType:        "PATCH",
			destinationURL: "/boards/updatesubscriber?id=" + board[0].ID.Hex(),
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, c := range cases {
		req, err := http.NewRequest(c.reqType, c.destinationURL, c.body)
		rr := httptest.NewRecorder()
		httpUpdateSubscriberHandler := http.HandlerFunc(cr.UpdateSubscriberCountHandler)
		// Error if request could not be made
		if err != nil {
			t.Errorf("%s Failed: Error %v", c.name, err)
		} else {
			httpUpdateSubscriberHandler.ServeHTTP(rr, req)
			if rr.Code != c.expectedStatus {
				t.Errorf("%s Failed. Expected %d but got %d.", c.name, c.expectedStatus, rr.Code)
			}
		}
	}
}

func TestCreateBoardHandler(t *testing.T) {
	cases := []struct {
		name           string
		destinationURL string
		reqType        string
		newBody        io.Reader
		expectedStatus int
	}{
		{
			name:           "Test Proper Board",
			destinationURL: "/boards/createboard",
			reqType:        "POST",
			newBody: bytes.NewBuffer([]byte(
				`{
				"title": "Test Proper Board",
				"description": "This is a test board. Do what you will!",
				"image": "https://storage.googleapis.com/gweb-uniblog-publish-prod/images/Chrome__logo.max-2800x2800.png"
			}`)),
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "Bad Url Test",
			destinationURL: "/boards/createboard",
			reqType:        "POST",
			newBody: bytes.NewBuffer([]byte(
				`{
				"title": "Test Bad URL",
				"description": "This is a test board. Do what you will!",
				"image": "good image url"
			}`)),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "No Title Test",
			destinationURL: "/boards/createboard",
			reqType:        "POST",
			newBody: bytes.NewBuffer([]byte(
				`{
				"title": "",
				"description": "Test No Titlte",
				"image": "https://storage.googleapis.com/gweb-uniblog-publish-prod/images/Chrome__logo.max-2800x2800.png"
			}`)),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "No Description Test",
			destinationURL: "/boards/createboard",
			reqType:        "POST",
			newBody: bytes.NewBuffer([]byte(
				`{
				"title": "Test No Description",
				"description": "",
				"image": "https://storage.googleapis.com/gweb-uniblog-publish-prod/images/Chrome__logo.max-2800x2800.png"
			}`)),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Bad Method Test",
			destinationURL: "/boards/createboard",
			reqType:        "GET",
			newBody: bytes.NewBuffer([]byte(
				`{
				"title": "Test No Description",
				"description": "",
				"image": "https://storage.googleapis.com/gweb-uniblog-publish-prod/images/Chrome__logo.max-2800x2800.png"
			}`)),
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "No Body Test",
			destinationURL: "/boards/createboard",
			reqType:        "POST",
			newBody:        nil,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Bad Body Test",
			destinationURL: "/boards/createboard",
			reqType:        "POST",
			newBody: bytes.NewBuffer([]byte(
				`{
				"image" "https://storage.googleapis.com/gweb-uniblog-publish-prod/images/Chrome__logo.max-2800x2800.png"
			}`)),
			expectedStatus: http.StatusBadRequest,
		},
	}
	for _, c := range cases {
		cr := prepTestCR()
		req, err := http.NewRequest(c.reqType, c.destinationURL, c.newBody)
		rr := httptest.NewRecorder()
		httpCreateBoardHandler := http.HandlerFunc(cr.CreateBoardHandler)
		// Error if request could not be made
		if err != nil {
			t.Errorf("%s Failed: Error %v", c.name, err)
		} else {
			httpCreateBoardHandler.ServeHTTP(rr, req)
			if rr.Code != c.expectedStatus {
				t.Errorf("%s Failed. Expected %d but got %d.", c.name, c.expectedStatus, rr.Code)
			}
		}
	}
}
