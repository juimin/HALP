package handlers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCommentsHandler tests the comments handler
func TestCommentsHandlerNoSession(t *testing.T) {
	// Get Context Instance
	cr := prepTestCR()
	// Generate the handlers
	commentsHandler := http.HandlerFunc(cr.CommentsHandler)

	cases := []struct {
		name     string
		endpoint string
		method   string
		body     io.Reader
		handler  http.HandlerFunc
		code     int
	}{
		{
			name:     "Test Comments POST - No Session",
			endpoint: "/comments",
			method:   "POST",
			body: bytes.NewBuffer([]byte(
				`{
				}`)),
			handler: commentsHandler,
			code:    http.StatusForbidden,
		},
		{
			name:     "Test Comments PATCH - No Session",
			endpoint: "/comments",
			method:   "PATCH",
			body: bytes.NewBuffer([]byte(
				`{
				}`)),
			handler: commentsHandler,
			code:    http.StatusForbidden,
		},
		{
			name:     "Test Comments DELETE - No Session",
			endpoint: "/comments",
			method:   "DELETE",
			body: bytes.NewBuffer([]byte(
				`{
				}`)),
			handler: commentsHandler,
			code:    http.StatusForbidden,
		},
		{
			name:     "Test Comments GET - No Session - Bad Request",
			endpoint: "/comments",
			method:   "GET",
			body: bytes.NewBuffer([]byte(
				`{
				}`)),
			handler: commentsHandler,
			code:    http.StatusBadRequest,
		},
		{
			name:     "Test Comments MYSTERY - No Session - Unsupported HTTP Method",
			endpoint: "/comments",
			method:   "OPTIONS",
			body: bytes.NewBuffer([]byte(
				`{
				}`)),
			handler: commentsHandler,
			code:    http.StatusMethodNotAllowed,
		},
	}

	for _, c := range cases {
		// Generate Request
		r, err := http.NewRequest(c.method, c.endpoint, c.body)
		if err != nil {
			t.Errorf("%s Failed: HTTP Request Generation Failed", c.name)
		}
		// Construct a response recorder
		w := httptest.NewRecorder()

		// Attempt Endpoint usage
		c.handler.ServeHTTP(w, r)
		if w.Code != c.code {
			t.Errorf("%s Failed: Expected Code to be %d but got %d", c.name, c.code, w.Code)
		}
	}
}

