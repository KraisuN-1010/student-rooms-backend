package main

import (
	"log"
	"net/http"

	"github.com/KraisuN-1010/student-rooms-backend/internal/websocket"
)

func main() {
	// Create WebSocket hub
	hub := websocket.NewHub()
	go hub.Run()

	// WebSocket endpoint
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWS(hub, w, r)
	})

	// Health check for WebSocket server
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("WebSocket server is running"))
	})

	log.Println("WebSocket server starting on port 8001")
	log.Fatal(http.ListenAndServe(":8001", nil))
}
