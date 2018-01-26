package sessions

import (
	"errors"
)

//ErrStateNotFound is returned from Store.Get() when the requested
//session id was not found in the store
var ErrStateNotFound = errors.New("no session state was found in the session store")

// Store interface for generating a store
type Store interface {
	// Save the sesion to the store
	Save(sid SessionID, sessionState interface{}) error

	//Get populates `sessionState` with the data previously saved
	//for the given SessionID
	Get(sid SessionID, sessionState interface{}) error

	//Delete deletes all state data associated with the SessionID from the store.
	Delete(sid SessionID) error
}
