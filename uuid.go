package uuid

import (
	"bytes"
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

// A Variant represents a UUIDs variant.
type Variant byte

// Constants returned by Variant.
const (
	INVALID   = iota // Invalid UUID
	RFC4122          // The variant specified in RFC4122
	RESERVED         // Reserved, NCS backward compatibility.
	MICROSOFT        // Reserved, Microsoft Corporation backward compatibility.
	FUTURE           // Reserved for future definition.
)

func (v Variant) String() string {
	switch v {
	case RFC4122:
		return "RFC4122"
	case RESERVED:
		return "RESERVED"
	case MICROSOFT:
		return "MICROSOFT"
	case FUTURE:
		return "FUTURE"
	case INVALID:
		return "INVALID"
	}
	return fmt.Sprintf("BAD_VARIANT_%d", v)
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

// Variant returns the variant encoded in uuid.  It returns INVALID if
// uuid is invalid.
func (uuid UUID) Variant() Variant {
	if len(uuid) != 16 {
		return INVALID
	}
	switch {
	case (uuid[8] & 0xc0) == 0x80:
		return RFC4122
	case (uuid[8] & 0xe0) == 0xc0:
		return MICROSOFT
	case (uuid[8] & 0xe0) == 0xe0:
		return FUTURE
	default:
		return RESERVED
	}
	panic("unreachable")
}

// Version returns the verison of uuid.  It returns false if uuid is not
// valid.
func (uuid UUID) Version() (Version, bool) {
	if len(uuid) != 16 {
		return 0, false
	}
	return Version(uuid[6] >> 4), true
}

// Equal returns true if uuid1 and uuid2 are equal.
func Equal(uuid1, uuid2 UUID) bool {
	return bytes.Equal(uuid1, uuid2)
}
