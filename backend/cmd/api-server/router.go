package main

import (
	"github.com/CzarSimon/clother/backend/pkg/httputil"
	"github.com/gin-gonic/gin"
)

// newRouter sets up a router add registers handlers for it.
func newRouter(config Config) *gin.Engine {
	r := httputil.DefaultRouter()

	apiV1 := r.Group("/api/v1", httputil.RequireAuthToken(config.AuthToken))
	apiV1.GET("/search", searchController)

	return r
}
