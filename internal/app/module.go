package app

import (
	"management/internal/app/repositories"
	"management/internal/app/rest"
	"management/internal/app/services"
	"management/internal/infra"
)

type Application struct {
}

func (m Application) Init(app *infra.Infra) error {
	orgRepository := repositories.NewOrganizationRepository(app.DB)

	orgService := services.NewOrganizationService(app.DB)

	rest.RegisterOrganizationRoutes(app, orgService, orgRepository)

	return nil
}
