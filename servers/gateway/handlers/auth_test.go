package handlers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test the users handler for creating a user
func TestUsersHandler(t *testing.T) {
	// Get Context Instance
	cr := prepTestCR()
	// Generate the handlers
	usersHandler := http.HandlerFunc(cr.UsersHandler)

	cases := []struct {
		name         string
		method       string
		expectedCode int
		handler      http.HandlerFunc
		body         io.Reader
		destination  string
	}{
		{
			name:         "Invalid Method to Users Handler",
			method:       "GET",
			expectedCode: http.StatusMethodNotAllowed,
			handler:      usersHandler,
			body:         nil,
			destination:  "/users",
		},
		{
			name:         "Post to Users Handler - Nil body",
			method:       "POST",
			expectedCode: http.StatusBadRequest,
			handler:      usersHandler,
			body:         nil,
			destination:  "/users",
		},
		{
			name:         "Post to Users Handler - Nil body",
			method:       "POST",
			expectedCode: http.StatusBadRequest,
			handler:      usersHandler,
			body:         nil,
			destination:  "/users",
		},
		{
			name:         "Post to Users Handler - Invalid new user",
			method:       "POST",
			expectedCode: http.StatusConflict,
			handler:      usersHandler,
			body: bytes.NewBuffer([]byte(
				`{
					"password": "potatopass",
					"passwordConf": "potatopass",
					"userName": "potat",
					"firstName":"firstPotato",
					"lastName": "lastPotato",
					"occupation": "vegetable"
				}`)),
			destination: "/users",
		},
		{
			name:         "Post to Users Handler - Valid new user",
			method:       "POST",
			expectedCode: http.StatusCreated,
			handler:      usersHandler,
			body: bytes.NewBuffer([]byte(
				`{
					"email":"goodpotato@potato.com",
					"password": "potatopass",
					"passwordConf": "potatopass",
					"userName": "potat",
					"firstName":"firstPotato",
					"lastName": "lastPotato",
					"occupation": "vegetable"
				}`)),
			destination: "/users",
		},
		{
			name:         "Post to Users Handler - Preexisting email",
			method:       "POST",
			expectedCode: http.StatusConflict,
			handler:      usersHandler,
			body: bytes.NewBuffer([]byte(
				`{
					"email":"goodpotato@potato.com",
					"password": "potatopass",
					"passwordConf": "potatopass",
					"userName": "potat",
					"firstName":"firstPotato",
					"lastName": "lastPotato",
					"occupation": "vegetable"
				}`)),
			destination: "/users",
		},
		{
			name:         "Post to Users Handler - Preexisting Username",
			method:       "POST",
			expectedCode: http.StatusConflict,
			handler:      usersHandler,
			body: bytes.NewBuffer([]byte(
				`{
					"email":"goodpotato2@potato.com",
					"password": "potatopass",
					"passwordConf": "potatopass",
					"userName": "potat",
					"firstName":"firstPotato",
					"lastName": "lastPotato",
					"occupation": "vegetable"
				}`)),
			destination: "/users",
		},
		{
			name:         "Post to Users Handler - Invalid Json",
			method:       "POST",
			expectedCode: http.StatusBadRequest,
			handler:      usersHandler,
			body: bytes.NewBuffer([]byte(
				`{
					"email":"goodpotato2@potato.com"
					"password": "potatopass"231
					"passwordConf": "potatopass",
					"userName": "potat",
					"firstName":"firstPotato",
					"lastName": "lastPotato",
					"occupation": "vegetable"
				}`)),
			destination: "/users",
		},
	}

	for _, c := range cases {
		// Generate a new recorder
		recorder := httptest.NewRecorder()
		// Generate the request
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

// Test the userse handler for updating the user
func TestUsersMeHandler(t *testing.T) {
	// Get Context Instance
	cr := prepTestCR()
	// Generate the handlers
	usersHandler := http.HandlerFunc(cr.UsersHandler)
	usersMeHandler := http.HandlerFunc(cr.UsersMeHandler)

	cases := []struct {
		name         string
		method       string
		expectedCode int
		handler      http.HandlerFunc
		body         io.Reader
		destination  string
	}{
		{
			name:         "Test GET for no/bad session",
			method:       "GET",
			expectedCode: http.StatusForbidden,
			handler:      usersMeHandler,
			body:         nil,
			destination:  "/users/me",
		},
		{
			name:         "Test PATCH for no/bad session",
			method:       "PATCH",
			expectedCode: http.StatusForbidden,
			handler:      usersMeHandler,
			body:         nil,
			destination:  "/users/me",
		},
		{
			name:         "Test Unallowed methods",
			method:       "DELETE",
			expectedCode: http.StatusUnauthorized,
			handler:      usersMeHandler,
			body:         nil,
			destination:  "/users/me",
		},
		{
			name:         "Generate Users for testing",
			method:       "POST",
			expectedCode: http.StatusCreated,
			handler:      usersHandler,
			body: bytes.NewBuffer([]byte(
				`{
				"email":"test@potato.com",
				"password": "potatopass",
				"passwordConf": "potatopass",
				"userName": "test",
				"firstName":"firstPotato",
				"lastName": "lastPotato",
				"occupation": "vegetable"
			}`)),
			destination: "/users",
		},
		{
			name:         "Test Get for valid session",
			method:       "GET",
			expectedCode: http.StatusAccepted,
			handler:      usersMeHandler,
			body:         nil,
			destination:  "/users/me",
		},
		{
			name:         "Test PATCH for invalid input",
			method:       "PATCH",
			expectedCode: http.StatusBadRequest,
			handler:      usersMeHandler,
			body: bytes.NewBuffer([]byte(
				`{
					"firstName": "newFirstName"sdsd
					"lastName": "newLastName",
					"occupation": "spud",
					"email": "test@potato.com"
			}`)),
			destination: "/users/me",
		},
		{
			name:         "Test PATCH for nil update",
			method:       "PATCH",
			expectedCode: http.StatusBadRequest,
			handler:      usersMeHandler,
			body:         nil,
			destination:  "/users/me",
		},
		{
			name:         "Test PATCH for valid input",
			method:       "PATCH",
			expectedCode: http.StatusAccepted,
			handler:      usersMeHandler,
			body: bytes.NewBuffer([]byte(
				`{
					"firstName": "newFirstName",
					"lastName": "newLastName",
					"occupation": "spud",
					"email": "test@potato.com"
			}`)),
			destination: "/users/me",
		},
	}

	authHeader := ""

	for _, c := range cases {
		// Generate a new recorder
		recorder := httptest.NewRecorder()
		// Generate the request
		req, err := http.NewRequest(c.method, c.destination, c.body)
		if authHeader != "" {
			req.Header.Add("Authorization", authHeader)
		}
		if err != nil {
			t.Errorf("%s Failed: Error %v", c.name, err)
		} else {
			c.handler.ServeHTTP(recorder, req)
			if recorder.Header().Get("Authorization") != "" {
				authHeader = recorder.Header().Get("Authorization")
			}
			if recorder.Code != c.expectedCode {
				t.Errorf("%s Failed. Expected %d but got %d", c.name, c.expectedCode, recorder.Code)
			}
		}
	}
}

// Test the Sessions handler for creating sessions
func TestSessionHandlers(t *testing.T) {
	// Get Context Instance
	cr := prepTestCR()
	// Generate the handlers
	usersHandler := http.HandlerFunc(cr.UsersHandler)
	sessionsHandler := http.HandlerFunc(cr.SessionsHandler)
	sessionsMineHandler := http.HandlerFunc(cr.SessionsMineHandler)

	cases := []struct {
		name         string
		method       string
		expectedCode int
		handler      http.HandlerFunc
		body         io.Reader
		destination  string
	}{
		{
			name:         "Invalid Session Method",
			method:       "GET",
			expectedCode: http.StatusMethodNotAllowed,
			handler:      sessionsHandler,
			body:         nil,
			destination:  "/sessions",
		},
		{
			name:         "Test POST No Body",
			method:       "POST",
			expectedCode: http.StatusUnauthorized,
			handler:      sessionsHandler,
			body:         nil,
			destination:  "/sessions",
		},
		{
			name:         "Test POST Non Existant User",
			method:       "POST",
			expectedCode: http.StatusUnauthorized,
			handler:      sessionsHandler,
			body: bytes.NewBuffer([]byte(`
				{
					"email": "nobody@email.com",
					"password": "something"	
				}
				`)),
			destination: "/sessions",
		},
		{
			name:         "Test Delete Method to sessions mine no auth",
			method:       "DELETE",
			expectedCode: http.StatusUnauthorized,
			handler:      sessionsMineHandler,
			body:         nil,
			destination:  "/sessions/mine",
		},
		{
			name:         "Generate Users for testing",
			method:       "POST",
			expectedCode: http.StatusCreated,
			handler:      usersHandler,
			body: bytes.NewBuffer([]byte(
				`{
				"email":"session@potato.com",
				"password": "potatopass",
				"passwordConf": "potatopass",
				"userName": "session",
				"firstName":"firstPotato",
				"lastName": "lastPotato",
				"occupation": "vegetable"
			}`)),
			destination: "/users",
		},
		{
			name:         "Test POST Non Existant User",
			method:       "POST",
			expectedCode: http.StatusUnauthorized,
			handler:      sessionsHandler,
			body: bytes.NewBuffer([]byte(`
				{
					"email": "nobody@email.com",
					"password": "something"	
				}
				`)),
			destination: "/sessions",
		},
		{
			name:         "Test POST Wrong Password to session",
			method:       "POST",
			expectedCode: http.StatusUnauthorized,
			handler:      sessionsHandler,
			body: bytes.NewBuffer([]byte(`
				{
					"email": "session@potato.com",
					"password": "wrongpass"	
				}
				`)),
			destination: "/sessions",
		},
		{
			name:         "Test POST Good Input",
			method:       "POST",
			expectedCode: http.StatusAccepted,
			handler:      sessionsHandler,
			body: bytes.NewBuffer([]byte(`
				{
					"email": "session@potato.com",
					"password": "potatopass"	
				}
				`)),
			destination: "/sessions",
		},
		{
			name:         "Test Bad Request Method to sessions mine",
			method:       "POST",
			expectedCode: http.StatusMethodNotAllowed,
			handler:      sessionsMineHandler,
			body:         nil,
			destination:  "/sessions/mine",
		},
		{
			name:         "Test Delete Method to sessions mine",
			method:       "DELETE",
			expectedCode: http.StatusOK,
			handler:      sessionsMineHandler,
			body:         nil,
			destination:  "/sessions/mine",
		},
	}

	authHeader := ""

	for _, c := range cases {
		// Generate a new recorder
		recorder := httptest.NewRecorder()
		// Generate the request
		req, err := http.NewRequest(c.method, c.destination, c.body)
		if authHeader != "" {
			req.Header.Add("Authorization", authHeader)
		}
		if err != nil {
			t.Errorf("%s Failed: Error %v", c.name, err)
		} else {
			c.handler.ServeHTTP(recorder, req)
			if recorder.Header().Get("Authorization") != "" {
				authHeader = recorder.Header().Get("Authorization")
			}
			if recorder.Code != c.expectedCode {
				t.Errorf("%s Failed. Expected %d but got %d", c.name, c.expectedCode, recorder.Code)
			}
		}
	}
}

// TestFavoritesHandler tests updates to the favorites
func TestFavoritesHandler(t *testing.T) {
	// Get Context Instance
	cr := prepTestCR()
	// Generate the handlers
	usersHandler := http.HandlerFunc(cr.UsersHandler)
	usersMeHandler := http.HandlerFunc(cr.UsersMeHandler)
	favoritesHandler := http.HandlerFunc(cr.FavoritesHandler)

	cases := []struct {
		name         string
		method       string
		expectedCode int
		handler      http.HandlerFunc
		body         io.Reader
		destination  string
	}{
		{
			name:         "Test PATCH for no/bad session",
			method:       "PATCH",
			expectedCode: http.StatusForbidden,
			handler:      favoritesHandler,
			body:         nil,
			destination:  "/favorites/update",
		},
		{
			name:         "Test bad method",
			method:       "POST",
			expectedCode: http.StatusMethodNotAllowed,
			handler:      favoritesHandler,
			body:         nil,
			destination:  "/favorites/update",
		},
		{
			name:         "Generate Users for testing",
			method:       "POST",
			expectedCode: http.StatusCreated,
			handler:      usersHandler,
			body: bytes.NewBuffer([]byte(
				`{
				"email":"r22d@potato.com",
				"password": "potatopass",
				"passwordConf": "potatopass",
				"userName": "dd3",
				"firstName":"firstPotato",
				"lastName": "lastPotato",
				"occupation": "vegetable"
			}`)),
			destination: "/users",
		},
		{
			name:         "Test Get for valid session",
			method:       "GET",
			expectedCode: http.StatusAccepted,
			handler:      usersMeHandler,
			body:         nil,
			destination:  "/users/me",
		},
		{
			name:         "Test PATCH for invalid input",
			method:       "PATCH",
			expectedCode: http.StatusBadRequest,
			handler:      favoritesHandler,
			body: bytes.NewBuffer([]byte(
				`{
					"firstName": "newFirstName"sdsd
					"lastName": "newLastName",
					"occupation": "spud",
					"email": "test@potato.com"
			}`)),
			destination: "/favorites/update",
		},
		{
			name:         "Test PATCH for nil update",
			method:       "PATCH",
			expectedCode: http.StatusBadRequest,
			handler:      favoritesHandler,
			body:         nil,
			destination:  "/favorites/update",
		},
		{
			name:         "Test PATCH for valid input (adding)",
			method:       "PATCH",
			expectedCode: http.StatusOK,
			handler:      favoritesHandler,
			body: bytes.NewBuffer([]byte(
				`{
					"adding": true,
					"updateID": "507f1f77bcf86cd799439011"
			}`)),
			destination: "/favorites/update",
		},
		{
			name:         "Test PATCH for valid input (remove)",
			method:       "PATCH",
			expectedCode: http.StatusOK,
			handler:      favoritesHandler,
			body: bytes.NewBuffer([]byte(
				`{
					"adding": false,
					"updateID": "507f1f77bcf86cd799439011"
			}`)),
			destination: "/favorites/update",
		},
	}

	authHeader := ""

	for _, c := range cases {
		// Generate a new recorder
		recorder := httptest.NewRecorder()
		// Generate the request
		req, err := http.NewRequest(c.method, c.destination, c.body)
		if authHeader != "" {
			req.Header.Add("Authorization", authHeader)
		}
		if err != nil {
			t.Errorf("%s Failed: Error %v", c.name, err)
		} else {
			c.handler.ServeHTTP(recorder, req)
			if recorder.Header().Get("Authorization") != "" {
				authHeader = recorder.Header().Get("Authorization")
			}
			if recorder.Code != c.expectedCode {
				t.Errorf("%s Failed. Expected %d but got %d", c.name, c.expectedCode, recorder.Code)
			}
		}
	}
}

// TestBookmarksHandler Tests for the updating of bookmarks for a user
func TestBookmarksHandler(t *testing.T) {
	// Get Context Instance
	cr := prepTestCR()
	// Generate the handlers
	usersHandler := http.HandlerFunc(cr.UsersHandler)
	usersMeHandler := http.HandlerFunc(cr.UsersMeHandler)
	bookmarksHandler := http.HandlerFunc(cr.BookmarksHandler)

	cases := []struct {
		name         string
		method       string
		expectedCode int
		handler      http.HandlerFunc
		body         io.Reader
		destination  string
	}{
		{
			name:         "Test PATCH for no/bad session",
			method:       "PATCH",
			expectedCode: http.StatusForbidden,
			handler:      bookmarksHandler,
			body:         nil,
			destination:  "/bookmarks/update",
		},
		{
			name:         "Test Bad Method",
			method:       "POST",
			expectedCode: http.StatusMethodNotAllowed,
			handler:      bookmarksHandler,
			body:         nil,
			destination:  "/bookmarks/update",
		},
		{
			name:         "Generate Users for testing",
			method:       "POST",
			expectedCode: http.StatusCreated,
			handler:      usersHandler,
			body: bytes.NewBuffer([]byte(
				`{
				"email":"peeople@potato.com",
				"password": "potatopass",
				"passwordConf": "potatopass",
				"userName": "dfds",
				"firstName":"firstPotato",
				"lastName": "lastPotato",
				"occupation": "vegetable"
			}`)),
			destination: "/users",
		},
		{
			name:         "Test Get for valid session",
			method:       "GET",
			expectedCode: http.StatusAccepted,
			handler:      usersMeHandler,
			body:         nil,
			destination:  "/users/me",
		},
		{
			name:         "Test PATCH for invalid input",
			method:       "PATCH",
			expectedCode: http.StatusBadRequest,
			handler:      bookmarksHandler,
			body: bytes.NewBuffer([]byte(
				`{
					"firstName": "newFirstName"sdsd
					"lastName": "newLastName",
					"occupation": "spud",
					"email": "test@potato.com"
			}`)),
			destination: "/bookmarks/update",
		},
		{
			name:         "Test PATCH for nil update",
			method:       "PATCH",
			expectedCode: http.StatusBadRequest,
			handler:      bookmarksHandler,
			body:         nil,
			destination:  "/bookmarks/update",
		},
		{
			name:         "Test PATCH for valid input (add)",
			method:       "PATCH",
			expectedCode: http.StatusOK,
			handler:      bookmarksHandler,
			body: bytes.NewBuffer([]byte(
				`{
					"adding": true,
					"updateID": "507f1f77bcf86cd799439011"
			}`)),
			destination: "/bookmarks/update",
		},
		{
			name:         "Test PATCH for valid input (remove)",
			method:       "PATCH",
			expectedCode: http.StatusOK,
			handler:      bookmarksHandler,
			body: bytes.NewBuffer([]byte(
				`{
					"adding": false,
					"updateID": "507f1f77bcf86cd799439011"
			}`)),
			destination: "/bookmarks/update",
		},
	}

	authHeader := ""

	for _, c := range cases {
		// Generate a new recorder
		recorder := httptest.NewRecorder()
		// Generate the request
		req, err := http.NewRequest(c.method, c.destination, c.body)
		if authHeader != "" {
			req.Header.Add("Authorization", authHeader)
		}
		if err != nil {
			t.Errorf("%s Failed: Error %v", c.name, err)
		} else {
			c.handler.ServeHTTP(recorder, req)
			if recorder.Header().Get("Authorization") != "" {
				authHeader = recorder.Header().Get("Authorization")
			}
			if recorder.Code != c.expectedCode {
				t.Errorf("%s Failed. Expected %d but got %d", c.name, c.expectedCode, recorder.Code)
			}
		}
	}
}
