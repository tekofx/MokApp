package configuration

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type AppConfiguration struct {
	Ip      string
	Port    string
	Version string
	ApiName string

	SumAppToken string

	CorsAccessControlAllowOrigin  string
	CorsAccessControlAllowMethods string
	CorsAccessControlAllowHeaders string
	CorsAccessControlMaxAge       string

	JWTSecret            string
	JWTRegisteredDomains []string
}

func LoadConfig(env_file string) AppConfiguration {
	err := godotenv.Load(env_file)
	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}
	config := AppConfiguration{
		SumAppToken: os.Getenv("SUM_APP_TOKEN"),

		Ip:      os.Getenv("IP"),
		Port:    os.Getenv("PORT"),
		Version: os.Getenv("VERSION"),
		ApiName: os.Getenv("API_NAME"),

		CorsAccessControlAllowOrigin:  os.Getenv("CORS_ORIGIN"),
		CorsAccessControlAllowMethods: os.Getenv("CORS_METHODS"),
		CorsAccessControlAllowHeaders: os.Getenv("CORS_HEADERS"),
		CorsAccessControlMaxAge:       os.Getenv("CORS_MAX_AGE"),

		JWTRegisteredDomains: strings.Split(os.Getenv("JWT_REGISTERED_DOMAINS"), ","),
		JWTSecret:            os.Getenv("JWT_SECRET"),
	}

	setDefaultVariablesIfNeeded(&config)
	return config

}
func setDefaultVariablesIfNeeded(configuration *AppConfiguration) {
	if configuration.Ip == "" {
		configuration.Ip = "127.0.0.1"
	}

	if configuration.Port == "" {
		configuration.Ip = "8080"
	}

	if configuration.ApiName == "" {
		configuration.Ip = "api"
	}

	if configuration.Version == "" {
		configuration.Ip = "v1"
	}
}
