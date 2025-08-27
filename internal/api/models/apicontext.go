package apimodels

import (
	"database/sql"

	"github.com/Itros97/MokApp/internal/configuration"
)

type APIContext struct {
	Configuration *configuration.AppConfiguration
	Trazability   Trazability
	Request       Request
	Response      Response
	Database      *sql.DB
}
