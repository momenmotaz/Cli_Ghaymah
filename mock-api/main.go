package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	mockToken = "test-token-123"
)

type DeployRequest struct {
	Name      string            `json:"name"`
	Image     string            `json:"image"`
	Resources Resources         `json:"resources"`
	EnvVars   map[string]string `json:"envVars"`
}

type Resources struct {
	CPU     string `json:"cpu"`
	Memory  string `json:"memory"`
	Storage string `json:"storage"`
}

type DeployResponse struct {
	AppID   string `json:"appId"`
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

	// Generate a unique ID for the deployment
	deploymentID := fmt.Sprintf("deploy-%d", time.Now().Unix())

	resp := DeployResponse{
		AppID:   deploymentID,
		Status:  "deploying",
		Message: fmt.Sprintf("Successfully deployed %s from image %s", req.Name, req.Image),
	}

	w.Header().Set("Content-Type", "application/json")
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

	w.Header().Set("Content-Type", "application/json")
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

	resp := LogsResponse{
		Logs: []string{
			fmt.Sprintf("[%s] Application started", time.Now().Format(time.RFC3339)),
			fmt.Sprintf("[%s] Listening on port 80", time.Now().Format(time.RFC3339)),
			fmt.Sprintf("[%s] Received incoming request", time.Now().Format(time.RFC3339)),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// Create a channel to handle graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Create server mux
	mux := http.NewServeMux()
	mux.HandleFunc("/apps", deployHandler)
	mux.HandleFunc("/apps/status", statusHandler)
	mux.HandleFunc("/apps/logs", logsHandler)

	// Create server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start server in a goroutine
	go func() {
		fmt.Println("\n=== Ghaymah Mock API Server ===")
		fmt.Println("Starting server on http://localhost:8080")
		fmt.Println("Press Ctrl+C to stop the server...")
		fmt.Printf("Test token: %s\n\n", mockToken)
		
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v\n", err)
		}
	}()

	// Wait for interrupt signal
	<-stop
	fmt.Println("\nReceived interrupt signal...")
	fmt.Println("Shutting down server gracefully...")

	// Create a deadline for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error during server shutdown: %v\n", err)
	}
	
	fmt.Println("Server stopped successfully!")
}
