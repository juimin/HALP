package sessions

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSessionGetSessionID(t *testing.T) {
	key := "test key"
	sid, err := NewSessionID(key)
	if err != nil {
		t.Fatalf("error generating SessionID: %v", err)
	}

	cases := []struct {
		name        string
		hint        string
		header      string
		expectError bool
	}{
		{
			"Valid SessionID and Scheme",
			"Remember to get the SessionID from the Authorization header (or `auth` query string parameter), validate it, and return it",
			schemeBearer + string(sid),
			false,
		},
		{
			"Invalid Scheme",
			"Remember to check the scheme prefix",
			"InvalidScheme " + string(sid),
			true,
		},
		{
			"No Scheme",
			"Remember to check the scheme prefix",
			string(sid),
			true,
		},
		{
			"Invalid SessionID",
			"Remember to validate the id before returning it",
			schemeBearer + "invalid",
			true,
		},
	}

	for _, c := range cases {
		//test using Authorization header
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Add(headerAuthorization, c.header)
		sidRet, err := GetSessionID(req, key)
		if err != nil && !c.expectError {
			t.Errorf("case %s: unexpected error: %v\nHINT: %s", c.name, err, c.hint)
		}
		if c.expectError && err == nil {
			t.Errorf("case %s: expected error but didn't get one\nHINT: %s", c.name, c.hint)
		}
		if !c.expectError && sidRet != sid {
			t.Errorf("case %s: incorrect SessionID returned: expected %s but got %s\nHINT: %s", c.name, sid, sidRet, c.hint)
		}
	}
}

func TestSessionGetSessionIDFromParam(t *testing.T) {
	key := "test key"
	sid, err := NewSessionID(key)
	if err != nil {
		t.Fatalf("error generating SessionID: %v", err)
	}

	URL := fmt.Sprintf("/?%s=%s%s", paramAuthorization, schemeBearer, string(sid))
	req, _ := http.NewRequest("GET", URL, nil)
	sidRet, err := GetSessionID(req, key)
	if err != nil {
		t.Errorf("error getting SessionID from query string parameter: %v", err)
	}
	if sidRet != sid {
		t.Errorf("incorrect SessionID returned:\nEXPECTED:\n%s\nACTUAL:\n%s", sid, sidRet)
	}
}

/*
TestSessionCyle is an integration test that runs through the full
cycle of session methods: BeginSession, GetState, EndSession. It
uses the MemStore as the session store.
*/
func TestSessionCycle(t *testing.T) {
	store := NewMemStore(time.Hour, time.Minute)
	key := "test key"

	//first try getting the session state before a session
	//has been started to ensure you get an error
	var state int
	req, _ := http.NewRequest("GET", "/", nil)
	_, err := GetState(req, key, store, &state)
	if err == nil {
		t.Error("no error returned when getting state before session has started")
	}

	//begin a new session
	state = 100
	respRec := httptest.NewRecorder()

	//try beginning a session with an empty session signing key
	//and ensure it fails
	_, err = BeginSession("", store, state, respRec)
	if err == nil {
		t.Error("expected error when beginning a new session with an empty signing key")
	}

	//then try with a valid signing key and make sure it works
	sid, err := BeginSession(key, store, state, respRec)
	if err != nil {
		t.Fatalf("error beginning session: %v", err)
	}
	if len(sid) == 0 {
		t.Error("SessionID returned from BeginSession was zero-length")
	}

	//ensure response recorder contains Authorization header
	token := respRec.Header().Get(headerAuthorization)
	if len(token) == 0 {
		t.Error("no token returned in Authorization header")
	}

	//get session state
	req, _ = http.NewRequest("GET", "/", nil)
	req.Header.Add(headerAuthorization, token)
	var state2 int
	sid2, err := GetState(req, key, store, &state2)
	if err != nil {
		t.Errorf("unexpected error getting session state: %v", err)
	}
	if sid2 != sid {
		t.Errorf("SessionID returned from GetState did not match SessionID returned from BeginSesion:\nEXPECTED:\n%s\nACTUAL:\n%s", sid, sid2)
	}
	if state2 != state {
		t.Errorf("incorrect session state: expected %d but got %d", state, state2)
	}

	//end the session
	sid2, err = EndSession(req, key, store)
	if err != nil {
		t.Errorf("unexpected error ending session: %v", err)
	}
	if sid2 != sid {
		t.Errorf("SessionID returned from EndSession did not match SessionID returned from BeginSesion:\nEXPECTED:\n%s\nACTUAL:\n%s", sid, sid2)
	}

	//try getting the session state with the same token to ensure
	//that we get back the correct error
	state2 = 0
	_, err = GetState(req, key, store, &state2)
	if err != ErrStateNotFound {
		t.Error("getting state after session end did not return ErrStateNotFound")
	}

	//try ending the session with no Authorization header in request
	//and ensure it generates an error
	req.Header.Del(headerAuthorization)
	_, err = EndSession(req, key, store)
	if err == nil {
		t.Error("expected error when attempting to end session with no Authorization header in request")
	}
}
