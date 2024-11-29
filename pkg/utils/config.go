package utils

import "time"

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
