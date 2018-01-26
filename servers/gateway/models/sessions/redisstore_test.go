package sessions

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"os"

	"github.com/go-redis/redis"
)

/*
TestRedisStore tests the RedisStore object
Because the redis.Client is a struct and not an interface,
this is really more of an integration than a unit test.
It tests the basic CRUD cycle, ensuring that session state
saved to redis can be retrieved again.

By default, the test will try to use a local instance of
redis running on its default port (6379). If you want to
use a different address, set the REDISADDR environment variable.
*/
func TestRedisStore(t *testing.T) {
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

	redisaddr := os.Getenv("REDISADDR")
	if len(redisaddr) == 0 {
		redisaddr = "127.0.0.1:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr: redisaddr,
	})

	store := NewRedisStore(client, time.Hour)

	if err := store.Get(sid, stateRet); err != ErrStateNotFound {
		t.Errorf("incorrect error when getting state that was never stored: expected %v but got %v", ErrStateNotFound, err)
	}

	if err := store.Save(sid, &state); err != nil {
		t.Fatalf("error saving state: %v", err)
	}

	//verify that trying to save an unmarshalable session state
	//generates an error (function values can't be encoded in JSON)
	if err := store.Save(sid, func() {}); err == nil {
		t.Error("expected erorr when attempting to save an unmarshalable session state")
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
