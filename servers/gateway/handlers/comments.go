package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/JuiMin/HALP/servers/gateway/models/comments"
	"github.com/JuiMin/HALP/servers/gateway/models/sessions"

	"gopkg.in/mgo.v2/bson"
)

// Provide a handler collection for interacting with comments

// CommentsHandler addresses requests that involve individual comments
func (cr *ContextReceiver) CommentsHandler(w http.ResponseWriter, r *http.Request) {
	// Check that the user is authenticated
	// Authenticate the user
	sessionActive := true
	status := http.StatusOK

	if r.Method != "GET" {
		sid, err := sessions.GetSessionID(r, cr.SigningKey)
		if err != nil {
			status = http.StatusForbidden
			sessionActive = false
		} else {
			// Session Get
			session := &SessionState{}
			err = cr.SessionStore.Get(sid, session)
			if err != nil {
				status = http.StatusInternalServerError
				sessionActive = false
			}
		}
	}

	// Check that the http request method is proper:
	switch r.Method {
	case "GET":
		// We should be getting three things in the request URL
		id := r.URL.Query().Get("id")
		queryType := r.URL.Query().Get("query")

		if len(id) == 0 || !bson.IsObjectIdHex(id) {
			status = http.StatusBadRequest
		} else {
			switch queryType {
			case "singleComment":
				comment, err := cr.CommentStore.GetByCommentID(bson.ObjectIdHex(id))
				if err != nil {
					status = http.StatusNotFound
				} else {
					json.NewEncoder(w).Encode(comment)
					status = http.StatusOK
				}
			case "singeSecondary":
				comment, err := cr.CommentStore.GetBySecondaryID(bson.ObjectIdHex(id))
				if err != nil {
					status = http.StatusNotFound
				} else {
					json.NewEncoder(w).Encode(comment)
					status = http.StatusOK
				}
			case "allByPost":
				comments, err := cr.CommentStore.GetCommentsByPostID(bson.ObjectIdHex(id))
				if err != nil {
					status = http.StatusNotFound
				} else {
					json.NewEncoder(w).Encode(comments)
					status = http.StatusOK
				}
			case "allByParent":
				comments, err := cr.CommentStore.GetByParentID(bson.ObjectIdHex(id))
				if err != nil {
					status = http.StatusNotFound
				} else {
					json.NewEncoder(w).Encode(comments)
					status = http.StatusOK
				}
			default:
				status = http.StatusBadRequest
			}
		}
	case "POST":
		if sessionActive {
			commentType := r.URL.Query().Get("type")
			if commentType == "primary" {
				comment := &comments.NewComment{}
				err := json.NewDecoder(r.Body).Decode(comment)
				if err != nil {
					status = http.StatusBadRequest
				} else {
					c, err := cr.CommentStore.InsertComment(comment)
					if err != nil {
						status = http.StatusBadRequest
					} else {
						// Return the comment
						status = http.StatusOK
						json.NewEncoder(w).Encode(&c)
					}
				}
			} else if commentType == "secondary" {
				secondaryComment := &comments.NewSecondaryComment{}
				err := json.NewDecoder(r.Body).Decode(secondaryComment)
				if err != nil {
					status = http.StatusBadRequest
				} else {
					sc, err := cr.CommentStore.InsertSecondaryComment(secondaryComment)
					if err != nil {
						status = http.StatusBadRequest
					} else {
						// Return the comment
						status = http.StatusOK
						json.NewEncoder(w).Encode(&sc)
					}
				}
			} else {
				status = http.StatusBadRequest
			}
		}
	case "PATCH":
		if sessionActive {
			commentType := r.URL.Query().Get("type")
			id := r.URL.Query().Get("id")
			if commentType == "primary" {
				comment := &comments.CommentUpdate{}
				err := json.NewDecoder(r.Body).Decode(comment)
				if err != nil {
					status = http.StatusBadRequest
				} else {
					c, err := cr.CommentStore.UpdateComment(bson.ObjectIdHex(id), comment)
					if err != nil {
						status = http.StatusBadRequest
					} else {
						// Return the comment
						status = http.StatusOK
						json.NewEncoder(w).Encode(&c)
					}
				}
			} else if commentType == "secondary" {
				secondaryComment := &comments.SecondaryCommentUpdate{}
				err := json.NewDecoder(r.Body).Decode(secondaryComment)
				if err != nil {
					status = http.StatusBadRequest
				} else {
					sc, err := cr.CommentStore.UpdateSecondaryComment(bson.ObjectIdHex(id), secondaryComment)
					if err != nil {
						status = http.StatusBadRequest
					} else {
						// Return the comment
						status = http.StatusOK
						json.NewEncoder(w).Encode(&sc)
					}
				}
			} else {
				status = http.StatusBadRequest
			}
		}
	case "DELETE":
		if sessionActive {
			id := r.URL.Query().Get("id")
			if bson.IsObjectIdHex(id) {
				err := cr.CommentStore.DeleteComment(bson.ObjectIdHex(id))
				if err != nil {
					status = http.StatusBadRequest
				} else {
					status = http.StatusOK
				}
			}
		}
	default:
		// We only accept GET, POST and PATCH here
		status = http.StatusMethodNotAllowed
	}
	// Set the status
	w.WriteHeader(status)
}

// VotesHandler handles voting
func (cr *ContextReceiver) VotesHandler(w http.ResponseWriter, r *http.Request) {
	// Authenticate the user
	sid, err := sessions.GetSessionID(r, cr.SigningKey)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
	}

	// Session Get
	session := &SessionState{}
	err = cr.SessionStore.Get(sid, session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Session is fine
	if err == nil {
		if r.Method == "PATCH" {
			commentType := r.URL.Query().Get("type")
			id := r.URL.Query().Get("id")
			if r.Body != nil {
				if commentType == "primary" {
					comment := &comments.CommentVote{}
					err := json.NewDecoder(r.Body).Decode(comment)
					if err != nil {
						w.WriteHeader(http.StatusBadRequest)
					} else {
						c, err := cr.CommentStore.CommentVote(bson.ObjectIdHex(id), comment)
						if err != nil {
							w.WriteHeader(http.StatusBadRequest)
						} else {
							// Return the comment
							w.WriteHeader(http.StatusOK)
							json.NewEncoder(w).Encode(&c)
						}
					}
				} else if commentType == "secondary" {
					secondaryComment := &comments.SecondaryCommentVote{}
					err := json.NewDecoder(r.Body).Decode(secondaryComment)
					if err != nil {
						w.WriteHeader(http.StatusBadRequest)
					} else {
						sc, err := cr.CommentStore.SecondaryCommentVote(bson.ObjectIdHex(id), secondaryComment)
						if err != nil {
							w.WriteHeader(http.StatusBadRequest)
						} else {
							// Return the comment
							w.WriteHeader(http.StatusOK)
							json.NewEncoder(w).Encode(&sc)
						}
					}
				} else {
					w.WriteHeader(http.StatusBadRequest)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
