package httputil

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Common HttpErrors
var (
	ErrUnauthorized        = NewError(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	ErrBadRequest          = NewError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	ErrInternalServerError = NewError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
)

// Error grouped status code and error message.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewError creates and returns a new error.
func NewError(code int, message string) Error {
	return Error{
		Code:    code,
		Message: message,
	}
}

// Error returns string representation of an httputil.Error,
// making the type compliant with the error interface
func (err Error) Error() string {
	return fmt.Sprintf("%d: %s", err.Code, err.Message)
}

// ErrorMiddleware wrapper function to deal with encountered errors
// during request handling.
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := getFirstError(c)
		if err == nil {
			return
		}

		var httpError Error
		switch err.(type) {
		case Error:
			httpError = err.(Error)
			break
		default:
			httpError = NewError(http.StatusInternalServerError, err.Error())
			break
		}

		c.AbortWithStatusJSON(httpError.Code, httpError)
	}
}

// getFirstError returns the first error in the gin.Context, nil if not present.
func getFirstError(c *gin.Context) error {
	allErrors := c.Errors
	if len(allErrors) == 0 {
		return nil
	}
	return allErrors[0].Err
}
