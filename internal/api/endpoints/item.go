package endpoints

import (
	apimodels "github.com/Itros97/MokApp/internal/api/models"
	mokuerrors "github.com/Itros97/MokApp/internal/errors"
	"github.com/Itros97/MokApp/internal/models"
)

var ItemGetEndpoint = apimodels.Endpoint{
	Path:     "items",
	Method:   apimodels.GetMethod,
	Listener: Get,
	Secured:  false,
	Database: true,
}

func Get(context *apimodels.APIContext) (*apimodels.Response, *mokuerrors.APIError) {

	return &apimodels.Response{
		Code: 200,
		Response: models.Item{
			ID:          1,
			Name:        "Sticker Miguel",
			Description: "Miguel woof woof bark",
			Stock:       1,
		},
	}, nil

}
