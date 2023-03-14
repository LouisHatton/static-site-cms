package main

import (
	"net/http"

	"github.com/LouisHatton/static-site-cms/cmd/admin-server/endpoints"
	"github.com/LouisHatton/static-site-cms/pkg/log"
)

var logger = log.Logger()

func main() {
	// Handle all endpoints
	for _, endpoint := range endpoints.AllEndpoints {
		endpoint.Handle()
	}

	// Listen on port 8080
	logger.Info("Starting server on 0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}
