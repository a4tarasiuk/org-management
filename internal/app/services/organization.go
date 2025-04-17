package services

import (
	"gorm.io/gorm"
	"management/internal/app/orm"
)

type Service interface {
	Create(*orm.Organization) error
}

func NewOrganizationService(db *gorm.DB) Service {
	return service{db: db}
}

type service struct {
	db *gorm.DB
}

func (s service) Create(org *orm.Organization) error {
	if result := s.db.Create(&org); result.Error != nil {
		return result.Error
	}

	return nil
}
