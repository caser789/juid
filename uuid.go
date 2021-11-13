package uuid

import (
	"crypto/rand"
	"fmt"
	"io"
)

var rander = rand.Reader

// SetRand sets the random number generator to r, which implents io.Reader.
// If r.Read returns an error when the package requests random data then
// a panic will be issued.
//
// Calling SetRand with nil sets the random number generator to the default
// generator.
func SetRand(r io.Reader) {
	if r == nil {
		rander = rand.Reader
		return
	}
	rander = r
}

// A Version represents a UUIDs version.
type Version byte

func (v Version) String() string {
	if v > 15 {
		return fmt.Sprintf("BAD_VERSION_%d", v)
	}
	return fmt.Sprintf("VERSION_%d", v)
}

// A UUID is a 128 bit (16 byte) Universal Unique IDentifier as defined in RFC
// 4122.
type UUID []byte

// New returns a new random (version 4) UUID as a string.  It is a convenience
// function for NewRandom().String().
func New() string {
	return NewRandom().String()
}

// String returns the string form of uuid, xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
// , or "" if uuid is invalid.
func (uuid UUID) String() string {
	if uuid == nil || len(uuid) != 16 {
		return ""
	}
	b := []byte(uuid)
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[:4], b[4:6], b[6:8], b[8:10], b[10:])
}
