package httputil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	userHeader = "X-UserID"
)

var ErrNoAuthHeader = NewError(http.StatusBadRequest, "No Authorization header provided")
var ErrUserIDMissing = NewError(http.StatusBadRequest, "UserId missing")

// RequireAuthToken middleware for putting authentication by bearer token
// in front of incomming requests
func RequireAuthToken(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, password, ok := c.Request.BasicAuth()
		if !ok {
			c.AbortWithStatusJSON(ErrNoAuthHeader.Code, ErrNoAuthHeader)
			return
		}

		if password != token {
			c.AbortWithStatusJSON(ErrUnauthorized.Code, ErrUnauthorized)
			return
		}

		SetContextUserID(c, userID)
		c.Next()
	}
}

// SetContextUserID sets the userID of the client that initated the request.
func SetContextUserID(c *gin.Context, userID string) {
	if userID != "" {
		c.Set(userHeader, userID)
	}
}

// GetUserID gets the user id set by the auth middleware.
func GetUserID(c *gin.Context) (string, error) {
	userID := c.GetString(userHeader)
	if userID == "" {
		return "", ErrUserIDMissing
	}
	return userID, nil
}
