package middleware

import (
	apimodels "github.com/Itros97/MokApp/internal/api/models"
	"github.com/Itros97/MokApp/internal/database"
	mokuerrors "github.com/Itros97/MokApp/internal/errors"
)

func Database(context *apimodels.APIContext) *mokuerrors.APIError {
	if !context.Trazability.Endpoint.Database || nil != context.Database {
		return nil
	}

	db, err := database.GetConnection()
	if nil != err {
		return mokuerrors.NewAPIError(mokuerrors.DatabaseError(mokuerrors.CannotConnectToDatabaseMessage))
	}

	context.Database = db
	if nil == context.Database {
		return mokuerrors.NewAPIError(mokuerrors.DatabaseError(mokuerrors.CannotConnectToDatabaseMessage))
	}

	return nil
}
