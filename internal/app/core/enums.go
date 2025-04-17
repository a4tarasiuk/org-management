package core

type OrganizationType uint8

const (
	_ OrganizationType = iota
	DistributorOrganization
	ClientOrganization
)
