package repositories

import "gorm.io/gorm"

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return organizationGORMRepository{db: db}
}
