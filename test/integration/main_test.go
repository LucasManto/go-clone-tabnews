package main_test

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"
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
	var responseBody map[string]any
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		t.Fatal(err)
	}

	updatedAt, ok := responseBody["updated_at"].(string)
	if !ok {
		t.Fatal("expected response body to contain \"updated_at\" as string value")
	}
	_, err = time.Parse(time.RFC3339, updatedAt)
	if err != nil {
		t.Error("expected \"updated_at\" to be a valid datetime", err)
	}

	dependencies, ok := responseBody["dependencies"].(map[string]any)
	if !ok {
		t.Fatal("expected response body to contain \"dependencies\" as object value")
	}
	database, ok := dependencies["database"].(map[string]any)
	if !ok {
		t.Fatal("expected \"dependencies\" to contain \"database\" as object value")
	}
	if database["database_version"] != "16.0" {
		t.Error("expected \"database_version\" to be \"16.0\", got", database["database_version"])
	}
	if database["max_connections"] != float64(100) {
		t.Error("expected \"max_connections\" to be 100, got", database["max_connections"])
	}
	if database["open_connections"] != float64(1) {
		t.Error("expected \"open_connections\" to be 1, got", database["open_connections"])
	}
}
