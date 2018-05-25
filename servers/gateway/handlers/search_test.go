package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JuiMin/HALP/servers/gateway/models/boards"
	"github.com/JuiMin/HALP/servers/gateway/models/comments"
	"github.com/JuiMin/HALP/servers/gateway/models/posts"
	"github.com/JuiMin/HALP/servers/gateway/models/users"
	"gopkg.in/mgo.v2/bson"
)

func TestRespond(t *testing.T) {
	rr := httptest.NewRecorder()
	respond(rr, &struct{}{})
	respond(rr, nil)
}

func TestSearchHandler(t *testing.T) {
	// Get Context Instance
	cr := prepTestCR()
	// Generate the handlers
	// Testing this handler
	searchHandler := http.HandlerFunc(cr.SearchHandler)
	// Include users handler for adding a user
	userHandler := http.HandlerFunc(cr.UsersHandler)

	// New User
	newUser := bytes.NewBuffer([]byte(
		`{
			"email":"gogurt@gibbera.dds",
			"password": "potatopass",
			"passwordConf": "potatopass",
			"userName": "tomatopotatoasdf",
			"firstName":"df",
			"lastName": "asdfdddds",
			"occupation": "vegetable"
		}`))

	// Insert a board
	board, err := cr.BoardStore.CreateBoard(&boards.NewBoard{
		Title:       "Potao",
		Description: "Poddddaadfda",
		Image:       "https://github.com",
	})

	if err != nil {
		t.Errorf("Problem adding board")
	}

	// Add board to trie
	err = cr.BoardTrie.Insert(board.Title, board.ID, 0)

	if err != nil {
		t.Errorf("Problem adding board")
	}

	// Insert a post
	post, err := cr.PostStore.Insert(&posts.NewPost{
		Title:    "somehig",
		ImageURL: "https://asdfadf.com",
		AuthorID: bson.NewObjectId(),
		BoardID:  board.ID,
		Caption:  "asdfasdf",
	})

	if err != nil {
		t.Errorf("Problem adding post")
	}

	cr.PostTrie.Insert(post.Title, post.ID, 0)

	if err != nil {
		t.Errorf("Problem adding post")
	}

	_, err = cr.CommentStore.InsertComment(&comments.NewComment{
		AuthorID: bson.NewObjectId(),
		Content:  "asdfadfasdf",
		PostID:   post.ID,
		ImageURL: "https://werr.comf",
	})

	if err != nil {
		t.Errorf("Problem adding comment seed")
	}

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
		userHandler.ServeHTTP(temp, req)
		if temp.Code != http.StatusCreated {
			t.Errorf("Error inserting user into the database: Expect %d but got %d", http.StatusCreated, temp.Code)
		}
		// Get the auth header
		if temp.Header().Get("Authorization") != "" {
			authHeader = temp.Header().Get("Authorization")
		}

		user := &users.User{}
		err := json.NewDecoder(temp.Body).Decode(user)

		if err != nil {
			t.Errorf("%v", err)
		} else {

			err = cr.UserTrie.Insert(user.UserName, user.ID, 0)

			if err != nil {
				t.Errorf("%v", err)
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
					name:     "Test search board",
					endpoint: "/Search?type=BOARD&search=p",
					method:   "GET",
					body:     nil,
					handler:  searchHandler,
					code:     http.StatusOK,
				},
				{
					name:     "Test search user",
					endpoint: "/Search?type=USER&search=t",
					method:   "GET",
					body:     nil,
					handler:  searchHandler,
					code:     http.StatusOK,
				},
				{
					name:     "Test search post",
					endpoint: "/Search?type=POST&search=som",
					method:   "GET",
					body:     nil,
					handler:  searchHandler,
					code:     http.StatusOK,
				},
				{
					name:     "Test no search term",
					endpoint: "/Search?type=BOARD",
					method:   "GET",
					body:     nil,
					handler:  searchHandler,
					code:     http.StatusBadRequest,
				},
				{
					name:     "Test no type term",
					endpoint: "/Search?search=BOARD",
					method:   "GET",
					body:     nil,
					handler:  searchHandler,
					code:     http.StatusBadRequest,
				},
				{
					name:     "Test post search term",
					endpoint: "/Search?type=POST&search=BOARD",
					method:   "GET",
					body:     nil,
					handler:  searchHandler,
					code:     http.StatusInternalServerError,
				},
				{
					name:     "Test user search term",
					endpoint: "/Search?type=USER&search=BOARD",
					method:   "GET",
					body:     nil,
					handler:  searchHandler,
					code:     http.StatusInternalServerError,
				},
				{
					name:     "Test search unknown type",
					endpoint: "/Search?type=asdf&search=BOARD",
					method:   "GET",
					body:     nil,
					handler:  searchHandler,
					code:     http.StatusBadRequest,
				},
				{
					name:     "Test search bad method",
					endpoint: "/Search?type=asdf&search=BOARD",
					method:   "POTATO",
					body:     nil,
					handler:  searchHandler,
					code:     http.StatusMethodNotAllowed,
				},
				{
					name:     "Test search board not found",
					endpoint: "/Search?type=BOARD&search=123123123?ASDF",
					method:   "GET",
					body:     nil,
					handler:  searchHandler,
					code:     http.StatusInternalServerError,
				},
				{
					name:     "Test search user not found",
					endpoint: "/Search?type=USER&search=fflfeflwefjakl;dfja;lsdf?",
					method:   "GET",
					body:     nil,
					handler:  searchHandler,
					code:     http.StatusInternalServerError,
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

			// Test with no auth
			// Generate Request
			r, err := http.NewRequest("GET", "/search", nil)

			if err != nil {
				t.Errorf("%s Failed: HTTP Request Generation Failed for no auth input", "No auth input")
			}
			// Construct a response recorder
			w := httptest.NewRecorder()

			// Attempt Endpoint usage
			searchHandler.ServeHTTP(w, r)
			if w.Code != http.StatusBadRequest {
				t.Errorf("%s Failed: Expected Code to be %d but got %d", "No Auth Test", http.StatusBadRequest, w.Code)
				t.Errorf("%s", w.Body)
			}

		}
	}
}
