package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey TOKEN_STRING")
	_, err := auth.GetAPIKey(header)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

func TestMalformedGetAPIKey(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "TOKEN_STRING")
	_, err := auth.GetAPIKey(header)
	if err == nil {
		t.Error("Expecting error: Malformed Token")
		return
	}
}
