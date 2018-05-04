package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JuiMin/HALP/servers/gateway/models/posts"

	"github.com/JuiMin/HALP/servers/gateway/models/boards"

	"gopkg.in/mgo.v2/bson"

	"github.com/JuiMin/HALP/servers/gateway/models/comments"
	"github.com/JuiMin/HALP/servers/gateway/models/users"
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
	userHandler := http.HandlerFunc(cr.UsersHandler)

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

	// Insert a board
	board, err := cr.BoardStore.CreateBoard(&boards.NewBoard{
		Title:       "Potao",
		Description: "Poddddaadfda",
		Image:       "https://github.com",
	})

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

	seedComment, err := cr.CommentStore.InsertComment(&comments.NewComment{
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

		// Insert Some Comments

		comment, err := cr.CommentStore.InsertComment(&comments.NewComment{
			AuthorID: user.ID,
			Content:  "test content",
			PostID:   bson.NewObjectId(),
			ImageURL: "https://github.com",
		})

		if err != nil {
			t.Errorf("Error inserting comment")
		}

		mysteryComment, err := cr.CommentStore.InsertComment(&comments.NewComment{
			AuthorID: bson.NewObjectId(),
			Content:  "test content",
			PostID:   bson.NewObjectId(),
			ImageURL: "https://github.com",
		})

		if err != nil {
			t.Errorf("Error inserting comment")
		}

		secondaryComment, err := cr.CommentStore.InsertSecondaryComment(&comments.NewSecondaryComment{
			AuthorID: user.ID,
			Content:  "test content",
			PostID:   bson.NewObjectId(),
			ImageURL: "https://github.com",
			Parent:   comment.ID,
		})

		if err != nil {
			t.Errorf("Error inserting comment")
		}

		mysterySecondaryComment, err := cr.CommentStore.InsertSecondaryComment(&comments.NewSecondaryComment{
			AuthorID: bson.NewObjectId(),
			Content:  "test content",
			PostID:   bson.NewObjectId(),
			ImageURL: "https://github.com",
			Parent:   comment.ID,
		})

		if err != nil {
			t.Errorf("Error inserting comment")
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
						"post_id": "` + post.ID.Hex() + `",
						"image_url": "www.something.comp"
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
						"image_url": "www.something.comp"
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
						"image_url": "www.something.comp"
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
						"image_url": "www.something.comp"
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
						"post_id": "` + post.ID.Hex() + `",
						"parent": "` + seedComment.ID.Hex() + `",
						"image_url": "www.something.comp"
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
						"image_url": "www.something.comp"
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
						"image_url": "www.something.comp"
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
			{
				name:     "Test Comment GET good",
				endpoint: "/comments?id=" + comment.ID.Hex() + "&query=singleComment",
				method:   "GET",
				body:     nil,
				handler:  commentsHandler,
				code:     http.StatusOK,
			},
			{
				name:     "Test SecondaryComment GET good",
				endpoint: "/comments?id=" + secondaryComment.ID.Hex() + "&query=singeSecondary",
				method:   "GET",
				body:     nil,
				handler:  commentsHandler,
				code:     http.StatusOK,
			},
			{
				name:     "Test SecondaryComment PATCH good",
				endpoint: "/comments?id=" + secondaryComment.ID.Hex() + "&type=secondary",
				method:   "PATCH",
				body: bytes.NewBuffer([]byte(
					`{
						"image_url": "https://githubd.com",
						"content": "I am a potato"
					}`)),
				handler: commentsHandler,
				code:    http.StatusOK,
			},
			{
				name:     "Test SecondaryComment PATCH good forbidden",
				endpoint: "/comments?id=" + mysterySecondaryComment.ID.Hex() + "&type=secondary",
				method:   "PATCH",
				body: bytes.NewBuffer([]byte(
					`{
						"image_url": "https://githubd.com",
						"content": "I am a potato"
					}`)),
				handler: commentsHandler,
				code:    http.StatusForbidden,
			},
			{
				name:     "Test SecondaryComment PATCH to not asecondary comment",
				endpoint: "/comments?id=" + comment.ID.Hex() + "&type=secondary",
				method:   "PATCH",
				body: bytes.NewBuffer([]byte(
					`{
						"image_url": "https://githubd.com",
						"content": "I am a potato"
					}`)),
				handler: commentsHandler,
				code:    http.StatusBadRequest,
			},
			{
				name:     "Test comment PATCH good request",
				endpoint: "/comments?id=" + comment.ID.Hex() + "&type=primary",
				method:   "PATCH",
				body: bytes.NewBuffer([]byte(
					`{
						"image_url": "https://githubd.com",
						"content": "I am a potato"
					}`)),
				handler: commentsHandler,
				code:    http.StatusOK,
			},
			{
				name:     "Test comment PATCH good request not allowed",
				endpoint: "/comments?id=" + mysteryComment.ID.Hex() + "&type=primary",
				method:   "PATCH",
				body: bytes.NewBuffer([]byte(
					`{
						"image_url": "https://githubd.com",
						"content": "I am a potato",
						"comments": ["` + secondaryComment.ID.Hex() + `"]
					}`)),
				handler: commentsHandler,
				code:    http.StatusForbidden,
			},
			{
				name:     "Test comment PATCH sub secondary for primary",
				endpoint: "/comments?id=" + secondaryComment.ID.Hex() + "&type=primary",
				method:   "PATCH",
				body: bytes.NewBuffer([]byte(
					`{
						"image_url": "https://githubd.com",
						"content": "I am a potato",
						"comments": ["` + secondaryComment.ID.Hex() + `"]
					}`)),
				handler: commentsHandler,
				code:    http.StatusBadRequest,
			},
			{
				name:     "Test comment PATCH bad body",
				endpoint: "/comments?id=" + secondaryComment.ID.Hex() + "&type=primary",
				method:   "PATCH",
				body: bytes.NewBuffer([]byte(
					`{
						"image_url": "https://githubd.com",
						"content": "I am a potato",,,
						"comments": ["` + secondaryComment.ID.Hex() + `"]
					}`)),
				handler: commentsHandler,
				code:    http.StatusBadRequest,
			},
			{
				name:     "Test comment PATCH bad body",
				endpoint: "/comments?id=" + secondaryComment.ID.Hex() + "&type=secondary",
				method:   "PATCH",
				body: bytes.NewBuffer([]byte(
					`{
						"image_url": "https://githubd.com",
						"content": "I am a potato",,,
						"comments": ["` + secondaryComment.ID.Hex() + `"]
					}`)),
				handler: commentsHandler,
				code:    http.StatusBadRequest,
			},
			{
				name:     "Test comment PATCH bad type",
				endpoint: "/comments?id=" + secondaryComment.ID.Hex() + "&type=asdfasdf",
				method:   "PATCH",
				body: bytes.NewBuffer([]byte(
					`{
						"image_url": "https://githubd.com",
						"content": "I am a potato",,,
						"comments": ["` + secondaryComment.ID.Hex() + `"]
					}`)),
				handler: commentsHandler,
				code:    http.StatusBadRequest,
			},
			{
				name:     "Test comment PATCH delete",
				endpoint: "/comments?id=" + secondaryComment.ID.Hex(),
				method:   "DELETE",
				body:     nil,
				handler:  commentsHandler,
				code:     http.StatusOK,
			},
			{
				name:     "Test comment PATCH repeat delete",
				endpoint: "/comments?id=" + secondaryComment.ID.Hex(),
				method:   "DELETE",
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
	// Get Context Instance
	cr := prepTestCR()
	// Generate the handlers
	// Testing this handler

	votesHandler := http.HandlerFunc(cr.VotesHandler)
	// Include users handler for adding a user
	usersHandler := http.HandlerFunc(cr.UsersHandler)

	// New User
	newUser := bytes.NewBuffer([]byte(
		`{
			"email":"asdfasdf@gibbera.comr",
			"password": "potatopass",
			"passwordConf": "potatopass",
			"userName": "sigmundfddfd",
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

	// Try bad session
	r, err := http.NewRequest("PATCH", "/sdfadsf/", nil)

	r.Header.Add("Authorization", authHeader)
	if err != nil {
		t.Errorf("%s Failed: HTTP Request Generation Failed", "No auth test")
	}
	// Construct a response recorder
	w := httptest.NewRecorder()

	// Attempt Endpoint usage
	votesHandler.ServeHTTP(w, r)
	if w.Code != http.StatusForbidden {
		t.Errorf("%s Failed: Expected Code to be %d but got %d", "No auth test", http.StatusForbidden, w.Code)
		t.Errorf("%s", w.Body)
	}

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

		if authHeader != "" {
			user := &users.User{}

			err := json.NewDecoder(temp.Body).Decode(user)

			if err != nil {
				t.Errorf("User Broken")
			}

			comment, err := cr.CommentStore.InsertComment(&comments.NewComment{
				AuthorID: user.ID,
				Content:  "test content",
				PostID:   bson.NewObjectId(),
				ImageURL: "https://github.com",
			})

			if err != nil {
				t.Errorf("Error inserting comment")
			}

			secondaryComment, err := cr.CommentStore.InsertSecondaryComment(&comments.NewSecondaryComment{
				AuthorID: user.ID,
				Content:  "test content",
				PostID:   bson.NewObjectId(),
				ImageURL: "https://github.com",
				Parent:   comment.ID,
			})

			if err != nil {
				t.Errorf("Error inserting comment")
			}

			if err == nil {
				cases := []struct {
					name     string
					endpoint string
					method   string
					body     io.Reader
					handler  http.HandlerFunc
					code     int
				}{
					{
						name:     "Test Vote wrong method",
						endpoint: "/vote",
						method:   "GET",
						body:     nil,
						handler:  votesHandler,
						code:     http.StatusMethodNotAllowed,
					},
					{
						name:     "Test Vote wrong method" + secondaryComment.Parent.String(),
						endpoint: "/vote",
						method:   "DELETE",
						body:     nil,
						handler:  votesHandler,
						code:     http.StatusMethodNotAllowed,
					},
					{
						name:     "Test Vote test bad type",
						endpoint: "/vote?type=234324",
						method:   "PATCH",
						body:     nil,
						handler:  votesHandler,
						code:     http.StatusBadRequest,
					},
					{
						name:     "Test Vote test good request",
						endpoint: "/vote?type=primary&id=" + comment.ID.Hex(),
						method:   "PATCH",
						body: bytes.NewBuffer([]byte(
							`{
								"upvote": 1,
								"downvote": 0
							}`)),
						handler: votesHandler,
						code:    http.StatusOK,
					},
					{
						name:     "Test Vote test bad body",
						endpoint: "/vote?type=primary&id=" + comment.ID.Hex(),
						method:   "PATCH",
						body: bytes.NewBuffer([]byte(
							`{
								"upvote": 1
								"downvote": 0
							}`)),
						handler: votesHandler,
						code:    http.StatusBadRequest,
					},
					{
						name:     "Test Vote test bad request",
						endpoint: "/vote?type=primary&id=" + comment.ID.Hex(),
						method:   "PATCH",
						body: bytes.NewBuffer([]byte(
							`{
								"upvote": 133,
								"downvote": 0
							}`)),
						handler: votesHandler,
						code:    http.StatusBadRequest,
					},
					{
						name:     "Test secondary Vote test good request",
						endpoint: "/vote?type=secondary&id=" + secondaryComment.ID.Hex(),
						method:   "PATCH",
						body: bytes.NewBuffer([]byte(
							`{
								"upvote": 1,
								"downvote": 0
							}`)),
						handler: votesHandler,
						code:    http.StatusOK,
					},
					{
						name:     "Test  secondary Vote test bad body",
						endpoint: "/vote?type=secondary&id=" + comment.ID.Hex(),
						method:   "PATCH",
						body: bytes.NewBuffer([]byte(
							`{
								"upvote": 1
								"downvote": 0
							}`)),
						handler: votesHandler,
						code:    http.StatusBadRequest,
					},
					{
						name:     "Test secondary Vote test bad request",
						endpoint: "/vote?type=secondary&id=" + comment.ID.Hex(),
						method:   "PATCH",
						body: bytes.NewBuffer([]byte(
							`{
								"upvote": 133,
								"downvote": 0
							}`)),
						handler: votesHandler,
						code:    http.StatusBadRequest,
					},
					{
						name:     "Test weird type",
						endpoint: "/vote?type=df&id=" + comment.ID.Hex(),
						method:   "PATCH",
						body: bytes.NewBuffer([]byte(
							`{
								"upvote": 133,
								"downvote": 0
							}`)),
						handler: votesHandler,
						code:    http.StatusBadRequest,
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
	}
}
