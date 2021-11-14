package uuid

import (
	"encoding/binary"
	"fmt"
)

// A Domain represents a Version 2 domain
type Domain byte

// Domain constants for DCE Security (Version 2) UUIDs.
const (
	DOMAIN_PERSON = 0
	DOMAIN_GROUP  = 1
	DOMAIN_ORG    = 2
)

func (d Domain) String() string {
	switch d {
	case DOMAIN_PERSON:
		return "DOMAIN_PERSON"
	case DOMAIN_GROUP:
		return "DOMAIN_GROUP"
	case DOMAIN_ORG:
		return "DOMAIN_ORG"
	}
	return fmt.Sprintf("DOMAIN_%d", d)
}

// Id returns the id for a Version 2 UUID or false.
func (uuid UUID) Id() (uint32, bool) {
	if v, _ := uuid.Version(); v != 2 {
		return 0, false
	}
	return binary.BigEndian.Uint32(uuid[0:4]), true
}
