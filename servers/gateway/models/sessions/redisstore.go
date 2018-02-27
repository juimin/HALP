package sessions

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis"
)

//RedisStore represents a session.Store backed by redis.
type RedisStore struct {
	//Redis client used to talk to redis server.
	Client *redis.Client
	//Used for key expiry time on redis.
	SessionDuration time.Duration
}

//NewRedisStore constructs a new RedisStore
func NewRedisStore(client *redis.Client, sessionDuration time.Duration) *RedisStore {
	//initialize and return a new RedisStore struct

	// If the arguments are nil, return nil
	if client == nil {
		return nil
	}

	// Construct he redis store and return it
	return &RedisStore{
		Client:          client,
		SessionDuration: sessionDuration,
	}
}

//Save saves the provided `sessionState` and associated SessionID to the store.
//The `sessionState` parameter is typically a pointer to a struct containing
//all the data you want to associated with the given SessionID.
func (rs *RedisStore) Save(sid SessionID, sessionState interface{}) error {
	if len(sid) == 0 || sessionState == nil {
		return errors.New("Invalid arguments for Save")
	}
	expiry, err := time.ParseDuration("1440s")
	if err != nil {
		return err
	}
	sessionJSON, err := json.Marshal(sessionState)
	if err != nil {
		return err
	}
	statcmd := rs.Client.Set(sid.getRedisKey(), sessionJSON, expiry)
	if statcmd.Err() != nil {
		return statcmd.Err()
	}
	return nil
}

//Get populates `sessionState` with the data previously saved
//for the given SessionID
func (rs *RedisStore) Get(sid SessionID, sessionState interface{}) error {
	// Reset the expire time
	val, err := rs.Client.Get(sid.getRedisKey()).Result()
	if err != nil {
		return ErrStateNotFound
	}
	expiry, err := time.ParseDuration(rs.SessionDuration.String())
	if err != nil {
		return err
	}
	rs.Client.Expire(sid.getRedisKey(), expiry)

	// Unmarshal to the session state variable
	err = json.Unmarshal([]byte(val), sessionState)
	if err != nil {
		return err
	}

	return nil
}

//Delete deletes all state data associated with the SessionID from the store.
func (rs *RedisStore) Delete(sid SessionID) error {
	_, err := rs.Client.Get(sid.getRedisKey()).Result()
	if err != nil {
		return ErrStateNotFound
	}
	err = rs.Client.Del(sid.getRedisKey()).Err()
	if err != nil {
		return err
	}
	return nil
}

//getRedisKey() returns the redis key to use for the SessionID
func (sid SessionID) getRedisKey() string {
	return "sid:" + sid.String()
}
