package types

import "net/http"

type Endpoint struct {
	Description string `json:"description"`
	Route       string `json:"route"`
	Handler     func(http.ResponseWriter, *http.Request)
}
