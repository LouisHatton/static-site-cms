package main

import (
	"net/http"
	"os"

	"github.com/LouisHatton/static-site-cms/pkg/log"
)

func main() {
	log.Logger().Info("Starting server on 0.0.0.0:3000")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(res http.ResponseWriter, req *http.Request) {
	fileDir := os.Getenv("SERVER_FILES")
	if fileDir == "" {
		fileDir = "./static"
	}
	file := http.FileServer(http.Dir(fileDir))
	file.ServeHTTP(res, req)
	log.Logger().Infow("Request received",
		"method", req.Method,
		"path", req.URL.Path,
	)
}
