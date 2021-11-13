package uuid

import (
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
