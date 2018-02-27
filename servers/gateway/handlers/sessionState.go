package handlers

import (
	"time"

	"github.com/JuiMin/HALP/servers/gateway/models/users"
)

// SessionState defines the struct for a session
type SessionState struct {
	StartTime time.Time  `json:"StartTime"`
	User      users.User `json:"User"`
}
