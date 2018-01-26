package handlers

import (
	"fmt"

	"github.com/JuiMin/HALP/servers/gateway/models/sessions"
	"github.com/JuiMin/HALP/servers/gateway/models/users"
)

// ContextReceiver keeps track of the Various Storage services and keeps a reference to each of them
type ContextReceiver struct {
	SigningKey string
	UserStore  users.Store
	RedisStore sessions.Store
}

// NewContextReceiver Creates a new context receiver with a session key and references to the session and user stores
func NewContextReceiver(key string, userStore users.Store, redisStore sessions.Store) (*ContextReceiver, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("No key set for signing key")
	}
	return &ContextReceiver{
		SigningKey: key,
		UserStore:  userStore,
		RedisStore: redisStore,
	}, nil
}
