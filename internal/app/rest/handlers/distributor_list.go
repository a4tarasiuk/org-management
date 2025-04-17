package handlers

import (
	"github.com/gin-gonic/gin"
	"management/internal/app/repositories"
)

type ListDistributorsHandler struct {
	repository repositories.OrganizationRepository
}

func NewListOrganizationHandler(repository repositories.OrganizationRepository) ListDistributorsHandler {
	return ListDistributorsHandler{repository: repository}
}

func (h ListDistributorsHandler) Handle(ctx *gin.Context) {

}
