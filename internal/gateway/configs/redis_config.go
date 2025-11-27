package configs

import (
	"log"
	"os"
	"strconv"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DBNumber int
}

func InitRedisConfig() *RedisConfig {
	dbNumber, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatalf("Error loading DBNumber")
	}

	cfg := RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DBNumber: dbNumber,
	}

	return &cfg
}
