package webid

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// A WebID is an ID that is cryptographicaly secure and unique, so that
// it can be used as a unique identifier in places like database keys.
// It's also URL-formatted, reasonably short, and uses human-readable
// Base 64 characters so it can appear in links and URLs safely
// and in a non-ugly way.
type WebID string

// New() returns a new WebID generated from an 8 byte random number.
func New() (string, error) {
	c := 8
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("generating random bytes; %w", err)
	}
	// The slice should now contain random bytes instead of only zeroes.
	s := base64.RawURLEncoding.EncodeToString(b)
	return s, nil
}

func MustNew() string {
	id, err := New()
	if err != nil {
		panic(err)
	}
	return id
}