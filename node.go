package uuid

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
