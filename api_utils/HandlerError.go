// Author: Paweł Konopko
// License: MIT

package api_utils

import "strconv"

type HandlerError struct {
	Code int
	Message string
}
func (e *HandlerError) Error() string {
	return "[" + strconv.Itoa(e.Code) + "] " + e.Message
}

func NewHandlerError(code int, message string) *HandlerError {
	return &HandlerError {
		Code: code,
		Message: message,
	}
}
