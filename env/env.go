package env

import (
	"log"
	"os"
)

var env string

const (
	DEVELOPMENT = "development"
	STAGING     = "staging"
	BETA        = "beta"
	PRODUCTION  = "production"
)

var envMap = map[string]bool{
	DEVELOPMENT: true,
	STAGING:     true,
	BETA:        true,
	PRODUCTION:  true,
}

func Init() {
	env = os.Getenv("env")
	if env == "" {
		env = DEVELOPMENT
	}
	if _, ok := envMap[env]; !ok {
		log.Fatalf("Invalid Environment: %s", env)
	}
}

func IsDevelopment() bool {
	return env == DEVELOPMENT
}

func IsStaging() bool {
	return env == STAGING
}

func IsBeta() bool {
	return env == BETA
}

func IsProduction() bool {
	return env == PRODUCTION
}
