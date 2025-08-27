package apimodels

import mokuerrors "github.com/Itros97/MokApp/internal/errors"

type Endpoint struct {
	Path   string     `json:"path,omitempty"`
	Method HTTPMethod `json:"method,omitempty"`

	RequestMimeType  MimeType `json:"requestMimeType,omitempty"`
	ResponseMimeType MimeType `json:"responseMimeType,omitempty"`

	Listener EndpointListener `json:"-"`

	IsMultipartForm bool `json:"containsFiles,omitempty"`
	Secured         bool `json:"secured,omitempty"`
	Database        bool `json:"-"`
}

type EndpointListener func(context *APIContext) (*Response, *mokuerrors.APIError)
