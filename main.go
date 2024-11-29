package main

import (
	"app/cmd/server"
	"app/internal/wire"
	"app/pkg/utils"

	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
)

func main() {
	appCfg := utils.GetAppConfig()

	logger := utils.InitLogger(appCfg.LogPath, appCfg.Debug)
	defer logger.Sync()

	logger.Info("app config", zap.Any("appCfg", appCfg))

	//init db connection
	dbCfg := utils.GetDBConfig()
	conn := utils.ConnectDB(dbCfg)
	db := utils.InitGorm(conn, dbCfg)

	//perform wiring/dependency injection
	router := wire.Wiring(db)

	//start server
	server.ApiServer(logger, appCfg.Port, appCfg.Name, router)
}
