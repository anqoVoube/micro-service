package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

type Address interface {
	Addr()
}

type postgreConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Poolsize int
}

//func (config *postgreConfig) Addr() string {
//	return fmt.Sprintf("%s:%s", config.Host, config.Port)
//}

type Config struct {
	Postgres postgreConfig
}

func (c Config) Addr() string {
	return fmt.Sprintf("%s:%d", config.Postgres.Host, config.Postgres.Port)
}

var (
	once   sync.Once
	config Config
)

// New returns a new Config struct
func New() *Config {
	once.Do(func() {
		config = Config{
			Postgres: postgreConfig{
				Host:     getEnv("PG_HOST", "postgres"),
				Port:     getIntEnv("PG_PORT", 5432),
				User:     getEnv("PG_USER", "postgres"),
				Password: getEnv("PG_PASSWORD", "password"),
				Database: getEnv("PG_DATABASE", "users"),
				Poolsize: getIntEnv("PG_POOLSIZE", 50),
			},
		}
	})
	return &config
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getIntEnv(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("Error occurred while trying to convert to int %s", err.Error())
		} else {
			return intValue
		}
	}

	return defaultVal
}
