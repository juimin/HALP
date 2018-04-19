package handlers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var validID = "5ad7f1bf9137241ece23152d"

func TestNewPostHandler(t *testing.T) {
	cr := prepTestCR()
	nph := http.HandlerFunc(cr.NewPostHandler)
	//objectid == 507f1f77bcf86cd799439011

	cases := []struct {
		name         string
		method       string
		expectedCode int
		handler      http.HandlerFunc
		body         io.Reader
		destination  string
	}{
		{
			name:         "invalid method to post handler",
			method:       "GET",
			expectedCode: http.StatusMethodNotAllowed,
			handler:      nph,
			body:         nil,
			destination:  "/posts/new",
		},
		{
			name:         "post to new post handler - nil body",
			method:       "POST",
			expectedCode: http.StatusBadRequest,
			handler:      nph,
			body:         nil,
			destination:  "/posts/new",
		},
		{
			name:         "new post handler - invalid new post",
			method:       "POST",
			expectedCode: http.StatusConflict,
			handler:      nph,
			body: bytes.NewBuffer([]byte(
				`{
					"title": "potatopass",
					"image_url": "potatopass",
					"caption": "potat",
					"author_id": "507f1f77bcf86cd799439011",
					"board_id": "507f1f77bcf86cd799439011"
				}`)),
			destination: "/posts/new",
		},
		{
			name:         "new post handler - valid new post",
			method:       "POST",
			expectedCode: http.StatusCreated,
			handler:      nph,
			body: bytes.NewBuffer([]byte(
				`{
					"title": "potatopass",
					"image_url": "http://google.com",
					"caption": "potat",
					"author_id": "507f1f77bcf86cd799439011",
					"board_id": "507f1f77bcf86cd799439011"
				}`)),
			destination: "/posts/new",
		},
		{
			name:         "new post handler - invalid json",
			method:       "POST",
			expectedCode: http.StatusBadRequest,
			handler:      nph,
			body: bytes.NewBuffer([]byte(
				`{
					"title": "potatopass"FUN,
					"image_url": "http://google.com",
					"caption": "potat",
					"author_id": "507f1f77bcf86cd799439011",
					"board_id": "507f1f77bcf86cd799439011"
				}`)),
			destination: "/posts/new",
		},
	}

	for _, c := range cases {
		recorder := httptest.NewRecorder()
		req, err := http.NewRequest(c.method, c.destination, c.body)
		if err != nil {
			t.Errorf("%s Failed: Error %v", c.name, err)
		} else {
			c.handler.ServeHTTP(recorder, req)
			if recorder.Code != c.expectedCode {
				t.Errorf("%s Failed. Expected %d but got %d", c.name, c.expectedCode, recorder.Code)
			}
		}
	}
}

