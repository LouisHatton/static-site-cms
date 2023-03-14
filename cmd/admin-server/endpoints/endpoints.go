package endpoints

import "net/http"

type Endpoint struct {
	Description string `json:"description"`
	Route       string `json:"route"`
	Handler     func(http.ResponseWriter, *http.Request)
}

func (e *Endpoint) Handle() {
	http.HandleFunc(e.Route, e.Handler)
}

var UploadZip = Endpoint{
	"When a user uploads a zip, fire this endpoint",
	"/upload/full-site",
	UploadHandler,
}

var AllEndpoints map[string]Endpoint = map[string]Endpoint{
	"uploadZip": UploadZip,
}
