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

func TestSetNodeID(t *testing.T) {
	v := SetNodeID([]byte{1, 2, 3, 4, 5})
	assert.Equal(t, v, false)
	assert.Equal(t, ifname, "")
	assert.Equal(t, nodeID, []byte(nil))

	v = SetNodeID([]byte{1, 2, 3, 4, 5, 6, 7})
	assert.Equal(t, v, true)
	assert.Equal(t, ifname, "user")
	assert.Equal(t, nodeID, []byte{1, 2, 3, 4, 5, 6})
}
