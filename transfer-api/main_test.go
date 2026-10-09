package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"
)

func TestHealthEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	logger, err := zap.NewProduction()
	if err != nil {
		t.Fatal("failure initializing the logger")
	}

	r := routes(logger.Sugar())

	r.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("want %d, have %d", http.StatusOK, res.StatusCode)
	}

	data := struct {
		Status string
	}{}

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		t.Errorf("%s: error expected to be nil", err.Error())
	}

	if data.Status != "Ok" {
		t.Errorf("want %s, have %s", "Ok", data.Status)
	}
}

func TestReadinessEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ready", nil)
	w := httptest.NewRecorder()
	logger, err := zap.NewProduction()
	if err != nil {
		t.Fatal("failure initializing the logger")
	}

	r := routes(logger.Sugar())
	r.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("want %d, have %d", http.StatusOK, res.StatusCode)
	}

	data := struct {
		Status string
	}{}

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		t.Errorf("%s: error expected to be nil", err.Error())
	}

	if data.Status != "Ready" {
		t.Errorf("want %s, have %s", "Ready", data.Status)
	}
}
