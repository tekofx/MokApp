package endpoints

import (
	"io"
	"net/http"

	apimodels "github.com/Itros97/MokApp/internal/api/models"
	mokuerrors "github.com/Itros97/MokApp/internal/errors"
	"github.com/Itros97/MokApp/internal/models"
	"github.com/Itros97/MokApp/internal/services"
	"github.com/Itros97/MokApp/internal/utils"
)

var ItemGetEndpoint = apimodels.Endpoint{
	Path:     "items/{id}",
	Method:   apimodels.GetMethod,
	Listener: Get,
	Secured:  false,
	Database: true,
}

func Get(context *apimodels.APIContext) (*apimodels.Response, *mokuerrors.APIError) {
	id, err := context.Request.GetParamInt64("id")
	if err != nil {
		return nil, mokuerrors.NewAPIError(mokuerrors.InvalidRequest(err.Error()))
	}

	item, mokuerr := services.GetItemById(context.Database, *id)
	if mokuerr != nil {
		return nil, mokuerrors.NewAPIError(mokuerr)
	}

	return &apimodels.Response{
		Code:     200,
		Response: item,
	}, nil

}

var ItemPostEndpoint = apimodels.Endpoint{
	Path:     "items",
	Method:   apimodels.PostMethod,
	Listener: Post,
	Secured:  false,
	Database: true,
}

func Post(context *apimodels.APIContext) (*apimodels.Response, *mokuerrors.APIError) {
	body := context.Request.Body.(io.ReadCloser)
	item := models.Item{}
	err := utils.ParseJSON(&body, &item)
	if nil != err {
		return nil, mokuerrors.NewAPIError(mokuerrors.New(mokuerrors.InvalidRequestErrorCode, "Request body has invalid format."))
	}

	context.Request.Body = item

	itemId, rerr := services.InsertItem(context.Database, context.Request.Body.(models.Item))

	if nil != rerr {
		return nil, mokuerrors.NewAPIError(rerr)
	}

	return &apimodels.Response{
		Code:     http.StatusOK,
		Response: itemId,
	}, nil

}
