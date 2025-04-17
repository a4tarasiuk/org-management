package persistence

import (
	"gorm.io/gorm"
	"management/internal/app/dto"
	"management/internal/app/orm"
)

func NewOrganizationRepository(db *gorm.DB) organizationGORMRepository {
	return organizationGORMRepository{db: db}
}

type organizationGORMRepository struct {
	db *gorm.DB
}

func (r organizationGORMRepository) GetMany(params dto.ListParams) ([]orm.Organization, error) {
	var organizations []orm.Organization

	result := r.db.Where("type = ?", params.Type).Find(&organizations)

	if result.Error != nil {
		return []orm.Organization{}, result.Error
	}

	return organizations, nil

}
