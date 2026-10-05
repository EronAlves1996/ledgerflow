package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthHandler).Methods("GET")
}
