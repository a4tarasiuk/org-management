package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "management/docs"
)

func Init(router *gin.Engine) {
	docs.SwaggerInfo.Title = "Management API"
	docs.SwaggerInfo.Version = "0.0.1"
	docs.SwaggerInfo.Description = "REST API for managing organizations"

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
