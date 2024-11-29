package wire

import (
	"app/internal/adaptor"
	"app/internal/data/repo"
	ucsystem "app/internal/usecase/system"
	"app/pkg/utils"

	"github.com/gin-gonic/gin"
)

func wireSystem(router *gin.RouterGroup, cfg *utils.AppConfig) {
	sadminRepo := repo.NewSadminRepo(cfg.SAdminUsername, cfg.SAdminPassword)
	systemUC := ucsystem.NewSystemUseCase(sadminRepo)
	systemAdaptor := adaptor.NewSystemAdaptor(systemUC)

	router.POST("/login", systemAdaptor.VerifySAdminLogin)
}
