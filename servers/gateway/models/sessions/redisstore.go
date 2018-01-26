package sessions

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis"
	cache "github.com/patrickmn/go-cache"
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
	redisStore := &RedisStore{
		Client:          client,
		SessionDuration: sessionDuration,
	}
	return redisStore
}

//Store implementation

//SaveEmail saves the counter for the provided email
func (rs *RedisStore) SaveEmail(email string, counter interface{}) error {
	buffer, err := json.Marshal(counter)
	if err != nil {
		return err
	}

	err1 := rs.Client.Set("email:"+email, buffer, 5*time.Minute).Err()
	if err1 != nil {
		return err1
	}

	return nil
}

//GetEmail will get the counter data for the given email
func (rs *RedisStore) GetEmail(email string, counter interface{}) error {
	found, err := rs.Client.Get("email:" + email).Bytes()
	if err != nil {
		return errors.New("no account of that email was found in the session store")
	}

	errMarshall := json.Unmarshal(found, counter)
	if errMarshall != nil {
		return errMarshall
	}

	return nil
}

//DeleteEmail deletes all the counter data associated with the Email from the store
func (rs *RedisStore) DeleteEmail(email string, counterState interface{}) error {
	//TODO: delete the data stored in redis for the provided SessionID
	err := rs.Client.Del("email:" + email).Err()
	if err != nil {
		return err
	}
	return nil
}

//Save saves the provided `sessionState` and associated SessionID to the store.
//The `sessionState` parameter is typically a pointer to a struct containing
//all the data you want to associated with the given SessionID.
func (rs *RedisStore) Save(sid SessionID, sessionState interface{}) error {
	//TODO: marshal the `sessionState` to JSON and save it in the redis database,
	//using `sid.getRedisKey()` for the key.
	//return any errors that occur along the way.
	buffer, err := json.Marshal(sessionState)
	if err != nil {
		return err
	}

	error1 := rs.Client.Set(sid.getRedisKey(), buffer, cache.DefaultExpiration).Err()
	if error1 != nil {
		return error1
	}

	return nil
}

//Get populates `sessionState` with the data previously saved
//for the given SessionID
func (rs *RedisStore) Get(sid SessionID, sessionState interface{}) error {
	//TODO: get the previously-saved session state data from redis,
	//unmarshal it back into the `sessionState` parameter
	//and reset the expiry time, so that it doesn't get deleted until
	//the SessionDuration has elapsed.

	found, err := rs.Client.Get(sid.getRedisKey()).Bytes()
	if err != nil {
		return ErrStateNotFound
	}

	//reset TTL
	rs.Client.Set(sid.getRedisKey(), found, 0)

	errMarshall := json.Unmarshal(found, sessionState)
	if errMarshall != nil {
		return errMarshall
	}
	//fmt.Println(sessionState)

	return nil
}

//Delete deletes all state data associated with the SessionID from the store.
func (rs *RedisStore) Delete(sid SessionID) error {
	//TODO: delete the data stored in redis for the provided SessionID
	err := rs.Client.Del(sid.getRedisKey()).Err()
	if err != nil {
		return err
	}
	return nil
}

//getRedisKey() returns the redis key to use for the SessionID
func (sid SessionID) getRedisKey() string {
	//convert the SessionID to a string and add the prefix "sid:" to keep
	//SessionID keys separate from other keys that might end up in this
	//redis instance
	return "sid:" + sid.String()
}
