package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUsersHandler(t *testing.T) {

	cases := []struct {
		name                string
		expectedContentType string
		expectedStatus      int
		expectedOutput      string
	}{
		{
			name:                "Passing Test",
			expectedContentType: contentTypeText,
			expectedStatus:      http.StatusOK,
			expectedOutput:      "Welcome to the gateway! There is no resource here right now!",
		},
	}

	for _, c := range cases {
		// Create http request for the root path
		req, err := http.NewRequest("GET", "/users", nil)
		// Fatal error report for test
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()

		// Serve the handler
		handler := http.HandlerFunc(RootHandler)
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("%s Failed: Testing Status Code: expected %v but got %v",
				c.name, c.expectedStatus, status)
		}

		if header := rr.Header().Get(headerContentType); header != c.expectedContentType {
			t.Errorf("%s Failed: Testing header, expected %s but got %s",
				c.name, c.expectedContentType, header)
		}

		// Check the response body is what we expect.
		if body := rr.Body.String(); body != c.expectedOutput {
			t.Errorf("%s Failed: Testing response body: expected %s but got %s",
				c.name, c.expectedOutput, body)
		}
	}
}
