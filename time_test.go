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
