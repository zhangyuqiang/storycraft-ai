package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"storycraft-ai/backend/config"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		resp := HealthResponse{Status: "ok", Service: "storycraft-ai"}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("encode response failed: %v", err)
		}
	})

	port := config.GetPort()
	log.Printf("storycraft-ai backend listening on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Printf("server error: %v", err)
		os.Exit(1)
	}
}
