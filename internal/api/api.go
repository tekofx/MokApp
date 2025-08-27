package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Itros97/MokApp/internal/api/controllers"
	"github.com/Itros97/MokApp/internal/api/endpoints"
	"github.com/Itros97/MokApp/internal/api/middleware"
	apimodels "github.com/Itros97/MokApp/internal/api/models"
	"github.com/Itros97/MokApp/internal/configuration"
	"github.com/Itros97/MokApp/internal/database"
	"github.com/Itros97/MokApp/internal/database/tables"
	mokuerrors "github.com/Itros97/MokApp/internal/errors"
	"github.com/Itros97/MokApp/internal/logger"
)

const (
	APIPath           = "/"
	ContentTypeHeader = "Content-Type"
	EnvFilePath       = ".env"
)

var APIMiddleware = []middleware.Middleware{
	middleware.Trazability,
	middleware.Database,
}

func startAPI(configuration *configuration.AppConfiguration, endpoints *[]apimodels.Endpoint) {
	// show log app title and start router
	log.Println("-----------------------------------")
	log.Println(" ", configuration.ApiName, " ")
	log.Println("-----------------------------------")

	// Add API path to endpoints
	newEndpoints := []apimodels.Endpoint{}
	for _, endpoint := range *endpoints {
		endpoint.Path = APIPath + configuration.ApiName + "/" + configuration.Version + "/" + endpoint.Path
		newEndpoints = append(newEndpoints, endpoint)
	}

	// Register endpoints
	registerEndpoints(newEndpoints, configuration)

	// Start listening HTTP requests
	log.Printf("API started on http://%s:%s%s", configuration.Ip, configuration.Port, APIPath)
	state := http.ListenAndServe(configuration.Ip+":"+configuration.Port, nil)
	log.Print(state.Error())
}

func registerEndpoints(endpoints []apimodels.Endpoint, configuration *configuration.AppConfiguration) {
	for _, endpoint := range endpoints {

		switch endpoint.Method {
		case apimodels.GetMethod:
			endpoint.Path = fmt.Sprintf("GET %s", endpoint.Path)
		case apimodels.PostMethod:
			endpoint.Path = fmt.Sprintf("POST %s", endpoint.Path)
		case apimodels.PutMethod:
			endpoint.Path = fmt.Sprintf("PUT %s", endpoint.Path)
		case apimodels.DeleteMethod:
			endpoint.Path = fmt.Sprintf("DELETE %s", endpoint.Path)
		case apimodels.PatchMethod:
			endpoint.Path = fmt.Sprintf("PATCH %s", endpoint.Path)
		}

		log.Printf("Endpoint %s registered. \n", endpoint.Path)

		// set defaults
		fmt.Println(endpoint.Path)
		setEndpointDefaults(&endpoint)

		http.HandleFunc(endpoint.Path, func(writer http.ResponseWriter, reader *http.Request) {

			// enable CORS
			writer.Header().Set("Access-Control-Allow-Origin", configuration.CorsAccessControlAllowOrigin)
			writer.Header().Set("Access-Control-Allow-Methods", configuration.CorsAccessControlAllowMethods)
			writer.Header().Set("Access-Control-Allow-Headers", configuration.CorsAccessControlAllowHeaders)
			writer.Header().Set("Access-Control-Max-Age", configuration.CorsAccessControlMaxAge)

			// calculate the time of the request
			start := time.Now()

			// create basic api context
			context := &apimodels.APIContext{
				Trazability: apimodels.Trazability{
					Endpoint: endpoint,
				},
				Configuration: configuration,
			}

			// Get request data
			err := middleware.Request(reader, context)
			if nil != err {
				logger.Error(
					context.Trazability.Endpoint.Path,
					time.Since(start).Microseconds(),
					"μs -",
					fmt.Sprintf("[%d]", err.Status),
					err.Message,
				)

				middleware.SendResponse(writer, err.Status, err, apimodels.MimeApplicationJSON)
				return
			}

			// Apply middleware to the request
			err = applyMiddleware(context)
			defer database.Close(context.Database)
			if nil != err {
				logger.Error(
					context.Trazability.Endpoint.Path,
					time.Since(start).Microseconds(),
					"μs -",
					fmt.Sprintf("[%d]", err.Status),
					err.Message,
				)
				middleware.SendResponse(writer, err.Status, err, apimodels.MimeApplicationJSON)
				return
			}

			// Execute the endpoint
			middleware.Response(context, writer)
		})
	}
}

func setEndpointDefaults(endpoint *apimodels.Endpoint) {
	if nil == endpoint.Listener {
		endpoint.Listener = controllers.NotImplemented
	}

	if endpoint.RequestMimeType == "" {
		endpoint.RequestMimeType = apimodels.MimeApplicationJSON
	}

	if endpoint.ResponseMimeType == "" {
		endpoint.ResponseMimeType = apimodels.MimeApplicationJSON
	}
}
func applyMiddleware(context *apimodels.APIContext) *mokuerrors.APIError {
	for _, middleware := range APIMiddleware {
		err := middleware(context)
		if nil != err {
			return err
		}
	}

	return nil
}
func Start() {

	config, merr := configuration.LoadConfig(".env")
	if merr != nil {
		fmt.Println(merr.Message)
		os.Exit(0)
	}
	db, err := database.GetConnection()
	if nil != err {
		logger.Error(err)
		return
	}
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = tables.UpdateDatabaseTablesToLatestVersion(pwd, db)
	if nil != err {
		logger.Error(err)
		return
	}
	database.Close(db)

	startAPI(config, &endpoints.EndpointRegistry)
}
