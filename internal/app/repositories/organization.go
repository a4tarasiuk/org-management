package repositories

import (
	"management/internal/app/dto"
	"management/internal/app/orm"
)

type OrganizationRepository interface {
	GetMany(params dto.ListParams) ([]orm.Organization, error)
}
