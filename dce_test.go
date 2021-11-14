package uuid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDomainString(t *testing.T) {
	assert.Equal(t, Domain(DOMAIN_PERSON).String(), "DOMAIN_PERSON")
	assert.Equal(t, Domain(DOMAIN_GROUP).String(), "DOMAIN_GROUP")
	assert.Equal(t, Domain(DOMAIN_ORG).String(), "DOMAIN_ORG")
	assert.Equal(t, Domain(4).String(), "DOMAIN_4")
	assert.Equal(t, Domain(5).String(), "DOMAIN_5")
}

func TestUUIDId(t *testing.T) {
	tests := []struct {
		name            string
		uuid            UUID
		expectedId      uint32
		expectedSuccess bool
	}{
		{
			name: "test invalid version 2 id",
			uuid: UUID([]byte{
				0, 0, 0, 0,
				1, 2, 0b11111111, 4,
				1, 2, 3, 4,
				1, 2, 3, 4,
			}),
			expectedId:      uint32(0),
			expectedSuccess: false,
		},
		{
			name: "test valid version 2 id",
			uuid: UUID([]byte{
				0, 0, 0, 1,
				1, 2, 0b00101111, 4,
				1, 2, 3, 4,
				1, 2, 3, 4,
			}),
			expectedId:      uint32(1),
			expectedSuccess: true,
		},
	}

	for _, tt := range tests {
		id, success := tt.uuid.Id()

		assert.Equal(t, tt.expectedId, id, tt.name)
		assert.Equal(t, tt.expectedSuccess, success, tt.name)
	}
}

func TestUUIDDomain(t *testing.T) {
	tests := []struct {
		name            string
		uuid            UUID
		expectedDomain  Domain
		expectedSuccess bool
	}{
		{
			name: "test invalid version 2 domain",
			uuid: UUID([]byte{
				0, 0, 0, 0,
				1, 2, 0b11111111, 4,
				1, 2, 3, 4,
				1, 2, 3, 4,
			}),
			expectedDomain:  DOMAIN_PERSON,
			expectedSuccess: false,
		},
		{
			name: "test valid version 2 domain - DOMAIN_PERSON",
			uuid: UUID([]byte{
				0, 0, 0, 0,
				1, 2, 0b00101111, 4,
				1, 0b00000000, 3, 4,
				1, 2, 3, 4,
			}),
			expectedDomain:  DOMAIN_PERSON,
			expectedSuccess: true,
		},
		{
			name: "test valid version 2 domain - DOMAIN_GROUP",
			uuid: UUID([]byte{
				0, 0, 0, 0,
				1, 2, 0b00101111, 4,
				1, 0b00000001, 3, 4,
				1, 2, 3, 4,
			}),
			expectedDomain:  DOMAIN_GROUP,
			expectedSuccess: true,
		},
		{
			name: "test valid version 2 domain - DOMAIN_ORG",
			uuid: UUID([]byte{
				0, 0, 0, 0,
				1, 2, 0b00101111, 4,
				1, 0b00000010, 3, 4,
				1, 2, 3, 4,
			}),
			expectedDomain:  DOMAIN_ORG,
			expectedSuccess: true,
		},
		{
			name: "test valid version 2 domain - Domain(3)",
			uuid: UUID([]byte{
				0, 0, 0, 0,
				1, 2, 0b00101111, 4,
				1, 0b00000011, 3, 4,
				1, 2, 3, 4,
			}),
			expectedDomain:  3,
			expectedSuccess: true,
		},
	}

	for _, tt := range tests {
		domain, success := tt.uuid.Domain()

		assert.Equal(t, tt.expectedDomain, domain, tt.name)
		assert.Equal(t, tt.expectedSuccess, success, tt.name)
	}
}
