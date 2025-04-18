package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"management/internal/app/core"
	"management/internal/app/repositories"
)

type ListDistributorsHandler struct {
	repository repositories.OrganizationRepository
}

func NewListOrganizationHandler(repository repositories.OrganizationRepository) ListDistributorsHandler {
	return ListDistributorsHandler{repository: repository}
}

func (h ListDistributorsHandler) Handle(ctx *gin.Context) {
	distributors, err := h.repository.GetMany(
		repositories.OrganizationGetManyParams{
			Type: core.DistributorOrganization,
		},
	)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}

	distributorListSchema := organizationListSchemaFromOrganizations(distributors)

	ctx.JSON(http.StatusOK, distributorListSchema)
}
