package sessions

import (
	"crypto/subtle"
	"errors"
	"net/http"
	"strings"
)

const headerAuthorization = "Authorization"
const paramAuthorization = "auth"
const schemeBearer = "Bearer "

//ErrNoSessionID is used when no session ID was found in the Authorization header
var ErrNoSessionID = errors.New("no session ID found in " + headerAuthorization + " header")

//ErrInvalidScheme is used when the authorization scheme is not supported
var ErrInvalidScheme = errors.New("authorization scheme not supported")

//BeginSession creates a new SessionID, saves the `sessionState` to the store, adds an
//Authorization header to the response with the SessionID, and returns the new SessionID
func BeginSession(signingKey string, store Store, sessionState interface{}, w http.ResponseWriter) (SessionID, error) {
	// Create new session ID
	id, err := NewSessionID(signingKey)
	if err != nil {
		return InvalidSessionID, err
	}
	// Save the session state to the store
	err = store.Save(id, sessionState)
	if err != nil {
		return InvalidSessionID, err
	}
	// Add header to the response writer
	w.Header().Add(headerAuthorization, schemeBearer+id.String())
	// Return the session id and nil error
	return id, nil
}

//GetSessionID extracts and validates the SessionID from the request headers
func GetSessionID(r *http.Request, signingKey string) (SessionID, error) {
	// Get the auth header value
	authorization := r.Header.Get(headerAuthorization)
	// If authorization doesn't exist, do the auth query
	// Now we have the correct scheme and one more item to look at. If it is empty,
	if len(authorization) == 0 {
		// Get the auth query string parameter
		authorization = r.URL.Query().Get(paramAuthorization)
	}

	// If the query param is also empty, then we can return an error
	if len(authorization) == 0 {
		return InvalidSessionID, ErrNoSessionID
	}

	auth := strings.Split(authorization, " ")
	// The bearer and id should be the elements of auth. If auth has two args, it passes, else it returns an error
	if len(auth) != 2 {
		return InvalidSessionID, ErrNoSessionID
	}
	// Since we now have two items in auth, item 1 at index 0 should be the scheme bearer. If not return error
	if subtle.ConstantTimeCompare([]byte(auth[0]), []byte(strings.Trim(schemeBearer, " "))) != 1 {
		return InvalidSessionID, ErrInvalidScheme
	}
	id, err := ValidateID(auth[1], signingKey)
	if err != nil {
		return InvalidSessionID, err
	}
	return id, nil
}

//GetState extracts the SessionID from the request,
//gets the associated state from the provided store into
//the `sessionState` parameter, and returns the SessionID
func GetState(r *http.Request, signingKey string, store Store, sessionState interface{}) (SessionID, error) {
	id, err := GetSessionID(r, signingKey)
	if err != nil {
		return InvalidSessionID, err
	}
	err = store.Get(id, sessionState)
	if err != nil {
		return InvalidSessionID, err
	}
	return id, nil
}

//EndSession extracts the SessionID from the request,
//and deletes the associated data in the provided store, returning
//the extracted SessionID.
func EndSession(r *http.Request, signingKey string, store Store) (SessionID, error) {
	id, err := GetSessionID(r, signingKey)
	if err != nil {
		return InvalidSessionID, err
	}
	err = store.Delete(id)
	if err != nil {
		return InvalidSessionID, err
	}
	return id, nil
}
