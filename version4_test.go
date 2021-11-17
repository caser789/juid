package uuid

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersion4NewRandom(t *testing.T) {
	for i := 0; i < 1000; i++ {
		u := New()

		assert.Equal(t, u[6]&0b11110000, uint8(0x46)&0b11110000)
		assert.Equal(t, u[8]&0b11000000, uint8(0x8e)&0b11000000)
	}
}

func TestRandomFromReader(t *testing.T) {
	myString := "8059ddhdle77cb52"
	r := bytes.NewReader([]byte(myString))
	r2 := bytes.NewReader([]byte(myString))
	uuid1, err := NewRandomFromReader(r)
	if err != nil {
		t.Errorf("failed generating UUID from a reader")
	}
	_, err = NewRandomFromReader(r)
	if err == nil {
		t.Errorf("expecting an error as reader has no more bytes. Got uuid. NewRandomFromReader may not be using the provided reader")
	}
	uuid3, err := NewRandomFromReader(r2)
	if err != nil {
		t.Errorf("failed generating UUID from a reader")
	}
	if uuid1 != uuid3 {
		t.Errorf("expected duplicates, got %q and %q", uuid1, uuid3)
	}
}
