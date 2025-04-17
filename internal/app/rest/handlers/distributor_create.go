package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"management/internal/app/core"
	"management/internal/app/orm"
	"management/internal/app/services"
)

type CreateDistributorHandler struct {
	service services.Service
}

func NewCreateDistributorHandler(service services.Service) CreateDistributorHandler {
	return CreateDistributorHandler{service: service}
}

func (h CreateDistributorHandler) Handle(ctx *gin.Context) {
	requestSchema := createOrganizationRequest{}

	if err := ctx.ShouldBindJSON(&requestSchema); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	organization := orm.Organization{
		Name: requestSchema.Name,
		Type: core.DistributorOrganization,
	}

	if err := h.service.Create(&organization); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	responseSchema := organizationResponseFromOrganization(organization)

	ctx.JSON(http.StatusCreated, responseSchema)
}
