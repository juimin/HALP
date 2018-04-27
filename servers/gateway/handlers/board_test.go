package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBoardsAllHandler(t *testing.T) {
	cases := []struct {
		name                string
		expectedContentType io.Reader
		reqType             string
		expectedStatus      int
		expectedOutput      string
	}{
		{
			name:                "Passing Test",
			expectedContentType: nil,
			reqType:             "GET",
			expectedStatus:      http.StatusOK,
			expectedOutput:      "",
		},
		{
			name:                "Should Fail",
			expectedContentType: nil,
			reqType:             "POST",
			expectedStatus:      http.StatusMethodNotAllowed,
			expectedOutput:      "",
		},
	}

	for _, c := range cases {
		req, err := http.NewRequest(c.reqType, "/boards", c.expectedContentType)
		rr := httptest.NewRecorder()
		cr := prepTestCR()
		httpBoardsAllHandler := http.HandlerFunc(cr.BoardsAllHandler)

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
	board, err := cr.BoardStore.CreateBoard()
	if err != nil {
		t.Errorf("Board not created: %v", err)
	}

	cases := []struct {
		name                string
		expectedContentType io.Reader
		reqType             string
		destinationURL      string
		expectedStatus      int
	}{
		{
			name:                "Passing Test",
			expectedContentType: nil,
			reqType:             "GET",
			destinationURL:      "/boards/single?id=" + board.ID.Hex(),
			expectedStatus:      http.StatusOK,
		},
		{
			name:                "Should Fail",
			expectedContentType: nil,
			reqType:             "GET",
			destinationURL:      "/boards/single/?id=ffdddfsfwefwecewcwec",
			expectedStatus:      http.StatusBadRequest,
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
				t.Errorf("This is the id: %s", c.destinationURL)
				t.Errorf(rr.Body.String())
			}
		}
	}
}

func TestUpdatePostHandler(t *testing.T) {

}

func TestUpdateSubscriberHandler(t *testing.T) {

}
