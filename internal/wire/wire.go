package wire

import (
	"app/pkg/middleware"
	"app/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Wiring(appCfg *utils.AppConfig) *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	router.Use(
		gin.Recovery(),
		middleware.RequestID(),
	)

	api := router.Group("/api/v1")
	wireSystem(api, appCfg)

	return router
}
