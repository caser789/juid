package uuid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_randomBits(t *testing.T) {
	data := make([]byte, 10, 10)
	randomBits(data)

	assert.NotEqual(t, make([]byte, 10, 10), data)
}

func Test_xvalues(t *testing.T) {
	assert.Equal(t, xvalues['a'], uint8(0xa))
	assert.Equal(t, xvalues['b'], uint8(0xb))
	assert.Equal(t, xvalues['c'], uint8(0xc))
	assert.Equal(t, xvalues['d'], uint8(0xd))
	assert.Equal(t, xvalues['e'], uint8(0xe))
	assert.Equal(t, xvalues['f'], uint8(0xf))

	assert.Equal(t, xvalues['A'], uint8(0xA))
	assert.Equal(t, xvalues['B'], uint8(0xB))
	assert.Equal(t, xvalues['C'], uint8(0xC))
	assert.Equal(t, xvalues['D'], uint8(0xD))
	assert.Equal(t, xvalues['E'], uint8(0xE))
	assert.Equal(t, xvalues['F'], uint8(0xF))

	assert.Equal(t, xvalues['0'], uint8(0x0))
	assert.Equal(t, xvalues['1'], uint8(0x1))
	assert.Equal(t, xvalues['2'], uint8(0x2))
	assert.Equal(t, xvalues['3'], uint8(0x3))
	assert.Equal(t, xvalues['4'], uint8(0x4))
	assert.Equal(t, xvalues['5'], uint8(0x5))
	assert.Equal(t, xvalues['6'], uint8(0x6))
	assert.Equal(t, xvalues['7'], uint8(0x7))
	assert.Equal(t, xvalues['8'], uint8(0x8))
	assert.Equal(t, xvalues['9'], uint8(0x9))
}
