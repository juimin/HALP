package sessions

import (
	"encoding/json"
	"time"

	"github.com/patrickmn/go-cache"
)

//MemStore represents an in-process memory session store.
//This should be used only for testing and prototyping.
//Production systems should use a shared server store like redis
type MemStore struct {
	entries *cache.Cache
}

//NewMemStore constructs and returns a new MemStore
func NewMemStore(sessionDuration time.Duration, purgeInterval time.Duration) *MemStore {
	return &MemStore{
		entries: cache.New(sessionDuration, purgeInterval),
	}
}

//Save saves the provided `sessionState` and associated SessionID to the store.
//The `sessionState` parameter is typically a pointer to a struct containing
//all the data you want to associated with the given SessionID.
func (ms *MemStore) Save(sid SessionID, state interface{}) error {
	j, err := json.Marshal(state)
	if nil != err {
		return err
	}
	ms.entries.Set(sid.String(), j, cache.DefaultExpiration)
	return nil
}

//Get populates `sessionState` with the data previously saved
//for the given SessionID
func (ms *MemStore) Get(sid SessionID, state interface{}) error {
	j, found := ms.entries.Get(sid.String())
	if !found {
		return ErrStateNotFound
	}
	//reset TTL
	ms.entries.Set(sid.String(), j, 0)
	return json.Unmarshal(j.([]byte), state)
}

//Delete deletes all state data associated with the SessionID from the store.
func (ms *MemStore) Delete(sid SessionID) error {
	ms.entries.Delete(sid.String())
	return nil
}
