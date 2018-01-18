package handlers

import (
	"fmt"
	"net/http"
)

// RootHandler handles a request to the root node
// This will probably send the root to the client
func RootHandler(w http.ResponseWriter, r *http.Request) {
	// Add the content type header
	w.Header().Add(headerContentType, contentTypeText)
	fmt.Fprint(w, "Welcome to the gateway! There is no resource here right now!")
}
