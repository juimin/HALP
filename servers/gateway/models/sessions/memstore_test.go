package sessions

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

/*
TestMemStore tests the MemStore object

Since a Store is like a database, you can't really test methods like Get()
or Delete() without also calling (and therefore testing) methods like Save(),
so instead of testing individual methods in isolation, this test runs through
a full CRUD cycle, ensuring the correct behavior occurs at each point in that
cycle. You should use a similar approach when testing your RedisStore implementation.
*/
func TestMemStore(t *testing.T) {
	type sessionState struct {
		Sval string
		Ival int
	}

	state := &sessionState{
		Sval: "testing",
		Ival: 99,
	}
	stateRet := &sessionState{}

	sid, err := NewSessionID("test key")
	if err != nil {
		t.Fatalf("error generating new SessionID: %v", err)
	}

	store := NewMemStore(time.Hour, time.Minute)

	if err := store.Get(sid, stateRet); err != ErrStateNotFound {
		t.Errorf("incorrect error when getting state that was never stored: expected %v but got %v", ErrStateNotFound, err)
	}

	if err := store.Save(sid, &state); err != nil {
		t.Fatalf("error saving state: %v", err)
	}

	if err := store.Get(sid, &stateRet); err != nil {
		t.Fatalf("error getting state: %v", err)
	}
	if !reflect.DeepEqual(state, stateRet) {
		jexp, _ := json.MarshalIndent(state, "", "  ")
		jact, _ := json.MarshalIndent(state, "", "  ")
		t.Errorf("incorrect state retrieved:\nEXPECTED\n%s\nACTUAL\n%s", string(jexp), string(jact))
	}

	if err := store.Delete(sid); err != nil {
		t.Errorf("error deleting state: %v", err)
	}

	if err := store.Get(sid, &stateRet); err != ErrStateNotFound {
		t.Fatalf("incorrect error when getting state that was deleted: expected %v but got %v", ErrStateNotFound, err)
	}
}

func TestMemStoreSaveUnmarshalble(t *testing.T) {
	//verify that saving an umarshalalbe session state
	//generates an error
	state := func() {} //function values can't be marshaled into JSON

	sid, err := NewSessionID("test key")
	if err != nil {
		t.Fatalf("error generating new SessionID: %v", err)
	}
	store := NewMemStore(time.Hour, time.Minute)
	if err := store.Save(sid, state); err == nil {
		t.Error("expected error when attempting to save a session state with an unmarshalable field")
	}
}