//Test the post update handler
func TestUpdatePostHandler(t *testing.T) {
	cr := prepTestCR()
	nph := http.HandlerFunc(cr.NewPostHandler)
	uph := http.HandlerFunc(cr.UpdatePostHandler)

	cases := []struct {
		name         string
		method       string
		expectedCode int
		handler      http.HandlerFunc
		body         io.Reader
		destination  string
	}{
		{
			name:         "test invalid method",
			method:       "GET",
			expectedCode: http.StatusMethodNotAllowed,
			handler:      uph,
			destination:  "/posts/update",
		},
		{
			name:         "generate post for testing",
			method:       "POST",
			expectedCode: http.StatusCreated,
			handler:      nph,
			body: bytes.NewBuffer([]byte(
				`{
					"title": "potatopass",
					"image_url": "http://google.com",
					"caption": "potat",
					"author_id": "507f1f77bcf86cd799439011",
					"board_id": "507f1f77bcf86cd799439011"
				}`)),
			destination: "/posts/new",
		},
		{
			name:         "test updating nonexistent post",
			method:       "PATCH",
			expectedCode: http.StatusBadRequest,
			handler:      uph,
			body: bytes.NewBuffer([]byte(
				`{
					"title": "potatopass",
					"image_url": "http://google.com",
					"caption": "",
					"upvotes": {"507f1f77bcf86cd799439011":true},
					"downvotes": {},
					"total_votes": 1
				}`)),
			destination: "/posts/update?id=507f1f77bcf86cd799439011",
		},
		{
			name:         "test updating posts for valid input",
			method:       "PATCH",
			expectedCode: http.StatusAccepted,
			handler:      uph,
			body: bytes.NewBuffer([]byte(
				`{
					"title": "potatopass",
					"image_url": "http://google.com",
					"caption": "",
					"upvotes": {"507f1f77bcf86cd799439011":true},
					"downvotes": {},
					"total_votes": 1
				}`)),
			destination: "/posts/update?id=" + validID,
		},
		{
			name:         "test updating posts for invalid input",
			method:       "PATCH",
			expectedCode: http.StatusInternalServerError,
			handler:      uph,
			body: bytes.NewBuffer([]byte(
				`{
					"title": "potatopass2",
					"image_url": "",
					"caption": "",
					"upvotes": {"507f1f77bcf86cd799439011":true},
					"downvotes": {},
					"total_votes": 1
				}`)),
			destination: "/posts/update?id=" + validID,
		},
		{
			name:         "test updating posts for invalid JSON",
			method:       "PATCH",
			expectedCode: http.StatusBadRequest,
			handler:      uph,
			body: bytes.NewBuffer([]byte(
				`{
					"title": "potatopass2"LoL,
					"image_url": "",
					"caption": "",
					"upvotes": {"507f1f77bcf86cd799439011":true},
					"downvotes": {},
					"total_votes": 1
				}`)),
			destination: "/posts/update?id=5ad7838d9137245e1228435d",
		},
		{
			name:         "test updating posts for invalid id in url",
			method:       "PATCH",
			expectedCode: http.StatusBadRequest,
			handler:      uph,
			body: bytes.NewBuffer([]byte(
				`{
					"title": "potatopass2",
					"image_url": "",
					"caption": "",
					"upvotes": {"507f1f77bcf86cd799439011":true},
					"downvotes": {},
					"total_votes": 1
				}`)),
			destination: "/posts/update?id=5ad7838d913",
		},
		{
			name:         "test updating posts for no id param in url",
			method:       "PATCH",
			expectedCode: http.StatusBadRequest,
			handler:      uph,
			body: bytes.NewBuffer([]byte(
				`{
					"title": "potatopass2",
					"image_url": "",
					"caption": "",
					"upvotes": {"507f1f77bcf86cd799439011":true},
					"downvotes": {},
					"total_votes": 1
				}`)),
			destination: "/posts/update",
		},
		{
			name:         "test updating posts for nil update",
			method:       "PATCH",
			expectedCode: http.StatusBadRequest,
			handler:      uph,
			body:         nil,
			destination:  "/posts/update?id=5ad7838d9137245e1228435d",
		},
	}

	for _, c := range cases {
		recorder := httptest.NewRecorder()
		req, err := http.NewRequest(c.method, c.destination, c.body)
		if err != nil {
			t.Errorf("%s Failed: Error %v", c.name, err)
		} else {
			c.handler.ServeHTTP(recorder, req)
			if recorder.Code != c.expectedCode {
				t.Errorf("%s Failed. Expected %d but got %d", c.name, c.expectedCode, recorder.Code)
			}
		}
	}
}

func TestGetPostHandler(t *testing.T) {
	cr := prepTestCR()
	gph := http.HandlerFunc(cr.GetPostHandler)

	cases := []struct {
		name         string
		method       string
		expectedCode int
		handler      http.HandlerFunc
		body         io.Reader
		destination  string
	}{
		{
			name:         "test getting a nonexistent post",
			method:       "GET",
			expectedCode: http.StatusBadRequest,
			handler:      gph,
			body:         nil,
			destination:  "/posts/get?id=507f1f77bcf86cd799439011",
		},
		{
			name:         "test getting a valid post",
			method:       "GET",
			expectedCode: http.StatusAccepted,
			handler:      gph,
			body:         nil,
			destination:  "/posts/get?id=" + validID,
		},
		{
			name:         "test getting an invalid id",
			method:       "GET",
			expectedCode: http.StatusBadRequest,
			handler:      gph,
			body:         nil,
			destination:  "/posts/get?id=5ad7838d9",
		},
		{
			name:         "test getting an invalid url parameter",
			method:       "GET",
			expectedCode: http.StatusBadRequest,
			handler:      gph,
			body:         nil,
			destination:  "/posts/get?hello=5ad7838d9",
		},
		{
			name:         "test getting a post with invalid method POST",
			method:       "POST",
			expectedCode: http.StatusMethodNotAllowed,
			handler:      gph,
			body:         nil,
			destination:  "/posts/get?id=5ad7838d9137245e1228435d",
		},
	}
	for _, c := range cases {
		recorder := httptest.NewRecorder()
		req, err := http.NewRequest(c.method, c.destination, c.body)
		if err != nil {
			t.Errorf("%s Failed: Error %v", c.name, err)
		} else {
			c.handler.ServeHTTP(recorder, req)
			if recorder.Code != c.expectedCode {
				t.Errorf("%s Failed. Expected %d but got %d", c.name, c.expectedCode, recorder.Code)
			}
		}
	}
}
