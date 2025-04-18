package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"management/internal/app/core"
	"management/internal/app/models"
	"management/internal/app/repositories"
	"management/internal/testutil"
)

func TestListDistributorsHandler(t *testing.T) {
	t.Run(
		"Get many 200 OK", func(t *testing.T) {
			w := httptest.NewRecorder()

			ctx := testutil.GetTestGinContext(w)

			testutil.MockJsonGet(ctx)

			distributorOrg1 := models.Organization{ID: 1, Type: core.DistributorOrganization}
			distributorOrg2 := models.Organization{ID: 2, Type: core.DistributorOrganization}

			repository := repositories.OrganizationInMemoryRepository{
				Organizations: []models.Organization{distributorOrg1, distributorOrg2},
			}

			handler := NewListOrganizationHandler(repository)
			handler.Handle(ctx)

			assert.Equal(t, w.Code, http.StatusOK)

			expectedBody := organizationListSchemaFromOrganizations(
				[]models.Organization{distributorOrg1, distributorOrg2},
			)

			var actualBody organizationListResponse
			err := json.Unmarshal(w.Body.Bytes(), &actualBody)
			assert.NoError(t, err)

			assert.Equal(t, actualBody, expectedBody)
		},
	)

	t.Run(
		"Only Distributor organizations are returned", func(t *testing.T) {
			w := httptest.NewRecorder()

			ctx := testutil.GetTestGinContext(w)

			testutil.MockJsonGet(ctx)

			distributorOrg1 := models.Organization{ID: 1, Type: core.DistributorOrganization, Name: "A"}
			distributorOrg2 := models.Organization{ID: 2, Type: core.DistributorOrganization, Name: "B"}
			clientOrg1 := models.Organization{ID: 3, Type: core.ClientOrganization, Name: "C"}
			clientOrg2 := models.Organization{ID: 4, Type: core.ClientOrganization, Name: "D"}

			repository := repositories.OrganizationInMemoryRepository{
				Organizations: []models.Organization{distributorOrg1, distributorOrg2, clientOrg1, clientOrg2},
			}

			handler := NewListOrganizationHandler(repository)
			handler.Handle(ctx)

			assert.Equal(t, w.Code, http.StatusOK)

			expectedBody := organizationListSchemaFromOrganizations(
				[]models.Organization{distributorOrg1, distributorOrg2},
			)

			var actualBody organizationListResponse
			err := json.Unmarshal(w.Body.Bytes(), &actualBody)
			assert.NoError(t, err)

			assert.Equal(t, actualBody, expectedBody)
		},
	)
}
