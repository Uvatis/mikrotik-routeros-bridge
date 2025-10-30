package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-routeros/routeros"
)

// MikroTikRequest represents the input parameters required for interacting with a MikroTik router through an HTTP request.
type MikroTikRequest struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Command  string `json:"command"`
}

// connectHandler handles HTTP requests to establish a connection to a MikroTik router using provided JSON payload.
// It expects a JSON containing host, port, user, and password to initiate the connection.
// Responds with a JSON confirming the connection status or an error with the appropriate HTTP status code.
func connectHandler(w http.ResponseWriter, r *http.Request) {
	var req MikroTikRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	address := req.Host + ":" + req.Port
	conn, err := routeros.Dial(address, req.User, req.Password)
	if err != nil {
		http.Error(w, "connection failed: "+err.Error(), http.StatusBadGateway)
		return
	}
	defer conn.Close()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "connected",
	})
}

// commandHandler handles HTTP requests to execute commands on a MikroTik router and returns the response in JSON format.
// It expects a JSON payload containing connection details (host, port, user, password) and a command to execute.
// In case of errors (e.g., invalid JSON, connection failure, command failure), it responds with an appropriate HTTP status code.
func commandHandler(w http.ResponseWriter, r *http.Request) {
	var req MikroTikRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	address := req.Host + ":" + req.Port
	conn, err := routeros.Dial(address, req.User, req.Password)
	if err != nil {
		http.Error(w, "connection failed: "+err.Error(), http.StatusBadGateway)
		return
	}
	defer conn.Close()

	reply, err := conn.Run(req.Command)
	if err != nil {
		http.Error(w, "command failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reply.Re)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/connect", connectHandler)
	mux.HandleFunc("/command", commandHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("MikroTik bridge running on port 8080")
	log.Fatal(server.ListenAndServe())
}
