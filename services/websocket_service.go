package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type WebSocketService struct {
	wsURL string
}

func NewWebSocketService() *WebSocketService {
	return &WebSocketService{
		wsURL: "http://localhost:8001", // WebSocket server URL
	}
}

type WebSocketMessage struct {
	Type    string      `json:"type"`
	RoomID  string      `json:"room_id"`
	UserID  string      `json:"user_id"`
	Content string      `json:"content"`
	Data    interface{} `json:"data"`
}

// BroadcastToRoom sends a message to all clients in a specific room
func (ws *WebSocketService) BroadcastToRoom(roomID, messageType, content string, data interface{}) error {
	message := WebSocketMessage{
		Type:    messageType,
		RoomID:  roomID,
		Content: content,
		Data:    data,
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	// Send to WebSocket server
	resp, err := http.Post(ws.wsURL+"/broadcast", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("WebSocket server returned status: %d", resp.StatusCode)
	}

	return nil
}

// BroadcastPostCreated notifies when a new post is created
func (ws *WebSocketService) BroadcastPostCreated(roomID string, post interface{}) {
	ws.BroadcastToRoom(roomID, "post_created", "New post created", post)
}

// BroadcastCommentCreated notifies when a new comment is created
func (ws *WebSocketService) BroadcastCommentCreated(roomID string, comment interface{}) {
	ws.BroadcastToRoom(roomID, "comment_created", "New comment added", comment)
}

// BroadcastDoubtCreated notifies when a new doubt is created
func (ws *WebSocketService) BroadcastDoubtCreated(roomID string, doubt interface{}) {
	ws.BroadcastToRoom(roomID, "doubt_created", "New doubt posted", doubt)
}

// BroadcastUserJoined notifies when a user joins a room
func (ws *WebSocketService) BroadcastUserJoined(roomID string, user interface{}) {
	ws.BroadcastToRoom(roomID, "user_joined", "User joined the room", user)
}

// BroadcastUserLeft notifies when a user leaves a room
func (ws *WebSocketService) BroadcastUserLeft(roomID string, user interface{}) {
	ws.BroadcastToRoom(roomID, "user_left", "User left the room", user)
}

// BroadcastTyping notifies when someone is typing
func (ws *WebSocketService) BroadcastTyping(roomID, userID string) {
	ws.BroadcastToRoom(roomID, "typing", userID+" is typing...", map[string]string{
		"user_id": userID,
		"room_id": roomID,
	})
}

// GetRoomStats gets statistics about a room
func (ws *WebSocketService) GetRoomStats(roomID string) (map[string]interface{}, error) {
	resp, err := http.Get(ws.wsURL + "/room/" + roomID + "/stats")
	if err != nil {
		return nil, fmt.Errorf("failed to get room stats: %v", err)
	}
	defer resp.Body.Close()

	var stats map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, fmt.Errorf("failed to decode stats: %v", err)
	}

	return stats, nil
}
