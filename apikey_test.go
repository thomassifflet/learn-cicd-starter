package main

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestApiKey(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey ")
	expected := "lolimanapikey"
	apiKey, err := auth.GetAPIKey(headers)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)

	}
	if apiKey != expected {
		t.Errorf("Expected %v, got %v", expected, apiKey)
	}
}
