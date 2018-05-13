package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/JuiMin/HALP/servers/gateway/models/boards"
	"github.com/JuiMin/HALP/servers/gateway/models/posts"
	"github.com/JuiMin/HALP/servers/gateway/models/users"
)

func TestNewPostHandler(t *testing.T) {
	cr := prepTestCR()
	nph := http.HandlerFunc(cr.NewPostHandler)
	//objectid == 507f1f77bcf86cd799439011

	// Insert a board
	board, err := cr.BoardStore.CreateBoard(&boards.NewBoard{
		Title:       "Potao",
		Description: "Poddddaadfda",
		Image:       "https://github.com",
	})

	if err != nil {
		t.Errorf("Problem adding board")
	}

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
					"board_id": "` + board.ID.Hex() + `"
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
					"board_id": "` + board.ID.Hex() + `"
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

	// Insert a board
	board, err := cr.BoardStore.CreateBoard(&boards.NewBoard{
		Title:       "Potao",
		Description: "Poddddaadfda",
		Image:       "https://github.com",
	})

	if err != nil {
		t.Errorf("Problem adding board")
	}

	// Include users handler for adding a user
	usersHandler := http.HandlerFunc(cr.UsersHandler)

	// New User
	newUser := bytes.NewBuffer([]byte(
		`{
				"email":"deepgold@gibbera.comr",
				"password": "potatopass",
				"passwordConf": "potatopass",
				"userName": "undfddfd",
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

	if err != nil {
		t.Errorf("Error generating request %v", err)
	} else {
		// Insert the user
		usersHandler.ServeHTTP(temp, req)
		if temp.Code != http.StatusCreated {
			t.Errorf("Error inserting user into the database: Expect %d but got %d", http.StatusCreated, temp.Code)
		}
		// Get the auth header
		if temp.Header().Get("Authorization") != "" {
			authHeader = temp.Header().Get("Authorization")
		}

		// Unload reponse with the user info
		user := &users.User{}

		err := json.NewDecoder(temp.Body).Decode(user)

		if err != nil {
			t.Errorf("User Broken")
		}

		if authHeader != "" {
			newPost := &posts.NewPost{
				Title:    "PP",
				ImageURL: "https://google.com",
				Caption:  "Hello there",
				AuthorID: user.ID,
				BoardID:  bson.NewObjectId(),
			}

			nph := http.HandlerFunc(cr.NewPostHandler)
			uph := http.HandlerFunc(cr.UpdatePostHandler)

			// Insert a test Post
			insertedPost, err := cr.PostStore.Insert(newPost)

			if err != nil {
				t.Errorf("Error inserting a test post")
			}

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
					"board_id": "` + board.ID.Hex() + `"
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
					"total_votes": 1
				}`)),
					destination: "/posts/update?id=" + insertedPost.ID.Hex(),
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
					"total_votes": 1
				}`)),
					destination: "/posts/update?id=" + insertedPost.ID.Hex(),
				},
				{
					name:         "test updating posts for invalid JSON",
					method:       "PATCH",
					expectedCode: http.StatusBadRequest,
					handler:      uph,
					body: bytes.NewBuffer([]byte(
						`{
					"title": "potatopass2"LoL,
					"image_udfrl": "",
					"captionasdfa": "",
					"upvotes": {"507f1f77bcf86cd799439011":true},
					"downvotes": {},
					"total_votes": 1
				}`)),
					destination: "/posts/update?id=" + insertedPost.ID.Hex(),
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
					"total_votes": 1
				}`)),
					destination: "/posts/update?id=5ad@?#?!@7838d913",
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
					destination:  "/posts/update?id=" + insertedPost.ID.Hex(),
				},
			}

			for _, c := range cases {
				recorder := httptest.NewRecorder()
				req, err := http.NewRequest(c.method, c.destination, c.body)
				// Add auth
				req.Header.Add("Authorization", authHeader)
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
	}
}

func TestGetPostHandler(t *testing.T) {

	newPost := &posts.NewPost{
		Title:    "PP",
		ImageURL: "https://google.com",
		Caption:  "Hello there",
		AuthorID: bson.NewObjectId(),
		BoardID:  bson.NewObjectId(),
	}

	cr := prepTestCR()

	// Insert a test Post
	insertedPost, err := cr.PostStore.Insert(newPost)

	if err != nil {
		t.Errorf("Error inserting a test post")
	}

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
			destination:  "/posts/get?id=" + insertedPost.ID.Hex(),
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
			method:       "PATCH",
			expectedCode: http.StatusMethodNotAllowed,
			handler:      gph,
			body:         nil,
			destination:  "/posts/get?id=5ad7838d9137245e1228435d",
		},
	}
	//CHANGE INVALID METHOD ABOVE to prevent test caching
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
