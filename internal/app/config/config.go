package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort   string
	DBUser    string
	DBPass    string
	DBName    string
	DBHost    string
	JWTKey    string
	RedisHost string
	RedisPass string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	// Try to load from Vault first
	vaultData := LoadVaultSecrets()

	get := func(key, fallback string) string {
		if vaultData != nil {
			if val, ok := vaultData[key]; ok {
				return val
			}
		}
		if val := os.Getenv(key); val != "" {
			return val
		}
		return fallback
	}

	cfg := &Config{
		AppPort:   get("APP_PORT", "8080"),
		DBUser:    get("DB_USER", "root"),
		DBPass:    get("DB_PASS", ""),
		DBName:    get("DB_NAME", "bumimedika"),
		DBHost:    get("DB_HOST", "127.0.0.1:3306"),
		JWTKey:    get("JWT_KEY", "supersecret"),
		RedisHost: get("REDIS_HOST", "localhost:6379"),
	}

	log.Printf("[CONFIG] Loaded with Vault=%v", vaultData != nil)
	return cfg
}
