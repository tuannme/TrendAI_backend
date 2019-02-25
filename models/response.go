package models

type Response interface{}

type ResponseWithError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewResponseWithError(code string, message string) ResponseWithError {
	return ResponseWithError{
		Code:    code,
		Message: message,
	}
}
