package main

import (
	"net/http"
	"os"

	"github.com/LouisHatton/static-site-cms/pkg/files"
	"github.com/LouisHatton/static-site-cms/pkg/log"
)

var logger = log.Logger()

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		logger.Info("Received file upload request")
		if err := files.UploadFile(r); err != nil {
			logger.Error("error parsing uploaded file: ", err.Error())
			http.Error(w, "error parsing uploaded file", http.StatusInternalServerError)
			return
		}

		// Clear the http directory
		os.RemoveAll("/http/")
		os.MkdirAll("/http/", 0777)

		if err := files.Unzip("/tmp/upload", "/http/"); err != nil {
			logger.Error("error unzipping uploaded file: ", err.Error())
			http.Error(w, "error unzipping uploaded file", http.StatusInternalServerError)
			return
		}
		logger.Info("New files unzipped and deployed in /http")
	}
}

func main() {
	// Upload route
	http.HandleFunc("/upload/full-site", uploadHandler)

	//Listen on port 8080
	logger.Info("Starting server on 0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}
