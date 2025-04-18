package handlers

import (
	"management/internal/app/core"
)

type createOrganizationRequest struct {
	Name string `json:"name" binding:"required"`
}

type createOrganizationResponse struct {
	ID uint `json:"id"`

	Name string `json:"name"`

	Type core.OrganizationType `json:"type"`
}

type organizationListItemSchema struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	ExternalID *uint  `json:"external_id"`
}

type organizationListResponse struct {
	Organizations []organizationListItemSchema `json:"organizations"`
}

type errorResponse struct {
	Error string `json:"error"`
}
