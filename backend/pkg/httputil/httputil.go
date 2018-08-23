package httputil

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DefaultRouter returns a gin router with the default middleware injected.
func DefaultRouter() *gin.Engine {
	r := gin.Default()
	r.Use(ErrorMiddleware())

	r.GET("/health", HandleHealthCheck)
	return r
}

// SendOK sends an ok status and message to the client.
func SendOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

// HandleHealthCheck responds possitively to a health check.
func HandleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}

// ParseQueryValue parses a query value from request.
func ParseQueryValue(c *gin.Context, key string) (string, error) {
	value, ok := c.GetQuery(key)
	if !ok {
		return "", NewError(http.StatusBadRequest,
			fmt.Sprintf("No value found for key: %s", key))
	}
	return value, nil
}

// ParseQueryValues parses query values from a request
func ParseQueryValues(c *gin.Context, key string) ([]string, error) {
	values, ok := c.GetQueryArray(key)
	if !ok {
		return nil, NewError(http.StatusBadRequest,
			fmt.Sprintf("No values found for key: %s", key))
	}
	return values, nil
}
