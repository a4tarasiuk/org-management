package rest

import (
	"management/internal/app/repositories"
	"management/internal/app/rest/handlers"
	"management/internal/app/services"
	"management/internal/infra"
)

func RegisterOrganizationRoutes(
	app *infra.Infra,
	orgService services.Service,
	orgRepository repositories.OrganizationRepository,
) {
	createDistributorHandler := handlers.NewCreateDistributorHandler(orgService)
	listDistributorsHandler := handlers.NewListOrganizationHandler(orgRepository)

	distributors := app.Router.Group("/organizations/distributors")

	distributors.POST("", createDistributorHandler.Handle)
	distributors.GET("", listDistributorsHandler.Handle)
}
