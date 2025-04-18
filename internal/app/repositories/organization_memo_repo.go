package repositories

import "management/internal/app/models"

type OrganizationInMemoryRepository struct {
	Organizations []models.Organization
}

func (r OrganizationInMemoryRepository) GetMany(params OrganizationGetManyParams) ([]models.Organization, error) {
	resultOrganizations := make([]models.Organization, 0)

	for _, org := range r.Organizations {
		if org.Type == params.Type {
			resultOrganizations = append(resultOrganizations, org)
		}
	}

	return resultOrganizations, nil
}
