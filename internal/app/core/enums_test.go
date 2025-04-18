package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationType_String(t *testing.T) {
	testCases := []struct {
		orgType            OrganizationType
		expectedStringName string
	}{
		{DistributorOrganization, "DISTRIBUTOR"},
		{ClientOrganization, "CLIENT"},
		{234, UnknownOrganizationTypeString},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.orgType.String(), tc.expectedStringName)
	}
}