// Test the comments handler using a set that includes a session
func TestCommentsHandlerSession(t *testing.T) {
	// Get Context Instance
	cr := prepTestCR()
	// Generate the handlers
	// Testing this handler
	commentsHandler := http.HandlerFunc(cr.CommentsHandler)
	// Include users handler for adding a user
	users := http.HandlerFunc(cr.UsersHandler)

	// New User
	newUser := bytes.NewBuffer([]byte(
		`{
			"email":"togurt@gibbera.comr",
			"password": "potatopass",
			"passwordConf": "potatopass",
			"userName": "sigmund",
			"firstName":"firstPotato",
			"lastName": "lastPotato",
			"occupation": "vegetable"
		}`))

	// Generate a new recorder
	temp := httptest.NewRecorder()
	// Generate the request
	req, err := http.NewRequest("POST", "/users", newUser)

	// Session
	authHeader := ""

	// Get the auth header
	if err != nil {
		t.Errorf("Failed Testing on Comments with Sessions - Generating Request")
	} else {
		// Insert the user
		users.ServeHTTP(temp, req)
		if temp.Code != http.StatusCreated {
			t.Errorf("Error inserting user into the database: Expect %d but got %d", http.StatusCreated, temp.Code)
		}
		// Get the auth header
		if temp.Header().Get("Authorization") != "" {
			authHeader = temp.Header().Get("Authorization")
		}

		cases := []struct {
			name     string
			endpoint string
			method   string
			body     io.Reader
			handler  http.HandlerFunc
			code     int
		}{
			{
				name:     "Test Comment Insertion - Primary",
				endpoint: "/comments?type=primary",
				method:   "POST",
				body: bytes.NewBuffer([]byte(
					`{
						"author_id": "507f1f77bcf86cd799439011",
						"content": "I am a potato",
						"post_id": "507f1f77bcf86cd799439011",
						"ImageURL": "www.something.comp"
					}`)),
				handler: commentsHandler,
				code:    http.StatusOK,
			},
			{
				name:     "Test Comment Insertion - Primary Bad Request Body",
				endpoint: "/comments?type=primary",
				method:   "POST",
				body: bytes.NewBuffer([]byte(
					`{
						"author_id": "",
						"content": "I am a potato",
						"post_id": "",
						"ImageURL": "www.something.comp"
					}`)),
				handler: commentsHandler,
				code:    http.StatusBadRequest,
			},
			{
				name:     "Test Comment Insertion - Primary Invalid Body",
				endpoint: "/comments?type=primary",
				method:   "POST",
				body: bytes.NewBuffer([]byte(
					`{
						"author_id": "",
						"content": "I am a potato"
						"post_id": "",
						"ImageURL": "www.something.comp"
					}`)),
				handler: commentsHandler,
				code:    http.StatusBadRequest,
			},
			{
				name:     "Test Comment Insertion - Secondary Invalid Body",
				endpoint: "/comments?type=secondary",
				method:   "POST",
				body: bytes.NewBuffer([]byte(
					`{
						"author_id": "",
						"content": "I am a potato"
						"post_id": "",
						"ImageURL": "www.something.comp"
					}`)),
				handler: commentsHandler,
				code:    http.StatusBadRequest,
			},
			{
				name:     "Test Comment Insertion - Secondary",
				endpoint: "/comments?type=secondary",
				method:   "POST",
				body: bytes.NewBuffer([]byte(
					`{
						"author_id": "507f1f77bcf86cd799439012",
						"content": "I am a potato",
						"post_id": "507f1f77bcf86cd799439012",
						"parent": "507f1f77bcf86cd799439013",
						"ImageURL": "www.something.comp"
					}`)),
				handler: commentsHandler,
				code:    http.StatusOK,
			},
			{
				name:     "Test Comment Insertion - Secondary Bad Request Body",
				endpoint: "/comments?type=secondary",
				method:   "POST",
				body: bytes.NewBuffer([]byte(
					`{
						"author_id": "",
						"content": "I am a potato",
						"post_id": "",
						"ImageURL": "www.something.comp"
					}`)),
				handler: commentsHandler,
				code:    http.StatusBadRequest,
			},
			{
				name:     "Test Comment Insertion - Secondary Unknown Type Query",
				endpoint: "/comments?type=hello",
				method:   "POST",
				body: bytes.NewBuffer([]byte(
					`{
						"author_id": "",
						"content": "I am a potato",
						"post_id": "",
						"ImageURL": "www.something.comp"
					}`)),
				handler: commentsHandler,
				code:    http.StatusBadRequest,
			},
			{
				name:     "Test Comment Retreival - Primary Comment Unknown",
				endpoint: "/comments?id=507f1f77bcf86cd799439012&query=singleComment",
				method:   "GET",
				body:     nil,
				handler:  commentsHandler,
				code:     http.StatusNotFound,
			},
			{
				name:     "Test Comment Retreival - Secondary Comment Unknown",
				endpoint: "/comments?id=507f1f77bcf86cd799439012&query=singeSecondary",
				method:   "GET",
				body:     nil,
				handler:  commentsHandler,
				code:     http.StatusNotFound,
			},
			{
				name:     "Test Comment Retreival - Post Comment Known",
				endpoint: "/comments?id=507f1f77bcf86cd799439012&query=allByPost",
				method:   "GET",
				body:     nil,
				handler:  commentsHandler,
				code:     http.StatusOK,
			},
			{
				name:     "Test Comment Retreival - Post Comment Unknown",
				endpoint: "/comments?id=4123&query=allByPost",
				method:   "GET",
				body:     nil,
				handler:  commentsHandler,
				code:     http.StatusBadRequest,
			},
			{
				name:     "Test Comment Retreival - Parent Comment known",
				endpoint: "/comments?id=507f1f77bcf86cd799439013&query=allByParent",
				method:   "GET",
				body:     nil,
				handler:  commentsHandler,
				code:     http.StatusOK,
			},
			{
				name:     "Test Comment Retreival - Parent Comment Unknown",
				endpoint: "/comments?id=2&query=allByParent",
				method:   "GET",
				body:     nil,
				handler:  commentsHandler,
				code:     http.StatusBadRequest,
			},
			{
				name:     "Test Comment GET bad query",
				endpoint: "/comments?id=407f1f77bcf86cd799439053&query=sss",
				method:   "GET",
				body:     nil,
				handler:  commentsHandler,
				code:     http.StatusBadRequest,
			},
		}

		for _, c := range cases {
			// Generate Request
			r, err := http.NewRequest(c.method, c.endpoint, c.body)

			r.Header.Add("Authorization", authHeader)
			if err != nil {
				t.Errorf("%s Failed: HTTP Request Generation Failed", c.name)
			}
			// Construct a response recorder
			w := httptest.NewRecorder()

			// Attempt Endpoint usage
			c.handler.ServeHTTP(w, r)
			if w.Code != c.code {
				t.Errorf("%s Failed: Expected Code to be %d but got %d", c.name, c.code, w.Code)
				t.Errorf("%s", w.Body)
			}

		}
	}
}

// TestVotesHandler tests dealing with votes
func TestVotesHandler(t *testing.T) {

}
