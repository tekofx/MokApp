package endpoints

import apimodels "github.com/Itros97/MokApp/internal/api/models"

var EndpointRegistry = []apimodels.Endpoint{
	ItemGetByIdEndpoint,
	ItemPostEndpoint,
	ItemGetAllpoint,
}
