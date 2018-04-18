package handlers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
