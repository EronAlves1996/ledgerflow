package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Status string
	}{
		Status: "Ok",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(body); err != nil {
		fmt.Println(fmt.Errorf("%w: error while encoding json", err))
	}
}

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Status string
	}{
		Status: "Ready",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(body); err != nil {
		fmt.Println(fmt.Errorf("%w: error while encoding json", err))
	}
}

type config struct {
	serverPort  int
	serviceName string
}

func BuildConfig() config {
	defaultConfig := config{
		serverPort:  8080,
		serviceName: "transfer-api",
	}

	if stringPort := os.Getenv("PORT"); stringPort != "" {
		if port, err := strconv.Atoi(stringPort); err != nil {
			log.Fatalf("%s", fmt.Errorf("%w: PORT env var invalid", err).Error())
		} else {
			if port > 65535 || port < 0 {
				log.Fatalf("PORT env var is out of range for valid ports")
			}
			defaultConfig.serverPort = port
		}
	}

	if serviceName := os.Getenv("SERVICE_NAME"); serviceName != "" {
		defaultConfig.serviceName = serviceName
	}

	return defaultConfig
}

func main() {
	config := BuildConfig()

	r := mux.NewRouter()
	r.HandleFunc("/health", HealthHandler).Methods("GET")
	r.HandleFunc("/readiness", ReadinessHandler).Methods("GET")
}
