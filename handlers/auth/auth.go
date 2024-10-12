package auth

import (
	"fmt"
	"net/http"
)

func LoginFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Logged in successfully")
}

func SignupFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Signed up successfully")
}
