package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment   string // develop, staging, production
	UserServiceHost string
	UserServicePort string
	ScheduleServiceHost string
	ScheduleServicePort string
	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func Load() Config {

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", ""))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", ""))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ""))

	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", ""))
	c.UserServicePort = cast.ToString(getOrReturnDefault("USER_GRPC_PORT", ""))
	c.ScheduleServiceHost = cast.ToString(getOrReturnDefault("SCHEDULE_SERVICE_HOST", ""))
	c.ScheduleServicePort = cast.ToString(getOrReturnDefault("SCHEDULE_GRPC_PORT", ""))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	if os.Getenv(key) == "" {
		return defaultValue
	}
	return os.Getenv(key)
}
