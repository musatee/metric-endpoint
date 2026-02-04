package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type ErrorResponse struct {
	ErrorRate float32 `json:"error_rate"`
	Message   string  `json:"message"`
}

func getErrorRate(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	response := ErrorResponse{
		ErrorRate: 70.0,
		Message:   "Current error rate",
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	log.Println(r.Method, r.URL.Path, time.Since(start))
	if err != nil {
		http.Error(w, "Failed to process response", http.StatusInternalServerError)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/apis/error", http.HandlerFunc(getErrorRate))
	port := ":8080"
	err := http.ListenAndServe(port, mux)
	if err != nil {
		fmt.Println("Something went wrong. Failed to start the server")
	}
}
