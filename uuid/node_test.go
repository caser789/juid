package uuid

import (
	"bytes"
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

func TestNodeID(t *testing.T) {
	nid := []byte{1, 2, 3, 4, 5, 6}
	SetNodeInterface("")
	s := NodeInterface()
	if s == "" || s == "user" {
		t.Errorf("NodeInterface %q after SetInteface\n", s)
	}
	node1 := NodeID()
	if node1 == nil {
		t.Errorf("NodeID %q nil after SetNodeInterface\n", s)
	}
	SetNodeID(nid)
	s = NodeInterface()
	if s != "user" {
		t.Errorf("Expected NodeInterface %q got %q\n", "user", s)
	}
	node2 := NodeID()
	if node2 == nil {
		t.Errorf("NodeID %q nil after SetNodeID\n", s)
	}
	if bytes.Equal(node1, node2) {
		t.Errorf("NodeID %q not changed after SetNodeID\n", s)
	} else if !bytes.Equal(nid, node2) {
		t.Errorf("NodeID is %x, expected %x\n", node2, nid)
	}
}
