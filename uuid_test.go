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

func TestVariant(t *testing.T) {
	assert.Equal(t, Variant(INVALID).String(), "INVALID")
	assert.Equal(t, Variant(RFC4122).String(), "RFC4122")
	assert.Equal(t, Variant(RESERVED).String(), "RESERVED")
	assert.Equal(t, Variant(MICROSOFT).String(), "MICROSOFT")
	assert.Equal(t, Variant(FUTURE).String(), "FUTURE")
	assert.Equal(t, Variant(5).String(), "BAD_VARIANT_5")
	assert.Equal(t, Variant(6).String(), "BAD_VARIANT_6")
}

func TestRandomUUID(t *testing.T) {
	m := make(map[string]bool)
	for x := 1; x < 32; x++ {
		uuid := NewRandom()
		s := uuid.String()
		if m[s] {
			t.Errorf("NewRandom returned duplicated UUID %s\n", s)
		}
		m[s] = true
		if uuid.Variant() != RFC4122 {
			t.Errorf("Random UUID is variant %d\n", uuid.Variant())
		}
		if v, _ := uuid.Version(); v != 4 {
			t.Errorf("Random UUID of version %s\n", v)
		}
	}
}
