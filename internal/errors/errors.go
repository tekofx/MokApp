package mokuerrors

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

func New(code MokuErrorCode, message string) *MokuError {
	return &MokuError{
		Code:    code,
		Message: message,
	}
}

type MokuErrorCode int

const (
	UnexpectedErrorCode                 MokuErrorCode = 0
	DatabaseErrorCode                   MokuErrorCode = 1
	NotImplementedErrorCode             MokuErrorCode = 2
	NothingChangedErrorCode             MokuErrorCode = 3
	CannotGenerateAuthTokenErrorCode    MokuErrorCode = 4
	CannotCreateValidationCodeErrorCode MokuErrorCode = 5
)

const (
	AccessDeniedMessage string = "access denied"

	CannotConnectToDatabaseMessage string = "cannot connect to database"
	DatabaseConnectionEmptyMessage string = "database connection cannot be empty"
	ServiceIDEmptyMessage          string = "service id cannot be empty"
	RegisteredDomainsEmptyMessage  string = "registered domains cannot be empty"
	SecretEmptyMessage             string = "secret cannot be empty"

	TokenEmptyMessage   string = "token cannot be empty"
	TokenInvalidMessage string = "invalid token"

	FileTooLargeMessage string = "%s is too long; the maximum size is %dMB"
