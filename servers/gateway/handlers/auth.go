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
		// Defaults for preflight checks
		errorMessage := ""
		canProceed := true
		status := http.StatusAccepted
		if r.Body == nil {
			errorMessage = "Error: Could not decode request body"
			status = http.StatusBadRequest
			canProceed = false
		}
		if canProceed {
			newUser := &users.NewUser{}
			err := json.NewDecoder(r.Body).Decode(newUser)
			// Preflight checks for the new user
			// Bad request for new user. Need the proper new user
			if err != nil {
				errorMessage = "Error: Could not decode request body"
				status = http.StatusBadRequest
				canProceed = false
			}
			if canProceed {
				// Validate the User as a valid new user
				err = newUser.Validate()
				if err != nil {
					errorMessage = "Error: Could not validate new user"
					status = http.StatusConflict
					canProceed = false
				}
			}
			if canProceed {
				emailTemp, _ := cr.UserStore.GetByEmail(newUser.Email)
				if emailTemp != nil {
					errorMessage = "Error: Email already Taken"
					status = http.StatusConflict
					canProceed = false
				}
			}
			if canProceed {
				usernameTemp, _ := cr.UserStore.GetByUserName(newUser.UserName)
				if usernameTemp != nil {
					errorMessage = "Error: Username already taken"
					status = http.StatusConflict
					canProceed = false
				}
			}
			if canProceed {
				// NOw that we have done our checks, we can try and insert the user into the store
				thisUser, err := cr.UserStore.Insert(newUser)
				if err != nil {
					fmt.Printf("Could not insert the user: %v", err)
					status = http.StatusInternalServerError
				}
				// Begin new session
				state := &SessionState{
					StartTime: time.Now(),
					User:      *thisUser,
				}
				_, err = sessions.BeginSession(cr.SigningKey, cr.SessionStore, &state, w)

				if err != nil {
					fmt.Printf(err.Error())
					status = http.StatusInternalServerError
				}

				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(&thisUser)
			}
		}
		// Something went wrong somewhere
		if !canProceed {
			fmt.Printf(errorMessage + "\n")
			w.WriteHeader(status)
			w.Write([]byte(errorMessage))
		}
	}
}

