package uuid

import (
	"encoding/binary"
	"fmt"
	"os"
)

// A Domain represents a Version 2 domain
type Domain byte

// Domain constants for DCE Security (Version 2) UUIDs.
const (
	Person Domain = 0
	Group         = 1
	Org           = 2
)

func (d Domain) String() string {
	switch d {
	case Person:
		return "Person"
	case Group:
		return "Group"
	case Org:
		return "Org"
	}
	return fmt.Sprintf("Domain%d", d)
}

// Id returns the id for a Version 2 UUID. Ids are only defined for Vrsion 2
// UUIDs.
func (uuid UUID) Id() uint32 {
	return binary.BigEndian.Uint32(uuid[0:4])
}

// Domain returns the domain for a Version 2 UUID.  Domains are only defined
// for Version 2 UUIDs.
func (uuid UUID) Domain() Domain {
	return Domain(uuid[9])
}

// NewDCESecurity returns a DCE Security (Version 2) UUID.
//
// The domain should be one of DOMAIN_PERSON, DOMAIN_GROUP or DOMAIN_ORG.
// On a POSIX system the id should be the users UID for the DOMAIN_PERSON
// domain and the users GID for the DOMAIN_GROUP.  The meaning of id for
// the domain DOMAIN_ORG or on non-POSIX systems is site defined.
//
// For a given domain/id pair the same token may be returned for up to
// 7 minutes and 10 seconds.
func NewDCESecurity(domain Domain, id uint32) (UUID, error) {
	uuid, err := NewUUID()
	if err == nil {
		uuid[6] = (uuid[6] & 0x0f) | 0x20 // Version 2
		uuid[9] = byte(domain)
		binary.BigEndian.PutUint32(uuid[0:], id)
	}
	return uuid, err
}

// NewDCEGroup returns a DCE Security (Version 2) UUID in the group
// domain with the id returned by os.Getuid.
//
//  NewDCESecurity(DOMAIN_PERSON, uint32(os.Getuid()))
func NewDCEPerson() (UUID, error) {
	return NewDCESecurity(Person, uint32(os.Getuid()))
}

// NewDCEPerson returns a DCE Security (Version 2) UUID in the group
// domain with the id returned by os.Getgid.
//
//  NewDCESecurity(DOMAIN_GROUP, uint32(os.Getgid()))
func NewDCEGroup() (UUID, error) {
	return NewDCESecurity(Group, uint32(os.Getgid()))
}
