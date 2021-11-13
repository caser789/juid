package uuid

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_randomBits(t *testing.T) {
	data := make([]byte, 10, 10)
	randomBits(data)
	fmt.Println(data)

	assert.NotEqual(t, make([]byte, 10, 10), data)
}