// FavoritesHandler allows the user to update their favorites
func (cr *ContextReceiver) FavoritesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" {
		// Get request gets the current user
		sid, err := sessions.GetSessionID(r, cr.SigningKey)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
		} else {
			state := &SessionState{}
			// Get the sesssion
			err = cr.SessionStore.Get(sid, state)
			// Check the session is in play
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				// Now that we have a valid session we can update the user
				updates := &users.FavoritesUpdate{}
				// Check if the r body is there
				if r.Body == nil {
					w.WriteHeader(http.StatusBadRequest)
				} else {
					err = json.NewDecoder(r.Body).Decode(updates)
					// Check if the object is the right format
					if err != nil {
						w.WriteHeader(http.StatusBadRequest)
					} else {
						err = state.User.UpdateFavorite(updates)
						if err != nil {
							w.WriteHeader(http.StatusInternalServerError)
						} else {
							_, err = cr.UserStore.FavoritesUpdate(state.User.ID, updates)
							if err != nil {
								w.WriteHeader(http.StatusInternalServerError)
							} else {
								// Create new session
								err = cr.SessionStore.Delete(sid)
								if err != nil {
									w.WriteHeader(http.StatusInternalServerError)
								} else {
									err = cr.SessionStore.Save(sid, state)
									if err != nil {
										w.WriteHeader(http.StatusInternalServerError)
									} else {
										// Everything was good
										w.WriteHeader(http.StatusOK)
										json.NewEncoder(w).Encode(state.User)
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		// We only handle updates here
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// BookmarksHandler allows the user to update their favorites
func (cr *ContextReceiver) BookmarksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" {
		// Get request gets the current user
		sid, err := sessions.GetSessionID(r, cr.SigningKey)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
		} else {
			state := &SessionState{}
			// Get the sesssion
			err = cr.SessionStore.Get(sid, state)
			// Check the session is in play
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				// Now that we have a valid session we can update the user
				updates := &users.BookmarksUpdate{}
				// Check if the r body is there
				if r.Body == nil {
					w.WriteHeader(http.StatusBadRequest)
				} else {
					err = json.NewDecoder(r.Body).Decode(updates)
					// Check if the object is the right format
					if err != nil {
						w.WriteHeader(http.StatusBadRequest)
					} else {
						err = state.User.UpdateBookmarks(updates)
						if err != nil {
							w.WriteHeader(http.StatusInternalServerError)
						} else {
							_, err = cr.UserStore.BookmarksUpdate(state.User.ID, updates)
							if err != nil {
								w.WriteHeader(http.StatusInternalServerError)
							} else {
								// Create new session
								err = cr.SessionStore.Delete(sid)
								if err != nil {
									w.WriteHeader(http.StatusInternalServerError)
								} else {
									err = cr.SessionStore.Save(sid, state)
									if err != nil {
										w.WriteHeader(http.StatusInternalServerError)
									} else {
										// Everything was good
										w.WriteHeader(http.StatusOK)
										json.NewEncoder(w).Encode(state.User)
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		// We only handle updates here
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// UsersMeHandler gets the current user or updates the current user
func (cr *ContextReceiver) UsersMeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		status := http.StatusAccepted
		// Get request gets the current user
		sid, err := sessions.GetSessionID(r, cr.SigningKey)
		if err != nil {
			// Could not get the user's session
			status = http.StatusForbidden
		}
		state := &SessionState{}
		err = cr.SessionStore.Get(sid, state)
		if err != nil {
			status = http.StatusForbidden
		}
		// Encode the state's user into the response
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(state.User)
	} else if r.Method == "PATCH" {
		canProceed := true
		// Ask the database if we have a session
		sid, err := sessions.GetSessionID(r, cr.SigningKey)
		if err != nil {
			// No we don't have a session so we can't do this
			w.WriteHeader(http.StatusForbidden)
			canProceed = false
		}
		if canProceed {
			// Get current user
			state := &SessionState{}
			err := cr.SessionStore.Get(sid, state)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				canProceed = false
			}
			if canProceed {
				// Decode the request body into updates for the user
				updates := &users.UserUpdate{}
				// Check if the patch body is nil
				if r.Body == nil {
					w.WriteHeader(http.StatusBadRequest)
					canProceed = false
				}
				// Apply the updates to the state and database
				// We should have updates parsed properly
				if canProceed {
					// Decode the updates into the correct format
					err = json.NewDecoder(r.Body).Decode(updates)
					if err != nil {
						w.WriteHeader(http.StatusBadRequest)
						canProceed = false
					}
					if canProceed {
						// Apply the updates to the struct
						err = state.User.ApplyUpdates(updates)
						if err != nil {
							fmt.Printf("Could not update user state\n")
						}
						// Update user in the user store
						err = cr.UserStore.UserUpdate(state.User.ID, updates)
						if err != nil {
							fmt.Printf("Could not update user in database\n")
						}
						// Update the session store
						// Replace the state
						err = cr.SessionStore.Delete(sid)
						if err != nil {
							fmt.Printf("Could not delete sid from session store\n")
						}
						err = cr.SessionStore.Save(sid, state)
						if err != nil {
							fmt.Printf("Could not save session state to sid\n")
							w.WriteHeader(http.StatusInternalServerError)
						} else {
							w.WriteHeader(http.StatusAccepted)
							// Output the updated user back
							json.NewEncoder(w).Encode(state.User)
						}
					}
				}
			}
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// SessionsHandler allows uses to begin a session with existing credentials
func (cr *ContextReceiver) SessionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		creds := &users.Credentials{}
		usr := &users.User{}
		err := fmt.Errorf("")
		if r.Body != nil {
			json.NewDecoder(r.Body).Decode(creds)
			usr, err = cr.UserStore.GetByEmail(creds.Email)
		} else {
			usr = nil
		}
		// User doesn't exist
		if usr == nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid Credentials: No usr"))
		} else {
			canProceed := true
			err = usr.Authenticate(creds.Password)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Invalid Credentials: not auth"))
				canProceed = false
			}
			if canProceed {
				_, err = sessions.BeginSession(cr.SigningKey, cr.SessionStore, &SessionState{
					StartTime: time.Now(),
					User:      *usr,
				}, w)
				if err != nil {
					// We somehow died while doing sesion insertion
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					w.WriteHeader(http.StatusAccepted)
					json.NewEncoder(w).Encode(usr)
				}
			}
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
