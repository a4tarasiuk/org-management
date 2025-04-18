package models

import (
	"management/internal/app/core"
)

type Organization struct {
	ID uint `gorm:"primarykey"`

	Name string `gorm:"not null;uniqueIndex;size:128"`

	Type core.OrganizationType `gorm:"not null"`

	ParentID *uint

	ExternalID *uint `gorm:"index"`

	GroupID *uint `gorm:"index"`
}
