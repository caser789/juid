package uuid

import (
	"crypto/rand"
	"io"
)

var rander = rand.Reader

// randomBits completely fills slice b with random data.
func randomBits(b []byte) {
	if _, err := io.ReadFull(rander, b); err != nil {
		panic(err) // rand should never fail
	}
}
