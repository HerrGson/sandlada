package main

import (
	"net/http"

	"github.com/CzarSimon/clother/backend/pkg/httputil"
	"github.com/gin-gonic/gin"
)

// SearchResponse group of a search query and a list of results.
type SearchResponse struct {
	Query   string   `json:"query"`
	Results []string `json:"results"`
}

// NewSearchResponse creates a new serach response.
func NewSearchResponse(query string) SearchResponse {
	return SearchResponse{
		Query:   query,
		Results: make([]string, 0),
	}
}

// searchController controller for handling search queries.
func searchController(c *gin.Context) {
	query, err := httputil.ParseQueryValue(c, "query")
	if err != nil {
		c.Error(httputil.ErrBadRequest)
		return
	}

	c.JSON(http.StatusOK, NewSearchResponse(query))
}
