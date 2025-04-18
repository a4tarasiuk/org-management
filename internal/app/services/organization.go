package services

import (
	"gorm.io/gorm"
	"management/internal/app/models"
)

type Service interface {
	Create(*models.Organization) error
}

func NewOrganizationService(db *gorm.DB) Service {
	return service{db: db}
}

type service struct {
	db *gorm.DB
}

func (s service) Create(org *models.Organization) error {
	if result := s.db.Create(&org); result.Error != nil {
		return result.Error
	}

	return nil
}
