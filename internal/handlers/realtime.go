package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/KraisuN-1010/student-rooms-backend/internal/websocket"
	"gofr.dev/pkg/gofr"
)

type RealtimeHandler struct {
	hub *websocket.Hub
}

func NewRealtimeHandler(hub *websocket.Hub) *RealtimeHandler {
	return &RealtimeHandler{hub: hub}
}

// BroadcastMessage broadcasts a message to all clients in a room
func (h *RealtimeHandler) BroadcastMessage(c *gofr.Context) (interface{}, error) {
	var req struct {
		RoomID  string      `json:"room_id" validate:"required"`
		Type    string      `json:"type" validate:"required"`
		Content string      `json:"content"`
		Data    interface{} `json:"data,omitempty"`
	}

	if err := c.Bind(&req); err != nil {
		return nil, err
	}

	// Create message
	message := websocket.Message{
		Type:    req.Type,
		RoomID:  req.RoomID,
		Content: req.Content,
		Data:    req.Data,
	}

	// Marshal message to JSON
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal message: %v", err)
	}

	// Broadcast to room
	h.hub.BroadcastToRoom(req.RoomID, messageBytes)

	return map[string]string{
		"status":  "success",
		"message": "Message broadcasted successfully",
	}, nil
}

// GetRoomStats returns statistics about a room
func (h *RealtimeHandler) GetRoomStats(c *gofr.Context) (interface{}, error) {
	roomID := c.PathParam("roomId")
	if roomID == "" {
		return nil, fmt.Errorf("room ID is required")
	}

	clientCount := h.hub.GetRoomClientCount(roomID)

	return map[string]interface{}{
		"room_id":       roomID,
		"active_users":  clientCount,
		"status":        "active",
	}, nil
}

// BroadcastPostCreated broadcasts when a new post is created
func (h *RealtimeHandler) BroadcastPostCreated(roomID string, post interface{}) {
	message := websocket.Message{
		Type:   "post_created",
		RoomID: roomID,
		Data:   post,
	}

	messageBytes, _ := json.Marshal(message)
	h.hub.BroadcastToRoom(roomID, messageBytes)
}

// BroadcastCommentCreated broadcasts when a new comment is created
func (h *RealtimeHandler) BroadcastCommentCreated(roomID string, comment interface{}) {
	message := websocket.Message{
		Type:   "comment_created",
		RoomID: roomID,
		Data:   comment,
	}

	messageBytes, _ := json.Marshal(message)
	h.hub.BroadcastToRoom(roomID, messageBytes)
}

// BroadcastDoubtCreated broadcasts when a new doubt is created
func (h *RealtimeHandler) BroadcastDoubtCreated(roomID string, doubt interface{}) {
	message := websocket.Message{
		Type:   "doubt_created",
		RoomID: roomID,
		Data:   doubt,
	}

	messageBytes, _ := json.Marshal(message)
	h.hub.BroadcastToRoom(roomID, messageBytes)
}

// BroadcastUserJoined broadcasts when a user joins a room
func (h *RealtimeHandler) BroadcastUserJoined(roomID string, user interface{}) {
	message := websocket.Message{
		Type:   "user_joined",
		RoomID: roomID,
		Data:   user,
	}

	messageBytes, _ := json.Marshal(message)
	h.hub.BroadcastToRoom(roomID, messageBytes)
}

// BroadcastUserLeft broadcasts when a user leaves a room
func (h *RealtimeHandler) BroadcastUserLeft(roomID string, user interface{}) {
	message := websocket.Message{
		Type:   "user_left",
		RoomID: roomID,
		Data:   user,
	}

	messageBytes, _ := json.Marshal(message)
	h.hub.BroadcastToRoom(roomID, messageBytes)
}
