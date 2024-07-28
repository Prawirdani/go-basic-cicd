package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	helloHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	greeting, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := "Hello World!"
	if string(greeting) != expected {
		t.Errorf("expected %q; got %q", expected, greeting)
	}
}

func TestSecretHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/secrets", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	secretHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	secret, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := fmt.Sprintf("The secret is: %s", os.Getenv("SECRET"))
	if string(secret) != expected {
		t.Errorf("expected %q; got %q", expected, secret)
	}
}
