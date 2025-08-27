package middleware

import (
	apimodels "github.com/Itros97/MokApp/internal/api/models"
	mokuerrors "github.com/Itros97/MokApp/internal/errors"
)

type Middleware func(context *apimodels.APIContext) *mokuerrors.APIError
