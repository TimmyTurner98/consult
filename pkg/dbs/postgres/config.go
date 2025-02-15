package postgres

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	SSLMode  string
}

func initConfigs() (*Config, error) {
	viper.SetConfigName("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
