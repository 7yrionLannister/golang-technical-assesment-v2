package env

import (
	"errors"
	"fmt"
	"io/fs"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// [Environment] is a struct that defines the environment variables that the application needs.
type Environment struct {
	ServerPort       string `env:"PORT,required"`
	DataBaseHost     string `env:"DB_HOST,required"`
	DataBaseEngine   string `env:"DB_ENGINE,required"`
	DataBaseUser     string `env:"DB_USER,required"`
	DataBasePassword string `env:"DB_PASSWORD,required"`
	DataBaseName     string `env:"DB_NAME,required"`
	DataBaseUrl      string
	LogLevel         string `env:"LOG_LEVEL" envDefault:"info"`
}

// Global variable that holds the environment variables that the application needs.
// Use only after calling [LoadEnv].
var Env Environment = Environment{}

// Loads the environment variables into [Env].
func LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	err = env.Parse(&Env)
	if err != nil {
		return fmt.Errorf("error parsing environment variables: %w", err)
	}
	baseURL := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable", Env.DataBaseEngine, Env.DataBaseUser, Env.DataBasePassword, Env.DataBaseHost, Env.DataBaseName)
	Env.DataBaseUrl = baseURL
	return nil
}
