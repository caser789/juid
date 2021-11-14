package uuid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUUIDNodeId(t *testing.T) {
	uuid := UUID(make([]byte, 6))
	node := uuid.NodeID()
	assert.Equal(t, node, []byte(nil))

	uuid = UUID([]byte{
		1, 2, 3, 4,
		1, 2, 3, 4,
		1, 2, 3, 4,
		1, 2, 3, 4,
	})
	node = uuid.NodeID()
	assert.Equal(t, node, []byte{3, 4, 1, 2, 3, 4})
}
