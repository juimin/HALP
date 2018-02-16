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
		// Preflight checks for the new user
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

// UsersMeHandler gets the current user or updates the current user
func (cr *ContextReceiver) UsersMeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Get request gets the current user
		sid, err := sessions.GetSessionID(r, cr.SigningKey)
		if err != nil {
			// Could not get the user's session
			w.WriteHeader(http.StatusForbidden)
		}
		state := &SessionState{}
		err = cr.SessionStore.Get(sid, state)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
		}
		// Encode the state's user into the response
		json.NewEncoder(w).Encode(state.User)
	} else if r.Method == "PATCH" {
		// Patch
		sid, err := sessions.GetSessionID(r, cr.SigningKey)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		// Get current user
		state := &SessionState{}
		cr.SessionStore.Get(sid, state)

		// Decode the request body into updates for the user
		updates := &users.UserUpdate{}
		json.NewDecoder(r.Body).Decode(updates)
		err = state.User.ApplyUpdates(updates)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		// Update user in the user store
		err = cr.UserStore.UserUpdate(state.User.ID, updates)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		// Update the session store
		err = cr.SessionStore.Delete(sid)
		if err != nil {
			fmt.Printf("Could not delete sid from session store")
		}
		err = cr.SessionStore.Save(sid, state)
		if err != nil {
			fmt.Printf("Could not save session state to sid")
		}
		// Write the new user out
		json.NewEncoder(w).Encode(state.User)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
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
