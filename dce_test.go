package uuid

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDomainString(t *testing.T) {
	assert.Equal(t, Domain(Person).String(), "Person")
	assert.Equal(t, Domain(Group).String(), "Group")
	assert.Equal(t, Domain(Org).String(), "Org")
	assert.Equal(t, Domain(4).String(), "Domain4")
	assert.Equal(t, Domain(5).String(), "Domain5")
}

func TestUUIDId(t *testing.T) {
	tests := []struct {
		name       string
		uuid       UUID
		expectedId uint32
	}{
		{
			name: "test invalid version 2 id",
			uuid: UUID{
				0, 0, 0, 0,
				1, 2, 0b11111111, 4,
				1, 2, 3, 4,
				1, 2, 3, 4,
			},
			expectedId: uint32(0),
		},
		{
			name: "test valid version 2 id",
			uuid: UUID{
				0, 0, 0, 1,
				1, 2, 0b00101111, 4,
				1, 2, 3, 4,
				1, 2, 3, 4,
			},
			expectedId: uint32(1),
		},
	}

	for _, tt := range tests {
		id := tt.uuid.Id()

		assert.Equal(t, tt.expectedId, id, tt.name)
	}
}

func TestUUIDDomain(t *testing.T) {
	tests := []struct {
		name           string
		uuid           UUID
		expectedDomain Domain
	}{
		{
			name: "test invalid version 2 domain",
			uuid: UUID{
				0, 0, 0, 0,
				1, 2, 0b11111111, 4,
				1, 8, 3, 4,
				1, 2, 3, 4,
			},
			expectedDomain: Domain(8),
		},
		{
			name: "test valid version 2 domain - Person",
			uuid: UUID{
				0, 0, 0, 0,
				1, 2, 0b00101111, 4,
				1, 0b00000000, 3, 4,
				1, 2, 3, 4,
			},
			expectedDomain: Person,
		},
		{
			name: "test valid version 2 domain - Group",
			uuid: UUID{
				0, 0, 0, 0,
				1, 2, 0b00101111, 4,
				1, 0b00000001, 3, 4,
				1, 2, 3, 4,
			},
			expectedDomain: Group,
		},
		{
			name: "test valid version 2 domain - Org",
			uuid: UUID{
				0, 0, 0, 0,
				1, 2, 0b00101111, 4,
				1, 0b00000010, 3, 4,
				1, 2, 3, 4,
			},
			expectedDomain: Org,
		},
		{
			name: "test valid version 2 domain - Domain(3)",
			uuid: UUID{
				0, 0, 0, 0,
				1, 2, 0b00101111, 4,
				1, 0b00000011, 3, 4,
				1, 2, 3, 4,
			},
			expectedDomain: 3,
		},
	}

	for _, tt := range tests {
		domain := tt.uuid.Domain()

		assert.Equal(t, tt.expectedDomain, domain, tt.name)
	}
}

func testDCE(t *testing.T, name string, uuid UUID, domain Domain, id uint32) {
	if uuid == NIL {
		t.Errorf("%s failed\n", name)
		return
	}
	if v := uuid.Version(); v != 2 {
		t.Errorf("%s: %s: expected version 2, got %s\n", name, uuid, v)
		return
	}
	if v := uuid.Domain(); v != domain {
		t.Errorf("%s: %s: expected domain %d, got %d\n", name, uuid, domain, v)
	}
	if v := uuid.Id(); v != id {
		t.Errorf("%s: %s: expected id %d, got %d\n", name, uuid, id, v)
	}
}

func TestDCE(t *testing.T) {
	u, _ := NewDCESecurity(42, 12345678)
	testDCE(t, "NewDCESecurity", u, 42, 12345678)
	u, _ = NewDCEPerson()
	testDCE(t, "NewDCEPerson", u, Person, uint32(os.Getuid()))
	u, _ = NewDCEGroup()
	testDCE(t, "NewDCEGroup", u, Group, uint32(os.Getgid()))
}
