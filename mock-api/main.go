package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	mockToken = "test-token-123"
)

type DeployRequest struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type DeployResponse struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type StatusResponse struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	CPU     string `json:"cpu"`
	Memory  string `json:"memory"`
	Uptime  string `json:"uptime"`
	Message string `json:"message"`
}

type LogsResponse struct {
	Logs []string `json:"logs"`
}

func validateToken(r *http.Request) bool {
	token := r.Header.Get("Authorization")
	return strings.TrimPrefix(token, "Bearer ") == mockToken
}

func deployHandler(w http.ResponseWriter, r *http.Request) {
	if !validateToken(r) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req DeployRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := DeployResponse{
		ID:      "app-123",
		Status:  "deploying",
		Message: fmt.Sprintf("Deploying %s from image %s", req.Name, req.Image),
	}

	json.NewEncoder(w).Encode(resp)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	if !validateToken(r) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Name parameter is required", http.StatusBadRequest)
		return
	}

	resp := StatusResponse{
		Name:    name,
		Status:  "running",
		CPU:     "25%",
		Memory:  "128MB",
		Uptime:  "2h 15m",
		Message: "Application is running normally",
	}

	json.NewEncoder(w).Encode(resp)
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	if !validateToken(r) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Name parameter is required", http.StatusBadRequest)
		return
	}

	// Generate some mock logs
	logs := []string{
		fmt.Sprintf("[%s] Application started", time.Now().Format(time.RFC3339)),
		fmt.Sprintf("[%s] Connecting to database", time.Now().Add(-1*time.Minute).Format(time.RFC3339)),
		fmt.Sprintf("[%s] Database connected successfully", time.Now().Add(-2*time.Minute).Format(time.RFC3339)),
		fmt.Sprintf("[%s] Listening on port 8080", time.Now().Add(-3*time.Minute).Format(time.RFC3339)),
	}

	resp := LogsResponse{
		Logs: logs,
	}

	json.NewEncoder(w).Encode(resp)
}

func main() {
	mux := http.NewServeMux()
	
	// API endpoints
	mux.HandleFunc("/apps", deployHandler)
	mux.HandleFunc("/apps/status", statusHandler)
	mux.HandleFunc("/apps/logs", logsHandler)

	port := ":8080"
	fmt.Printf("Mock API server starting on http://localhost%s\n", port)
	fmt.Printf("Use token: %s\n", mockToken)
	log.Fatal(http.ListenAndServe(port, mux))
}
