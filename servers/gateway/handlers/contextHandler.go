package handlers

import (
	"fmt"

	"github.com/JuiMin/HALP/servers/gateway/models/users"
)

type ContextReceiver struct {
	SigningKey string
	UserStore  users.Store
}

// NewContextReceiver Creates a new context receiver with a session key and references to the session and user stores
func NewContextReceiver(key string, userStore users.Store) (*ContextReceiver, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("No key set for signing key")
	}
	return &ContextReceiver{
		SigningKey: key,
		UserStore:  userStore,
	}, nil
}
