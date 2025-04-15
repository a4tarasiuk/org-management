package infra

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Application struct {
	DB *gorm.DB

	Router *gin.Engine
}

func (app *Application) Run() {
	app.Router.Run()
}
