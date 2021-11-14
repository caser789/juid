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
