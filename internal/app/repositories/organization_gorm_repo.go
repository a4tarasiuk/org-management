package repositories

import (
	"gorm.io/gorm"
	"management/internal/app/models"
)

type organizationGORMRepository struct {
	db *gorm.DB
}

func (r organizationGORMRepository) GetMany(params OrganizationGetManyParams) ([]models.Organization, error) {
	var organizations []models.Organization

	result := r.db.Where("type = ?", params.Type).Find(&organizations)

	if result.Error != nil {
		return []models.Organization{}, result.Error
	}

	return organizations, nil

}
