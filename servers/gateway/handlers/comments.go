package handlers

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

// Provide a handler collection for interacting with comments

// CommentsHandler addresses requests that involve individual comments
func (cr *ContextReceiver) CommentsHandler(w http.ResponseWriter, r *http.Request) {
	// Check that the user is authenticated

	// Check that the http request method is proper:

	switch r.Method {
	case "GET":
		// We should be getting three things in the request URL
		id := r.URL.Query().Get("id")
		queryType := r.URL.Query().Get("queryType")

		if len(id) == 0 || !bson.IsObjectIdHex(id) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			switch queryType {
			case "singleComment":
				comment, err := cr.CommentStore.GetByCommentID(bson.ObjectIdHex(id))
				if err != nil {
					w.WriteHeader(http.StatusNotFound)
				} else {
					json.NewEncoder(w).Encode(comment)
					w.WriteHeader(http.StatusOK)
				}
			case "singeSecondary":
				comment, err := cr.CommentStore.GetBySecondaryID(bson.ObjectIdHex(id))
				if err != nil {
					w.WriteHeader(http.StatusNotFound)
				} else {
					json.NewEncoder(w).Encode(comment)
					w.WriteHeader(http.StatusOK)
				}
			case "allByPost":
				comments, err := cr.CommentStore.GetCommentsByPostID(bson.ObjectIdHex(id))
				if err != nil {
					w.WriteHeader(http.StatusNotFound)
				} else {
					json.NewEncoder(w).Encode(comments)
					w.WriteHeader(http.StatusOK)
				}
			case "allByParent":
				comments, err := cr.CommentStore.GetByParentID(bson.ObjectIdHex(id))
				if err != nil {
					w.WriteHeader(http.StatusNotFound)
				} else {
					json.NewEncoder(w).Encode(comments)
					w.WriteHeader(http.StatusOK)
				}
			default:
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	case "POST":
	case "PATCH":
	case "DELETE":
		id := r.URL.Query().Get("id")
		if bson.IsObjectIdHex(id) {
			err := cr.CommentStore.DeleteComment(bson.ObjectIdHex(id))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		}
	default:
		// We only accept GET, POST and PATCH here
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
