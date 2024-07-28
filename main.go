package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func secretHandler(w http.ResponseWriter, r *http.Request) {
	m := fmt.Sprintf("The secret is: %s", os.Getenv("SECRET"))
	w.Write([]byte(m))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloHandler)

	mux.HandleFunc("/secrets", secretHandler)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
}
