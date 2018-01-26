package sessions

import (
	"encoding/base64"
	"testing"
)

// TestNewID tests the creation of a new session ID
func TestNewID(t *testing.T) {
	cases := []struct {
		name        string
		hint        string
		signingKey  string
		expectError bool
	}{
		{
			"Empty Signing Key",
			"Remember to return an error if `signingKey` is zero-length",
			"",
			true,
		},
		{
			"Valid Signing Key",
			"Remember to return a valid base64-url-encoded SessionID if the `signingKey` is non-zero-length",
			"test key",
			false,
		},
	}

	for _, c := range cases {
		sid, err := NewSessionID(c.signingKey)
		if err != nil && !c.expectError {
			t.Errorf("case %s: unexpected error generating new SessionID: %v\nHINT: %s", c.name, err, c.hint)
		}
		if err == nil {
			//ensure sid is non-zero-length
			if len(sid) == 0 {
				t.Errorf("case %s: new SessionID is zero-length\nHINT: %s", c.name, c.hint)
			}
			//ensure sid is base64-url-encoded
			_, err := base64.URLEncoding.DecodeString(string(sid))
			if err != nil {
				t.Errorf("case %s: new SessionID failed base64-url-decoding: %v\nHINT: %s", c.name, err, c.hint)
			}
		}
	}
}

// TestToString tests making the session ID a string from the sid format
func TestToString(t *testing.T) {
	sid, err := NewSessionID("test key")
	if err != nil {
		t.Errorf("unexpected error generating new SessionID: %v", err)
	}
	s := sid.String()
	if len(s) == 0 {
		t.Error(".String() method returned zero-length string")
	}
}

func TestValidateID(t *testing.T) {
	cases := []struct {
		name          string
		hint          string
		signingKey    string
		validationKey string
		sidMutator    func(SessionID) SessionID
		expectError   bool
	}{
		{
			"Valid Key",
			"Remember to return the input `id` parameter as a SessionID with no error if `id` is valid",
			"test key",
			"test key",
			nil,
			false,
		},
		{
			"Different Key",
			"If the key used for validation is different than the key used to sign, it should return an error",
			"test key",
			"different key",
			nil,
			true,
		},
		{
			"Mutated ID Portion",
			"If the ID portion of the SessionID was mutated after it was generated, it should return an error",
			"test key",
			"test key",
			func(sid SessionID) SessionID {
				buf, _ := base64.URLEncoding.DecodeString(string(sid))
				buf[0] = buf[0] + 1
				return SessionID(base64.URLEncoding.EncodeToString(buf))
			},
			true,
		},
		{
			"Mutated Signature Portion",
			"If the signature portion of the SessionID was mutated after it was generated, it should return an error",
			"test key",
			"test key",
			func(sid SessionID) SessionID {
				buf, _ := base64.URLEncoding.DecodeString(string(sid))
				buf[len(buf)-2] = buf[len(buf)-2] + 1
				return SessionID(base64.URLEncoding.EncodeToString(buf))
			},
			true,
		},
		{
			"Invalid Base64 Encoding",
			"If the base64 decoding fails, it should return an error",
			"test key",
			"test key",
			func(sid SessionID) SessionID {
				buf, _ := base64.URLEncoding.DecodeString(string(sid))
				buf[0] = byte('+') // + is not part of base64-URL alphabet
				return SessionID(base64.URLEncoding.EncodeToString(buf))
			},
			true,
		},
		{
			"Incorrect Length",
			"If the base64 decoding fails, it should return an error",
			"test key",
			"test key",
			func(sid SessionID) SessionID {
				buf, _ := base64.URLEncoding.DecodeString(string(sid))
				return SessionID(base64.URLEncoding.EncodeToString(buf[0 : len(buf)-2]))
			},
			true,
		},
	}

	for _, c := range cases {
		sid, err := NewSessionID(c.signingKey)
		if err != nil {
			t.Errorf("case %s: unexpected error generating new SessionID: %v", c.name, err)
			continue
		}

		if c.sidMutator != nil && len(sid) > 0 {
			sid = c.sidMutator(sid)
		}

		sid2, err := ValidateID(string(sid), c.validationKey)
		if err != nil && !c.expectError {
			t.Errorf("case %s: unexpected error validating SessionID: %v\nHINT: %s", c.name, err, c.hint)
		}
		if c.expectError && err == nil {
			t.Errorf("case %s: expected error but didn't get one\nHINT: %s", c.name, c.hint)
		}

		if err == nil && sid2 != sid {
			t.Errorf("case %s: validated SessionID does not equal original SessionID\nHINT: %s", c.name, c.hint)
		}
	}
}
