package uuid

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestClockSequence(t *testing.T) {
	clockSeq = 0b0011111111111111
	v := ClockSequence()
	assert.Equal(t, v, 0b0011111111111111)

	clockSeq = 0b0111111111111111
	v = ClockSequence()
	assert.Equal(t, v, 0b0011111111111111)

	clockSeq = 0b1011111111111111
	v = ClockSequence()
	assert.Equal(t, v, 0b0011111111111111)

	clockSeq = 0b1111111111111111
	v = ClockSequence()
	assert.Equal(t, v, 0b0011111111111111)

	clockSeq = 0
	v = ClockSequence()
	assert.NotEqual(t, v, 0)
}

func TestSetClockSequence(t *testing.T) {
	lasttime = 111
	SetClockSequence(0b0011111111111111)
	assert.Equal(t, clockSeq, uint16(0b1011111111111111))
	assert.Equal(t, lasttime, uint64(0))

	clockSeq = 0b1011111111111111
	lasttime = 222
	SetClockSequence(0b0011111111111111)
	assert.Equal(t, clockSeq, uint16(0b1011111111111111))
	assert.Equal(t, lasttime, uint64(222))

	SetClockSequence(-1)
	assert.Equal(t, lasttime, uint64(0))
}

func TestClockSeq(t *testing.T) {
	// Fake time.Now for this test to return a monotonically advancing time; restore it at end.
	defer func(orig func() time.Time) { timeNow = orig }(timeNow)
	monTime := time.Now()
	timeNow = func() time.Time {
		monTime = monTime.Add(1 * time.Second)
		return monTime
	}

	SetClockSequence(-1)
	uuid1 := Must(NewUUID())
	uuid2 := Must(NewUUID())

	if _clockSeq(t, uuid1) != _clockSeq(t, uuid2) {
		t.Errorf("clock sequence %d != %d\n", _clockSeq(t, uuid1), _clockSeq(t, uuid2))
	}

	SetClockSequence(-1)
	uuid2 = Must(NewUUID())

	// Just on the very off chance we generated the same sequence
	// two times we try again.
	if _clockSeq(t, uuid1) == _clockSeq(t, uuid2) {
		SetClockSequence(-1)
		uuid2 = Must(NewUUID())
	}
	if _clockSeq(t, uuid1) == _clockSeq(t, uuid2) {
		t.Errorf("Duplicate clock sequence %d\n", _clockSeq(t, uuid1))
	}

	SetClockSequence(0x1234)
	uuid1 = Must(NewUUID())
	if seq := _clockSeq(t, uuid1); seq != 0x1234 {
		t.Errorf("%s: expected seq 0x1234 got 0x%04x\n", uuid1, seq)
	}
}

func _clockSeq(t *testing.T, uuid UUID) int {
	seq := uuid.ClockSequence()
	return seq
}
