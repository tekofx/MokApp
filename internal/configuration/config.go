package configuration

import (
	"fmt"
	"log"
	"os"
	"strings"

	mokuerrors "github.com/Itros97/MokApp/internal/errors"
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

func LoadConfig(env_file string) (*AppConfiguration, *mokuerrors.MokuError) {
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
	merr := checkRequiredConfig(&config)
	if merr != nil {
		return nil, merr
	}
	return &config, nil

}
func setDefaultVariablesIfNeeded(configuration *AppConfiguration) {
	if configuration.Ip == "" {
		configuration.Ip = "127.0.0.1"
	}

	if configuration.Port == "" {
		configuration.Port = "8080"
	}

	if configuration.ApiName == "" {
		configuration.ApiName = "api"
	}

	if configuration.Version == "" {
		configuration.Version = "v1"
	}
}

func checkRequiredConfig(configuration *AppConfiguration) *mokuerrors.MokuError {
	if configuration.SumAppToken == "" {
		return mokuerrors.New(mokuerrors.MissingRequiredConfigParameterErrorCode, fmt.Sprintf(mokuerrors.MissingRequiredConfigParameterMessage, "SUM_APP_TOKEN"))
	}

	return nil

}
