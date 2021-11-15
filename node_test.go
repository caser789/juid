package uuid

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUUIDNodeId(t *testing.T) {
	uuid := UUID{
		1, 2, 3, 4,
		1, 2, 3, 4,
		1, 2, 3, 4,
		1, 2, 3, 4,
	}
	node := uuid.NodeID()
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

func TestNode(t *testing.T) {
	// This test is mostly to make sure we don't leave nodeMu locked.
	ifname = ""
	if ni := NodeInterface(); ni != "" {
		t.Errorf("NodeInterface got %q, want %q", ni, "")
	}
	if SetNodeInterface("xyzzy") {
		t.Error("SetNodeInterface succeeded on a bad interface name")
	}
	if !SetNodeInterface("") {
		t.Error("SetNodeInterface failed")
	}
	if ni := NodeInterface(); ni == "" {
		t.Error("NodeInterface returned an empty string")
	}

	ni := NodeID()
	if len(ni) != 6 {
		t.Errorf("ni got %d bytes, want 6", len(ni))
	}
	hasData := false
	for _, b := range ni {
		if b != 0 {
			hasData = true
		}
	}
	if !hasData {
		t.Error("nodeid is all zeros")
	}

	id := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	SetNodeID(id)
	ni = NodeID()
	if !bytes.Equal(ni, id[:6]) {
		t.Errorf("got nodeid %v, want %v", ni, id[:6])
	}

	if ni := NodeInterface(); ni != "user" {
		t.Errorf("got inteface %q, want %q", ni, "user")
	}
}
