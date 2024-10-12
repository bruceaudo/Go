package upload

import (
	"encoding/json"
	"net/http"
	"path"
	"slices"

	"github.com/bruceaudo/app/utils"
	"github.com/bruceaudo/app/utils/parsers"
)

const (
	maxUploadSize = 50 << 20 // 50 MB
)

var (
	allowedFileExtensions = []string{".pdf", ".docx", ".txt"}
)

func UploadFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.JsonError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		utils.JsonError(w, "Error parsing multi-part form", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		utils.JsonError(w, "Error parsing file => "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	if handler.Size > maxUploadSize {
		utils.JsonError(w, "File too large", http.StatusBadRequest)
		return
	}

	fileExt := path.Ext(handler.Filename)

	if !slices.Contains(allowedFileExtensions, fileExt) {
		utils.JsonError(w, "Unsupported file type", http.StatusUnsupportedMediaType)
		return
	}

	switch fileExt {
	case ".pdf":

		parsers.ParsePDF()

		sendResponse(w, ".pdf file uploaded successfully")
	case ".docx":

		parsers.ParseDOCX()

		sendResponse(w, ".docx file uploaded successfully")
	case ".txt":

		parsers.ParseTXT(w, file)
		sendResponse(w, ".txt file uploaded successfully")
	}

}

func sendResponse(w http.ResponseWriter, message string) {
	response := utils.Response{
		Message: message,
		Status:  http.StatusOK,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.JsonError(w, "File upload failed", http.StatusInternalServerError)
		return
	}
}
