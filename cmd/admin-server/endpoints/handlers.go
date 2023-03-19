package endpoints

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/LouisHatton/static-site-cms/pkg/files"
	"github.com/LouisHatton/static-site-cms/pkg/log"
	"github.com/LouisHatton/static-site-cms/pkg/utils"
)

var logger = log.Logger()

// Main route for when user hits /
func MainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		logger.Info("Received main route, for testing")
	}
}

// Upload Zip endpoint for when a user uploads their
// zip containing the HTML files
func UploadHandler(w http.ResponseWriter, r *http.Request) {
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

func GetSiteStructure(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	logger.Info("Received get site structure request")
	siteMap, err := files.GenerateDirectoryMap(utils.GetEnv("SERVER_FILES", "/http"))
	if err != nil {
		logger.Error("error generating directory map: ", err.Error())
		http.Error(w, "error generating dir map", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(siteMap)
	if err != nil {
		logger.Error("error marshaling data: ", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}
