package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-routeros/routeros/v3"
)

// MikroTikRequest represents the input parameters required for interacting with a MikroTik router through an HTTP request.
type MikroTikRequest struct {
	Host     string            `json:"host"`
	Port     string            `json:"port"`
	User     string            `json:"user"`
	Password string            `json:"password"`
	Command  string            `json:"command"`
	Payload  map[string]string `json:"payload,omitempty"`
}

// connectHandler handles HTTP requests to establish a connection to a MikroTik router using provided JSON payload.
// It expects a JSON containing host, port, user, and password to initiate the connection.
// Responds with a JSON confirming the connection status or an error with the appropriate HTTP status code.
func connectHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting to connect to MikroTik router...")
	var req MikroTikRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errMsg := fmt.Sprintf("JSON decode error: %v", err)
		log.Println(errMsg)
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	log.Printf("Connection requested - Host: %s, Port: %s, User: %s", req.Host, req.Port, req.User)

	address := req.Host + ":" + req.Port
	conn, err := routeros.Dial(address, req.User, req.Password)
	if err != nil {
		http.Error(w, "connection failed: "+err.Error(), http.StatusBadGateway)
		return
	}
	defer conn.Close()

	w.Header().Set("Content-Type", "application/json")
	log.Println("Sending successful connection response")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "connected",
	})
}

// commandHandler handles HTTP requests to execute commands on a MikroTik router and returns the response in JSON format.
// It expects a JSON payload containing connection details (host, port, user, password) and a command to execute.
// In case of errors (e.g., invalid JSON, connection failure, command failure), it responds with an appropriate HTTP status code
func commandHandler(w http.ResponseWriter, r *http.Request) {
	var req MikroTikRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		log.Printf("[ERROR] invalid JSON: %v", err)
		return
	}

	address := req.Host + ":" + req.Port
	conn, err := routeros.Dial(address, req.User, req.Password)
	if err != nil {
		http.Error(w, "connection failed", http.StatusBadGateway)
		log.Printf("[ERROR] connection to %s failed: %v", address, err)
		return
	}
	defer conn.Close()

	var reply *routeros.Reply
	if len(req.Payload) > 0 {
		reply, err = conn.RunArgs(buildArgs(req.Command, req.Payload))
	} else {
		reply, err = conn.Run(req.Command)
	}

	if err != nil {
		http.Error(w, "command failed", http.StatusInternalServerError)
		payloadJSON, _ := json.Marshal(req.Payload)
		log.Printf("[ERROR] command failed: %v | command=%s | payload=%s", err, req.Command, string(payloadJSON))
		return
	}

	var results []map[string]string
	for _, re := range reply.Re {
		if re.Word == "!re" && re.Map != nil {
			results = append(results, re.Map)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func buildArgs(command string, payload map[string]string) []string {
	args := []string{command}
	for k, v := range payload {
		args = append(args, "="+k+"="+v)
	}
	return args
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

	log.Println("Starting MikroTik Bridge server...")
	log.Printf("Server listening on port %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
