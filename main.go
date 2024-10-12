package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bruceaudo/app/handlers/auth"
	"github.com/bruceaudo/app/handlers/upload"
)

const (
	address string = "localhost:8080"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/login", auth.LoginFunc)
	mux.HandleFunc("/signup", auth.SignupFunc)
	mux.HandleFunc("/upload", upload.UploadFunc)

	server := &http.Server{
		Addr:           address,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Server starting on http://%s\n", address)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server => ", err)
	}
}
