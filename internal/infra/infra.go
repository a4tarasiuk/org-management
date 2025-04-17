package infra

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"management/internal/infra/conf"
	"management/internal/infra/postgres"
)

type Infra struct {
	DB *gorm.DB

	Router *gin.Engine
}

func Init(config conf.Config) (*Infra, error) {
	db, err := postgres.InitDB(config)
	if err != nil {
		return nil, err
	}

	router := gin.Default()

	app := &Infra{
		DB:     db,
		Router: router,
	}

	return app, nil
}

func (app *Infra) Run() error {
	err := app.Router.Run()
	return err
}
