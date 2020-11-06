package config


import (
	"fmt"
	"github.com/subosito/gotenv"
	"os"
)

// Config represents app configuration
type Config struct {
	HTTPPort   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBName     string
	DBPassword string
}

// Init initialize
func New() *Config {
	return &Config{
		HTTPPort:   getEnvWithDefault("PORT", "8080"),
		DBHost:     getEnvWithDefault("DB_HOST", "127.0.0.1"),
		DBPort:     getEnvWithDefault("DB_PORT", "3306"),
		DBUser:     getEnvWithDefault("DB_USER", "root"),
		DBPassword: getEnvWithDefault("DB_PASSWORD", ""),
		DBName:     getEnvWithDefault("DB_NAME", "todo_app"),
	}
}

// DBUrl returns connection string for DB connection
func (c *Config) DBUrl() string {

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)
}

// DBUrlWithoutDBName returns connection string for DB connection omitting the db name
func (c *Config) DBUrlWithoutDBName() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
	)
}

func getEnvWithDefault(key, defaultVal string) string {
	gotenv.Load()
	val := os.Getenv(key)

	if val == "" {
		return defaultVal
	}

	return val
}
