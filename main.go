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

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	mux.HandleFunc("/secrets", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(
			fmt.Sprintf("This came from secrets: %s", os.Getenv("SECRET")),
		))
	})

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
}
