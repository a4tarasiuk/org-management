package handlers

import (
	"management/internal/app/core"
	"management/internal/app/orm"
)

type createOrganizationRequest struct {
	Name string `json:"name" binding:"required"`
}

type createOrganizationResponse struct {
	ID uint `json:"id"`

	Name string `json:"name"`

	Type core.OrganizationType `json:"type"`
}

func organizationResponseFromOrganization(org orm.Organization) createOrganizationResponse {
	return createOrganizationResponse{
		ID:   org.ID,
		Name: org.Name,
		Type: org.Type,
	}
}

type ErrorResponse struct {
	Error string `json:"error"`
}
