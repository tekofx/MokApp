// Package controllers provides functions to handle http requests in the api
package controllers

import (
	"net/http"

	apimodels "github.com/Itros97/MokApp/internal/api/models"
	mokuerrors "github.com/Itros97/MokApp/internal/errors"
)

func Health(context *apimodels.APIContext) (*apimodels.Response, *mokuerrors.APIError) {
	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: "OK",
	}, nil
}

func EmptyCheck(context *apimodels.APIContext) *mokuerrors.APIError {
	return nil
}

func NotImplemented(context *apimodels.APIContext) (*apimodels.Response, *mokuerrors.APIError) {
	return nil, &mokuerrors.APIError{
		Status:    http.StatusNotImplemented,
		MokuError: mokuerrors.TODO(),
	}
}
