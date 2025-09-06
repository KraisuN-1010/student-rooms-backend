package websocket

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// Hub maintains the set of active clients and broadcasts messages to the clients
type Hub struct {
	// Registered clients
	clients map[*Client]bool

	// Inbound messages from the clients
	broadcast chan []byte

	// Register requests from the clients
	register chan *Client

	// Unregister requests from clients
	unregister chan *Client

	// Room-specific clients
	roomClients map[string]map[*Client]bool

	// Mutex for thread-safe operations
	mutex sync.RWMutex
}

// Client represents a websocket client
type Client struct {
	// The websocket connection
	conn *websocket.Conn

	// Buffered channel of outbound messages
	send chan []byte

	// User ID
	userID string

	// Room ID the client is currently in
	roomID string

	// Hub reference
	hub *Hub
}

// Message represents a real-time message
type Message struct {
	Type    string      `json:"type"`
	RoomID  string      `json:"room_id,omitempty"`
	UserID  string      `json:"user_id,omitempty"`
	Data    interface{} `json:"data"`
	Content string      `json:"content,omitempty"`
}

// NewHub creates a new WebSocket hub
func NewHub() *Hub {
	return &Hub{
		clients:     make(map[*Client]bool),
		broadcast:   make(chan []byte),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		roomClients: make(map[string]map[*Client]bool),
	}
}

// Run starts the hub
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mutex.Lock()
			h.clients[client] = true
			if client.roomID != "" {
				if h.roomClients[client.roomID] == nil {
					h.roomClients[client.roomID] = make(map[*Client]bool)
				}
				h.roomClients[client.roomID][client] = true
			}
			h.mutex.Unlock()
			log.Printf("Client connected. Total clients: %d", len(h.clients))

		case client := <-h.unregister:
			h.mutex.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				if client.roomID != "" {
					if roomClients, exists := h.roomClients[client.roomID]; exists {
						delete(roomClients, client)
						if len(roomClients) == 0 {
							delete(h.roomClients, client.roomID)
						}
					}
				}
				close(client.send)
			}
			h.mutex.Unlock()
			log.Printf("Client disconnected. Total clients: %d", len(h.clients))

		case message := <-h.broadcast:
			h.mutex.RLock()
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
			h.mutex.RUnlock()
		}
	}
}

// BroadcastToRoom sends a message to all clients in a specific room
func (h *Hub) BroadcastToRoom(roomID string, message []byte) {
	h.mutex.RLock()
	if roomClients, exists := h.roomClients[roomID]; exists {
		for client := range roomClients {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(h.clients, client)
				delete(roomClients, client)
			}
		}
	}
	h.mutex.RUnlock()
}

// GetRoomClientCount returns the number of clients in a room
func (h *Hub) GetRoomClientCount(roomID string) int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	if roomClients, exists := h.roomClients[roomID]; exists {
		return len(roomClients)
	}
	return 0
}
