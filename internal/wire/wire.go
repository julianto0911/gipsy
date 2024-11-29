package wire

import (
	"app/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Wiring(db *gorm.DB) *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	router.Use(
		gin.Recovery(),
		middleware.RequestID(),
	)

	api := router.Group("/api/v1")
	wireProduct(api, db)

	return router
}
