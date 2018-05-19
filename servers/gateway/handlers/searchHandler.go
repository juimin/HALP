package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JuiMin/HALP/servers/gateway/models/sessions"
)

// NSEARCH is a constant for max search terms returned
const NSEARCH = 20

// Respond handles the response
func respond(w http.ResponseWriter, value interface{}) {
	w.Header().Add(headerContentType, contentTypeJSON)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		http.Error(w, fmt.Sprintf("error encoding response value to JSON: %v", err), http.StatusInternalServerError)
	}
}

// SearchHandler Define the ServeHTTP Function for the databsae
func (cr *ContextReceiver) SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Get the state
	state := &SessionState{}
	_, errGetState := sessions.GetState(r, cr.SigningKey, cr.SessionStore, state)
	if errGetState != nil {
		http.Error(w, fmt.Sprintf("error getting session: %v", errGetState), http.StatusUnauthorized)
		return
	}
	switch r.Method {
	// YOU NEED TO MAKE AN POST, DELETE, AND PATCH for adding things to the trie

	// You can't actually update a trie entry because it's just the room name and id
	// the ID never changes and the room shouldn't either. (YOu would never switch the locks on a door)
	// Assume that you will remove the key then add a new one
	case "GET":
		// If get a get request, then we should look at the search term and find out what we have
		searchTerm := r.URL.Query().Get("search")
		searchField := r.URL.Query().Get("type")
		// Check if the search term is soemthing
		if len(searchTerm) == 0 {
			// Return empty JSON object
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "{}")
			return
		}

		if len(searchField) == 0 {
			// Return empty JSON object
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "{}")
			return
		}

		switch searchField {
		case "POST":
			// Get results from the trie
			results, err := cr.PostTrie.NValues(searchTerm, NSEARCH, 0)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error retreiving search results for %s, Error: %v", searchTerm, err),
					http.StatusInternalServerError)
				return
			}

			// Respond with the results
			w.Header().Add(headerContentType, contentTypeJSON)
			respond(w, results)
		case "BOARD":
			// Get results from the trie
			results, err := cr.BoardTrie.NValues(searchTerm, NSEARCH, 0)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error retreiving search results for %s, Error: %v", searchTerm, err),
					http.StatusInternalServerError)
				return
			}

			// Respond with the results
			w.Header().Add(headerContentType, contentTypeJSON)
			respond(w, results)
		case "USER":
			// Get results from the trie
			results, err := cr.UserTrie.NValues(searchTerm, NSEARCH, 0)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error retreiving search results for %s, Error: %v", searchTerm, err),
					http.StatusInternalServerError)
				return
			}

			// Respond with the results
			w.Header().Add(headerContentType, contentTypeJSON)
			respond(w, results)
		default:
			// Add the content type header
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
		// Add the content type header
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
