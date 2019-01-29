package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Config represent application configuration vars
type Config struct {
	Host         string
	Port         string
	AccessToken  string
	Account      string
	Repositories []string
	PausedLabels []string
}

// InitConfig create new config struct
// From environment variables
func InitConfig() *Config {
	// Init .env file
	_ = godotenv.Load()

	return &Config{
		Host:         os.Getenv("APP_HOST"),
		Port:         os.Getenv("APP_PORT"),
		AccessToken:  os.Getenv("GH_ACCESS_TOKEN"),
		Account:      os.Getenv("GH_ACCOUNT"),
		Repositories: divide(os.Getenv("GH_REPOSITORIES")),
		PausedLabels: divide(os.Getenv("PAUSED_LABELS")),
	}
}

// divide divides given string
func divide(str string) []string {
	return strings.Split(str, "|")
}
