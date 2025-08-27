package configuration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfiguration struct {
	SumAppToken string
}

func LoadConfig(env_file string) AppConfiguration {
	err := godotenv.Load(env_file)
	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}
	config := AppConfiguration{
		SumAppToken: os.Getenv("SUM_APP_TOKEN"),
	}

	return config

}
