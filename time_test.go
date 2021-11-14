package uuid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClockSequence(t *testing.T) {
	clock_seq = 0b0011111111111111
	v := ClockSequence()
	assert.Equal(t, v, 0b0011111111111111)

	clock_seq = 0b0111111111111111
	v = ClockSequence()
	assert.Equal(t, v, 0b0011111111111111)

	clock_seq = 0b1011111111111111
	v = ClockSequence()
	assert.Equal(t, v, 0b0011111111111111)

	clock_seq = 0b1111111111111111
	v = ClockSequence()
	assert.Equal(t, v, 0b0011111111111111)

	clock_seq = 0
	v = ClockSequence()
	assert.NotEqual(t, v, 0)
}

func TestSetClockSequence(t *testing.T) {
	lasttime = 111
	SetClockSequence(0b0011111111111111)
	assert.Equal(t, clock_seq, uint16(0b1011111111111111))
	assert.Equal(t, lasttime, uint64(0))

	clock_seq = 0b1011111111111111
	lasttime = 222
	SetClockSequence(0b0011111111111111)
	assert.Equal(t, clock_seq, uint16(0b1011111111111111))
	assert.Equal(t, lasttime, uint64(222))

	SetClockSequence(-1)
	assert.Equal(t, lasttime, uint64(0))
}

func TestClockSeq(t *testing.T) {
	SetClockSequence(-1)
	uuid1 := NewUUID()
	uuid2 := NewUUID()

	if clockSeq(t, uuid1) == clockSeq(t, uuid2) {
		t.Errorf("clock sequence %d == %d\n", clockSeq(t, uuid1), clockSeq(t, uuid2))
	}

	SetClockSequence(-1)
	uuid2 = NewUUID()

	// Just on the very off chance we generated the same sequence
	// two times we try again.
	if clockSeq(t, uuid1) == clockSeq(t, uuid2) {
		SetClockSequence(-1)
		uuid2 = NewUUID()
	}
	if clockSeq(t, uuid1) == clockSeq(t, uuid2) {
		t.Errorf("Duplicate clock sequence %d\n", clockSeq(t, uuid1))
	}

	SetClockSequence(0x1234)
	uuid1 = NewUUID()
	if seq := clockSeq(t, uuid1); seq != 0x1234 {
		t.Errorf("%s: expected seq 0x1234 got 0x%04x\n", uuid1, seq)
	}
}

func clockSeq(t *testing.T, uuid UUID) int {
	seq, ok := uuid.ClockSequence()
	if !ok {
		t.Fatalf("%s: invalid clock sequence\n", uuid)
	}
	return seq
}
