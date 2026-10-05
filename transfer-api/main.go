package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
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

func BuildConfig(l *zap.SugaredLogger) config {
	defaultConfig := config{
		serverPort:  8080,
		serviceName: "transfer-api",
	}

	if stringPort := os.Getenv("PORT"); stringPort != "" {
		if port, err := strconv.Atoi(stringPort); err != nil {
			l.Fatalf("%s", fmt.Errorf("%w: PORT env var invalid", err).Error())
		} else {
			if port > 65535 || port < 0 {
				l.Fatalf("PORT env var is out of range for valid ports")
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
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("%w: unable to initialize logger", err)
	}

	l := logger.Sugar()
	config := BuildConfig(l)

	r := mux.NewRouter()
	r.HandleFunc("/health", HealthHandler).Methods("GET")
	r.HandleFunc("/ready", ReadinessHandler).Methods("GET")

	l.Infof("Starting service %s at %d", config.serviceName, config.serverPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.serverPort), r); err != nil {
		l.Fatalf("%w: unable to start service", err)
	}

	l.Infof("Shuting down service %s", config.serviceName)
}
