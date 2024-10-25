package app

import (
	"os"
	"pagination/common/postgresql"
)

type ConfigurationManager struct {
	PostgreSqlConfig postgresql.Config
}

func NewConfigurationManager() *ConfigurationManager {
	postgreSqlConfig := getPostgreSqlConfig()
	return &ConfigurationManager{
		PostgreSqlConfig: postgreSqlConfig,
	}
}

func getPostgreSqlConfig() postgresql.Config {
	return postgresql.Config{
		Host:                  getEnv("DB_HOST", "localhost"),
		Port:                  getEnv("DB_PORT", "5432"),
		UserName:              getEnv("DB_USER", "postgres"),
		Password:              getEnv("DB_PASSWORD", "153515"),
		DbName:                getEnv("DB_NAME", "workshops"),
		MaxConnections:        getEnv("DB_MAX_CONNECTIONS", "10"),
		MaxConnectionIdleTime: getEnv("DB_MAX_IDLE_TIME", "10s"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
