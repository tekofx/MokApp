package api

import "github.com/Itros97/MokApp/internal/configuration"

func Start() {
	config := configuration.LoadConfig(".env")

}
