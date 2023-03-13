package files

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/LouisHatton/static-site-cms/pkg/log"
)

var logger = log.Logger()

// Takes a http request and uploads and saves the files to the /tmp/uploads/upload directory
func UploadFile(r *http.Request) error {
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)
	defer r.Body.Close()

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("file")
	if err != nil {
		return fmt.Errorf("unable to get form file with name 'file': " + err.Error())
	}

	defer file.Close()
	logger.Debugf("Uploaded File: %+v\n", handler.Filename)
	logger.Debugf("File Size: %+v\n", handler.Size)
	logger.Debugf("MIME Header: %+v\n", handler.Header)

	// Create file
	dst, err := os.Create("/tmp/upload")
	if err != nil {
		return fmt.Errorf("unable to create a file at '/tmp/uploads/upload': " + err.Error())
	}
	defer dst.Close()

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		return fmt.Errorf("unable to copy the uploaded file onto the file system %s: "+err.Error(), handler.Filename)
	}

	return nil
}
