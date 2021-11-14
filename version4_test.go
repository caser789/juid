package uuid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersion4NewRandom(t *testing.T) {
	for i := 0; i < 1000; i++ {
		u := NewRandom()

		assert.Equal(t, u[6]&0b11110000, uint8(0x46)&0b11110000)
		assert.Equal(t, u[8]&0b11000000, uint8(0x8e)&0b11000000)
	}
}
