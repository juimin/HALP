package sessions

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
)

//InvalidSessionID represents an empty, invalid session ID
const InvalidSessionID SessionID = ""

//idLength is the length of the ID portion
const idLength = 32

//signedLength is the full length of the signed session ID
//(ID portion plus signature)
const signedLength = idLength + sha256.Size

//SessionID represents a valid, digitally-signed session ID.
//This is a base64 URL encoded string created from a byte slice
//where the first `idLength` bytes are crytographically random
//bytes representing the unique session ID, and the remaining bytes
//are an HMAC hash of those ID bytes (i.e., a digital signature).
//The byte slice layout is like so:
//+-----------------------------------------------------+
//|...32 crypto random bytes...|HMAC hash of those bytes|
//+-----------------------------------------------------+
type SessionID string

//ErrInvalidID is returned when an invalid session id is passed to ValidateID()
var ErrInvalidID = errors.New("Invalid Session ID")

//NewSessionID creates and returns a new digitally-signed session ID,
//using `signingKey` as the HMAC signing key. An error is returned only
//if there was an error generating random bytes for the session ID
func NewSessionID(signingKey string) (SessionID, error) {
	//TODO: if `signingKey` is zero-length, return InvalidSessionID
	//and an error indicating that it may not be empty

	if len(signingKey) == 0 {
		return InvalidSessionID, ErrInvalidID
	}

	// Create random bytes slice of id length
	b := make([]byte, idLength)
	_, err := rand.Read(b)

	// If there was an error producing the bytes, return the invalid session id
	// error
	if err != nil {
		return InvalidSessionID, err
	}

	// Set the signing key of the HMAC hasher to be the signing key
	h := hmac.New(sha256.New, []byte(signingKey))
	h.Write(b)

	// Append the 32 bytes of random with the hmac hash at the end
	session := append(b, h.Sum(nil)...)

	// Make a new session id out of the base 64 url encoded byte slice
	if len(session) != signedLength {
		return InvalidSessionID, ErrInvalidID
	}
	return SessionID(base64.URLEncoding.EncodeToString(session)), nil
}

//ValidateID validates the string in the `id` parameter
//using the `signingKey` as the HMAC signing key
//and returns an error if invalid, or a SessionID if valid
func ValidateID(id string, signingKey string) (SessionID, error) {

	// If the length of id is nothing, return error
	if len(id) == 0 || len(signingKey) == 0 {
		return InvalidSessionID, ErrInvalidID
	}

	// Generate the hmac
	h := hmac.New(sha256.New, []byte(signingKey))

	// Decode the signature based on the signing key and the id
	decoded, err := base64.URLEncoding.DecodeString(id)

	// Check if there was a decoding error
	if err != nil {
		return InvalidSessionID, ErrInvalidID
	}

	// Write the id to the hmac and get the hash back
	h.Write(decoded[0:idLength])
	sig := h.Sum(nil)

	// Compare using constant time compare
	if subtle.ConstantTimeCompare(sig, decoded[idLength:len(decoded)]) == 1 {
		return SessionID(id), nil
	}
	// If we didn't return before, the id is invalid so we return that here
	return InvalidSessionID, ErrInvalidID
}

//String returns a string representation of the sessionID
func (sid SessionID) String() string {
	return string(sid)
}
