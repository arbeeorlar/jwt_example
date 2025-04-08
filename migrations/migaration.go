package migrations

import (
	"github.com/arbeeorlar/jwt-example/initializers"
	"github.com/arbeeorlar/jwt-example/models"
	"os"
)

func init() {
	initializers.LoadEnvVariable()
	config := initializers.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_INSTANCE"),
		SSLMode:  "disable",
	}
	initializers.DatabaseConfig(config)
}

func AutoMigrate() {
	initializers.DB.AutoMigrate(&models.JWTUser{})

}
