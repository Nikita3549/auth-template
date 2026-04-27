package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiPort string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file, using system env")
	}

	return &Config{ApiPort: getEnv("API_PORT")}
}

func getEnv(key string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	panic(fmt.Sprintf("Can't resolve %q .env variable", key))
}
