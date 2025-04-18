package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"management/internal/app/core"
	"management/internal/app/models"
	"management/internal/app/services"
)

type CreateDistributorHandler struct {
	service services.Service
}

func NewCreateDistributorHandler(service services.Service) CreateDistributorHandler {
	return CreateDistributorHandler{service: service}
}

// Handle Create Distributor 	godoc
// @Summary		Create Distributor organization
// @Description	Creates a new Distributor organization.
// @Tags		distributors
// @Produce		json
// @Param       organization  body      createOrganizationRequest  true  "Organization JSON"
// @Success 	201 {object} createOrganizationResponse
// @Failure 	500 {object} errorResponse "Internal Server Error"
// @Router		/organizations/distributors [post]
func (h CreateDistributorHandler) Handle(ctx *gin.Context) {
	requestSchema := createOrganizationRequest{}

	if err := ctx.ShouldBindJSON(&requestSchema); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}

	organization := models.Organization{
		Name: requestSchema.Name,
		Type: core.DistributorOrganization,
	}

	if err := h.service.Create(&organization); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	responseSchema := organizationResponseFromOrganization(organization)

	ctx.JSON(http.StatusCreated, responseSchema)
}
