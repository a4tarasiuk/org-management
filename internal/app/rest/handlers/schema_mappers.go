package handlers

import "management/internal/app/models"

func organizationResponseFromOrganization(org models.Organization) createOrganizationResponse {
	return createOrganizationResponse{
		ID:   org.ID,
		Name: org.Name,
		Type: org.Type.String(),
	}
}

func organizationListSchemaFromOrganizations(organizations []models.Organization) organizationListResponse {
	organizationItemSchemas := make([]organizationListItemSchema, 0)

	for _, org := range organizations {
		organizationItemSchemas = append(
			organizationItemSchemas, organizationListItemSchema{
				ID:         org.ID,
				Name:       org.Name,
				Type:       org.Type.String(),
				ExternalID: org.ExternalID,
			},
		)
	}

	return organizationListResponse{Organizations: organizationItemSchemas}
}
