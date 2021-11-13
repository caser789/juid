package uuid

var (
	ifname string // name of interface being used
	nodeID []byte // hardware for version 1 UUIDs
)

// NodeID returns the 6 byte node id encoded in uuid.  It returns nil if uuid is
// not valid.  The NodeID is only well defined for version 1 and 2 UUIDs.
func (uuid UUID) NodeID() []byte {
	if len(uuid) != 16 {
		return nil
	}
	node := make([]byte, 6)
	copy(node, uuid[10:])
	return node
}

// SetNodeID sets the Node ID to be used for Version 1 UUIDs.  The first 6 bytes
// of id are used.  If id is less than 6 bytes then false is returned and the
// Node ID is not set.
func SetNodeID(id []byte) bool {
	if setNodeID(id) {
		ifname = "user"
		return true
	}
	return false
}

func setNodeID(id []byte) bool {
	if len(id) < 6 {
		return false
	}
	if nodeID == nil {
		nodeID = make([]byte, 6)
	}
	copy(nodeID, id)
	return true
}
