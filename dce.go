package uuid

import (
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
