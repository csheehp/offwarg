package errors

import "encoding/json"

type errorCode string

const (
	// ErrorCodeBindError is the error code for binding errors.
	ErrorCodeBindError    errorCode = "WARG-400-001"
	ErrorJSONMarshalError errorCode = "WARG-400-002"
)

type wargError struct {
	ErrorCode       errorCode `json:"error_code"`
	Message         string    `json:"message"`
	DetailedMessage string    `json:"detailed_message"`
}

func (e *wargError) Error() string {
	res, err := json.MarshalIndent(e, "", "4")
	if err != nil {
		return newJSONMarhsalError(err.Error()).Error()
	}
	return string(res)
}

func newJSONMarhsalError(message string) *wargError {
	return &wargError{
		ErrorCode:       ErrorJSONMarshalError,
		Message:         "unable to marhsall output",
		DetailedMessage: message,
	}
}

func NewBindError(message string) *wargError {
	return &wargError{
		ErrorCode:       "bind_error",
		Message:         "failed to bind request body",
		DetailedMessage: message,
	}
}
