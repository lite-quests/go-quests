package env_vars

import (
	"fmt"
	"os"
	"strconv"
)

// TODO: Implement LoadConfig
// Read README.md for the instructions

type Config struct {
	DBHost    string
	DBPort    int
	DebugMode bool
}

func LoadConfig() (Config, error) {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return Config{}, fmt.Errorf("DB_HOST is not set or empty")
	}
	dbPort := 5432
	debugMode := false
	if val, exists := os.LookupEnv("DB_PORT"); exists && val != "" {
		var err error
		dbPort, err = strconv.Atoi(val)
		if err != nil {
			return Config{}, fmt.Errorf("DB_PORT is not a valid integer")
		}
	}
	if val, exists := os.LookupEnv("DEBUG_MODE"); exists && val != "" {
		var err error
		debugMode, err = strconv.ParseBool(val)
		if err != nil {
			return Config{}, fmt.Errorf("DEBUG_MODE is not a valid boolean")
		}
	}

	return Config{
		DBHost:    dbHost,
		DBPort:    dbPort,
		DebugMode: debugMode,
	}, nil
}
