package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"management/internal/app/core"
	"management/internal/app/orm"
	"management/internal/testutil"
)

type organizationInMemoryService struct {
	m mock.Mock
}

func (s organizationInMemoryService) Create(organization *orm.Organization) error {
	args := s.m.Called(organization)

	organization.ID++

	return args.Error(0)
}

func Test_CreateDistributorOrganizationHandler(t *testing.T) {
	t.Run(
		"When created successfully 201", func(t *testing.T) {
			w := httptest.NewRecorder()

			ctx := testutil.GetTestGinContext(w)

			createOrgSchema := createOrganizationRequest{Name: "Test Organization"}

			testutil.MockJsonPost(ctx, createOrgSchema)

			orgToBeCreated := orm.Organization{Name: createOrgSchema.Name, Type: core.DistributorOrganization}

			serviceMock := mock.Mock{}
			serviceMock.On("Create", &orgToBeCreated).Return(nil)

			handler := CreateDistributorHandler{service: organizationInMemoryService{m: serviceMock}}
			handler.Handle(ctx)

			serviceMock.AssertExpectations(t)

			assert.Equal(t, w.Code, http.StatusCreated)

			expectedResponseSchema := createOrganizationResponse{
				ID:   1,
				Name: createOrgSchema.Name,
				Type: core.DistributorOrganization,
			}

			expectedBody, err := json.Marshal(expectedResponseSchema)
			assert.NoError(t, err)

			assert.Equal(t, w.Body.Bytes(), expectedBody)
		},
	)

	t.Run(
		"When failed 500", func(t *testing.T) {
			w := httptest.NewRecorder()

			ctx := testutil.GetTestGinContext(w)

			createOrgSchema := createOrganizationRequest{Name: "Test Organization"}

			testutil.MockJsonPost(ctx, createOrgSchema)

			err := errors.New("test error")

			serviceMock := mock.Mock{}
			serviceMock.On("Create", mock.Anything).Return(err)

			handler := CreateDistributorHandler{service: organizationInMemoryService{m: serviceMock}}
			handler.Handle(ctx)

			assert.Equal(t, w.Code, http.StatusInternalServerError)

			expectedBody, err := json.Marshal(ErrorResponse{Error: err.Error()})
			assert.NoError(t, err)

			assert.Equal(t, w.Body.Bytes(), expectedBody)
		},
	)

	t.Run(
		"When invalid request 400", func(t *testing.T) {
			w := httptest.NewRecorder()

			ctx := testutil.GetTestGinContext(w)

			createOrgSchema := struct {
				field string
			}{field: "value"}

			testutil.MockJsonPost(ctx, createOrgSchema)

			handler := CreateDistributorHandler{service: organizationInMemoryService{m: mock.Mock{}}}
			handler.Handle(ctx)

			assert.Equal(t, w.Code, http.StatusBadRequest)

			assert.Contains(t, w.Body.String(), "Field validation")
			assert.Contains(t, w.Body.String(), "'Name'")

			responseBody := ErrorResponse{}
			err := json.Unmarshal(w.Body.Bytes(), &responseBody)
			assert.NoError(t, err)
		},
	)
}
