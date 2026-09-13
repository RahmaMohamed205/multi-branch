package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	branch := os.Getenv("BRANCH_NAME")
	if branch == "" {
		branch = "unknown"
	}

	response := map[string]string{
		"message": "Hello from the multibranch pipeline app!",
		"branch":  branch,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)

	log.Println("App running on port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
