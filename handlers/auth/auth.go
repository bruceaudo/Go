package auth

import (
	"fmt"
	"net/http"

	"github.com/bruceaudo/app/utils"
)

type User struct {
	FullName string `json:"fullName"` 
	Email string `json:"email"` 
	Password string `json:"password"` 
	IsVerified string `json:"isVerified"`
}

func LoginFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Logged in successfully")
}

func SignupFunc(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(5 << 20); err != nil{
		utils.JsonError(w, "Failed to parse multi-part form", http.StatusInternalServerError)
	}

	fmt.Fprint(w, "Signed up successfully")
}
