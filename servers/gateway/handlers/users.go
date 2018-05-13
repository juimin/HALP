package handlers

// FavoritesHandler allows the user to update their favorites
import (
	"encoding/json"
	"net/http"

	"github.com/JuiMin/HALP/servers/gateway/models/sessions"
	"github.com/JuiMin/HALP/servers/gateway/models/users"
)

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
						err = state.User.UpdateFavorites(updates)
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
