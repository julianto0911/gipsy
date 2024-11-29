package utils

import (
	"database/sql"
	"log"
	"time"

	"github.com/julianto0911/tools"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitLogger(path string, debug bool) *zap.Logger {
	l, err := tools.NewLogger(path, debug)
	if err != nil {
		log.Fatal("can't init logger %w", zap.Error(err))
	}
	return l
}

func ConnectDB(cfg tools.DBConfiguration) *sql.DB {
	cfg.DbType = tools.Postgresql
	sqlConn, err := tools.ConnectDB(cfg)
	if err != nil {
		log.Fatal("can't connect to database ", err)
	}
	return sqlConn
}

func InitGorm(sqlConn *sql.DB, cfg tools.DBConfiguration) *gorm.DB {
	dbLogger := tools.CreateLogger(cfg.Logging)
	db, err := tools.NewGormDB(cfg.DbType, sqlConn, dbLogger, cfg.Logging)
	if err != nil {
		log.Fatal("can't init gorm database ", err)
	}
	return db
}

func GetAppConfig() AppConfig {
	cfg := AppConfig{
		LogPath:        tools.EnvString("LOG_PATH"),
		Debug:          tools.EnvBool("DEBUG"),
		Timezone:       tools.EnvString("TIMEZONE"),
		Port:           tools.EnvString("PORT"),
		Name:           tools.EnvString("APP_NAME"),
		NetTimeOut:     time.Duration(tools.EnvInt("NET_TIMEOUT")) * time.Second,
		PublicKey:      tools.EnvString("PUBLIC_KEY"),
		PrivateKey:     tools.EnvString("PRIVATE_KEY"),
		Environment:    tools.EnvString("ENVIRONMENT"),
		SAdminUsername: tools.EnvString("SADMIN_USERNAME"),
		SAdminPassword: tools.EnvString("SADMIN_PASSWORD"),
	}

	//default port
	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	//default logging path
	if cfg.LogPath == "" {
		cfg.LogPath = "/logs/"
	}

	//load location
	var err error
	cfg.Location, err = time.LoadLocation(cfg.Timezone)
	if err != nil {
		log.Fatal("can't load location :" + cfg.Timezone + ",error :" + err.Error())
	}

	return cfg
}

func GetDBConfig() tools.DBConfiguration {
	cfg := tools.DBConfiguration{
		Host:           tools.EnvString("DB_HOST"),
		DBName:         tools.EnvString("DB_NAME"),
		Username:       tools.EnvString("DB_USERNAME"),
		Password:       tools.EnvString("DB_PASSWORD"),
		Logging:        tools.EnvBool("DB_DEBUG"),
		Port:           tools.EnvString("DB_PORT"),
		Schema:         tools.EnvString("DB_SCHEMA"),
		SessionName:    tools.EnvString("DB_SESSION_NAME"),
		ConnectTimeOut: tools.EnvInt("DB_CONNECT_TIMEOUT"),
		MaxOpenConn:    tools.EnvInt("DB_MAX_OPEN_CONN"),
		MaxIdleConn:    tools.EnvInt("DB_MAX_IDLE_CONN"),
	}

	//default db connection time out
	if cfg.ConnectTimeOut == 0 {
		cfg.ConnectTimeOut = 30
	}
	//default db maximum open connection
	if cfg.MaxOpenConn == 0 {
		cfg.MaxOpenConn = 50
	}
	//default db maximum idle connection
	if cfg.MaxIdleConn == 0 {
		cfg.MaxIdleConn = 10
	}

	return cfg
}

func GetRedisConfig() tools.RedisConfiguration {
	cfg := tools.RedisConfiguration{
		Host:     tools.EnvString("REDIS_HOST"),
		Port:     tools.EnvString("REDIS_PORT"),
		Password: tools.EnvString("REDIS_PASSWORD"),
		Prefix:   tools.EnvString("REDIS_PREFIX"),
		UseMock:  tools.EnvBool("REDIS_USE_MOCK"),
	}

	return cfg
}
