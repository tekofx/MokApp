package configuration

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type AppConfiguration struct {
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

		CorsAccessControlAllowOrigin:  os.Getenv("CORS_ORIGIN"),
		CorsAccessControlAllowMethods: os.Getenv("CORS_METHODS"),
		CorsAccessControlAllowHeaders: os.Getenv("CORS_HEADERS"),
		CorsAccessControlMaxAge:       os.Getenv("CORS_MAX_AGE"),

		JWTRegisteredDomains: strings.Split(os.Getenv("JWT_REGISTERED_DOMAINS"), ","),
		JWTSecret:            os.Getenv("JWT_SECRET"),
	}

	return config

}
