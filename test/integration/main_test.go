package main_test

import (
	"net/http"
	"testing"
)

func TestV1Status(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/api/v1/status")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
