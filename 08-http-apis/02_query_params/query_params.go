package query_params

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// SearchResult represents a search query with pagination
type SearchResult struct {
	Query string `json:"query"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

// FilterResult represents filtered tags
type FilterResult struct {
	Tags []string `json:"tags"`
}

// SearchHandler handles search requests with pagination
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
	params := r.URL.Query()
	query := params.Get("q")
	if query == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// var err error
	res := SearchResult{
		Query: query,
	}
	page := params.Get("page")

	pageVal, pageErr := strconv.Atoi(page)
	if pageErr != nil {
		res.Page = 1
	} else {
		res.Page = pageVal
	}

	limit := params.Get("limit")
	limitVal, limitErr := strconv.Atoi(limit)
	if limitErr != nil {
		res.Limit = 10
	} else {
		res.Limit = limitVal
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// FilterHandler handles filtering by multiple tags
func FilterHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(human): Implement
	result := FilterResult{
		Tags: make([]string, 0),
	}

	tags := r.URL.Query()["tag"]
	result.Tags = append(result.Tags, tags...)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
