package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/JuiMin/HALP/servers/gateway/models/sessions"
	"github.com/JuiMin/HALP/servers/gateway/models/users"
)

// UsersHandler handlers requests for the users resource and facilitates
func (cr *ContextReceiver) UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// We only accept post to this handler
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		newUser := &users.NewUser{}
		err := json.NewDecoder(r.Body).Decode(newUser)
		errorMessage := ""
		canProceed := true
		if err != nil {
			errorMessage = "Error: could not decode request body"
		}
		err = newUser.Validate()
		if err != nil {
			errorMessage = "Error: Could not validate new user"
			canProceed = false
		}
		emailTemp, err := cr.UserStore.GetByEmail(newUser.Email)
		if emailTemp != nil {
			errorMessage = "Email Taken"
			canProceed = false
		}
		usernameTemp, err := cr.UserStore.GetByUserName(newUser.UserName)
		if usernameTemp != nil {
			errorMessage = "Error: Username taken"
			canProceed = false
		}
		if canProceed {
			// NOw that we have done our checks, we can try and insert the user into the store
			thisUser, err := cr.UserStore.Insert(newUser)
			if err != nil {
				fmt.Printf("Could not insert the user: %v", err)
			}
			// Begin new session
			state := &SessionState{
				StartTime: time.Now(),
				User:      *thisUser,
			}
			_, err = sessions.BeginSession(cr.SigningKey, cr.SessionStore, &state, w)

			if err != nil {
				fmt.Printf(err.Error())
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(&thisUser)
		} else {
			fmt.Printf(errorMessage)
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte(errorMessage))
		}
	}
}

// SessionsHandler allows uses to begin a session with existing credentials
func (cr *ContextReceiver) SessionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		creds := &users.Credentials{}
		json.NewDecoder(r.Body).Decode(creds)
		usr, err := cr.UserStore.GetByEmail(creds.Email)
		// User doesn't exist
		if usr == nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid Credentials: No usr"))
		} else {
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Invalid Credentials: Err getting user"))
			}
			err = usr.Authenticate(creds.Password)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Invalid Credentials: not auth"))
			}
			_, err = sessions.BeginSession(cr.SigningKey, cr.SessionStore, &SessionState{
				StartTime: time.Now(),
				User:      *usr,
			}, w)
			if err != nil {
				w.WriteHeader(http.StatusExpectationFailed)
			}
			// Encode JSON
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(usr)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// SessionsMineHandler ends the session
func (cr *ContextReceiver) SessionsMineHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		_, err := sessions.EndSession(r, cr.SigningKey, cr.SessionStore)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Signed Out"))
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
