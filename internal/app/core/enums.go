package core

type OrganizationType uint8

const UnknownOrganizationTypeString string = "unknown organization type"

const (
	_ OrganizationType = iota
	DistributorOrganization
	ClientOrganization
)

func (t OrganizationType) String() string {
	switch t {
	case DistributorOrganization:
		return "DISTRIBUTOR"
	case ClientOrganization:
		return "CLIENT"
	default:
		return UnknownOrganizationTypeString
	}
}
