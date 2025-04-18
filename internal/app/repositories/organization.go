package repositories

import (
	"management/internal/app/core"
	"management/internal/app/models"
)

type OrganizationGetManyParams struct {
	Type core.OrganizationType
}

type OrganizationRepository interface {
	GetMany(params OrganizationGetManyParams) ([]models.Organization, error)
}
