package uuid

import (
	"fmt"
)

// MarshalText implements encoding.TextMarshaler.
func (u UUID) MarshalText() ([]byte, error) {
	var js [36]byte
	encodeHex(js[:], u)
	return js[:], nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (u *UUID) UnmarshalText(data []byte) error {
	// See comment in ParseBytes why we do this.
	// id, err := ParseBytes(data)
	id, err := ParseBytes(data)
	if err == nil {
		*u = id
	}
	return err
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (u UUID) MarshalBinary() ([]byte, error) {
	return u[:], nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (u *UUID) UnmarshalBinary(data []byte) error {
	if len(data) != 16 {
		return fmt.Errorf("invalid UUID (got %d bytes)", len(data))
	}
	copy(u[:], data)
	return nil
}
