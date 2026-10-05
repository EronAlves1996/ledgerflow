package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

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

func routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthHandler).Methods("GET")
	r.HandleFunc("/ready", ReadinessHandler).Methods("GET")

	return r
}

func startServer(l *zap.SugaredLogger, config config, r *mux.Router) {
	l.Infof("Starting service %s at %d", config.serviceName, config.serverPort)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.serverPort),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			l.Fatalf(fmt.Errorf("unable to start service:%w", err).Error())
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	l.Info("Shuting down the server")
	if err := server.Shutdown(shutdownCtx); err != nil {
		l.Fatalf(fmt.Errorf("unable to shutdown server: %w", err).Error())
	}

	l.Info("Server shutdown")
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("%s: unable to initialize logger", err.Error())
	}

	l := logger.Sugar()
	config := BuildConfig(l)

	r := routes()
	startServer(l, config, r)
}
