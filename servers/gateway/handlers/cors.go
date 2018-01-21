package handlers

import (
	"net/http"
)

// CORSHandler is an instance of an http handler that handles CORS headers
type CORSHandler struct {
	Handler http.Handler
}

// ServeHTTP defines what the CORS handler does when an http request is passed through it
func (ch *CORSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Check if the origin of the request is from somewhere we don't like
	if r.Header.Get("Origin") == "http://evil.com" {
		http.Error(w, "Sorry, bad request blocked", http.StatusUnauthorized)
		return
	}

	//set the various CORS response headers depending on
	//what you want your server to allow
	w.Header().Add(accessControlAllowOrigin, accessControlValue)
	//...more CORS response headers...
	w.Header().Add(accessControlAllowMethods, accessControlMethods)
	w.Header().Add(exposeHeaders, exposedHeaders)
	w.Header().Add(allowHeaders, allowedHeaders)
	w.Header().Add(accessControlAllowAge, accessControlAge)

	//if this is preflight request, the method will
	//be OPTIONS, so call the real handler only if
	//the method is something else
	if r.Method != "OPTIONS" {
		ch.Handler.ServeHTTP(w, r)
	}
}

// NewCORSHandler takes in an http handler and wraps it as a CORS handler
func NewCORSHandler(handlerToWrap http.Handler) *CORSHandler {
	return &CORSHandler{handlerToWrap}
}
