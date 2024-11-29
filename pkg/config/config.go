package config

import (
	"fmt"
	"time"

	"github.com/julianto0911/tools"
)

type Configuration struct {
	APP   AppConfig
	REDIS tools.RedisConfiguration
	DB    tools.DBConfiguration
}

type AppConfig struct {
	Debug          bool
	Timezone       string
	Port           string
	Location       *time.Location `anonymous:"true"`
	LogPath        string
	Name           string
	NetTimeOut     time.Duration
	Environment    string
	PublicKey      string
	PrivateKey     string
	SAdminUsername string
	SAdminPassword string
}

func ReadConfiguration() (Configuration, error) {
	cfg := Configuration{
		APP: AppConfig{
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
		},
		DB: tools.DBConfiguration{
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
		},
		REDIS: tools.RedisConfiguration{
			Host:     tools.EnvString("REDIS_HOST"),
			Port:     tools.EnvString("REDIS_PORT"),
			Password: tools.EnvString("REDIS_PASSWORD"),
			Prefix:   tools.EnvString("REDIS_PREFIX"),
			UseMock:  tools.EnvBool("REDIS_USE_MOCK"),
		},
	}

	//default port
	if cfg.APP.Port == "" {
		cfg.APP.Port = "8080"
	}
	//default logging path
	if cfg.APP.LogPath == "" {
		cfg.APP.LogPath = "/logs/"
	}
	//default db connection time out
	if cfg.DB.ConnectTimeOut == 0 {
		cfg.DB.ConnectTimeOut = 30
	}
	//default db maximum open connection
	if cfg.DB.MaxOpenConn == 0 {
		cfg.DB.MaxOpenConn = 50
	}
	//default db maximum idle connection
	if cfg.DB.MaxIdleConn == 0 {
		cfg.DB.MaxIdleConn = 10
	}
	//load location
	var err error
	cfg.APP.Location, err = time.LoadLocation(cfg.APP.Timezone)
	if err != nil {
		return Configuration{}, fmt.Errorf("can't load location :" + cfg.APP.Timezone + ",error :" + err.Error())
	}

	return cfg, nil
}
