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
	uuid := UUID([]byte{
		0, 0, 0, 0,
		1, 2, 0b11111111, 4,
		1, 2, 3, 4,
		1, 2, 3, 4,
	})
	id, yes := uuid.Id()
	assert.Equal(t, yes, false)
	assert.Equal(t, id, uint32(0))

	uuid = UUID([]byte{
		0, 0, 0, 1,
		1, 2, 0b00101111, 4,
		1, 2, 3, 4,
		1, 2, 3, 4,
	})
	id, yes = uuid.Id()
	assert.Equal(t, yes, true)
	assert.Equal(t, id, uint32(1))
}
