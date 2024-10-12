package parsers

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/bruceaudo/app/utils"
)

func ParsePDF() {}

func ParseDOCX() {}


// Function that parses files with .txt extension
func ParseTXT(w http.ResponseWriter, file multipart.File) {
	var buf bytes.Buffer

	var err error

	if _, err = io.Copy(&buf, file); err != nil {
		utils.JsonError(w, "Error reading multi-part file contents into buffer", http.StatusInternalServerError)
		return 
	}

	fmt.Printf("File contents: %s\n", buf.String())
}
