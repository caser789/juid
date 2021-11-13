package uuid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type badRand struct{}

func (r badRand) Read(buf []byte) (int, error) {
	for i, _ := range buf {
		buf[i] = byte(i)
	}
	return len(buf), nil
}

func TestBadRand(t *testing.T) {
	SetRand(badRand{})
	uuid1 := New()
	uuid2 := New()
	if uuid1 != uuid2 {
		t.Errorf("expected duplicates, got %q and %q\n", uuid1, uuid2)
	}

	SetRand(nil)
	uuid1 = New()
	uuid2 = New()
	if uuid1 == uuid2 {
		t.Errorf("unexpected duplicates, got %q and %q\n", uuid1, uuid2)
	}
}

func TestVersionString(t *testing.T) {
	assert.Equal(t, Version(0b00000000).String(), "VERSION_0")
	assert.Equal(t, Version(0b00000001).String(), "VERSION_1")
	assert.Equal(t, Version(0b00000010).String(), "VERSION_2")
	assert.Equal(t, Version(0b00000011).String(), "VERSION_3")
	assert.Equal(t, Version(0b00000100).String(), "VERSION_4")
	assert.Equal(t, Version(0b00000101).String(), "VERSION_5")
	assert.Equal(t, Version(0b00000110).String(), "VERSION_6")
	assert.Equal(t, Version(0b00000111).String(), "VERSION_7")
	assert.Equal(t, Version(0b00001000).String(), "VERSION_8")
	assert.Equal(t, Version(0b00001001).String(), "VERSION_9")
	assert.Equal(t, Version(0b00001010).String(), "VERSION_10")
	assert.Equal(t, Version(0b00001011).String(), "VERSION_11")
	assert.Equal(t, Version(0b00001100).String(), "VERSION_12")
	assert.Equal(t, Version(0b00001101).String(), "VERSION_13")
	assert.Equal(t, Version(0b00001110).String(), "VERSION_14")
	assert.Equal(t, Version(0b00001111).String(), "VERSION_15")
	assert.Equal(t, Version(0b00010000).String(), "BAD_VERSION_16")
	assert.Equal(t, Version(0b00010001).String(), "BAD_VERSION_17")
}
