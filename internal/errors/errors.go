package mokuerrors

import "net/http"

type MokuError struct {
	Code    MokuErrorCode `json:"code,omitempty"`
	Message string        `json:"message,omitempty"`
}

type APIError struct {
	Status int `json:"status,omitempty"`
	MokuError
}

func Unexpected(message string) *MokuError {
	return &MokuError{
		Code:    UnexpectedErrorCode,
		Message: message,
	}
}
func TODO() MokuError {
	return MokuError{
		Code:    NotImplementedErrorCode,
		Message: "Not yet implemented",
	}
}
func Unauthorized(message string) *MokuError {
	return &MokuError{
		Code:    UnauthorizedErrorCode,
		Message: message,
	}
}

func DatabaseError(message string) *MokuError {
	return &MokuError{
		Code:    InvalidRequestErrorCode,
		Message: message,
	}
}
func New(code MokuErrorCode, message string) *MokuError {
	return &MokuError{
		Code:    code,
		Message: message,
	}
}

func InvalidRequest(message string) *MokuError {
	return &MokuError{
		Code:    InvalidRequestErrorCode,
		Message: message,
	}
}

func NotFound(message string) *MokuError {
	return &MokuError{
		Code:    NotFoundErrorCode,
		Message: message,
	}
}

func NewAPIError(mokuError *MokuError) *APIError {
	var status int

	if 0 <= mokuError.Code && mokuError.Code <= 999 {
		status = http.StatusInternalServerError
	} else if 1000 <= mokuError.Code && mokuError.Code <= 3999 {
		status = http.StatusBadRequest
	} else if 4000 <= mokuError.Code && mokuError.Code <= 4999 {
		status = http.StatusNotFound
	} else if 5000 <= mokuError.Code && mokuError.Code <= 5999 {
		status = http.StatusUnauthorized
	} else if 6000 <= mokuError.Code && mokuError.Code <= 6999 {
		status = http.StatusForbidden
	} else {
		status = http.StatusTeapot
	}

	return &APIError{
		Status:    status,
		MokuError: *mokuError,
	}
}

type MokuErrorCode int

const (
	// 0 --> 999 | SYSTEM UNEXPECTED ERRORS
	UnexpectedErrorCode                     MokuErrorCode = 0
	DatabaseErrorCode                       MokuErrorCode = 1
	NotImplementedErrorCode                 MokuErrorCode = 2
	NothingChangedErrorCode                 MokuErrorCode = 3
	CannotGenerateAuthTokenErrorCode        MokuErrorCode = 4
	CannotCreateValidationCodeErrorCode     MokuErrorCode = 5
	MissingRequiredConfigParameterErrorCode MokuErrorCode = 6

	// 1000 -> 3999 | VALIDATION ERRORS
	InvalidRequestErrorCode MokuErrorCode = 1000

	// 1100 -> 1299 | ITEM RELATED VALIDATION ERRORS
	ItemAlreadyExistsErrorCode    MokuErrorCode = 1100
	ItemAlreadyValidatedErrorCode MokuErrorCode = 1101

	// 4000 -> 4999 | LOOKUP ERRORS
	NotFoundErrorCode MokuErrorCode = 4000

	// 5000 -> 5999 | AUTHORITATION ERRORS
	UnauthorizedErrorCode MokuErrorCode = 5000
)

const (
	AccessDeniedMessage string = "access denied"

	CannotConnectToDatabaseMessage        string = "cannot connect to database"
	DatabaseConnectionEmptyMessage        string = "database connection cannot be empty"
	ServiceIDEmptyMessage                 string = "service id cannot be empty"
	RegisteredDomainsEmptyMessage         string = "registered domains cannot be empty"
	SecretEmptyMessage                    string = "secret cannot be empty"
	MissingRequiredConfigParameterMessage string = "missing config parameter %s"

	TokenEmptyMessage   string = "token cannot be empty"
	TokenInvalidMessage string = "invalid token"

	FileTooLargeMessage string = "%s is too long; the maximum size is %dMB"

	// DAL operations
	ItemInvalidMessage       string = "item cannot be empty"
	ItemIdNegativeMessage    string = "item id must be greater than 0"
	ItemNotFoundMessage      string = "item not found"
	ItemAlreadyExistsMessage string = "item already exists"
)
