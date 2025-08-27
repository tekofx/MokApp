package middleware

import (
	"time"

	apimodels "github.com/Itros97/MokApp/internal/api/models"
	mokuerrors "github.com/Itros97/MokApp/internal/errors"
)

func Trazability(context *apimodels.APIContext) *mokuerrors.APIError {
	time := time.Now().UnixMilli()

	context.Trazability = apimodels.Trazability{
		Endpoint:  context.Trazability.Endpoint,
		Timestamp: &time,
	}

	return nil
}
