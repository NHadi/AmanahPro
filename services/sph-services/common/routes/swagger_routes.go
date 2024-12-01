package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	httpSwagger "github.com/swaggo/gin-swagger"
)

// RegisterSwaggerRoutes registers Swagger routes for API documentation
func RegisterSwaggerRoutes(router *gin.Engine) {
	router.GET("/swagger/*any", httpSwagger.WrapHandler(swaggerFiles.Handler))
}
