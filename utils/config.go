package utils

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config hold application configuration
type Config struct {
	Port     int    `envconfig:"http_port"`
	Token    string `envconfig:"github_token"`
	Username string `envconfig:"github_username"`
}

// LoadConfig provide routine for loading configuration
func LoadConfig() (Config, error) {
	var c Config
	err := godotenv.Load()
	if err != nil {
		return c, err
	}

	err = envconfig.Process("kleng", &c)
	return c, err
}
